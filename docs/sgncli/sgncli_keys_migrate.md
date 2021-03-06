## sgncli keys migrate

Migrate keys from the legacy (db-based) Keybase

### Synopsis

Migrate key information from the legacy (db-based) Keybase to the new keyring-based Keybase.
For each key material entry, the command will prompt if the key should be skipped or not. If the key
is not to be skipped, the passphrase must be entered. The key will only be migrated if the passphrase
is correct. Otherwise, the command will exit and migration must be repeated.

It is recommended to run in 'dry-run' mode first to verify all key migration material.


```
sgncli keys migrate [flags]
```

### Options

```
      --dry-run   Run migration without actually persisting any changes to the new Keybase
  -h, --help      help for migrate
```

### Options inherited from parent commands

```
      --config string            config path (default "./config.json")
  -e, --encoding string          Binary encoding (hex|b64|btc) (default "hex")
      --home string              directory for config and data (default "$HOME/.sgncli")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
  -o, --output string            Output format (text|json) (default "text")
      --trace                    print out full stack trace on errors
```

### SEE ALSO

* [sgncli keys](sgncli_keys.md)	 - Add or view local private keys

###### Auto generated by spf13/cobra
