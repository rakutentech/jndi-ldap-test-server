package routes

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vjeantet/ldapserver"
)

func handleBind(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	r := m.GetBindRequest()
	log.Info().
		Str("component", "ldap").
		Str("event", "request").
		Str("client_ip", m.Client.Addr().String()).
		Dict("request", zerolog.Dict().
			Str("type", "bind").
			Str("auth_type", r.AuthenticationChoice()).
			Str("user", string(r.Name())),
		).
		Msg("Incoming LDAP Bind Request")
	res := ldapserver.NewBindResponse(ldapserver.LDAPResultSuccess)
	w.Write(res)
}
