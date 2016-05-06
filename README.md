Git backup package in Go
========================

go-git-backup offers a convenient way to back up remote GitHub/BitBucket
users/organizations with a collection of repositories for each.


Code status
-----------

[![Build Status](https://travis-ci.org/guillaumeaubert/go-git-backup.svg?branch=master)](https://travis-ci.org/guillaumeaubert/go-git-backup)
[![Go Report Card](https://goreportcard.com/badge/github.com/guillaumeaubert/go-git-backup)](https://goreportcard.com/report/github.com/guillaumeaubert/go-git-backup)
[![GoDoc](https://godoc.org/github.com/guillaumeaubert/go-git-backup?status.svg)](https://godoc.org/github.com/guillaumeaubert/go-git-backup)


Setup
-----

You first need to define a configuration file, for example ~/gitbackup.yml,
with the following content:

```
  backup_directory: /where/your/backups/will/be/stored
  targets:
    - name: github-guillaumeaubert
      source: github
      type: users
      entity: aubertg
      token: mysecrettoken
    - name: bitbucket-aubertg
      source: bitbucket
      type: users
      entity: aubertg
      password: mysecretpassword
```


Configuration options
---------------------

You can define as many targets as your config file as you would like. Each
target should have the following information:

* `name`: an internal name, used as the top level directory in your backup
  directory to group all the repositories belonging to this target.
* `source`: "github" or "bitbucket". Other sources are not yet supported.
* `type`: "users" or "orgs", depending on what type of entity you are backing
  up.
* `entity`: the name of the entity being backed up, either a username or an
  organization name.
* `token`: for GitHub, generate a token that gives access to this user or
  organization.
* `password`: BitBucket doesn't support tokens yet, so you will need to use
  your normal password.


Usage
-----

```
go run gitbackup.go -config ~/gitbackup.yml

gitbackup -config ~/gitbackup.yml
```


Copyright
---------

Copyright (C) 2016 Guillaume Aubert


License
-------

This software is released under the MIT license. See the LICENSE file for
details.
