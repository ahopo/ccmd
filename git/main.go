package git

import (
	"fmt"
	"os/exec"
)

//base
type Config struct {
	_type       int
	repository  string
	gQuery      []string
	rootFolder  string
	superBranch string
}
type extension struct {
	gQuery []string
}

const (
	clone    int = 1
	CHECKOUT int = 2
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
func (g *Config) GotoSuperBranch() {
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, "-C", g.rootFolder, "switch", g.superBranch)
}

//Fetch
func (g *Config) Fetch() {
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, "-C", g.rootFolder, "fetch")
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
	g._type = clone
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, []string{"-C", g.rootFolder, "checkout", g.repository}...)
	return g
}

// List all tags
func (g *Config) GetAllTags() *Config {
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, "-C", g.rootFolder, "tag")
	return g
}

// List all Branchs
func (g *Config) GetAllBranchs() *Config {
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, "-C", g.rootFolder, "branch", "-r")
	return g
}

//  tag extension
//if called the branch will be void
func (g *Config) Tag(tagname string) *extension {
	x := new(extension)

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
func (g *Config) Branch(branchname string) *extension {
	x := new(extension)

	remoteTagOrBranch(&g.gQuery) // to empty extension attached
	if len(branchname) > 0 {
		g.gQuery = append(g.gQuery, []string{"--branch", branchname}...)
	}
	x.gQuery = g.gQuery
	return x
}

//execute command
func (x *extension) Exec() (string, error) {
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
