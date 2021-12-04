package resolvers

import (
	"context"
	"hello_world/src/graph/model"
	"io/ioutil"

	"github.com/99designs/gqlgen/graphql"
)

func SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	content, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}
	return &model.File{
		ID:          1,
		Name:        file.Filename,
		Content:     string(content),
		ContentType: file.ContentType,
	}, nil
}
