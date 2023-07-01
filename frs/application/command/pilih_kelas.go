package command

type PilihKelasRequest struct {
}

type PilihKelasCommand struct {
}

func NewPilihKelasCommand() (*PilihKelasCommand, error) {
	return nil, nil
}

func (c *PilihKelasCommand) Execute(r *PilihKelasRequest) error {
	return nil
}
