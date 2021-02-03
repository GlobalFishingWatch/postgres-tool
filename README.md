# postgres-tool

## Description

postgres-tool is an agnostic CLI to expose commands which allows you to manage command using Postgres.

Format:
```
postgres [command] [--flags]
```

### Tech Stack:
* [Golang](https://golang.org/doc/)
* [Cobra Framework](https://github.com/spf13/cobra#working-with-flags)
* [Viper](https://github.com/spf13/viper)
* [Docker](https://docs.docker.com/)

### Git
* Repository:
  https://github.com/GlobalFishingWatch/postgres-tool

## Usage

There are available the following commands:
* create-view

---

### Command: [create-view]

The create-view command allows you to create a new view.

#### Flags
##### Required flags
- `--table-name=` The name of the table to link with the view
- `--view-name=` The name of the destination view
- `--postgres-address=` The database address and port.
- `--postgres-user=` The database user.
- `--postgres-password=` The database password.
- `--postgres-database=` The destination name database.

##### Optional flags

#### Example
Here an example of this command:
```
postgres create-view \
  --table-name="vessels_2021_02_01" \
  --view-name="vessels" \
  --postgres-address="localhost:5432" \
  --postgres-user="postgres" \
  --postgres-password="XaD2sd$34Sdas1$ae" \
  --postgres-database="postgres" 
```


### Command: [delete-view]

The delete-view command allows you to delete a view.

#### Flags
##### Required flags
- `--view-name=` The name of the destination view
- `--postgres-address=` The database address and port.
- `--postgres-user=` The database user.
- `--postgres-password=` The database password.
- `--postgres-database=` The destination name database.

##### Optional flags

#### Example
Here an example of this command:
```
postgres delete-view \
  --view-name="vessels" \
  --postgres-address="localhost:5432" \
  --postgres-user="postgres" \
  --postgres-password="XaD2sd$34Sdas1$ae" \
  --postgres-database="postgres" 
```

### Command: [delete-table]

The delete-table command allows you to delete a table.

#### Flags
##### Required flags
- `--table-name=` The name of the table to delete
- `--postgres-address=` The database address and port.
- `--postgres-user=` The database user.
- `--postgres-password=` The database password.
- `--postgres-database=` The destination name database.

##### Optional flags

#### Example
Here an example of this command:
```
postgres delete-table \
  --table-name="vessels" \
  --postgres-address="localhost:5432" \
  --postgres-user="postgres" \
  --postgres-password="XaD2sd$34Sdas1$ae" \
  --postgres-database="postgres" 
```

### Command: [create-index]

The create-index command allows you to create a index for an existing table.

#### Flags
##### Required flags
- `--table-name=` The name of the destination table
- `--index-name=` The name of the new index
- `--column=` The name of the column
- `--postgres-address=` The database address and port.
- `--postgres-user=` The database user.
- `--postgres-password=` The database password.
- `--postgres-database=` The destination name database.

##### Optional flags

#### Example
Here an example of this command:
```
postgres delete-table \
  --table-name="vessels" \
  --index-name="vessels_id_index" \
  --column="id" \
  --postgres-address="localhost:5432" \
  --postgres-user="postgres" \
  --postgres-password="XaD2sd$34Sdas1$ae" \
  --postgres-database="postgres" 
```
