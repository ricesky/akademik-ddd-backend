package command

import (
	"errors"

	"github.com/google/uuid"
	"its.id/akademik/presensi/domain/dosen"
	"its.id/akademik/presensi/domain/pertemuan"
)

type MulaiPertemuanRequest struct {
	PertemuanId              uuid.UUID
	DosenId                  uuid.UUID
	ModePertemuan            string
	BentukKehadiran          string
	MenitBerlakuKodePresensi int
}

type MulaiPertemuanCommand struct {
	pertemuanRepository pertemuan.PertemuanRepository
	dosenRepository     dosen.DosenRepository
}

func NewMulaiPertemuanCommand(
	pertemuanRepo pertemuan.PertemuanRepository,
	dosenRepo dosen.DosenRepository,
) (*MulaiPertemuanCommand, error) {

	if pertemuanRepo == nil {
		return nil, errors.New("pertemuan repository tidak boleh nil")
	}

	if dosenRepo == nil {
		return nil, errors.New("dosen repository tidak boleh nil")
	}

	return &MulaiPertemuanCommand{
		pertemuanRepo,
		dosenRepo,
	}, nil

}

func (c *MulaiPertemuanCommand) Execute(r MulaiPertemuanRequest) error {

	pertemuanId := pertemuan.PertemuanId(r.PertemuanId)

	p, err := c.pertemuanRepository.FindById(pertemuanId)

	if err != nil {
		return err
	}

	mode, err := pertemuan.NewModePertemuan(r.ModePertemuan)

	if err != nil {
		return err
	}

	bentukHadir, err := pertemuan.NewBentukKehadiran(r.BentukKehadiran)

	if err != nil {
		return err
	}

	err = p.Mulai(mode, bentukHadir, r.MenitBerlakuKodePresensi)

	if err != nil {
		return err
	}

	return c.pertemuanRepository.Update(p)
}
