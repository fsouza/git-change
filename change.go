package main

import (
	"flag"
	"github.com/fsouza/gogit/git"
	"launchpad.net/goetveld/rietveld"
)

type Rietveld interface {
	AddComment(*rietveld.Issue, *rietveld.Comment) error
	Issue(int) (*rietveld.Issue, error)
	IssueURL(*rietveld.Issue) string
	SendDelta(*rietveld.Issue, rietveld.Delta, bool) error
}

type ChangeManager struct {
	repo      *git.Repository
	defaultCc string
	server    string
	crclient  Rietveld
}

func (c *ChangeManager) cr() Rietveld {
	if c.crclient == nil {
		c.crclient = rietveld.CodeReview
	}
	return c.crclient
}

func (c *ChangeManager) New(rev string) (*rietveld.Issue, error) {
	return nil, nil
}

type Command func([]string) error

var commands = map[string]Command{}

func diff(from string) {

}

func init() {
	flag.Parse()
}

func main() {
	var name string
	args := flag.Args()
	if len(args) < 1 {
		name = "default"
	} else {
		name = args[0]
		args = args[1:]
	}
	c := commands[name]
	c(args)
}
