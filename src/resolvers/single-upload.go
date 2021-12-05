package resolvers

import (
	"context"
	"hello_world/src/graph/model"
	"hello_world/src/utils"

	"github.com/99designs/gqlgen/graphql"
)

func SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	authError := utils.ShouldAuth(ctx)
	if authError != nil {
		return nil, authError
	}

	upload := utils.UploadFile(file.File)

	return &model.File{
		ID:          upload.Key,
		Name:        file.Filename,
		Content:     upload.Output.Location,
		ContentType: file.ContentType,
	}, nil
}
