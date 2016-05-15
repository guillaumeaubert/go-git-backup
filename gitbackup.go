package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/guillaumeaubert/go-git-backup/gitbackup"
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
