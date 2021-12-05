package resolvers

import (
	"context"
	"hello_world/src/graph/model"
	"hello_world/src/utils"

	"github.com/99designs/gqlgen/graphql"
)

func SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	upload := utils.UploadFile(file.File, file.ContentType)

	return &model.File{
		ID:          1,
		Name:        file.Filename,
		Content:     upload.Location,
		ContentType: file.ContentType,
	}, nil
}
