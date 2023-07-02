package command

import (
	"errors"

	"github.com/google/uuid"
	"its.id/akademik/frs/domain/frs"
	"its.id/akademik/frs/domain/kelas"
)

type PilihKelasByMahasiswaRequest struct {
	KelasId     uuid.UUID
	FrsId       uuid.UUID
	MahasiswaId uuid.UUID
}

type PilihKelasByMahasiswaCommand struct {
	kelasRepository kelas.KelasRepository
	frsRepository   frs.FrsRepository
}

func NewPilihKelasByMahasiswaCommand(kelasRepo kelas.KelasRepository, frsRepo frs.FrsRepository) (*PilihKelasByMahasiswaCommand, error) {

	if kelasRepo == nil {
		return nil, errors.New("kelas repository tidak bisa nil")
	}

	if frsRepo == nil {
		return nil, errors.New("frs repository tidak bisa nil")
	}

	return nil, nil
}

func (c *PilihKelasByMahasiswaCommand) Execute(r *PilihKelasByMahasiswaRequest) error {

	kelasId := kelas.KelasId(r.KelasId)

	kelas, err := c.kelasRepository.FindById(kelasId)

	if err != nil {
		return err
	}

	if kelas == nil {
		return errors.New("kelas tidak ditemukan")
	}

	frsId := frs.FrsId(r.FrsId)

	frs, err := c.frsRepository.FindById(frsId)

	if err != nil {
		return err
	}

	if frs == nil {
		return errors.New("frs tidak ditemukan")
	}

	err = frs.TambahKuliahByMahasiswa(kelas)

	c.frsRepository.Update(frs)

	return nil
}
