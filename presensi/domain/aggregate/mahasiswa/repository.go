package mahasiswa

type MahasiswaRepository interface {
	ById(id MahasiswaId) (*Mahasiswa, error)
}
