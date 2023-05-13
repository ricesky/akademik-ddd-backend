package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/microsoft/go-mssqldb"

	"github.com/gin-gonic/gin"
	"its.id/akademik/presensi/infrastructure/handler"
	"its.id/akademik/presensi/infrastructure/sqlserver"
)

const (
	host     string = "10.199.14.47"
	port     int    = 1433
	database string = "PRESENSI_DDD"
	user     string = "presensi_ddd_app"
	password string = "presensi_ddd_app123#"
)

func main() {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", host, user, password, port, database)

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

	dosenQuery := sqlserver.NewSqlServerDosenQuery(db, ctx)
	dosenHandler := handler.NewDosenHandler(dosenQuery)

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
