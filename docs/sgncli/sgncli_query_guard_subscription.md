## sgncli query guard subscription

query subscription info associated with the eth address

### Synopsis

query subscription info associated with the eth address

```
sgncli query guard subscription [ethAddress] [flags]
```

### Options

```
      --height int   Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help         help for subscription
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

* [sgncli query guard](sgncli_query_guard.md)	 - Querying commands for the guard module

###### Auto generated by spf13/cobra
