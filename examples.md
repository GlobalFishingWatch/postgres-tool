
## Local Examples

### Create a view
```
go run main.go create-view --table-name=global_vessels_2020_02_02 --view-name=vessels_3 --postgres-address=localhost:5432 --postgres-user=postgres --postgres-password=a1234567 --postgres-database=postgres
```

### Delete a view
```
go run main.go delete-view --view-name=vessels_3 --postgres-address=localhost:5432 --postgres-user=postgres --postgres-password=a1234567 --postgres-database=postgres
```