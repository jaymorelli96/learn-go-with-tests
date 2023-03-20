package blogpost_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/jaymorelli96/learn-go-with-tests/reading-files"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("failing stub for opening new file")
}

func TestBlogPosts(t *testing.T) {
	t.Run("error while opening file", func(t *testing.T) {
		_, err := blogposts.NewPosts(StubFailingFS{})

		if err == nil {
			t.Error("expected error")
		}
	})

	t.Run("get posts", func(t *testing.T) {
		const (
			firstMetaData  = "Title: Post 1\nDescription: Description 1\nTags: tdd, go\n---\nHello World\nThis is the body!"
			secondMetaData = "Title: Post 1\nDescription: Description 2\nTags: oop, grpc"
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstMetaData)},
			"hello-world2.md": {Data: []byte(secondMetaData)},
		}

		posts, err := blogposts.NewPosts(fs)

		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := blogposts.Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: "Hello World\nThis is the body!"}

		assertPost(t, got, want)
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
