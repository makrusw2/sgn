## sgncli tx validator edit-candidate-description

Edit candidate description for the eth address

### Synopsis

Edit candidate description for the eth address

```
sgncli tx validator edit-candidate-description [eth-addr] [flags]
```

### Options

```
      --details string            The candidate's (optional) details (default "[do-not-modify]")
      --dry-run                   ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --gas string                gas limit to set per-transaction; set to "auto" to calculate required gas automatically (default 200000) (default "200000")
  -h, --help                      help for edit-candidate-description
      --identity string           The (optional) identity signature (ex. UPort or Keybase) (default "[do-not-modify]")
      --indent                    Add indent to JSON response
      --moniker string            The candidate's name (default "[do-not-modify]")
      --security-contact string   The candidate's (optional) security contact email (default "[do-not-modify]")
      --trust-node                Trust connected full node (don't verify proofs for responses) (default true)
      --website string            The candidate's (optional) website (default "[do-not-modify]")
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

* [sgncli tx validator](sgncli_tx_validator.md)	 - Validator transaction subcommands

###### Auto generated by spf13/cobra
