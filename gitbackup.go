// go-git-backup offers a convenient way to back up remote GitHub/BitBucket
// users/organizations with a collection of repositories for each.
//
// You first need to define a configuration file, for example ~/gitbackup.yml,
// with the following content:
//
//     backup_directory: /where/your/backups/will/be/stored
//     targets:
//       - name: github-guillaumeaubert
//         source: github
//         type: users
//         entity: aubertg
//         token: mysecrettoken
//       - name: bitbucket-aubertg
//         source: bitbucket
//         type: users
//         entity: aubertg
//         password: mysecretpassword
//
// You can define as many targets as your config file as you would like. Each
// target should have the following information:
//   * name: an internal name, used as the top level directory in your backup
//     directory to group all the repositories belonging to this target.
//   * source: "github" or "bitbucket". Other sources are not yet supported.
//   * type: "users" or "orgs", depending on what type of entity you are
//     backing up.
//   * entity: the name of the entity being backed up, either a username or an
//     organization name.
//   * token: for GitHub, generate a token that gives access to this user or
//     organization.
//   * password: BitBucket doesn't support tokens yet, so you will need to use
//     your normal password.
//
// Usage:
//
//     gitbackup --config=~/gitbackup.yml
//
package main

import (
	"flag"
	"fmt"
	"log"

	gitbackup "github.com/guillaumeaubert/go-git-backup/lib"
)

// Main function.
func main() {
	// Configure logging.
	log.SetPrefix("gitbackup: ")

	// Parse command-line flags.
	configPathPtr := flag.String("config", "", "Path to the configuration file holding hosts and credentials information")
	flag.Parse()

	// Get the config.
	config := gitbackup.GetConfig(*configPathPtr)

	// Back up each target.
	for _, target := range config.Targets {
		err := gitbackup.BackupTarget(target, config.BackupDirectory)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Done.
	fmt.Println("Backups completed.")
}
