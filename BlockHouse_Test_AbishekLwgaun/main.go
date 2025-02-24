package main

import (
	"BlockHouse_Test_AbishekLwgaun/db-connect"
	"BlockHouse_Test_AbishekLwgaun/routers"
	"log"
	"net/http"
)

func main() {
	//connects to the db
	db_connect.ConnectDB()

	//gin routr with all routes
	r := routers.Create_routes()

	// logs for server
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	// starting the gin server
	r.Run(":8080")
}
