# altima

## Goals

- Performance
- Tool for SysAdmins
- Easy to install and configure
- Consistent interface
- Flexible module repos
- Different arch releases
- Modules built in anything, mainly Shell

## Example commands

```sh
# Install altima itself
brew install altima

# Initialization
altima init

# Manage repos
altima repo add altima-rc https://dartgo.org/altima-rc
altima repo remove altima-rc
altima repo list
altima repo update

# Manage packages
altima search
altima install cyberark
altima list
```

## Build

```sh
go build
```