package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.ClearUser(context.Background()); err != nil {
		return fmt.Errorf("error clearing users table: %v\n", err)
	}
	fmt.Println("Successfully cleared users table")
	return nil
}
