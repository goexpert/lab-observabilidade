package pkg

type CepDto struct {
	Cep string
}

type DtoError struct {
	Message string `json:"msg"`
}

type LogradouroDto struct {
	Cep        string `json:"cep"`
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}
