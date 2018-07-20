package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koyuta/go-application-sample/infrastructure"
	"github.com/koyuta/go-application-sample/interfaces"
	"github.com/koyuta/go-application-sample/interfaces/middleware"
	"github.com/koyuta/go-application-sample/registry"
)

var (
	DBUser = flag.String("dbuser", "defaultuser", "User name for database")
	DBPass = flag.String("dbpass", "defaultpasswd", "")
	DBHost = flag.String("dbhost", "127.0.0.1", "")
	DBPort = flag.Int("dbport", 3306, "")
	DBName = flag.String("dbname", "database_name", "")

	fd   = flag.Uint("fd", 0, "")
	port = flag.Int("Port", 8000, "")
)

func init() {
	flag.Parse()
}

func main() {
	db, err := infrastructure.NewMySQLHandler(*DBUser, *DBPass, *DBHost, *DBPort, *DBName)
	if err != nil {
		log.Fatal(err)
	}

	u := registry.NewUserRegistry(db).Registry()

	mux := interfaces.NewRouter()
	mux.ContextHandle("/", middleware.CommonMiddleware(u.Get))

	listener, err := makeListener(fd, port)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.Serve(listener, mux); err != nil {
		log.Fatal(err)
	}
}

func makeListener(fd *uint, port *int) (net.Listener, error) {
	if fd != nil {
		return net.FileListener(os.NewFile(uintptr(*fd), ""))
	}
	if port != nil {
		return net.Listen("tcp", fmt.Sprintf(":%d", *port))
	}
	return nil, errors.New("")
}
