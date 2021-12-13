package args

import (
	"github.com/urfave/cli/v2"
)

func RunWithArgs(args []string, runApp func(ctx *cli.Context) error) error {
	app := &cli.App{
		Name: "jndi-ldap-test-server",
		Description:
		"A minimalistic LDAP server that is meant for test vulnerability to JNDI+LDAP injection attacks" +
			"in Java, especially CVE-2021-44228.",
		Usage: "jndi-ldap-test-server [options]",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "port to listen on",
				Value:   1389,
			},
			&cli.StringFlag{
				Name:  "listen-address",
				Usage: "network address to listen on",
				Value: "0.0.0.0",
			},
			&cli.StringFlag{
				Name:  "payload",
				Usage: "set the default exploit payload string",
				Value: "!!! VULNERABLE !!!",
			},
			&cli.GenericFlag{
				Name:    "dynamic-payloads",
				Aliases: []string{"d"},
				Usage:   "allow dynamic payloads (allowed: string)",
				Value:   &EnumValueSet{
					Enum:       []string{"string"},
				},
			},
			&cli.GenericFlag{
				Name:  "color",
				Usage: "force console color settings (allowed: auto, always, never)",
				Value: EnumValues("auto", "always", "never"),
			},
			&cli.GenericFlag{
				Name:  "log-level",
				Usage: "log level",
				Value: EnumValues("info", "debug", "warn", "err", "fatal"),
			},
		},
		Action: runApp,
	}

	return app.Run(args)
}
