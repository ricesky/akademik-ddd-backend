package frs

type FrsRepository interface {
	FindById(id FrsId) (*Frs, error)
	Update(frs *Frs) error
}
