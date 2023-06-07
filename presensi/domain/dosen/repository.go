package dosen

type DosenRepository interface {
	GetById(id DosenId) (*Dosen, error)
}
