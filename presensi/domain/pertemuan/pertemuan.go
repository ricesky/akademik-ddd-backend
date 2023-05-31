package pertemuan

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"its.id/akademik/presensi/domain/kelas"
)

const (
	jedaBolehPresensi              int = 30
	minimalMasaBerlakuKodePresensi int = 15
)

type PertemuanId uuid.UUID
type RuanganId uuid.UUID

type Pertemuan struct {
	id           PertemuanId
	kelas        *kelas.Kelas
	urutan       UrutanPertemuan
	ruanganId    RuanganId
	jadwal       JadwalPertemuan
	topik        TopikPerkuliahan
	mode         ModePertemuan
	status       StatusPertemuan
	kodePresensi KodePresensi
}

func NewPertemuan(
	id PertemuanId,
	kelas *kelas.Kelas,
	urutan UrutanPertemuan,
	ruanganId RuanganId,
	jadwal JadwalPertemuan,
	topik TopikPerkuliahan,
	mode ModePertemuan,
	status StatusPertemuan,
	kodePresensi KodePresensi,
) (*Pertemuan, error) {

	if kelas.RencanaPertemuan() <= 0 {
		errors.New("tidak_dapat_buat_pertemuan_baru_karena_belum_ada_rencana_pertemuan")
	}

	if kelas.IsSelesai() {
		errors.New("tidak_dapat_membuat_pertemuan_baru_karena_kelas_sudah_selesai")
	}

	if mode.IsOffline() && ruanganId == RuanganId(uuid.Nil) {
		errors.New("mode_tatap_muka_offline_harus_memiliki_ruangan")
	}

	if mode.IsHybrid() && ruanganId == RuanganId(uuid.Nil) {
		errors.New("mode_tatap_muka_hybrid_harus_memiliki_ruangan")
	}

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

func (p *Pertemuan) ID() PertemuanId {
	return p.id
}

func (p *Pertemuan) Ubah() {

}

func (p *Pertemuan) Mulai(
	mode ModePertemuan,
	bentukKehadiran BentukKehadiran,
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

	kodePresensiBaru, err := GenerateRandomKodePresensi(masaBerlakuKodePresensi)

	if err != nil {
		return err
	}

	p.kodePresensi = kodePresensiBaru
	p.status = NewStatusPertemuanSedangBerlangsung()

	return nil
}

func (p *Pertemuan) Akhiri() error {

	if p.kelas.IsSelesai() {
		return errors.New("tidak_dapat_mengakhiri_pertemuan_karena_kelas_sudah_selesai")
	}

	if p.status.IsBelumDimulai() {
		return errors.New("tidak_dapat_mengakhiri_pertemuan_yang_belum_dimulai")
	}

	if p.status.IsTerlewat() {
		return errors.New("tidak_dapat_mengakhiri_pertemuan_yang_terlewat")
	}

	p.status = NewStatusPertemuanSelesai()

	return nil
}

func (p *Pertemuan) Lupa(mode ModePertemuan) error {

	if p.kelas.IsSelesai() {
		return errors.New("tidak_dapat_mengakhiri_pertemuan_karena_kelas_sudah_selesai")
	}

	if p.status.IsSedangBerlangsung() || p.status.IsSelesai() {
		return errors.New("tidak_dapat_menandai_lupa_presensi_pada_pertemuan_yang_sedang_berlangsung_atau_selesai")
	}

	p.status = NewStatusPertemuanSelesai()
	p.mode = mode

	return nil
}

func (p *Pertemuan) isBolehMulaiPertemuan() bool {
	jeda := time.Duration(jedaBolehPresensi) * time.Minute

	waktuBolehMulai := p.jadwal.WaktuMulai().Add(-jeda)

	return waktuBolehMulai.After(time.Now())
}

func (p *Pertemuan) isTerlewat() bool {
	return p.status.IsBelumDimulai() && p.jadwal.WaktuSelesai().Before(time.Now())
}
