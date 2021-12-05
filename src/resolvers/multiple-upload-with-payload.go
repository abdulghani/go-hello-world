package resolvers

import (
	"context"
	"errors"
	"hello_world/src/graph/model"
	"io/ioutil"
	"strconv"
)

func MultipleUploadWithPayload(ctx context.Context, req []*model.UploadFile) ([]*model.File, error) {
	if len(req) == 0 {
		return nil, errors.New("empty list")
	}
	var resp []*model.File
	for i := range req {
		content, err := ioutil.ReadAll(req[i].File.File)
		if err != nil {
			return []*model.File{}, err
		}
		resp = append(resp, &model.File{
			ID:      strconv.Itoa(i + 1),
			Name:    req[i].File.Filename,
			Content: string(content),
		})
	}
	return resp, nil
}
