package pertemuan_test

import (
	"testing"

	vo "its.id/akademik/presensi/domain/value_object"
)

func Test_kode_mode_pertemuan_sesuai(t *testing.T) {

	online := vo.NewModePertemuanOnline()
	offline := vo.NewModePertemuanOffline()
	hybrid := vo.NewModePertemuanHybrid()

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
