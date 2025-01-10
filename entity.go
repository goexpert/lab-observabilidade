package labobservalidadade

import (
	"errors"
	"log/slog"
	"regexp"
)

type CepEntity struct {
	cep string
}

func NewCep(cep string) (*CepDto, error) {

	var _cep = &CepEntity{
		cep: cep,
	}

	err := _cep.IsValid()
	if err != nil {
		slog.Error("cep inválido", "error", err.Error())
		return nil, err
	}

	return &CepDto{
		Cep: _cep.cep,
	}, nil
}

func (c *CepEntity) IsValid() error {

	var re = regexp.MustCompile(`^[0-9]{8}$`)

	if !re.MatchString(c.cep) {
		return errors.New("cep inválido")
	}
	return nil
}
