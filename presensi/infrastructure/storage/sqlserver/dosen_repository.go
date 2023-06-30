package sqlserver

import (
	"context"
	"database/sql"

	"its.id/akademik/presensi/domain/dosen"
)

type SqlServerDosenRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewSqlServerDosenRepository(db *sql.DB, ctx context.Context) dosen.DosenRepository {
	return &SqlServerDosenRepository{db, ctx}
}

func (s *SqlServerDosenRepository) FindById(id dosen.DosenId) (*dosen.Dosen, error) {

	tsql := `SELECT d.id_dosen, d.nama
			 FROM dosen d
			 WHERE d.id_dosen = @DosenId AND d.deleted_at IS NULL`

	rows, err := s.db.QueryContext(s.ctx, tsql, sql.Named("DosenId", id))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id dosen.DosenId
		var nama string
		err = rows.Scan(&id, &nama)
		if err != nil {
			return nil, err
		}

		return dosen.NewDosen(id, nama), nil
	}

	return nil, nil

}
