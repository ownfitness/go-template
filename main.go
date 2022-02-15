package main

import (
	"fmt"

	"github.com/ownfitness/template-go/pkg/gcp"

	"github.com/ownfitness/template-go/pkg/router"

	"github.com/rs/zerolog"

	"github.com/ownfitness/template-go/pkg/cfg"

	"github.com/rs/zerolog/log"
)

func main() {
	// Generate new config object
	config, err := cfg.New()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	// Set zerolog global default level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	clients, err := gcp.FirebaseClient(config.Project)

	// Generate new gin.Engine
	r := router.New(config.Debug, clients)

	if err := r.Run(fmt.Sprintf(":%s", config.Port)); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
