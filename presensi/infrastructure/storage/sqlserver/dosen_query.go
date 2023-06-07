package sqlserver

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"its.id/akademik/presensi/application/query"
)

type SqlServerDosenQueryHandler struct {
	db  *sql.DB
	ctx context.Context
}

func NewSqlServerDosenQueryHandler(db *sql.DB, ctx context.Context) query.DosenQueryHandler {
	return &SqlServerDosenQueryHandler{db, ctx}
}

func (s *SqlServerDosenQueryHandler) GetByUserId(userId string) (*query.DosenQueryResult, error) {

	tsql := `SELECT d.id_dosen, d.nama
			 FROM dosen d
			 WHERE d.id_user_sso = @UserId AND d.deleted_at IS NULL`

	rows, err := s.db.QueryContext(s.ctx, tsql, sql.Named("UserId", userId))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id uuid.UUID
		var nama string
		err = rows.Scan(&id, &nama)
		if err != nil {
			return nil, err
		}

		return &query.DosenQueryResult{DosenId: id, Nama: nama}, nil
	}

	return nil, nil
}
