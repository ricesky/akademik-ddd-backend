package command

import (
	"errors"

	"github.com/google/uuid"
	"its.id/akademik/presensi/domain/pertemuan"
)

type LupaPresensiRequest struct {
	PertemuanId   uuid.UUID
	ModePertemuan string
}

type LupaPresensiCommand struct {
	pertemuanRepository pertemuan.PertemuanRepository
}

func NewLupaPresensiCommand(pertemuanRepo pertemuan.PertemuanRepository) (*LupaPresensiCommand, error) {

	if pertemuanRepo == nil {
		return nil, errors.New("pertemuan repository tidak boleh nil")
	}

	return &LupaPresensiCommand{pertemuanRepo}, nil
}

func (c *LupaPresensiCommand) Execute(r LupaPresensiRequest) error {

	pertemuanId := pertemuan.PertemuanId(r.PertemuanId)

	p, err := c.pertemuanRepository.FindById(pertemuanId)

	if err != nil {
		return err
	}

	mode, err := pertemuan.NewModePertemuan(r.ModePertemuan)

	if err != nil {
		return err
	}

	err = p.Lupa(mode)

	if err != nil {
		return err
	}

	return c.pertemuanRepository.Update(p)
}
