package input

import (
	"bytes"
	"fmt"

	"../model"
	"gopkg.in/yaml.v2"
)

func Unmarshall(data []byte) (model.Diagram, error) {
	if len(data) == 0 {
		return model.Diagram{}, fmt.Errorf("data cannot be empty")
	}

	r := bytes.NewReader(data)
	dec := yaml.NewDecoder(r)

	var d model.Diagram

	if err := dec.Decode(&d); err != nil {
		return model.Diagram{}, err
	}

	return d, nil
}
