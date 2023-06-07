package sqlserver

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"its.id/akademik/presensi/application/query"
)

type SqlServerRuanganQueryHandler struct {
	db  *sql.DB
	ctx context.Context
}

func NewSqlServerRuanganQueryHandler(db *sql.DB, ctx context.Context) query.RuanganQueryHandler {
	return &SqlServerRuanganQueryHandler{db, ctx}
}

func (s *SqlServerRuanganQueryHandler) GetAll() ([]query.RuanganQueryResult, error) {

	tsql := `SELECT r.id_ruangan, r.kode, r.nama
			FROM ruangan r
			ORDER BY r.kode ASC`

	rows, err := s.db.QueryContext(s.ctx, tsql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []query.RuanganQueryResult

	for rows.Next() {
		var (
			ruanganId uuid.UUID
			kode      string
			nama      string
		)

		err = rows.Scan(&ruanganId, &kode, &nama)
		if err != nil {
			return nil, err
		}

		result = append(result, query.RuanganQueryResult{RuanganId: ruanganId, Kode: kode, Nama: nama})
	}

	return result, nil
}
