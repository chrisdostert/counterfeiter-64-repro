package counterfeitertests

import . "gopkg.in/src-d/go-git.v4"

type JustAnExample interface {
	DoesStuffWithGit(options CloneOptions)
}
