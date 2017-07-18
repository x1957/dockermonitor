package falcon

import (
	"encoding/json"
	"log"
	"testing"
)

func Test_ToJson(t *testing.T) {
	tmp := Falcon{
		URL:     "123",
		Metrics: "456",
	}
	bs, _ := json.Marshal(tmp)
	log.Printf(string(bs))
}
