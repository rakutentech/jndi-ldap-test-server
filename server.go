package main

import (
	"fmt"
	"github.com/rakuten-tech/jndi-ldap-test-server/routes"
	"github.com/rs/zerolog/log"
	"github.com/vjeantet/ldapserver"
	ldap "github.com/vjeantet/ldapserver"
)

func startServer(listenAddress string, port int) *ldapserver.Server {
	server := ldapserver.NewServer()
	server.Handle(routes.AllRoutes())

	go listen(server, listenAddress, port)

	return server
}

func listen(server *ldap.Server, listenAddress string, port int) {
	err := server.ListenAndServe(fmt.Sprintf("%s:%d", listenAddress, port), func(server *ldap.Server) {
		// Called if server is listening successfully
		log.Info().
			Str("component", "server").
			Str("event", "listen").
			Str("listen_address", listenAddress).
			Int("port", port).
			Msgf("Listening on %s:%d", listenAddress, port)
	})
	if err != nil {
		log.Fatal().
			Str("component", "server").
			Str("event", "listen").
			Str("listen_address", listenAddress).
			Int("port", port).
			Err(err).
			Msgf("Cannot listen on %s:%d", listenAddress, port)
	}
}
