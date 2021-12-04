package resolvers

import (
	"context"
	"errors"
	"hello_world/src/graph/model"
	"io/ioutil"

	"github.com/99designs/gqlgen/graphql"
)

func MultipleUpload(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
	if len(files) == 0 {
		return nil, errors.New("empty list")
	}
	var resp []*model.File
	for i := range files {
		content, err := ioutil.ReadAll(files[i].File)
		if err != nil {
			return []*model.File{}, err
		}
		resp = append(resp, &model.File{
			ID:          i + 1,
			Name:        files[i].Filename,
			Content:     string(content),
			ContentType: files[i].ContentType,
		})
	}
	return resp, nil
}
