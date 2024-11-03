package env

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var AdminManagementPort = "7500"

// Port end
var AdminManagementURL = "/api/v1/"

var SysEmail = ""
var SysEmailPassword = ""

// FileDirectory variable
var FileDirectory = "/mnt/disks/static_files/uploads/"

// CROS origin setuped
var HeadersOk = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "x-token", "current-time", "utc-time"})
var OriginsOk = handlers.AllowedOrigins([]string{("*")})
var MethodsOk = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

// ServerStartList func used to Server start init
func ServerStartList(r *mux.Router, port string) *http.Server {
	// set timeouts to avoid Slowloris attacks.
	var srv = &http.Server{
		Addr:         "localhost:" + port,
		WriteTimeout: time.Second * 45,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(OriginsOk, HeadersOk, MethodsOk)(r), // Pass our instance of gorilla/mux in.
	}
	log.Println("Server location:", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
	return srv
}
