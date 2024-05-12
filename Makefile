build:
	@templ generate view
	@go build -tags dev -o bin/dreampicai main.go 

css:
	@tailwindcss -i view/css/app.css -o public/styles.css --watch 

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D daisyui@latest

migratedown:
	@go run cmd/migrate/main.go down

migrateup:
	@go run cmd/migrate/main.go up

reset:
	@go run cmd/reset/main.go up

run: build
	@./bin/dreampicai

templ:
	@templ generate --watch --proxy=http://localhost:3000

seed:
	@go run cmd/seed/main.go
