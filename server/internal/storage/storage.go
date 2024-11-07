package storage

import (
	"context"
	"errors"
	"log/slog"
	"plants/internal/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage struct {
	Client   *mongo.Client
	DataBase *mongo.Database
}

func New(ctx context.Context, uri, db string) (*Storage, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		slog.Error("unable to connect")
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		slog.Error("ping doesnt work")
		return nil, err
	}
	database := client.Database(db)
	return &Storage{
		DataBase: database,
		Client:   client,
	}, nil
}

func (s Storage) Disconnect(ctx context.Context) error {
	return s.Client.Disconnect(ctx)
}

func (s *Storage) GetPlantsWithCareRules(ctx context.Context) ([]*models.Plant, error) {
	collection := s.DataBase.Collection("plants")
	filter := bson.M{"care_rules": bson.M{"$ne": nil}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	plants := make([]*models.Plant, 0)
	for cursor.Next(ctx) {
		var plant models.Plant
		cursor.Decode(&plant)
		plants = append(plants, &plant)
	}
	return plants, nil
}

func (s *Storage) CreateNewCareRule(ctx context.Context, rule *models.CareRules) error {
	collection := s.DataBase.Collection("care_rules")
	_, err := collection.InsertOne(ctx, rule)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetCareRulesForPlant(ctx context.Context, species string) (*models.CareRules, error) {
	collection := s.DataBase.Collection("care_rules")
	filter := bson.M{"species": species}
	cursor := collection.FindOne(ctx, filter)

	var result models.CareRules
	err := cursor.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *Storage) GetPlants(ctx context.Context) ([]*models.Plant, error) {
	collection := s.DataBase.Collection("plants")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	plants := make([]*models.Plant, 0)
	for cursor.Next(ctx) {
		var plant models.Plant
		if err := cursor.Decode(&plant); err != nil {
			return nil, err
		}
		plants = append(plants, &plant)
	}
	return plants, nil
}

func (s *Storage) AddPlant(ctx context.Context, plant *models.Plant) error {
	collection := s.DataBase.Collection("plants")
	careRules, err := s.GetCareRulesForPlant(ctx, plant.Species)
	if err == nil {
		plant.CareRules = careRules.ID
	} else {
		plant.CareRules = primitive.NewObjectID()
	}
	_, err = collection.InsertOne(ctx, plant)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) SearchUser(ctx context.Context, login string, password string) (string, int32, error) {
	collection := s.DataBase.Collection("users")
	var filter bson.M
	if strings.Contains(login, "@") {
		filter = bson.M{"email": login, "password": password}
	} else {
		filter = bson.M{"phone_number": login, "password": password}
	}
	cursor := collection.FindOne(ctx, filter)
	var result models.User
	err := cursor.Decode(&result)
	if err != nil {
		return "", 0, err
	}
	return result.ID.Hex(), result.Role, nil
}

func (s *Storage) AddUser(ctx context.Context, user *models.User) error {
	collection := s.DataBase.Collection("users")
	var filter bson.M

	if user.Email != "" {
		filter = bson.M{"email": user.Email}
	} else {
		filter = bson.M{"phone_number": user.PhoneNumber}
	}

	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			_, err = collection.InsertOne(ctx, user)
			return err
		}
		return err
	}
	return nil
}

func (s *Storage) GetUser(ctx context.Context, id string) (*models.User, error) {
	collection := s.DataBase.Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"id": objID}
	cursor := collection.FindOne(ctx, filter)

	var result models.User
	err = cursor.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *Storage) GetTrade(ctx context.Context, id string, mode int32) (*models.Trade, error) {
	collection := s.DataBase.Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	var filter bson.M
	if mode == 1 {
		filter = bson.M{
			"offerer": bson.M{"id": objID},
		}
	} else {
		filter = bson.M{
			"accepter": bson.M{"id": objID},
		}
	}
	cursor := collection.FindOne(ctx, filter)

	var result models.Trade
	err = cursor.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}