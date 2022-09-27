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

//Checkout
func (g *Git) Checkout() *Git {
	g._type = clone
	g.gQuery = []string{}
	g.gQuery = append(g.gQuery, []string{"-C", g.rootFolder, "checkout", g.repository}...)
	return g
}

//  tag extension
//if called the branch will be void
func (g *Git) Tag(tagname string) *extension {
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
func (g *Git) Branch(branchname string) *extension {
	x := new(extension)

	remoteTagOrBranch(&g.gQuery) // to empty extension attached
	if len(branchname) > 0 {
		g.gQuery = append(g.gQuery, []string{"--branch", branchname}...)
	}
	x.gQuery = g.gQuery
	return x
}

//execute command
func (x *extension) Execute() (string, error) {
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
