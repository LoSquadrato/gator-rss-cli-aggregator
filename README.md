# Welcome to the Blog Aggregator

We're going to build an RSS feed aggregator in Go! We'll call it "Gator", you know, because aggreGATOR 🐊. Anyhow, it's a CLI tool that allows users to:

* Add RSS feeds from across the internet to be collected
* Store the collected posts in a PostgreSQL database
* Follow and unfollow RSS feeds that other users have added
* View summaries of the aggregated posts in the terminal, with a link to the full post

### A learn-by-doing project
Gator is a multi-user CLI application. There's no server (other than the database), so it's only intended for local use.

### Commands
Gator is a CLI application, and like many CLI applications, it has a set of valid commands. For example:

* `gator login` - sets the current user in the config
* `gator register` - adds a new user to the database
* `gator users` - lists all the users in the database
* etc...

We'll be hand-rolling our CLI rather than using a framework like [Cobra](https://github.com/spf13/cobra) or [Bubble Tea](https://github.com/charmbracelet/bubbletea). This will give us a better understanding of how CLI applications work.
