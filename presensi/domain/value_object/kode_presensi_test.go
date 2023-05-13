package value_object

import (
	"testing"
	"time"
)

func Test_panjang_kode_presensi_tidak_sesuai(t *testing.T) {

	_, err := NewKodePresensi("1", time.Time{})

	if err == nil {
		t.Fatal("seharusnya muncul error")
	}

	t.Log(err.Error())
}

func Test_waktu_berlaku_kode_presensi_tidak_sesuai(t *testing.T) {

	_, err := NewKodePresensi("123456", time.Time{})

	if err == nil {
		t.Fatal("seharusnya muncul error")
	}

	t.Log(err.Error())
}

func Test_kode_presensi_sesuai(t *testing.T) {

	kode := "123456"
	time := time.Date(2023, 5, 13, 1, 0, 0, 0, time.Local)

	kp, err := NewKodePresensi(kode, time)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("seharusnya tidak muncul error")
	}

	if kp.Kode() != kode {
		t.Fatalf("kode presensi seharusnya %s tetapi %s.", kp.Kode(), kode)
	}

	if kp.BerlakuSampai() != time {
		t.Fatalf("waktu berlaku kode presensi seharusnya %s tetapi %s.", kp.BerlakuSampai(), time)
	}

}
