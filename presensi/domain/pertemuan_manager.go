package domain

import (
	"errors"

	"its.id/akademik/presensi/domain/kelas"
	"its.id/akademik/presensi/domain/pertemuan"
)

type PertemuanManagerService struct {
	pertemuanRepository pertemuan.PertemuanRepository
}

func NewPertemuanManagerService(pertemuanRepo pertemuan.PertemuanRepository) *PertemuanManagerService {
	return &PertemuanManagerService{pertemuanRepo}
}

func (s *PertemuanManagerService) TambahPertemuan(pertemuan *pertemuan.Pertemuan) error {

	if pertemuan.Kelas().IsSelesai() {
		return errors.New("tidak dapat menambahkan pertemuan karena kelas telah selesai")
	}

	listPertemuanKelas, err := s.pertemuanRepository.FindByKelasId(pertemuan.Kelas().Id())

	if err != nil {
		return err
	}

	for _, pk := range listPertemuanKelas {

		if pertemuan.IsUrutanPertemuanSama(&pk) {
			return errors.New("urutan pertemuan sama dengan pertemuan lainnya")
		}
	}

	err = s.pertemuanRepository.Save(pertemuan)

	if err != nil {
		return err
	}

	return nil
}

func (s *PertemuanManagerService) GenerateSemuaPertemuan(kelas *kelas.Kelas) error {

	if kelas.IsSelesai() {
		return errors.New("tidak dapat menambahkan pertemuan karena kelas telah selesai")
	}

	if kelas.RencanaPertemuan() == 0 {
		return errors.New("tidak dapat mengenerate pertemuan karena jumlah rencana pertemuan belum diisi")
	}

	listPertemuanKelas, err := s.pertemuanRepository.FindByKelasId(kelas.Id())

	if err != nil {
		return err
	}

	for i := 1; i <= kelas.RencanaPertemuan(); i++ {
		for _, pk := range listPertemuanKelas {

			urutan, err := pertemuan.NewUrutanPertemuan(i)

			if err != nil {
				return err
			}

			if urutan.EqualTo(pk.Urutan()) {
				break
			}

			// TODO: buat pertemuan sesuai jadwal kelas

		}
	}

	return nil
}
