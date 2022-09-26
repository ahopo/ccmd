package git

import (
	"fmt"
	"os/exec"
)

//base
type Git struct {
	_type      int
	repository string
	gQuery     []string
	rootFolder string
}
type extension struct {
	gQuery []string
}

const (
	clone    int = 1
	CHECKOUT int = 2
)

//set repo
func (g *Git) SetRepo(repo string) {
	g.repository = repo
}

//set root folder
func (g *Git) SetRootFolder(rootfolder string) {
	g.rootFolder = rootfolder
}

//Clone
func (g *Git) Clone() *Git {
	g._type = clone
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, []string{"-C", g.rootFolder, "clone", g.repository}...)
	return g
}

//  tag extension
//if called the branch will be void
func (g *Git) Tag(tagname string) *extension {
	x := extension{}
	if len(tagname) == 0 {
		return &x
	}
	g.gQuery[3] = "" // to empty extension attached
	switch g._type {
	case clone:
		g.gQuery = append(g.gQuery, []string{"--branch", tagname}...)
	default:
		g.gQuery = append(g.gQuery, fmt.Sprintf("tags/%s", tagname))
	}

	x.gQuery = g.gQuery
	return &x
}

//  branch extension
//if called the tag will be void
func (g *Git) Branch(branchname string) *extension {
	x := extension{}
	if len(branchname) == 0 {
		return &x
	}
	g.gQuery[3] = "" // to empty extension attached
	g.gQuery = append(g.gQuery, []string{"--branch", branchname}...)
	x.gQuery = g.gQuery
	return &x
}

//execute command
func (x *extension) Execute() (string, error) {
	cmd := exec.Command("git", x.gQuery...)
	o, err := cmd.CombinedOutput()
	return string(o), err
}
