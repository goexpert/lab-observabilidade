package labobservalidadade

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

type LocaleWeatherDto struct {
	Locale string  `json:"city"`
	TempC  float64 `json:"temp_c"`
	TempF  float64 `json:"temp_f"`
	TempK  float64 `json:"temp_k"`
}
