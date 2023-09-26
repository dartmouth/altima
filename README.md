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
altima repo add altima-core https://git.dartmouth.edu/research-itc-public/altima-core-modules/-/raw/main/
altima repo list
altima update
altima repo remove altima-core

# Manage packages
altima search
altima install cyberark
altima list
```

## Build

```sh
go build
```