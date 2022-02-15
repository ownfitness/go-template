package cfg

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config struct to define the environment variables to parse
// to create a config object to use in the application
type Config struct {
	Debug   bool   `envconfig:"debug" default:"false"`
	Port    string `envconfig:"port" default:"8080"`
	Project string `envconfig:"project" required:"true"`
}

// New config object that contains the logic to
// process the environment variables and also
// enables the ability to manipulate the object
// before sending it back to the application
func New() (Config, error) {
	var c Config

	err := envconfig.Process("", &c)
	if err != nil {
		return c, fmt.Errorf("error processing environment config: %s", err.Error())
	}

	return c, nil
}
