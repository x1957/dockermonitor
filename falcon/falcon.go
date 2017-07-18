package falcon

type Falcon struct {
	URL     string   `json:"url"`
	Metrics string   `json:"metrics"`
	Tests   []string `json:"tesing,omitempty"`
}

func Run() {

}
