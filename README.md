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
