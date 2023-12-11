# RetailGo
RetailGo


Contributors: Colby Frey, Apramay Singh, David Nguyen, Giridhar Vadhul, Leo Asatoorian, Jonathan Michel


### Getting Started with the Backend!

```console foo@bar:~$ git clone git@github.com:hktrib/RetailGo.git```


To lessen the install burden, we'll work with postgres on a docker container via PORT# 5432

1. Ensure that you have docker desktop installed. If not download it here https://www.docker.com/products/docker-desktop/

4. Once your project has been cloned navigate to the **server/Makefile** file.


5. Makefile commands are pretty self-explanatory. 
    - ```foo@bar:RetailGo/server$ make postgres``` spins up a docker containers
    - ```foo@bar:RetailGo/server$ make createdb``` creates a database -> retail_go
    - ```foo@bar:RetailGo/server$ make drodb``` deletes a database -> retail_go
    - ```console foo@bar:RetailGo/server$ make sqlc_delete``` deletes the **db/sqlc** folder.
6. run **make postgres** **make createdb** and Configure TablePlus to connect to the retail_go database via PORT# 5432 (Password: secret)
7. Run ```console foo@bar:RetailGo/server$ go run main.go``` 

##### Working With Ent
https://entgo.io/docs/getting-started

##### Docs
Go-Chi -> https://go-chi.io/
