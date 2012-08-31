package main

import (
	"github.com/fsouza/gogit/git"
	"launchpad.net/goetveld/rietveld"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

func createRepository() (*git.Repository, string) {
	path, err := filepath.Abs(".repo")
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(path, 0700)
	if err != nil {
		panic(err)
	}
	repo, err := git.NewRepository(path, false)
	if err != nil {
		panic(err)
	}
	return repo, path
}

func TestCreateCL(t *testing.T) {
	repo, path := createRepository()
	defer repo.Free()
	defer os.RemoveAll(path)
	config, err := repo.Config()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer config.Free()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	os.Chdir(path)
	defer os.Chdir(cwd)
	err = exec.Command("touch", "readme").Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = exec.Command("git", "add", "readme").Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = exec.Command("git", "commit", "-m", "added readme").Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = exec.Command("git", "checkout", "-b", "patch").Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = exec.Command("touch", "readme2").Run()
	if err != nil {
		t.Error(nil)
		t.FailNow()
	}
	err = exec.Command("git", "add", "readme2").Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = exec.Command("git", "commit", "-m", "added readme2").Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fake := fakerietveld{}
	ch := ChangeManager{
		crclient: &fake,
		repo:     repo,
	}
	issue, err := ch.New("master")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	expected := map[string]*rietveld.Issue{
		"upload": issue,
	}
	if !reflect.DeepEqual(fake.actions, expected) {
		t.Errorf("Should upload issue to rietveld server, but did not. Got %+v, want %+v.", fake.actions, expected)
	}
}
