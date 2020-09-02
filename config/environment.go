package config

import (
	"fmt"
	"ms.api/log"
	"strings"
)

type Environment string

const (
	Local      Environment = "local"
	Staging    Environment = "staging"
	Production Environment = "production"
)

func (e Environment) IsValid() error {
	switch e {
	case Local, Staging, Production:
		return nil
	default:
		return fmt.Errorf("unknwon environment : %s", e)
	}
}

func (e Environment) String() string { return strings.ToLower(string(e)) }

func (e Environment) LogLevel() log.Level {
	switch e {
	case Staging, Production:
		return log.LevelInfo
	case Local:
		fallthrough
	default:
		return log.LevelDebug
	}
}
