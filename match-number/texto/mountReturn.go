package texto

import "encoding/json"

//GetRetorno .
func (retorno Retorno) GetRetorno() []byte {
	payload, _ := json.Marshal(retorno)
	return payload
}

//GetRetorno .
func (retorno RetornoArray) GetRetorno() []byte {
	payload, _ := json.Marshal(retorno)
	return payload
}
