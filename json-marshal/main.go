package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First      string
	Last       string `json:"-"`
	Age        int    `json:"wisdom-score"`
	unexported int
}

func main() {
	//p1 := Person{
	//	First:      "Michael",
	//	Last:       "Jackson",
	//	Age:        900,
	//	unexported: 1,
	//}
	//var p2 Person
	//fmt.Println(p1.First)
	//fmt.Println(p1.unexported)
	//bs, _ := json.Marshal(p1)
	//fmt.Println("%T\n", bs)
	//fmt.Println(string(bs))
	////Unmarshalling
	//_ = []byte(`{"First": "Patrick", "Last": "Logan", "wisdom-score": 2}`)
	//json.Unmarshal(bs, &p2)
	//fmt.Println(p2.First)
	//fmt.Println(p2.Last)
	//fmt.Println(p2.Age)

	// Encoder
	//pEncode := Person{"James", "Anderson", 30, 1}
	//json.NewEncoder(os.Stdout).Encode(pEncode)

	//Decoder
	var p4 Person
	rdr := strings.NewReader(`{"First": "Hanuman", "Last": "Vayu", "Age": 3}`)
	json.NewDecoder(rdr).Decode(&p4)
	fmt.Println(p4.First)
	fmt.Println(p4)

}
