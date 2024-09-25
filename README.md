# Blog Aggregator

This is a command-line blog aggregator application written in Go. It allows users to register, login, manage RSS feeds, and aggregate blog posts from various sources.

## Features

- User authentication (register, login, reset)
- RSS feed management (add, list, follow, unfollow)
- Blog post aggregation
- Database integration with PostgreSQL

## Prerequisites

- Go 1.x
- PostgreSQL database
- Required Go packages (see `import` statements in the code)

## Configuration

The application uses a configuration file (format not specified in the provided code). Ensure you have a proper configuration file with the following details:

- Database URL
- Current username (for authentication)

## Usage

Run the application with a command and its arguments:

```
go run main.go <command> [arguments]
```

Available commands:

- `login`: Authenticate user
- `register`: Register a new user
- `reset`: Reset user data (specifics not provided in the code)
- `users`: List users (requires authentication)
- `agg`: Aggregate blog posts (specifics not provided in the code)
- `addfeed`: Add a new RSS feed (requires authentication)
- `feeds`: List all feeds
- `follow`: Follow a feed (requires authentication)
- `following`: List followed feeds
- `unfollow`: Unfollow a feed (requires authentication)

## Database Schema

The application uses a PostgreSQL database with the following tables (inferred from the code):

- Users
- Feeds
- Posts

Exact schema details are not provided in the given code.

## Main Components

1. `main()`: Entry point of the application, handles command parsing and execution.
2. `middlewareLoggedIn()`: Middleware to check if a user is logged in before executing certain commands.
3. `scrapFeeds()`: Function to fetch and store blog posts from RSS feeds.
4. Helper functions: `StringToNullString()` and `DateToNullDate()` for database operations.

## Error Handling

The application uses Go's standard error handling mechanisms. Fatal errors are logged, and the program exits with a non-zero status code.

## Extensibility

New commands can be added by implementing handler functions and registering them in the `main()` function using the `c.register()` method.

## Note

This README is based on the provided code snippet. Some details about specific functionalities, database schema, and configuration might be missing or inaccurate without access to the complete codebase.
