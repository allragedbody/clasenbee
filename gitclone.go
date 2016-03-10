package main

import (
	"fmt"
	git "github.com/libgit2/git2go"
	"io/ioutil"
	"runtime"
	// "log"
)

type Repo struct {
	GitRepo *git.Repository
}

// type Git struct{
//     GitConfig *git
// }

//     GitConfig() (bool, error)

type GitService interface {
	CreateRepo() (bool, error)
	AddToRepo() (bool, error)
	CommitToRepo() (bool, error)
	pullFromRemote() (bool, error)
	pushToRemote() (bool, error)
}

func checkFatal(err error) {
	if err == nil {
		return
	}

	// The failure happens at wherever we were called, not here
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Unable to get caller\n")
	}
	fmt.Printf("Fail at %v:%v; %v", file, line, err)
}

func createRepo() *git.Repository {
	// figure out where we can create the test repo
	// path, err := ioutil.TempDir(".", "git2go")
	// checkFatal(err)
	path := "./git2go"
	repo, err := git.InitRepository(path, false)
	checkFatal(err)

	tmpfile := "README"
	err = ioutil.WriteFile(path+"/"+tmpfile, []byte("foo\n"), 0644)

	checkFatal(err)

	return repo
}

func main() {
	createRepo()
}
