# Octopus

<p align="center"><img src="https://user-images.githubusercontent.com/14275767/34849223-d3f7c896-f708-11e7-8df1-33e2cb60a0ba.png" /></p>

Manages git projects locally.

## Instalation

Download the [latest](https://github.com/B0go/octopus/releases/latest) version compatible with your OS

Unpack the contents anywhere and execute the install.sh script:

```sh
~/blah
λ tar -xvzf octopus_0.1.3_macOS_64-bit.tar.gz
x install.sh
x octopus

~/blah
λ ./install.sh

Installing octopus...


Octopus successfully installed!


~/blah
λ octopus
NAME:
   octopus - Manages Github projects locally

USAGE:
   octopus [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     install, i  octopus install julius, octopus install --team=fortknox, octopus install --all
     get, g      octopus get SUBCOMMAND
     run, r      octopus run julius
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

## Usage

Octopus manages the projects configured at the config.yaml file located at `~/.octopus/config.yaml`. You must create one.

### Functionality

#### Install

- Specific project

```sh
octopus install myproject
```

- All team's project

```sh
octopus install --team=myteam
```

- All configured projects

```sh
octopus install --all
```

#### Get

- Configured projects

```sh
octopus get projects
```

- Configured teams

```sh
octopus get teams
```

### Run

- Run a project

```sh
octopus run myproject
```

- Run a project on a specific branch

```sh
octopus run myproject --branch=mybranch
```
