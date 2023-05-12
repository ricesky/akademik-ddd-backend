package repository

import (
	"github.com/google/uuid"
	"its.id/akademik/presensi/domain/entity"
)

type MahasiswaRepository interface {
	ById(id uuid.UUID) (entity.Mahasiswa, error)
}
