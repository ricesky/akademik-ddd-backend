package command

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"its.id/akademik/presensi/domain"
	"its.id/akademik/presensi/domain/kelas"
	"its.id/akademik/presensi/domain/pertemuan"
)

type TambahPertemuanRequest struct {
	KelasId      uuid.UUID
	PertemuanKe  int
	RuanganId    uuid.NullUUID
	WaktuMulai   time.Time
	WaktuSelesai time.Time
	Topik        string
	TopikEn      string
	Mode         string
}

type TambahPertemuanCommand struct {
	kelasRepository         kelas.KelasRepository
	pertemuanManagerService *domain.PertemuanManagerService
}

func NewTambahPertemuanCommand(
	kelasRepo kelas.KelasRepository,
	pertemuanGeneratorService *domain.PertemuanManagerService,
) (*TambahPertemuanCommand, error) {

	if kelasRepo == nil {
		return nil, errors.New("KelasRepository tidak boleh nil")
	}

	if pertemuanGeneratorService == nil {
		return nil, errors.New("PertemuanManagerService tidak boleh nil")
	}

	return &TambahPertemuanCommand{
		kelasRepo,
		pertemuanGeneratorService,
	}, nil

}

func (c *TambahPertemuanCommand) Execute(r TambahPertemuanRequest) error {

	// TODO: command request validation
	// TODO: check dosen yang menambah pertemuan harus dosen pengajar kelas tersebut

	kelas, err := c.kelasRepository.FindById(r.KelasId)

	if err != nil {
		return err
	}

	if kelas == nil {
		return errors.New("kelas tidak ditemukan")
	}

	var ruanganId pertemuan.RuanganId
	if r.RuanganId.Valid {
		ruanganId = pertemuan.RuanganId(r.RuanganId.UUID)
	} else {
		ruanganId = pertemuan.RuanganId(uuid.Nil)
	}

	jadwal, err := pertemuan.NewJadwalPertemuan(r.WaktuMulai, r.WaktuSelesai)

	if err != nil {
		return err
	}

	urutan, err := pertemuan.NewUrutanPertemuan(r.PertemuanKe)

	if err != nil {
		return err
	}

	topik, err := pertemuan.NewTopikPerkuliahan(r.Topik, r.TopikEn)

	if err != nil {
		return err
	}

	mode, err := pertemuan.NewModePertemuan(r.Mode)

	if err != nil {
		return err
	}

	pertemuan, err := pertemuan.NewPertemuan(
		kelas,
		urutan,
		ruanganId,
		jadwal,
		topik,
		mode,
	)

	if err != nil {
		return err
	}

	return c.pertemuanManagerService.TambahPertemuan(pertemuan)
}
