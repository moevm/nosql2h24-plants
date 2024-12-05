package data

import (
	"context"
	"io/ioutil"
	"os"

	api "plants/internal/pkg/pb/data/v1"
)

func (h *Handler) ExportDBV1(
	ctx context.Context,
	req *api.ExportDBV1Request,
) (*api.ExportDBV1Response, error) {
	res, err := h.storage.ExportDB(ctx)
	if err != nil {
		return &api.ExportDBV1Response{}, err
	}
	filePath := "exported_db.json"
	err = ioutil.WriteFile(filePath, res, 0644)
	if err != nil {
		return nil, err
	}
	defer os.Remove(filePath)
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &api.ExportDBV1Response{Db: res}, nil
}
