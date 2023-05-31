package pertemuan_test

import (
	"testing"

	"its.id/akademik/presensi/domain/pertemuan"
)

func Test_urutan_pertemuan_sesuai(t *testing.T) {

	u, err := pertemuan.NewUrutanPertemuan(1)

	if u.Urutan() != 1 {
		t.Fatal("urutan tidak sesuai")
	}

	if err != nil {
		t.Fatal("error seharusnya nil")
	}

}

func Test_urutan_pertemuan_tidak_boleh_sama_dengan_nol(t *testing.T) {

	_, err := pertemuan.NewUrutanPertemuan(0)

	if err == nil {
		t.Fatal("seharusnya muncul error")
	}

}

func Test_urutan_pertemuan_tidak_boleh_kurang_dari_nol(t *testing.T) {

	_, err := pertemuan.NewUrutanPertemuan(-1)

	if err == nil {
		t.Fatal("seharusnya muncul error")
	}

}

func Test_urutan_pertemuan_equal(t *testing.T) {

	u1, _ := pertemuan.NewUrutanPertemuan(1)
	u2, _ := pertemuan.NewUrutanPertemuan(1)

	if u1 != u2 {
		t.Fatal("seharusnya equal")
	}

}
