# Go, GORM & Gin CRUD Example

## Install
Best use Go 1.11 or newer for Go module dependency management

1. Clone this repository:
   ```
   git clone https://github.com/ppxl/gorm_crud_example.git
   ```
2. CD to `gorm_crud_example` folder:
   ```
   cd $GOPATH/src/github.com/ppxl/gorm_crud_example
   ```
3. Open `main.go` and modify this variable values:
   ```
   dbUser, dbPassword, dbName := "postgres", "postgres", "mydb"
   ```
4. Start `PostgreSQL` and create the database:
   ```
   docker-compose up -d
   docker exec -it gorm_crud_example_db_1 psql --username postgres
     create database mydb;
     \q
   ```
5. Run `main.go`:
   ```
   go run main.go
   ```

## Features

- \[x] Database Migration
- \[x] Create Data
- \[x] Read All Data
- \[x] Find One Data By ID
- \[x] Update Data
- \[x] Delete One Data By ID
- \[x] Delete Multiple Data By IDs
- \[x] Sort & Paginate Data
- \[x] Search Data
