package sqlserver

import (
	"context"
	"database/sql"

	"its.id/akademik/presensi/application/query"
)

type SqlServerSemesterQueryHandler struct {
	db  *sql.DB
	ctx context.Context
}

func NewSqlServerSemesterQueryHandler(db *sql.DB, ctx context.Context) query.SemesterQueryHandler {
	return &SqlServerSemesterQueryHandler{db, ctx}
}

func (s *SqlServerSemesterQueryHandler) GetAktif() (*query.Semester, error) {

	tsql := `SELECT s.id_semester, s.nama, s.nama_en
			FROM semester s
			WHERE s.is_smt_aktif = 1`

	rows, err := s.db.QueryContext(s.ctx, tsql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id int
		var nama string
		err = rows.Scan(&id, &nama)
		if err != nil {
			return nil, err
		}

		return &query.Semester{SemesterId: id, Nama: nama}, nil
	}

	return nil, nil

}
