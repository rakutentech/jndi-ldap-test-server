package main

import (
	"fmt"
	"github.com/rakuten-tech/jndi-ldap-test-server/args"
	"github.com/rakuten-tech/jndi-ldap-test-server/routes"
	"github.com/rakuten-tech/jndi-ldap-test-server/util/logging"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	ldap "github.com/vjeantet/ldapserver"
)

func main() {
	logging.InitializeLogger()
	err := args.RunWithArgs(os.Args, runApp)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to launch server")
	}
}

func runApp(c *cli.Context) error {
	logging.UpdateLoggerWithFlags(
		&logging.Flags{
			Color: c.String("color"),
			Level: c.String("log-level"),
		},
	)
	ldapLogger := log.With().Str("component", "server").Logger()
	ldap.Logger = logging.NewStdAdapter(&ldapLogger)

	payload := c.String("exploit-payload")
	if payload != "" {
		routes.SetVulnerablePayload(payload)
	}

	server := ldap.NewServer()
	server.Handle(routes.AllRoutes())

	listenAddress := c.String("listen-address")
	port := c.Int("port")

	go func() {
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
	}()

	// Quit gracefully on SIGINT (CTRL-C) and SIGTERM
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	close(ch)

	server.Stop()

	return nil
}
