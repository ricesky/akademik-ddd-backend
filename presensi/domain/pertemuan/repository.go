package pertemuan

import "its.id/akademik/presensi/domain/kelas"

type PertemuanRepository interface {
	FindById(id PertemuanId) (*Pertemuan, error)
	FindByKelasId(kelasId kelas.KelasId) ([]Pertemuan, error)
	Save(pertemuan *Pertemuan) error
	Update(pertemuan *Pertemuan) error
}
