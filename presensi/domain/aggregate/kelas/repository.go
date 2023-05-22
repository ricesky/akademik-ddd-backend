package kelas

import "github.com/google/uuid"

type KelasRepository interface {
	ById(id uuid.UUID) (*Kelas, error)
}
