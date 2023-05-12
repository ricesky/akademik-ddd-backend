package repository

import (
	"github.com/google/uuid"
	"its.id/akademik/presensi/domain/entity"
)

type DosenRepository interface {
	ById(id uuid.UUID) (entity.Dosen, error)
}
