package stake

import "gopkg.in/yaml.v2"

// A Claim made by a claimant about a subject, e.g. '12345' is Joe's userid.
type Claim struct {
	Subject string `json:"subject"`
	Claim   string `json:"claim"`
}

func (claim *Claim) String() string {
	yaml, err := yaml.Marshal(&claim)
	if err != nil {
		panic(err)
	}
	return string(yaml)
}
