package json_

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/yezihack/colorlog"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var pln = fmt.Println

func TestJson(t *testing.T) {
	bolB, _ := json.Marshal(true)
	pln(string(bolB))

	intB, _ := json.Marshal(1)
	pln(string(intB))

	fltB, _ := json.Marshal(2.34)
	pln(string(fltB))

	strB, _ := json.Marshal("gopher")
	pln(string(strB))

	slcD := []string{"apple", "pear", "orange"}
	slcB, _ := json.Marshal(slcD)
	pln(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	pln(string(mapB))

	res1D := &response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	pln(string(res1B))

	res2D := &response2{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	pln(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	pln(dat)
	num := dat["num"].(float64)
	pln(num)
	strs := dat["strs"].([]interface{})
	pln(strs)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res1, res2 := response1{}, response2{}
	json.Unmarshal([]byte(str), &res1)
	json.Unmarshal([]byte(str), &res2)
	pln(res1)
	pln(res2)
	pln(res1.Fruits[0])
	pln(res2.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

func get(key string) interface{} {
	return nil
}

func TestTest(t *testing.T) {
	a := get("a").(int)
	colorlog.Debug("a: %d", a)
}
