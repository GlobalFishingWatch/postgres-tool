
## Local Examples

### Create a view
```
go run main.go create-view --table-name=global_vessels_2020_02_02 --view-name=vessels_3 --postgres-address=localhost:5432 --postgres-user=postgres --postgres-password=a1234567 --postgres-database=postgres
```

### Delete a view
```
go run main.go delete-view --view-name=vessels_3 --postgres-address=localhost:5432 --postgres-user=postgres --postgres-password=a1234567 --postgres-database=postgres
```

### Delete table
```
go run main.go delete-table --table-name=vessels_2020_02_01 --postgres-address=localhost:5432 --postgres-user=postgres --postgres-password=a1234567 --postgres-database=postgres
```

### Execute raw sql
```
go run main.go execute-raw-sql --sql="CREATE TABLE IF NOT EXISTS public.event_test (id BIGSERIAL, event_start timestamp without time zone NOT NULL, event_end timestamp without time zone, primary key (id, event_start)) PARTITION BY RANGE (event_start); CREATE TABLE IF NOT EXISTS public.event_test_2021  PARTITION OF public.event_test FOR VALUES FROM ('2021-01-01') TO ('2020-01-01');" --postgres-address=localhost:5432 --postgres-user=postgres --postgres-password=a1234567 --postgres-database=postgres
```


