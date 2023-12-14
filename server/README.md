# Server

API Client built using [go-chi](https://github.com/go-chi/chi).

## Getting Started

To lessen the install burden, we'll work with postgres on a docker container via PORT# 5432

1. Ensure that you have docker desktop installed. If not download it here https://www.docker.com/products/docker-desktop/
2. Once your project has been cloned navigate to the **server/Makefile** file.
3. Makefile commands are pretty self-explanatory.
   - `foo@bar:RetailGo/server$ make postgres` spins up a docker container
   - `foo@bar:RetailGo/server$ make createdb` creates a database -> retail_go
   - `foo@bar:RetailGo/server$ make drodb` deletes a database -> retail_go
   - `console foo@bar:RetailGo/server$ make sqlc_delete` deletes the **db/sqlc** folder.
4. run **make postgres** **make createdb** and Configure TablePlus to connect to the retail_go database via PORT# 5432 (Password: secret)
5. Run `console foo@bar:RetailGo/server$ go run main.go`

##### Docs

- [Go-Chi](https://go-chi.io/)
- [Ent](https://entgo.io/docs/getting-started)
