# `jrnl`

CLI for managing my personal journal implemented in go.

## Installation

```
go get github.com/matt-hoiland/jrnl
go install github.com/matt-hoiland/jrnl
```

## Global flags

* `--version`
* `--help`

## Primary Commands

* `init <name>`
* `entry`
  * `ls [-t <tag> ...]`
  * `add <name> [-t <tag> ...]`
* `tag`
  * `ls [-e <entry> ...]`
  * `add <tag> [-e <entry> ...]`
* `update`
