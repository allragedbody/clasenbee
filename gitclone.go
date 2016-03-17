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

func CreateRepo() *git.Repository {
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

// func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
// 	ret, cred := git.NewCredSshKey("git", "/var/www/testgo/git/id_rsa.pub", "/var/www/testgo/git/id_rsa", "")
// 	return git.ErrorCode(ret), &cred
// }

// func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
// 	return 0
// }

// func CloneRepo() *git.Repository {
// 	// figure out where we can create the test repo
// 	// path, err := ioutil.TempDir(".", "git2go")
// 	// checkFatal(err)
// 	path := "./git2go"
// 	repo, err := git.InitRepository(path, false)
// 	checkFatal(err)

// 	tmpfile := "README"
// 	err = ioutil.WriteFile(path+"/"+tmpfile, []byte("foo\n"), 0644)

// 	checkFatal(err)

// 	return repo
// }

func main() {
	// CreateRepo()
	repo, err := git.Clone("root@192.168.150.73:ats.git", "ats", &git.CloneOptions{})
	if err != nil {
		checkFatal(err)
	}
	fmt.Println(repo)

}
