# Stellar Vanity Address Generator - `svag`

CLI tool to generate a vanity Stellar Public Key address containing the word you are looking.

## What is a Stellar vanity addresses?

If you know about Stellar and Lumens you know there is a public key/address. For example Lumenaut has the following public key/address that ends in `NAUT`

```
GCCD6AJOYZCUAQLX32ZJF2MKFFAUJ53PVCFQI3RHWKL3V47QYE2BNAUT
```

Using this CLI tool you can generate a vanity URL that begins with or ends with a specific word you have in mind.

## How to generate vanity addresses?

First install the CLI tool using `go`

1. `go get github.com/ankurp/fast-stellar-vanity-address-generator/svag`
1. `$GOPATH/bin/svag codr`

### Search for multiple words

You can also specify multiple words in the command line and it will find the first address that contains any of the words you are looking for.

For example: `svag cat dog`

## Why use this CLI tool?

1. You can run this offline after installing the CLI command to ensure it does not sending any public or secret key.
1. You can find addresses that end with or begin with a specific word
1. Supports searching of multiple words instead of one at a time
1. It maximizes the use of all cores on your machine to find a word as fast as possible

# Donate

If you like this project, consider making a donation via Stellar to support its development:
```
GAH5N24RWKLBBT7MLKSHRNAIKV45GUCBZZQ2PF4L6NLJXRWLY3DONATE
```
