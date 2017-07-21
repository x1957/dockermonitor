package falcon

import (
	"github.com/x1957/dockermonitor/agent"
	"time"
	"os"
	"encoding/json"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
)

type Falcon struct {
	URL string
	Step int32
	Client *http.Client
}

type Metrics struct {
	Metric     string   `json:"metric"`
	Endpoint string   `json:"endpoint"`
	Timestamp   int64 `json:"timestamp"`
	Step int32 `json:"step"`
	Value int64 `json:"value"`
	CounterType string `json:"counterType"`
	Tags string `json:"tags,omitempty"`
}

const AGENT_URL string = "http://127.0.0.1:1988/v1/push"

func NewFalcon() agent.Agent {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    3 * time.Minute,
		DisableCompression: true,
	}
	client := &http.Client{
		Transport: tr,
		Timeout: 1 * time.Minute,
	}
	return &Falcon{
		URL: AGENT_URL,
		Step: 120,
		Client: client,
	}
}

func (f *Falcon) RecordGauge(name string, gauge int64) {
	jsonData, err := f.buildJson(name, gauge)
	if err != nil {
		log.Printf("error in build json %v", err)
		return
	}
	f.push(jsonData)
}

// this two methods need refactor
func (f *Falcon) RecordDuration(name string, duration time.Duration) {
	jsonData, err := f.buildJson(name, int64(duration))
	if err != nil {
		log.Printf("error in build json %v", err)
		return
	}
	f.push(jsonData)
}
func (f *Falcon) Run() {
	// do nothing
}

func (f *Falcon) buildJson(name string, data int64) (string, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return "", err
	}
	metrics := Metrics{
		Metric: name,
		Endpoint: hostName,
		Timestamp: time.Now().Unix(),
		Step: f.Step,
		Value: data,
		CounterType: "GAUGE",
	}
	var mlist []Metrics = make([]Metrics, 1)
	mlist[0] = metrics
	bs, err1 :=  json.Marshal(mlist)
	if err1 != nil {
		return "", err
	}
	jsonData := string(bs)
	return jsonData, nil
}

func (f *Falcon) push(data string) {
	resp, err := f.Client.Post(f.URL, "", strings.NewReader(data))
	if err != nil {
		log.Printf("Push error: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Push response error: %v", err)
		return
	}
	log.Printf("Agent response %s", string(body))
}
