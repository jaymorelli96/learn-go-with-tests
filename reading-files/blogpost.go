package blogpost

import (
	"io/fs"
)

func NewPosts(fileSystem fs.FS) ([]Post, error) {

	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err //todo
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}
