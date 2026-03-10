package main

import (
    "fmt"
	"errors"
)

type ConfigError struct {
	Line int
	Detail string
}

// Error() メソッドを実装して、ConfigErrorがerrorインターフェースを満たすようにする
func (e *ConfigError) Error() string {
	return fmt.Sprintf("line %d: %s", e.Line, e.Detail)
}

var ErrInvalidPath = errors.New("invalid configuration path")

func ParseConfig(path string) error {
	if path == "" {
		return ErrInvalidPath
	}
	err := &ConfigError{Line: 3, Detail: "missing required field 'port'"}
	return fmt.Errorf("failed to parse config: %w", err)
}

func main() {
	err := ParseConfig("config.yaml")
	if err != nil {
		if errors.Is(err, ErrInvalidPath) {
			fmt.Println("Configuration path is invalid")
		}
		var configErr *ConfigError
		if errors.As(err, &configErr) {
			fmt.Printf("Config error at line %d: %s\n", configErr.Line, configErr.Detail)	
		}
	}
}