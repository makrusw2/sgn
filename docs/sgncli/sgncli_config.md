## sgncli config

Create or query an application CLI configuration file

### Synopsis

Create or query an application CLI configuration file

```
sgncli config <key> [value] [flags]
```

### Options

```
      --get    print configuration value or its default if unset
  -h, --help   help for config
```

### Options inherited from parent commands

```
      --config string     config path (default "./config.json")
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "$HOME/.sgncli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors
```

### SEE ALSO

* [sgncli](sgncli.md)	 - SGN node command line interface

###### Auto generated by spf13/cobra
