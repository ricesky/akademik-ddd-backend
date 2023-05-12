package value_object_test

import (
	"testing"
	"time"

	vo "its.id/akademik/presensi/domain/value_object"
)

func Test_JadwalPertemuan(t *testing.T) {

	wm := time.Now()
	ws := time.Now()

	j, _ := vo.NewJadwalPertemuan(wm, ws)

	if j.WaktuMulai() != j.WaktuSelesai() {
		t.Fatal("a and  b were not equal")
	}
}
