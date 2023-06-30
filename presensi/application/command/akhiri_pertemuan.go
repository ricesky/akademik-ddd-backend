package command

import (
	"errors"

	"github.com/google/uuid"
	"its.id/akademik/presensi/domain/pertemuan"
)

type AkhiriPertemuanRequest struct {
	PertemuanId uuid.UUID
}

type AkhiriPertemuanCommand struct {
	pertemuanRepository pertemuan.PertemuanRepository
}

func NewAkhiriPertemuanCommand(pertemuanRepo pertemuan.PertemuanRepository) (*AkhiriPertemuanCommand, error) {

	if pertemuanRepo == nil {
		return nil, errors.New("pertemuan repository tidak boleh nil")
	}

	return &AkhiriPertemuanCommand{pertemuanRepo}, nil
}

func (c *AkhiriPertemuanCommand) Execute(r AkhiriPertemuanRequest) error {

	pertemuanId := pertemuan.PertemuanId(r.PertemuanId)

	p, err := c.pertemuanRepository.FindById(pertemuanId)

	if err != nil {
		return err
	}

	err = p.Akhiri()

	if err != nil {
		return err
	}

	return c.pertemuanRepository.Update(p)
}
