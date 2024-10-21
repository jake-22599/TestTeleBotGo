package tools

// third-party libraries
import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

// Only visible to the local file
const defaultRemoteId string = "origin"

func GetGitRepoReferences(sshUser string, sshFilepath string, sshPassword string, repoPath string) ([]string, error) {
	pubKeys, err := ssh.NewPublicKeysFromFile(sshUser, sshFilepath, sshPassword)
	if err != nil {
		log.Printf("default auth builder: %v", err)
		return nil, err
	}

	// load the git repo at path
	rem, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// get the default remote address for this git repo
	repoRemote, err := rem.Remote(defaultRemoteId)
	if err == git.ErrRemoteNotFound {
		log.Printf("Git repo (%s) does not have said remote", repoPath)
		return nil, err
	}

	// We can then use every Remote functions to retrieve wanted information
	refs, err := repoRemote.List(&git.ListOptions{
		// Returns all references, including peeled references.
		PeelingOption: git.AppendPeeled,
		// Auth credentials
		Auth: pubKeys,
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// Filters the references list and only keeps tags
	var tags []string
	var refStrs []string
	for _, ref := range refs {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		}
		refStrs = append(tags, ref.Name().Short())
	}

	if len(tags) == 0 {
		log.Println("No tags!")
	} else {
		log.Printf("Tags found: %v", tags)
	}

	log.Printf("Refs found: %v", refStrs)
	return refStrs, nil
}
