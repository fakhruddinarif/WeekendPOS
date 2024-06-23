# WeekendPOS

### Migrate Install
```bash
go install -tags 'postgres,mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
### Add Migration
```bash
migrate create -ext sql -dir db/migrations create_table_name
```
### Run Migration
```bash
migrate -path db/migrations -database "mysql://user:pass@tcp(localhost:3306)/name_database" -path db/migrations up
```