package ccmd

import "github.com/ahopo/ccmd/git"

func InitGit() git.Git {
	return *new(git.Git)
}
