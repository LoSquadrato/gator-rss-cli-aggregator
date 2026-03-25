# 🐊 Gator — RSS CLI Aggregator

Gator is a simple, multi-user RSS feed aggregator built in Go. It allows you to collect, manage, and browse RSS feeds directly from your terminal.

Unlike traditional web-based aggregators, Gator is a **local-first CLI application** — there’s no backend server, just a PostgreSQL database.

---

## ✨ Features

- 📡 Add RSS feeds from across the internet
- 🗃️ Store posts in a PostgreSQL database
- 👥 Multi-user support (locally)
- ➕ Follow and unfollow feeds
- 📖 Browse aggregated posts in the terminal
- ⚡ Lightweight and fast CLI workflow

---

## 🖥️ Example Usage

```bash
$ gator-rss-cli-aggregator register alice
User "alice" created successfully!

$ gator-rss-cli-aggregator login alice

$ gator-rss-cli-aggregator addfeed https://example.com/rss
Feed added!

$ gator-rss-cli-aggregator follow https://example.com/rss
Now following feed!

$ gator-rss-cli-aggregator browse 3

1. "Go Concurrency Patterns"  
   https://example.com/post1

2. "Understanding Interfaces in Go"  
   https://example.com/post2

3. "Building CLI Apps in Go"  
   https://example.com/post3
````

---

## 🧰 Tech Stack

* **Go** — CLI application
* **PostgreSQL** — data storage
* **Goose** — database migrations

---

## 🚀 Getting Started

### 1. Prerequisites

* Go **1.26+**
* PostgreSQL **15+**

---

### 2. Clone the Repository

```bash
git clone https://github.com/LoSquadrato/gator-rss-cli-aggregator.git
cd gator-rss-cli-aggregator
```

---

### 3. Set Up PostgreSQL

Start the PostgreSQL service:

```bash
# Mac
brew services start postgresql@15

# Linux
sudo service postgresql start
```

Open the PostgreSQL shell:

```bash
# Mac
psql postgres

# Linux
sudo -u postgres psql
```

Create and connect to the database:

```sql
CREATE DATABASE gator;
\c gator;
```

(Optional for Linux/WSL)

```sql
ALTER USER postgres PASSWORD 'postgres';
```

---

### 4. Install Goose (Migration Tool)

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

---

### 5. Run Database Migrations

Navigate to the schema directory:

```bash
cd sql/schema
```

Run migrations:

```bash
goose postgres "<connection_string>" up
```

Example:

```bash
goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

---

### 6. Configure the App

Create a config file:

```bash
~/.gatorconfig.json
```

Add:

```json
{
  "db_url": "<connection_string>"
}
```

---

### 7. Install the CLI

```bash
go install github.com/LoSquadrato/gator-rss-cli-aggregator@latest
```

---

### 8. Run the Application

```bash
gator-rss-cli-aggregator <command>
```

---

## 📜 Commands

| Command               | Description           |
| --------------------- | --------------------- |
| `register <username>` | Create a new user     |
| `login <username>`    | Set the current user  |
| `users`               | List all users        |
| `addfeed <url>`       | Add a new RSS feed    |
| `feeds`               | List all feeds        |
| `follow <url>`        | Follow a feed         |
| `following`           | List followed feeds   |
| `unfollow <url>`      | Unfollow a feed       |
| `agg <interval>`      | Fetch and store posts |
| `browse <n>`          | View latest posts     |
| `reset`               | ⚠️ Reset the database |

---

## ⚠️ Warning

The `reset` command will permanently delete **all data** in the database.

---

## 🧠 How It Works

* Users and feeds are stored in PostgreSQL
* RSS feeds are fetched using the `agg` command
* Posts are parsed and saved to the database
* Users can follow feeds to personalize their content
* The `browse` command displays posts from followed feeds

---

## 🔮 Future Improvements

* ⏱️ Background worker for automatic feed polling
* 🌐 Web interface
* 🐳 Docker support
* 🏷️ Feed categories and tagging

---

## 🙏 Acknowledgements

Special thanks to [Boot.dev](https://bootdev.com/) and their excellent learning resources.

Check out their GitHub: [https://github.com/bootdev](https://github.com/bootdev)

---

## 📄 License

This project is open-source and available under the MIT License. See [LICENSE](LICENSE) for details.