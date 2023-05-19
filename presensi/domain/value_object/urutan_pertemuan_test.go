package value_object_test

import (
	"testing"

	vo "its.id/akademik/presensi/domain/value_object"
)

func Test_urutan_pertemuan_sesuai(t *testing.T) {

	u, err := vo.NewUrutanPertemuan(1)

	if u.Urutan() != 1 {
		t.Fatal("urutan tidak sesuai")
	}

	if err != nil {
		t.Fatal("error seharusnya nil")
	}

}

func Test_urutan_pertemuan_tidak_boleh_sama_dengan_nol(t *testing.T) {

	_, err := vo.NewUrutanPertemuan(0)

	if err == nil {
		t.Fatal("seharusnya muncul error")
	}

}

func Test_urutan_pertemuan_tidak_boleh_kurang_dari_nol(t *testing.T) {

	_, err := vo.NewUrutanPertemuan(-1)

	if err == nil {
		t.Fatal("seharusnya muncul error")
	}

}

func Test_urutan_pertemuan_equal(t *testing.T) {

	u1, _ := vo.NewUrutanPertemuan(1)
	u2, _ := vo.NewUrutanPertemuan(1)

	if u1 != u2 {
		t.Fatal("seharusnya equal")
	}

}
