package routes

import ldap "github.com/vjeantet/ldapserver"

// AllRoutes returns an ldap.RouteMux with all our routes
func AllRoutes() *ldap.RouteMux {
	routes := ldap.NewRouteMux()
	routes.Bind(handleBind)
	routes.Search(handleSearch)
	return routes
}
