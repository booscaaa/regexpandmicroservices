package texto

//Texto .
type Texto struct {
	Texto string `json:"texto"`
}

//Retorno .
type Retorno struct {
	Texto   string `json:"texto"`
	IsValid bool   `json:"isValid"`
}

//RetornoArray .
type RetornoArray struct {
	Texto   string   `json:"texto"`
	Valores []string `json:"valores"`
}
