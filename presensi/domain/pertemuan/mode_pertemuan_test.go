package pertemuan_test

import (
	"testing"

	"its.id/akademik/presensi/domain/aggregate/pertemuan"
)

func Test_kode_mode_pertemuan_sesuai(t *testing.T) {

	online := pertemuan.NewModePertemuanOnline()
	offline := pertemuan.NewModePertemuanOffline()
	hybrid := pertemuan.NewModePertemuanHybrid()

	if !online.IsOnline() {
		t.Fatal("status online tidak sesuai")
	}

	if !offline.IsOffline() {
		t.Fatal("status offline tidak sesuai")
	}

	if !hybrid.IsHybrid() {
		t.Fatal("status hybrid tidak sesuai")
	}

}
