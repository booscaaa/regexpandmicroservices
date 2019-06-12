package texto

import (
	"fmt"
	"regexp"
)

//MatchNumber .
func (t Texto) MatchNumber() []byte {
	fmt.Println(t.Texto)
	match := regexp.MustCompile("^[0-9]+$").MatchString(t.Texto)
	fmt.Println(match)

	// r, _ := regexp.Compile("p([a-z]+)ch")
	// fmt.Println(r.MatchString("peach"))
	// fmt.Println(r.FindString("peach punch"))
	// fmt.Println(r.FindStringIndex("peach punch"))
	// fmt.Println(r.FindStringSubmatch("peach punch"))
	// fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// fmt.Println(r.FindAllString("peach punch pinch", -1))
	// fmt.Println(r.FindAllStringSubmatchIndex(
	// 	"peach punch pinch", -1))
	// fmt.Println(r.FindAllString("peach punch pinch", 2))
	// fmt.Println(r.Match([]byte("peach")))
	// r = regexp.MustCompile("p([a-z]+)ch")
	// fmt.Println(r)
	// fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	// in := []byte("a peach")
	// out := r.ReplaceAllFunc(in, bytes.ToUpper)
	// fmt.Println(string(out))
	// jsonResp := simplejson.New()
	// jsonResp.Set("certo", t)

	// payload, _ := jsonResp.MarshalJSON()
	// json.Marshal(t)
	retorno := Retorno{
		t.Texto,
		match,
	}
	return retorno.GetRetorno()
}
