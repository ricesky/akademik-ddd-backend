package query

type Ruangan struct {
	RuanganId string
	Kode      string
}

type RuanganQueryHandler interface {
	GetAll() []Ruangan
}
