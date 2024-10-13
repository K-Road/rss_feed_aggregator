package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/K-Road/rss_feed_aggregator/internal/config"
	"github.com/K-Road/rss_feed_aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("Error opening db")
	}
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cliCommands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cliCommands.register("login", handlerLogin)
	cliCommands.register("register", handlerRegister)
	cliCommands.register("reset", handlerReset)
	cliCommands.register("users", handlerUsers)
	cliCommands.register("agg", handlerAgg)
	cliCommands.register("addfeed", handlerAddFeed)
	cliCommands.register("feeds", handlerListFeeds)
	cliCommands.register("follow", handlerFollow)
	cliCommands.register("following", handlerListFeedFollows)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	commandName := os.Args[1]
	args := os.Args[2:]

	err = cliCommands.run(programState, command{Name: commandName, Arguments: args})
	if err != nil {
		log.Fatal(err)
	}

	// cfg, err = config.GetConfig()
	// if err != nil {
	// 	log.Fatalf("error reading config: %v", err)
	// }
	// fmt.Printf("Read config again: %+v\n", cfg)

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Fatal("PORT env not set")
	// }

	// mux := http.NewServeMux()
	// srv := &http.Server{
	// 	Addr:    ":" + port,
	// 	Handler: mux,
	// }

	// mux.HandleFunc("GET /v1/healthz", handlerhealthz)
	// mux.HandleFunc("GET /v1/err", handlererr)

	// if err := srv.ListenAndServe(); err != nil {
	// 	fmt.Println("Server failed:", err)
	// }
}
