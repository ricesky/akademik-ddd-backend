package kelas

type KelasRepository interface {
	FindById(id KelasId) (*Kelas, error)
}
