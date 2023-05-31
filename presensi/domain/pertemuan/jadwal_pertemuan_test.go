package pertemuan_test

import (
	"fmt"
	"testing"
	"time"

	"its.id/akademik/presensi/domain/pertemuan"
)

func Test_waktu_mulai_selesai_tidak_boleh_sama(t *testing.T) {

	tt := time.Date(2023, 05, 13, 0, 0, 0, 0, time.Local)

	j, err := pertemuan.NewJadwalPertemuan(tt, tt)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err == nil {
		t.Fatal("seharusnya ada pesan error")
	}

	if !j.WaktuMulai().IsZero() {
		t.Fatal("waktu mulai seharusnya zero")
	}

	if !j.WaktuSelesai().IsZero() {
		t.Fatal("waktu selesai seharusnya zero")
	}

}

func Test_waktu_mulai_tidak_boleh_lebih_dari_waktu_selesai(t *testing.T) {

	waktuMulai := time.Date(2023, 05, 13, 1, 0, 0, 0, time.Local)
	waktuSelesai := time.Date(2023, 05, 13, 0, 0, 0, 0, time.Local)

	j, err := pertemuan.NewJadwalPertemuan(waktuMulai, waktuSelesai)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err == nil {
		t.Fatal("seharusnya ada pesan error")
	}

	if !j.WaktuMulai().IsZero() {
		t.Fatal("waktu mulai seharusnya zero")
	}

	if !j.WaktuSelesai().IsZero() {
		t.Fatal("waktu selesai seharusnya zero")
	}

}

func Test_argumen_waktu_mulai_tidak_boleh_kosong(t *testing.T) {

	tt := time.Time{}

	j, err := pertemuan.NewJadwalPertemuan(tt, tt)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err == nil {
		t.Fatal("seharusnya ada pesan error")
	}

	if !j.WaktuMulai().IsZero() {
		t.Fatal("waktu mulai seharusnya zero")
	}

	if !j.WaktuSelesai().IsZero() {
		t.Fatal("waktu selesai seharusnya zero")
	}
}

func Test_argumen_waktu_selesai_tidak_boleh_kosong(t *testing.T) {

	waktuMulai := time.Date(2023, 05, 13, 1, 0, 0, 0, time.Local)
	waktuSelesai := time.Time{}

	j, err := pertemuan.NewJadwalPertemuan(waktuMulai, waktuSelesai)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err == nil {
		t.Fatal("seharusnya ada pesan error")
	}

	if !j.WaktuMulai().IsZero() {
		t.Fatal("waktu mulai seharusnya zero")
	}

	if !j.WaktuSelesai().IsZero() {
		t.Fatal("waktu selesai seharusnya zero")
	}
}

func Test_equality(t *testing.T) {

	waktuMulai := time.Date(2023, 05, 13, 1, 0, 0, 0, time.Local)
	waktuSelesai := time.Date(2023, 05, 13, 2, 0, 0, 0, time.Local)

	j1, _ := pertemuan.NewJadwalPertemuan(waktuMulai, waktuSelesai)
	j2, _ := pertemuan.NewJadwalPertemuan(waktuMulai, waktuSelesai)

	if j1 != j2 {
		t.Fatal("hasil perbandingan tidak sama")
	}

}
