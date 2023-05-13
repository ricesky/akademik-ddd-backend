package sqlserver

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"its.id/akademik/presensi/application/query"
)

type SqlServerDosenQuery struct {
	db  *sql.DB
	ctx context.Context
}

func NewSqlServerDosenQuery(db *sql.DB, ctx context.Context) query.DosenQuery {
	return &SqlServerDosenQuery{db, ctx}
}

func (s *SqlServerDosenQuery) Execute(userId string) (*query.Dosen, error) {

	tsql := `SELECT d.id_dosen, d.nama
			 FROM dosen d
			 WHERE d.id_user_sso = @UserId AND d.deleted_at IS NULL`

	rows, err := s.db.QueryContext(s.ctx, tsql, sql.Named("UserId", userId))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id_dosen uuid.UUID
		var nama string
		err = rows.Scan(&id_dosen, &nama)
		if err != nil {
			return nil, err
		}

		return &query.Dosen{DosenId: id_dosen, Nama: nama}, nil
	}

	return nil, nil
}
