package httpserver

import (
	"github.com/thebravebyte/findr/internals"
)



// AddRoutes adds app loggeer function to help keep track of the http request pipeline, and also
// middleware, cors middleware methods to handle the http request pipeline from the server to 
// to the client
func (srv *Server) AddRoutes(fa internals.FindrAppService) {

	srv.LoggingAppRequest(srv.logger)

	// srv.AddMiddleware()

	srv.StaticFile("/", "assests")

	srv.mux.HandleFunc("GET /", fa.FindrHomePage())
	srv.mux.HandleFunc("POST /sign-up", fa.Register())

}
