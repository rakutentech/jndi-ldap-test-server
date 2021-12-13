package main

import (
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
	// Setup according to CLI flags
	setupLogger(c)
	c.StringSlice("dynamic-payloads")
	routes.SetExploit(routes.ParseExploitSettings(
		c.String("payload"),
		args.GetEnumValueSet(c, "dynamic-payloads"),
	))
	server := startServer(c.String("listen-address"), c.Int("port"))

	// Quit gracefully on SIGINT (CTRL-C) and SIGTERM
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	close(ch)

	server.Stop()

	return nil
}

func setupLogger(c *cli.Context) {
	logging.UpdateLoggerWithFlags(
		&logging.Flags{
			Color: c.String("color"),
			Level: c.String("log-level"),
		},
	)
	ldapLogger := log.With().Str("component", "server").Logger()
	ldap.Logger = logging.NewStdAdapter(&ldapLogger)
}
