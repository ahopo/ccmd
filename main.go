package ccmd

import "github.com/ahopo/ccmd/git"

func InitGit() git.Config {
	return *new(git.Config)
}
