package main

import (
	"log"

	"github.com/asvinicius/actnsgo/internal/client"
	"github.com/asvinicius/actnsgo/internal/config"
	"github.com/asvinicius/actnsgo/internal/db"
	"github.com/asvinicius/actnsgo/internal/repository"
	"github.com/asvinicius/actnsgo/internal/service/marketstatus"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	pool, err := db.Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	msrepo := repository.NewMarketStatusRepository(pool)
	mscli := client.NewClient(cfg.Cartola.CartolaURL)
	mss := marketstatus.NewMarketStatusService(msrepo, mscli)

	closed, err := mss.HasMarketClosed()

	if err != nil {
		log.Fatal(err)
	}

	if !closed {
		return
	}

	// chama a execução do backup
}
