package pertemuan_test

import (
	"testing"

	vo "its.id/akademik/presensi/domain/value_object"
)

func Test_kode_status_pertemuan_sesuai(t *testing.T) {

	statusBelumDimulai := vo.NewStatusPertemuanBelumDimulai()
	statusSedangBerlangsung := vo.NewStatusPertemuanSedangBerlangsung()
	statusSelesai := vo.NewStatusPertemuanSelesai()
	statusTerlewat := vo.NewStatusPertemuanTerlewat()

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
