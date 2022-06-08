module api/host

go 1.18

require (
	api/handlers v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.4.0
)

replace api/handlers => ../handlers
