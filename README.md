# Dreampicai
Web application for avatar image generation with the use of generative artificial intelligence.

### Development tools
- [DBeaver](https://dbeaver.io)
- [Docker](https://www.docker.com)
- [Golang](https://golang.org)
- [Homebrew](https://brew.sh)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
    ```zsh
    brew install golang-migrate
    ```
- [Sqlc](https://github.com/sqlc-dev/sqlc#installation)
    ```zsh
    brew install sqlc
    ```

### Local environment setup
- Create the local container network
    ``` zsh
    make create-local-container-network
    ```
- Create docker container with PostgreSQL
    ```zsh
    make create-local-container-postgresql
    ```
- Create PosgtreSQL database
    ```zsh
    make create-local-container-postgresql-database
    ```
- Run database migration up all versions
    ```zsh
    make migrate-up
    ```
- Run database migration up 1 version
    ```zsh
    make migrate-up-1
    ```
- Run database migration down all versions
    ```zsh
    make migrate-down
    ```
- Run database migration down 1 version
    ```zsh
    make migrate-down-1
    ```
- Drop database
    ```zsh
    make drop-local-postgresql-database
    ```

### Code generation
- Create a new database migration
    ```zsh
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```
- Generate golang code for SQL queries
    ```zsh
    make sqlc
    ```

### Application run
- Run application
    ```zsh
    make run
    ```
    