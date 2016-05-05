// "gitbackup" holds the functions that do the actual backing up of git
// repositories.
package gitbackup

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// BackupTarget backs up an entity that holds one or more git repositories and
// has an interface to retrieve that list of repositories.
// Examples of entities include:
//   - A GitHub user.
//   - A BitBucket user.
//   - A GitHub organization.
func BackupTarget(target map[string]string, backupDirectory string) error {
	log.Printf(`Backing up target "%s"`, target["name"])

	// TODO: replace with a factory pattern?
	switch target["source"] {
	case "github":
		return backupGitHub(target, backupDirectory)
	case "bitbucket":
		return backupBitBucket(target, backupDirectory)
	default:
		return fmt.Errorf(`"%s" is not a recognized source type`, target["source"])
	}
}

// backupGitHub finds all the repositories under a given user or organization
// before backing up each one.
func backupGitHub(target map[string]string, backupDirectory string) error {
	// Create URL to request list of repos.
	var requestURL string = fmt.Sprintf(
		"https://api.github.com/%s/%s/repos?access_token=%s&per_page=200",
		target["type"],
		target["entity"],
		target["token"],
	)

	// Retrieve list of repositories.
	response, err := http.Get(requestURL)
	if err != nil {
		return fmt.Errorf("Failed to connect with the source to retrieve the list of repositories: %s", err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Failed to retrieve the list of repositories: %s", err)
	}

	// Parse JSON response.
	var dat []map[string]interface{}
	if err := json.Unmarshal(contents, &dat); err != nil {
		return fmt.Errorf("Failed to parse JSON: %s", err)
	}

	// Back up each repository.
	for _, repo := range dat {
		repoName, _ := repo["name"].(string)
		cloneURL, _ := repo["clone_url"].(string)
		cloneURL = strings.Replace(
			cloneURL,
			"https://",
			fmt.Sprintf("https://%s:%s@", target["entity"], target["token"]),
			1,
		)
		backupRepository(
			target["name"],
			repoName,
			cloneURL,
			backupDirectory,
		)
	}

	// No errors.
	return nil
}

func backupBitBucket(target map[string]string, backupDirectory string) error {
	// TODO: implement.

	// No errors.
	return nil
}

// backupRepository takes a remote git repository and backs it up locally.
// Note that this makes a mirror repository - in other words, the backup only
// contains the content of a normal .git repository but no working directory,
// which saves space. You can always get a normal repository from the backup by
// doing a normal git clone of the backup itself.
func backupRepository(targetName string, repoName string, cloneURL string, backupDirectory string) {
	var cloneDirectory string = filepath.Join(backupDirectory, targetName, repoName)
	fmt.Println(fmt.Sprintf("#> %s", repoName))
	log.Printf(`Backing up repo "%s"`, repoName)

	if _, err := os.Stat(cloneDirectory); os.IsNotExist(err) {
		// The repo doesn't exist locally, clone it.
		log.Printf("Cloning %s to %s", cloneURL, cloneDirectory)

		cmd := exec.Command("git", "clone", "--mirror", cloneURL, cloneDirectory)
		cmdOut, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error cloning the repository:", err)
		} else {
			fmt.Println("Cloned repository.")
			if len(cmdOut) > 0 {
				fmt.Printf(string(cmdOut))
			}
		}
	} else {
		// The repo already exists, pull updates.
		log.Printf("Pulling git repo in %s", cloneDirectory)

		cmd := exec.Command("git", "fetch", "-p", cloneURL)
		cmd.Dir = cloneDirectory
		cmdOut, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error pulling in the repository:", err)
		} else {
			// Display pulled information.
			fmt.Println("Pulled latest updates in the repository.")
			if len(cmdOut) > 0 {
				fmt.Printf(string(cmdOut))
			}
		}
	}
}
