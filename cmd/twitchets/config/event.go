package config

import "gopkg.in/yaml.v3"

type Event struct {
	Name       string   `json:"name"`
	Similarity *float64 `json:"Similarity"`
}

func (e *Event) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {

	case yaml.ScalarNode:
		var name string
		err := node.Decode(&name)
		if err != nil {
			return err
		}

		e.Name = name

	case yaml.MappingNode:
		type eventAlias Event
		var event eventAlias
		err := node.Decode(&event)
		if err != nil {
			return err
		}

		*e = Event(event)
	}

	return nil
}
