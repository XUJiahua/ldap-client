## ldap-client

A simple ldap client to create account and reset account password

## Usage

```bash
$ ./ldap-client
a simple ldap client to create account and reset account password

Usage:
  ldap-client [command]

Available Commands:
  completion    Generate the autocompletion script for the specified shell
  createAccount create account in ldap
  help          Help about any command
  resetPassword reset password

Flags:
  -h, --help                   help for ldap-client
      --ldap-password string   ldap password (default "123456")
      --ldap-url string        ldap url (default "ldap://localhost:389")
      --ldap-user string       ldap user (default "cn=admin,dc=example,dc=com")
      --people-dn string       ldap people dn (default "ou=people,dc=example,dc=com")

Use "ldap-client [command] --help" for more information about a command.

```

If you don’t want to pollute your command line, you can setting flags by using environment variables.

```bash
$ export LDAP_URL=ldap://localhost:389
$ export LDAP_USER=cn=admin,dc=example,dc=com
$ export LDAP_PASSWORD=123456
$ export PEOPLE_DN=ou=people,dc=example,dc=com
```



```bash
$ ./ldap-client createAccount jack.xu@example.com
login user: jack.xu
login password: woj2A8pV

$ ./ldap-client resetPassword john.xu@example.com 
login user: john.xu
login password: QyhjrwB7
```

