
# Build from code
## Requirment to build this code 
- Golang 1.20

## Building instructions
- **Step 1.1.** Change directory to the root of this project
```bash
cd <root_folder>
```
- **Step 1.2.** Build the API Server
```bash
go build -o build/short_url_api_server ./cmd/api-server
```

- **Step 1.3.** Build the redirection server
```bash
go build -o build/short_url_redirection_server ./cmd/redirection-server
```

# Install
The binaries can be installed seperately, in two different machines, the only requirement is that they should be able to access the same database.

- **Step 2.1.** Copy the binary of the API Server (`build/short_url_api_server`) to `/usr/local/bin` (or any other folder in your `PATH`).
```bash
sudo cp build/short_url_api_server /usr/local/bin
```

- **Step 2.2.** Copy the binary of the redirection Server (`build/short_url_redirection_server`) to `/usr/local/bin` (or any other folder in your `PATH`).
```bash
sudo cp build/short_url_redirection_server /usr/local/bin
```

- **Step 2.3.** Copy the `redirection_404.html` and `redirection_500.html` files to a dedicated folder (for example `/opt/short_url/`).
```bash
sudo mkdir /opt/short_url
sudo cp redirection_404.html /opt/short_url/
sudo cp redirection_500.html /opt/short_url/
```

# Database preparation
- **Step 3.1.** Create a dedicated user for the application.
```bash
sudo -u postgres createuser <username>
```

- **Step 3.2.** Create a new database in Postgres.
```bash
sudo -u postgres createdb <dbname>
```
- **Step 3.3.** Give the user a password.
```bash
sudo -u postgres psql
psql=# alter user <username> with encrypted password '<password>';
```

- **Step 3.4.** Grant the user access to the database.
```bash
psql=# grant all privileges on database <dbname> to <username>;
```

- **Step 3.5.** Run the following script to create the tables: `./db/migration/000001_init_schema.up.sql`
- **Step 3.6.** Take a note of the following information: 
    - Hostname : The IP adress or the hostname of the Postgres server.  
    - Port : The network port on which the Postgres server is listening (usually *5432*)
    - Database : The name of the database created in step 3.2.
    - Username : The name of the user created in step 3.1.
    - Password : The password of the user created in step 3.3.

# Configuration 
Create the following environment variables:

Please ensure that you set the required environment variables accordingly, while the optional ones can be adjusted as per your specific needs. 
The default values, if applicable, will be used when the optional variables are not explicitly provided.

## Common configuration 
This is a list of common environment variables that are used by both servers (if the servers are installed on different machines, those environment variables should be set on both machines)

| Name                        | Description                                                  | Required/Optional | Default Value | Example                                                                  |
|-----------------------------|--------------------------------------------------------------|-------------------|---------------|--------------------------------------------------------------------------|
| `DB_DRIVER`                 | Database driver name.                                        | Required          |               | `postgres`                                                               |
| `DB_SOURCE`                 | URL of the PostgreSQL database. More detail about the format of this parameter can be found here: https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters                              | Required          |               | `postgresql://postgres:password@localhost:5433/postgres?sslmode=disable`1 |
| `DB_MAX_IDLE_CONNS`         | Maximum number of idle connections in the connection pool.    | Optional          | 5             | `5`                                                                      |
| `DB_MAX_OPEN_CONNS`         | Maximum number of open connections in the connection pool.    | Optional          | 10            | `10`                                                                     |
| `DB_CONN_MAX_IDLE_TIME`     | Maximum time (in minutes) a connection can be idle.           | Optional          | 1             | `1`                                                                      |
| `DB_CONN_MAX_LIFE_TIME`     | Maximum time (in minutes) a connection can be kept open.      | Optional          | 30            | `30`                                                                     |

## API Server configuration
This is a list of environment variables that are used by the API Server only.
| Name                        | Description                                                  | Required/Optional | Default Value | Example                                                                  |
|-----------------------------|--------------------------------------------------------------|-------------------|---------------|--------------------------------------------------------------------------|
| `ADMIN_API_KEY`             | API key used by the admin.                                   | Required          |               | `1234`                                                                   |
| `REDIRECTION_SERVER_BASE_URL` | Base URL of the redirection server, this must be a public accessible URL                           | Required          |               | `https://domain.name`                                                  |

## Redirection Server configuration
This is a list of environment variables that are used by the Redirection Server only.
| Name                        | Description                                                  | Required/Optional | Default Value | Example                                                                  |
|-----------------------------|--------------------------------------------------------------|-------------------|---------------|--------------------------------------------------------------------------|
| `REDIRECTION_404_PAGE`      | Path to the 404 error page for the redirection server.       | Required          |               | `/opt/short_url/redirection_404.html`                                    |
| `REDIRECTION_500_PAGE`      | Path to the 500 error page for the redirection server.       | Required          |               | `/opt/short_url/redirection_500.html`                                    |

# Launch

## API Server
Launch the API Server using the following command, by specifyging port and host:

```bash
short_url_api_server --port 8080 --host localhost 
```

### Parameters : 
    - `--port` : The port on which the API Server will listen to.
    - `--host` : The hostname or IP address of the host machine.

### API Documentation: 
A detailed documentation of the API can be found [here](api.md)

## Redirection Server
Launch the Redirection Server using the following command, by specifyging port and host:

```bash
short_url_redirection_server --port 8085 --host localhost 
```
### Parameters : 
    - `--port` : The port on which the Redirection Server will listen to.
    - `--host` : The hostname or IP address of the host machine.
