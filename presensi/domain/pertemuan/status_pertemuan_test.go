package pertemuan_test

import (
	"testing"

	"its.id/akademik/presensi/domain/pertemuan"
)

func Test_kode_status_pertemuan_sesuai(t *testing.T) {

	statusBelumDimulai := pertemuan.NewStatusPertemuanBelumDimulai()
	statusSedangBerlangsung := pertemuan.NewStatusPertemuanSedangBerlangsung()
	statusSelesai := pertemuan.NewStatusPertemuanSelesai()
	statusTerlewat := pertemuan.NewStatusPertemuanTerlewat()

	if !statusBelumDimulai.IsBelumDimulai() {
		t.Fatal("status belum dimulai tidak sesuai")
	}

	if !statusSedangBerlangsung.IsSedangBerlangsung() {
		t.Fatal("status sedang berlangsung tidak sesuai")
	}

	if !statusSelesai.IsSelesai() {
		t.Fatal("status selesai tidak sesuai")
	}

	if !statusTerlewat.IsTerlewat() {
		t.Fatal("status terlewat tidak sesuai")
	}

}
