package counterfeitertests_test

import (
	"testing"

	fakes "github.com/tjarratt/counterfeitertests/counterfeitertestsfakes"
	git "gopkg.in/src-d/go-git.v4"
)

func TestItDoesStuff(t *testing.T) {
	options := git.CloneOptions{}

	myFake := &fakes.FakeJustAnExample{}
	myFake.DoesStuffWithGit(options)
}
