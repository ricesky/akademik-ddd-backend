package frs

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"its.id/akademik/frs/domain/dosen"
	"its.id/akademik/frs/domain/kelas"
	"its.id/akademik/frs/domain/mahasiswa"
)

const (
	totalSksAmbilMaks   int     = 24
	totalSksAmbilAwal   int     = 0
	bebanStudiMaks18Sks int     = 18
	bebanStudiMaks20Sks int     = 20
	bebanStudiMaks22Sks int     = 22
	bebanStudiMaks24Sks int     = 24
	ips2_50             float64 = 2.500
	ips3_00             float64 = 3.000
	ips3_50             float64 = 3.500
	ips4_00             float64 = 4.000
)

type FrsId = uuid.UUID

type Frs struct {
	id          FrsId
	semesterId  SemesterId
	mahasiswaId mahasiswa.MahasiswaId
	sksAmbil    int
	sksBatas    int
	sksSisa     int
	kuliahAmbil []*Kuliah
}

func NewFrs(id FrsId, semesterId SemesterId, mahasiswaId mahasiswa.MahasiswaId, ips float64) (*Frs, error) {

	var bebanStudiMaks int = 0

	if ips < ips2_50 {
		bebanStudiMaks = bebanStudiMaks18Sks
	} else if ips >= ips2_50 && ips < ips3_00 {
		bebanStudiMaks = bebanStudiMaks20Sks
	} else if ips >= ips3_00 && ips < ips3_50 {
		bebanStudiMaks = bebanStudiMaks22Sks
	} else if ips >= ips3_50 && ips <= ips4_00 {
		bebanStudiMaks = bebanStudiMaks24Sks
	}

	frsId := uuid.New()

	return &Frs{
		id:         frsId,
		semesterId: semesterId,
		sksAmbil:   totalSksAmbilAwal,
		sksBatas:   bebanStudiMaks,
		sksSisa:    bebanStudiMaks,
	}, nil

}

func (f *Frs) TambahKuliahByDosen(kelas *kelas.Kelas) error {

	mk := kelas.MataKuliah()

	totalSksAmbilBaru := f.sksAmbil + mk.Sks()

	if totalSksAmbilBaru > totalSksAmbilMaks {
		return errors.New("tidak dapat melebihi batas total 24 sks")
	}

	f.kuliahAmbil = append(f.kuliahAmbil, &Kuliah{kelas.Id(), kelas.MataKuliah()})
	f.sksAmbil = totalSksAmbilBaru

	return nil
}

func (f *Frs) TambahKuliahByMahasiswa(kelas *kelas.Kelas) error {

	mk := kelas.MataKuliah()

	totalSksAmbilBaru := f.sksAmbil + mk.Sks()

	if totalSksAmbilBaru > totalSksAmbilMaks {
		return errors.New("tidak dapat melebihi batas total 24 sks")
	}

	if totalSksAmbilBaru > f.sksBatas {
		return fmt.Errorf("tidak dapat melebihi batas total %d sks sesuai indeks prestasi semester", f.sksBatas)
	}

	f.kuliahAmbil = append(f.kuliahAmbil, &Kuliah{kelas.Id(), kelas.MataKuliah()})
	f.sksAmbil = totalSksAmbilBaru

	return nil
}

func (f *Frs) Setuju(dosenId dosen.DosenId) error {
	return nil
}

func (f *Frs) BatalSetuju(dosenId dosen.DosenId) error {
	return nil
}

func (f *Frs) Id() FrsId {
	return f.id
}
