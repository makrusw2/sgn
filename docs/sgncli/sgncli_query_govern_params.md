## sgncli query govern params

Query the parameters of the governance process

### Synopsis

Query the all the parameters for the governance process.

Example:
$ <appcli> query gov params

```
sgncli query govern params [flags]
```

### Options

```
      --height int   Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help         help for params
      --indent       Add indent to JSON response
      --trust-node   Trust connected full node (don't verify proofs for responses)
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

* [sgncli query govern](sgncli_query_govern.md)	 - Querying commands for the governance module

###### Auto generated by spf13/cobra
