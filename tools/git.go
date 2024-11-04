package tools

// third-party libraries
import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

// Only visible to the local file
const defaultRemoteId string = "origin"

type GitHandler struct {
	// * private * //
	sshKeys *ssh.PublicKeys

	// * public * //

}

func InitializeGitHandler(sshUser string, sshFilepath string, sshPassword string) GitHandler {
	pubKeys, err := ssh.NewPublicKeysFromFile(sshUser, sshFilepath, sshPassword)
	if err != nil {
		log.Panic(err)
	}

	// load the access control system
	return GitHandler{
		sshKeys: pubKeys,
	}
}

func (handler GitHandler) FetchGitRepoRemoteReferences(repoPath string) ([]*plumbing.Reference, error) {
	// load the git repo at path
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// get the default remote address for this git repo
	repoRemote, err := repo.Remote(defaultRemoteId)
	if err == git.ErrRemoteNotFound {
		log.Printf("Git repo (%s) does not have said remote", repoPath)
		return nil, err
	}

	// We can then use every Remote functions to retrieve wanted information
	refs, err := repoRemote.List(&git.ListOptions{
		// Returns all references, including peeled references.
		PeelingOption: git.AppendPeeled,
		// Auth credentials
		Auth: handler.sshKeys,
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// Filters the references list and only keeps tags
	var tags []string
	var brancheRefs []*plumbing.Reference
	for _, ref := range refs {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		} else if ref.Name().IsBranch() {
			brancheRefs = append(brancheRefs, ref)
		}
	}

	if len(tags) == 0 {
		log.Println("No tags!")
	} else {
		log.Printf("Tags found: %v", tags)
	}

	log.Printf("Refs found: %v", brancheRefs)
	return brancheRefs, nil
}

func (handler GitHandler) FetchOrigin(repo *git.Repository, refSpecStr string) error {
	remote, err := repo.Remote(defaultRemoteId)
	if err != nil {
		log.Print(err)
		return err
	}

	var refSpecs []config.RefSpec
	if refSpecStr != "" {
		refSpecs = []config.RefSpec{config.RefSpec(refSpecStr)}
	}

	if err = remote.Fetch(&git.FetchOptions{
		RefSpecs: refSpecs,
		// Auth credentials
		Auth: handler.sshKeys,
	}); err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Print("refs already up to date")
		} else {
			return fmt.Errorf("fetch origin failed: %v", err)
		}
	}

	return nil
}

func (handler GitHandler) CheckoutGitRepoBranch(repoPath string, branchRef *plumbing.Reference) error {
	// load the git repo at path
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Print(err)
		return err
	}

	tree, err := repo.Worktree()
	if err != nil {
		log.Print(err)
		return err
	}

	// reset all uncommited files
	err = tree.Reset(&git.ResetOptions{
		Mode: git.HardReset,
	})
	if err != nil {
		log.Print(err)
		return err
	}

	err = tree.Checkout(&git.CheckoutOptions{
		Branch: branchRef.Name(),
		Force:  true,
	})

	// if there is an error checking out locally, try to fetch the branch from remote (Origin)
	if err != nil {
		mirrorRemoteBranchRefSpec := fmt.Sprintf("refs/heads/%s:refs/heads/%s", branchRef.Name().Short(), branchRef.Name().Short())
		err = handler.FetchOrigin(repo, mirrorRemoteBranchRefSpec)
		if err != nil {
			return err
		}

		err = tree.Checkout(&git.CheckoutOptions{
			Branch: branchRef.Name(),
			Force:  true,
		})
	}

	return err
}

func (handler GitHandler) CheckoutGitRepoCommitHash(repoPath string, commitHash plumbing.Hash) error {
	// load the git repo at path
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Print(err)
		return err
	}

	tree, err := repo.Worktree()
	if err != nil {
		log.Print(err)
		return err
	}

	err = tree.Checkout(&git.CheckoutOptions{
		Hash:  commitHash,
		Force: true,
	})

	return err
}
