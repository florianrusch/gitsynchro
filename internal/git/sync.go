package git

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/florianrusch/gitsynchro/internal"
	"github.com/florianrusch/gitsynchro/internal/config"
	"github.com/florianrusch/gitsynchro/internal/log"
	g "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func HandleRepo(repo config.Repository) error {
	log.Infof("Checking repo: %s (%s)", repo.Name, repo.Path)

	// We instantiate a new repository targeting the given path (the .git folder)
	gitRepo, err := g.PlainOpen(repo.Path)
	if err != nil {
		return err
	}

	err = checkPreconditions(gitRepo, repo.DefaultBranch)
	if err != nil {
		return err
	}

	for _, destination := range repo.Destinations {
		log.Infof("Pushing to remote: %s", destination.RemoteName)
		err = gitRepo.Push(&g.PushOptions{
			RemoteName: destination.RemoteName,
			Progress:   os.Stdout,
		})

		if err != nil && !errors.Is(err, g.NoErrAlreadyUpToDate) {
			return err
		} else if errors.Is(err, g.NoErrAlreadyUpToDate) {
			log.Debugf("Remote is already up to date")
		}
	}

	log.Infof("Push finished")

	return nil
}

func checkPreconditions(gitRepo *g.Repository, defaultBranch string) error {
	err := printCurrentBranch(gitRepo)
	if err != nil {
		return err
	}

	head, _ := gitRepo.Head()
	isDefault := isBranchADefault(head, defaultBranch)

	if isDefault {
		log.Debugf("Current branch is default branch")
	} else {
		return fmt.Errorf("current branch is NOT default branch")
	}

	err = printAllBranches(gitRepo)
	if err != nil {
		return err
	}

	err = printAllRemotes(gitRepo)
	if err != nil {
		return err
	}

	isStateClean, err := isStateClean(gitRepo)
	if err != nil {
		return err
	}

	if isStateClean {
		log.Debugf("State is clean")
	} else {
		return fmt.Errorf("state is NOT clean")
	}

	return nil
}

func isStateClean(r *g.Repository) (bool, error) {
	worktree, err := r.Worktree()
	if err != nil {
		return false, err
	}

	status, err := worktree.Status()
	if err != nil {
		return false, err
	}

	return status.IsClean(), nil
}

func printCurrentBranch(r *g.Repository) error {
	head, err := r.Head()
	if err != nil {
		return err
	}

	if !head.Name().IsBranch() {
		return fmt.Errorf("current branch is not a branch")
	}

	log.Infof("Current Branch: %s", head.Name().Short())

	return nil
}

func isBranchADefault(reference *plumbing.Reference, configuredDefaultBranch string) bool {
	defaults := []string{"main", "master"}

	if reference.Name().Short() == configuredDefaultBranch {
		return true
	}

	for _, v := range defaults {
		if v == reference.Name().Short() {
			return true
		}
	}

	return false
}

func printAllBranches(r *g.Repository) error {
	branches, err := r.Branches()
	if err != nil {
		return err
	}

	var outputString string

	_ = branches.ForEach(func(ref *plumbing.Reference) error {
		outputString += fmt.Sprintf(", %s", ref.Name().Short())

		return nil
	})

	outputString, _ = strings.CutPrefix(outputString, ", ")
	log.Infof("Branches: %s", outputString)

	return nil
}

func printAllRemotes(r *g.Repository) error {
	remotes, err := r.Remotes()
	if err != nil {
		return err
	}

	var outputString = internal.NestedJoin[*g.Remote](
		", ",
		remotes,
		func(remote *g.Remote) string { return remote.Config().Name },
	)

	outputString, _ = strings.CutPrefix(outputString, ", ")
	log.Infof("Remotes: %s", outputString)

	return nil
}
