package falcon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_ToJson(t *testing.T) {
	tmp := Metrics{
		Metric:      "docker-test",
		Endpoint:    "1957",
		Timestamp:   12345,
		Value:       1957,
		CounterType: "GAUGE",
		Tags:        "kubernetes",
	}
	bs, err := json.Marshal(tmp)
	if err != nil {
		log.Printf("error: %v", err)
		t.Fail()
	}
	log.Printf(string(bs))
	if string(bs) != `{"metric":"docker-test","endpoint":"1957","timestamp":12345,"step":0,"value":1957,"counterType":"GAUGE","tags":"kubernetes"}` {
		t.Fail()
	}
	var mlist []Metrics
	mlist = append(mlist, tmp)
	bs, err = json.Marshal(mlist)
	if err != nil {
		log.Printf("error: %v", err)
		t.Fail()
	}
	log.Printf(string(bs))
	if string(bs) != `[{"metric":"docker-test","endpoint":"1957","timestamp":12345,"step":0,"value":1957,"counterType":"GAUGE","tags":"kubernetes"}]` {
		t.Fail()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", string(body))
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Test_HTTPServer(t *testing.T) {
	http.HandleFunc("/v1/push", handler)
	http.ListenAndServe(":1988", nil)
}
