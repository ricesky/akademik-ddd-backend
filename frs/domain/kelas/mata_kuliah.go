package kelas

type MataKuliah struct {
	kode string
	nama string
	sks  int
}

func NewMataKuliah(kode string, nama string, sks int) MataKuliah {
	return MataKuliah{
		kode: kode,
		nama: nama,
		sks:  sks,
	}
}

func (m *MataKuliah) Kode() string {
	return m.kode
}

func (m *MataKuliah) Nama() string {
	return m.nama
}

func (m *MataKuliah) Sks() int {
	return m.sks
}

func (m *MataKuliah) EqualTo(other MataKuliah) bool {
	return m.kode == other.kode
}
