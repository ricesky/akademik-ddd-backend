package dosen

type DosenRepository interface {
	FindById(id DosenId) (*Dosen, error)
}
