package git

import (
	"fmt"
	"os/exec"
)

//Config of git
type Config struct {
	_type       int
	repository  string
	gQuery      []string
	rootFolder  string
	superBranch string
}

//for execution of git command
type execute struct {
	gQuery []string
}

const (
	clone    int = 1
	checkout int = 2
)

//set repo
func (g *Config) SetRepo(repo string) {
	g.repository = repo
}

//set root folder
func (g *Config) SetRootFolder(rootfolder string) {
	g.rootFolder = rootfolder
}

//	set super branch
//it could be master or main
func (g *Config) SetSuperBranch(superbranch string) {
	g.superBranch = superbranch
}

//Goto to super branch
func (g *Config) GotoSuperBranch() *execute {
	x := new(execute)
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, "-C", g.rootFolder, "switch", g.superBranch)
	x.gQuery = g.gQuery
	return x
}

//Fetch
func (g *Config) Fetch() *execute {
	x := new(execute)
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, "-C", g.rootFolder, "fetch")
	x.gQuery = g.gQuery
	return x
}

//Clone
func (g *Config) Clone() *Config {
	g._type = clone
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, []string{"-C", g.rootFolder, "clone", g.repository}...)
	return g
}

//Checkout
func (g *Config) Checkout() *Config {
	g._type = checkout
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, []string{"-C", g.rootFolder, "checkout", g.repository}...)
	return g
}

// List all tags
func (g *Config) GetAllTags() *execute {
	x := new(execute)
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, "-C", g.rootFolder, "tag")
	x.gQuery = g.gQuery
	return x
}

// List all Branchs
func (g *Config) GetAllBranchs() *execute {
	x := new(execute)
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, "-C", g.rootFolder, "branch", "-r")
	x.gQuery = g.gQuery
	return x
}

//  tag extension
//if called the branch will be void
func (g *Config) Tag(tagname string) *execute {
	x := new(execute)

	remoteTagOrBranch(&g.gQuery) // to empty extension attached
	if len(tagname) > 0 {
		switch g._type {
		case clone:
			g.gQuery = append(g.gQuery, []string{"--branch", tagname}...)
		default:
			g.gQuery = append(g.gQuery, fmt.Sprintf("tags/%s", tagname))
		}
	}

	x.gQuery = g.gQuery
	return x
}

//  branch extension
//if called the tag will be void
func (g *Config) Branch(branchname string) *execute {
	x := new(execute)

	remoteTagOrBranch(&g.gQuery) // to empty extension attached
	if len(branchname) > 0 && g._type == clone {
		g.gQuery = append(g.gQuery, []string{"--branch", branchname}...)
	}
	x.gQuery = g.gQuery
	return x
}

//execute command
func (x *execute) Exec() (string, error) {
	fmt.Println(x.gQuery)
	cmd := exec.Command("git", x.gQuery...)
	o, err := cmd.CombinedOutput()
	return string(o), err
}

//remove tag and branch
func remoteTagOrBranch(data *[]string) {
	emptystr := ""
	if len(*data) == 5 {
		(*data)[4] = emptystr
	}
}
