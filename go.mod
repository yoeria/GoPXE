module github.com/yoeria/gopxe

go 1.18

require (
	github.com/coreos/bbolt v1.3.0
	github.com/gorilla/mux v1.6.2
)

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	golang.org/x/sys v0.0.0-20181011152604-fa43e7bc11ba // indirect
)

replace (
	github.com/yoeria/gopxe/conf => ./conf
	github.com/yoeria/gopxe/db => ./db
	github.com/yoeria/gopxe/handlers => ./handlers
	github.com/yoeria/gopxe/routes => ./routes
)
