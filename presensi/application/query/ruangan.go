package query

type Ruangan struct {
	RuanganId string
	Kode      string
}

type DaftarRuanganQueryHandler interface {
	Execute() []Ruangan
}
