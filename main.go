package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/database"

	_ "github.com/lib/pq"

	"github.com/LoSquadrato/gator-rss-cli-aggregator/internal/config"
)

// State holds the database connection and the config
type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	fmt.Println("Starting GATOR")
	// Read config .gatorconfig.json file
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config file: %v\n", err)
		os.Exit(1)
	}
	// Open database connection
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("error opening database connection: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	// Create database queries
	dbQueries := database.New(db)
	// Create state and commands
	state := &state{db: dbQueries, cfg: cfg}
	cmds := &commands{}
	// Register commands
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	// ERROR if no command is provided
	if len(os.Args) < 2 {
		fmt.Println("error too less args in commands")
		os.Exit(1)
	}
	// Run command
	cmd := command{name: os.Args[1], arguments: os.Args[2:]}
	err = cmds.run(state, cmd)
	if err != nil {
		fmt.Printf("error run command:\n %s", err)
		os.Exit(1)
	}
}
