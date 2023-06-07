package kelas

type KelasRepository interface {
	GetById(kelasId KelasId) (*Kelas, error)
}
