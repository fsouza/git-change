#git-change

git-change is a git extension for interaction with `rietveld
<https://code.google.com/p/rietveld>`_ code revie tool.

It is inspired by [git-cl](https://github.com/martine/git-cl) and the mercurial
plugin provided by Google available at the [Go programming language repository]
(http://code.google.com/p/go/source/browse/lib/codereview/codereview.py).

##Install

You can download pre-compiled binaries and add it to your `PATH`, or run `go
get` (if you have [Go](http://golang.org/doc/install) installed):

    % go get github.com/fsouza/git-change

##Configure

To configure `git-change`, run the config command:

    % git config git-change.server <server> (default is codereview.appspot.com)
    % git config git-change.default-cc <default-cc-email> (default is none)

##Usage

###Create a new CL

Make a new branch, commit your changes and run:

    % git change master

It will generate the new CL within Rietveld.

###Update a CL

To update your CL, change make other commits and just run `git change` again.
It will detect that there is already a CL associated with your current branch
and upload a new patch-set.

More instructions soon :)
