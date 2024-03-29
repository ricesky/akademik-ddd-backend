package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"

	"github.com/gin-gonic/gin"
	"its.id/akademik/presensi/infrastructure/storage/sqlserver"
	ginhandler "its.id/akademik/presensi/presentation"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var (
		dbhost     string = os.Getenv("DB_HOST")
		dbport     string = os.Getenv("DB_PORT")
		dbname     string = os.Getenv("DB_NAME")
		dbuser     string = os.Getenv("DB_USER")
		dbpassword string = os.Getenv("DB_PASSWORD")
	)

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", dbhost, dbuser, dbpassword, dbport, dbname)

	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	dosenRepo := sqlserver.NewSqlServerDosenRepository(db, ctx)

	dosenId, _ := uuid.Parse("fb02200b-1ec8-4c24-a10c-000a9430d47e")
	d, _ := dosenRepo.FindById(dosenId)

	log.Println(d.Nama())

	dosenQuery := sqlserver.NewSqlServerDosenQueryHandler(db, ctx)
	dosenHandler := ginhandler.NewDosenHandler(dosenQuery)

	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/dosen", dosenHandler.GetDosen)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(router.Run(":" + port))
}
