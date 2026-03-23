package main

import "context"

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.ListUser(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			println(user.Name + " (current)")
			continue
		}
		println(user.Name)
	}
	return nil
}
