package responses

import "fmt"

type Gateway struct {
	Data   *map[string]interface{} `json:"data"`
	Errors []*struct {
		Message   string `json:"message"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Extensions struct {
			Code string `json:"code"`
		} `json:"extensions"`
	} `json:"errors"`
}

func (g *Gateway) ErrorsToString() string {
	if g.Errors == nil {
		return ""
	}

	var errors string
	for i, err := range g.Errors {
		errors += fmt.Sprintf("Error %d: %s ", i, err.Message)
	}
	return errors
}
