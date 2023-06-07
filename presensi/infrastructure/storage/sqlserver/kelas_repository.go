package sqlserver

import (
	"context"
	"database/sql"

	"its.id/akademik/presensi/domain/kelas"
)

type SqlServerKelasRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewSqlServerKelasRepository(db *sql.DB, ctx context.Context) kelas.KelasRepository {
	return &SqlServerKelasRepository{db, ctx}
}

func (s *SqlServerKelasRepository) GetById(id kelas.KelasId) (*kelas.Kelas, error) {

	tsql := `SELECT k.id_kelas, k.rencana_tm, k.is_nilai_final
			 FROM kelas k
			 WHERE k.id_kelas = @KelasId AND k.deleted_at IS NULL`

	rows, err := s.db.QueryContext(s.ctx, tsql, sql.Named("KelasId", id))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var (
			id               kelas.KelasId
			rencanaPertemuan int
			selesai          bool
		)

		err = rows.Scan(&id, &rencanaPertemuan, &selesai)
		if err != nil {
			return nil, err
		}

		return kelas.NewKelas(id, rencanaPertemuan, selesai)
	}

	return nil, nil
}
