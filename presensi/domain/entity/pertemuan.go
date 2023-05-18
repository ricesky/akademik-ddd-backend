package entity

import (
	"errors"
	"fmt"
	"time"

	vo "its.id/akademik/presensi/domain/value_object"
)

const (
	jedaBolehPresensi              int = 30
	minimalMasaBerlakuKodePresensi int = 15
)

type Pertemuan struct {
	id           vo.PertemuanId
	kelas        *Kelas
	urutan       vo.UrutanPertemuan
	ruanganId    vo.RuanganId
	jadwal       vo.JadwalPertemuan
	topik        vo.TopikPerkuliahan
	mode         vo.ModePertemuan
	status       vo.StatusPertemuan
	kodePresensi vo.KodePresensi
}

func NewPertemuan(
	id vo.PertemuanId,
	kelas *Kelas,
	urutan vo.UrutanPertemuan,
	ruanganId vo.RuanganId,
	jadwal vo.JadwalPertemuan,
	topik vo.TopikPerkuliahan,
	mode vo.ModePertemuan,
	status vo.StatusPertemuan,
	kodePresensi vo.KodePresensi,
) (*Pertemuan, error) {

	return &Pertemuan{
		id:           id,
		kelas:        kelas,
		urutan:       urutan,
		ruanganId:    ruanganId,
		jadwal:       jadwal,
		topik:        topik,
		mode:         mode,
		status:       status,
		kodePresensi: kodePresensi,
	}, nil
}

func (p *Pertemuan) ID() vo.PertemuanId {
	return p.id
}

func (p *Pertemuan) Mulai(
	mode vo.ModePertemuan,
	bentukKehadiran vo.BentukKehadiran,
	menitBerlaku int,
) error {

	if p.kelas.IsSelesai() {
		return errors.New("tidak_dapat_memulai_pertemuan_karena_kelas_sudah_selesai")
	}

	if !p.isBolehMulaiPertemuan() {
		return errors.New("pertemuan_belum_boleh_dimulai")
	}

	if p.isTerlewat() {
		return errors.New("pertemuan_sudah_terlewati")
	}

	if p.mode.IsOnline() && !bentukKehadiran.IsHadirOnline() {
		return errors.New("kehadiran_dosen_harus_online_untuk_mode_pertemuan_online")
	}

	if p.mode.IsOffline() && !bentukKehadiran.IsHadirOffline() {
		return errors.New("kehadiran_dosen_harus_offline_untuk_mode_pertemuan_offline")
	}

	if menitBerlaku > 0 && menitBerlaku < minimalMasaBerlakuKodePresensi {
		return fmt.Errorf("menit_berlaku_kode_presensi_tidak_boleh_kurang_dari_%d_menit", minimalMasaBerlakuKodePresensi)
	}

	var masaBerlakuKodePresensi time.Time

	if menitBerlaku <= 0 {
		masaBerlakuKodePresensi = p.jadwal.WaktuSelesai()
	} else {
		durasiBerlaku := time.Duration(menitBerlaku) * time.Minute
		masaBerlakuKodePresensi = p.jadwal.WaktuMulai().Add(durasiBerlaku)
	}

	kodePresensiBaru, err := vo.GenerateRandomKodePresensi(masaBerlakuKodePresensi)

	if err != nil {
		return err
	}

	p.kodePresensi = kodePresensiBaru
	p.status = vo.NewStatusPertemuanSedangBerlangsung()

	return nil
}

func (p *Pertemuan) Akhiri() {
	// TODO: implement function akhiri
}

func (p *Pertemuan) Lupa(mode vo.ModePertemuan) {
	// TODO: implement function lupa
}

func (p *Pertemuan) GantiKodePresensi() {

}

func (p *Pertemuan) isBolehMulaiPertemuan() bool {
	jeda := time.Duration(jedaBolehPresensi) * time.Minute

	waktuBolehMulai := p.jadwal.WaktuMulai().Add(-jeda)

	return waktuBolehMulai.After(time.Now())
}

func (p *Pertemuan) isTerlewat() bool {
	return p.status.IsBelumDimulai() && p.jadwal.WaktuSelesai().Before(time.Now())
}
