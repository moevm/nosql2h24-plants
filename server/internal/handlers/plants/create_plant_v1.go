package plants

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"plants/internal/models"
	api "plants/internal/pkg/pb/plants/v1"
)

func (h *Handler) CreatePlantV1(
	ctx context.Context,
	req *api.CreatePlantV1Request,
) (*api.CreatePlantV1Response, error) {
	userId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}
	err = h.storage.AddPlant(ctx, &models.Plant{
		ID:                primitive.NewObjectID(),
		UserID:            userId,
		Image:             req.Image,
		Size:              req.Size,
		Price:             req.Price,
		LightCondition:    req.LightCondition,
		WateringFrequency: req.WateringFrequency,
		TemperatureRegime: req.TemperatureRegime,
		CareComplexity:    req.CareComplexity,
		Description:       req.Description,
		Type:              req.Type,
		Species:           req.Species,
		CareRules:         primitive.NilObjectID,
		CreatedAt:         time.Now().UTC(),
		Place:             req.Place,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create plant")
	}
	return &api.CreatePlantV1Response{}, nil
}