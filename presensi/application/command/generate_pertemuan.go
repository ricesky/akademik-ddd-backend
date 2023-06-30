package command

import (
	"errors"

	"github.com/google/uuid"
	"its.id/akademik/presensi/domain"
	"its.id/akademik/presensi/domain/kelas"
)

type GeneratePertemuanRequest struct {
	kelasId uuid.UUID
}

type GeneratePertemuanCommand struct {
	kelasRepository           kelas.KelasRepository
	pertemuanGeneratorService *domain.PertemuanManagerService
}

func NewGeneratePertemuanCommand(
	kelasRepo kelas.KelasRepository,
	pertemuanManagerService *domain.PertemuanManagerService,
) (*GeneratePertemuanCommand, error) {

	if kelasRepo == nil {
		return nil, errors.New("KelasRepository tidak boleh nil")
	}

	if pertemuanManagerService == nil {
		return nil, errors.New("PertemuanManagerService tidak boleh nil")
	}

	return &GeneratePertemuanCommand{
		kelasRepo,
		pertemuanManagerService,
	}, nil

}

func (c *GeneratePertemuanCommand) Execute(r GeneratePertemuanRequest) error {

	// TODO: command request validation
	// TODO: check dosen yang menambah pertemuan harus dosen pengajar kelas tersebut

	kelas, err := c.kelasRepository.FindById(r.kelasId)

	if err != nil {
		return err
	}

	if kelas == nil {
		return errors.New("kelas tidak ditemukan")
	}

	return c.pertemuanGeneratorService.GenerateSemuaPertemuan(kelas)
}
