package kelas

type KelasRepository interface {
	FindById(kelasId KelasId) (*Kelas, error)
}
