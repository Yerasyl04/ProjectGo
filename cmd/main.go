package main

import "Projects/Go/internal/config"
import "fmt"
import "golang.org/x/exp/slog"
import "log"
import  "github.com/Yerasyl04/ProjectGo/internal/config"



func main() {
	// TODO: init config: cleanenv
	// TODO: init logger: slog
	// TODO: init storage: sqlite
	// TODO: init router: chi,"chi render"
	// TODO: run server:
	cfg := config.Mustload()
	fmt.Println(cfg)
	log:= setupLogger(cfg.env)
	log.Info("Starting url-shortener",slog.String("env",cfg.Env))
	log.Debug("Debug messages are enabled")
}
func setupLogger(env string){
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New {
			slog.NewTextHandler(os.Stdout,&slog.HandlerOptions{Level:slog.LevelDebug})
		}
	case envDev :
		log = slog.New {
			slog.NewJSONHandler(os.stdout,&slog.HandlerOptions{Level:slog.LevelDebug})
		}
	case envProd :
		log = slog.New {
			slog.NewJSONHandler(os.stdout,&slog.HandlerOptions{Level:slog.LevelInfo})
		}
		rerurn log



	}
}