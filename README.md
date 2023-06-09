# altima

Goals
- Performance
- Tool for SysAdmins
- Easy to install and configure
- Consistent interface
- Flexible module repos
- Different arch releases

Location
- GitLab
- GitHub
- https://github.com/elijahgagne/altima
- https://github.com/dartmouth/altima

Name
- altima
- altima.sh

How to work
- Git repo issues

Core
- Build in Go

Modules
- Build in anything, mainly Shell

Roles
- People writing Go for the core
- People writing Shell for the modules
- Documentation
- Packaging and installation
- Testers

Examples

```sh
# Install altima itself
brew install altima

# Initialization
altima init

# Manage repos
altima repo add https://dartgo.org/altima-rc
altima repo remove https://dartgo.org/altima-rc
altima repo list

# Manage packages
altima search
altima install cyberark
altima list
```
