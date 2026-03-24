package main

import (
	"context"
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
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting user: %v", err)
		}
		return handler(s, cmd, user)
	}
}
