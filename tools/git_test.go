package tools

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func TestGitStuff(t *testing.T) {
	path := ""
	sshFilepath := ""

	// default values
	// TODO: to be moved to a config in future
	sshUser := "git"
	sshPassword := ""

	// load git handler
	gitHandler := InitializeGitHandler(sshUser, sshFilepath, sshPassword)

	// get all the references for the git repo
	refs, err := gitHandler.FetchGitRepoRemoteReferences(path)
	if err != nil {
		t.Fatalf("Error occured while trying to fetch remote references; %s", err.Error())
	} else {
		refNames := ""
		for _, ref := range refs {
			refNames += fmt.Sprintf("\n- %s", ref.Name().Short())
		}
		fmt.Println(refNames)
	}

	///////////////////////////////////////////////////
	// working to a certain degree,
	// issue is related to branch may not be found locally
	/// solution: to fetch and track the remote branch locally
	///////////////////////////////////////////////////

	// TODO:
	targetBranchName := "master"
	// choose a random ref
	branchRef := refs[rand.IntN(len(refs))]
	for _, ref := range refs {
		if ref.Name().Short() == targetBranchName {
			branchRef = ref
			break
		}
	}

	err = gitHandler.CheckoutGitRepoBranch(path, branchRef)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Error occured while trying to checkout git branch; %s", err.Error())
	} else {

		fmt.Println("")
	}

	///////////////////////////////////////////////////
	// branchRef := refs[rand.IntN(len(refs))]
	// err = gitHandler.CheckoutGitRepoCommitHash(path, branchRef.Hash())
	// if err != nil {
	// 	t.Fatalf("Error occured while trying to checkout commit; %s", err.Error())
	// } else {

	// 	fmt.Println("")
	// }
}
