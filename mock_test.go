package main

import (
	"launchpad.net/goetveld/rietveld"
	"strconv"
)

type fakerietveld struct {
	actions map[string]*rietveld.Issue
}

func (r *fakerietveld) addaction(key string, issue *rietveld.Issue) {
	if r.actions == nil {
		r.actions = make(map[string]*rietveld.Issue)
	}
	r.actions[key] = issue
}

func (r *fakerietveld) AddComment(issue *rietveld.Issue, comment *rietveld.Comment) error {
	r.addaction("comment", issue)
	return nil
}

func (r *fakerietveld) Issue(issue int) (*rietveld.Issue, error) {
	r.addaction("getissue", &rietveld.Issue{Id: issue})
	return nil, nil
}

func (r *fakerietveld) IssueURL(issue *rietveld.Issue) string {
	r.addaction("geturl", issue)
	return "http://codereview.appspot.com/" + strconv.Itoa(issue.Id)
}

func (r *fakerietveld) SendDelta(issue *rietveld.Issue, delta rietveld.Delta, sendMail bool) error {
	r.addaction("upload", issue)
	return nil
}
