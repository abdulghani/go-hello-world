package resolvers

import (
	"context"
	"hello_world/src/graph/model"
	"io/ioutil"
	"strconv"
)

func SingleUploadWithPayload(ctx context.Context, req model.UploadFile) (*model.File, error) {
	content, err := ioutil.ReadAll(req.File.File)
	if err != nil {
		return nil, err
	}
	return &model.File{
		ID:          strconv.Itoa(req.ID),
		Name:        req.File.Filename,
		Content:     string(content),
		ContentType: req.File.ContentType,
	}, nil
}
