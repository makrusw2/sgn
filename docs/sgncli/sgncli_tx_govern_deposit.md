## sgncli tx govern deposit

Deposit tokens for an active proposal

### Synopsis

Submit a deposit for an active proposal. You can
find the proposal-id by running "<appcli> query gov proposals".

Example:
$ <appcli> tx gov deposit 1 10

```
sgncli tx govern deposit [proposal-id] [deposit] [flags]
```

### Options

```
      --dry-run      ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --gas string   gas limit to set per-transaction; set to "auto" to calculate required gas automatically (default 200000) (default "200000")
  -h, --help         help for deposit
      --indent       Add indent to JSON response
      --trust-node   Trust connected full node (don't verify proofs for responses) (default true)
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

* [sgncli tx govern](sgncli_tx_govern.md)	 - Governance transactions subcommands

###### Auto generated by spf13/cobra
