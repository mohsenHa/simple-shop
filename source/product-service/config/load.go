package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"strings"
)

func Load(configPath string) (Config, error) {

	// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
	var k = koanf.New(".")

	// Load default values using the confmap provider.
	// We provide a flat map with the "." delimiter.
	// A nested map can be loaded by setting the delimiter to an empty string "".
	err := k.Load(confmap.Provider(defaultConfig, "."), nil)
	if err != nil {
		return Config{}, err
	}

	// Load YAML config and merge into the previously loaded config (because we can).
	err = k.Load(file.Provider(configPath), yaml.Parser())
	if err != nil {
		return Config{}, err
	}

	err = k.Load(env.Provider("MESSENGER_", ".", func(s string) string {
		str := strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, "MESSENGER_")), "_", ".")

		return strings.ReplaceAll(str, "..", "_")
	}), nil)

	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}

	return cfg, nil
}
