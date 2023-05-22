package dosen

import (
	"github.com/google/uuid"
)

type DosenRepository interface {
	ById(id uuid.UUID) (*Dosen, error)
}
