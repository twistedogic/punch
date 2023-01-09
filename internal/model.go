package internal

import (
	"net/url"
	"time"

	"gopkg.in/yaml.v3"
)

type Duration time.Duration

func (d *Duration) UnmarshalYAML(node *yaml.Node) error {
	dur, err := time.ParseDuration(node.Value)
	if err != nil {
		return err
	}
	*d = dur
	return nil
}

type URL url.URL

func (u *URL) UnmarshalYAML(node *yaml.Node) error {
	endpoint, err := url.Parse(node.Value)
	if err != nil {
		return err
	}
	*u = *endpoint
	return nil
}

type Bytes []byte

func (b *Bytes) UnmarshalYAML(node *yaml.Node) error {
	*b = []byte(node.Value)
	return nil
}
