package floz

import "encoding/json"

type Config struct {
	JSONEncoder JSONMarshal
}

func NewConfig() *Config {
	return &Config{
		JSONEncoder: json.Marshal,
	}
}
