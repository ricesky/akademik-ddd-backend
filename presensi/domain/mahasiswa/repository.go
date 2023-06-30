package mahasiswa

type MahasiswaRepository interface {
	FindById(id MahasiswaId) (*Mahasiswa, error)
}
