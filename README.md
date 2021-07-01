# gotop
A cpu usage viewer made using go for terminal

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

# Installation
```
git clone https://github.com/krishpranav/gotop
cd gotop
go build gotop
./gotop
```

# Usage
```
gotop is a system and resource monitor written in golang.

While using a TUI based command, press ? to get information about key bindings (if any) for that command.

Usage:
  gotop [flags]
  gotop [command]

Available Commands:
  about       about is a command that gives information about the project in a cute way
  completion  Generate completion script
  container   container command is used to get information related to docker containers
  export      Used to export profiled data.
  help        Help about any command
  proc        proc command is used to get per-process information

Flags:
      --config string   config file (default is $HOME/.gotop.yaml)
  -c, --cpuinfo         Info about the CPU Load over all CPUs
  -h, --help            help for gotop
  -r, --refresh uint    Overall stats UI refreshes rate in milliseconds greater than 1000 (default 1000)

Use "gotop [command] --help" for more information about a command.

```