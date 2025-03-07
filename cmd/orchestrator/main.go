package main

import (
	"log"

	"github.com/Danilka776/web_go_calc_2/internal/orchestrator/api"
)

func main() {
	log.Println("Запуск оркестратора...")
	if err := api.SetupRouter(); err != nil {
		log.Fatalf("Ошибка при запуске оркестратора: %v", err)
	}
}
