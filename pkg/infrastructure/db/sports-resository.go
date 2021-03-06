package db

import (
	"context"
	"time"

	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SportsModel struct {
	ID        	primitive.ObjectID 	`bson:"_id" json:"_id,omitempty"`
	Key 		string				`bson:"key"`
	Active		bool				`bson:"active"`
	Group		string				`bson:"group"`
	Details		string				`bson:"details"`
	Title		string				`bson:"title"`
	CreatedAt 	time.Time          	`bson:"created_at"`
	UpdatedAt 	time.Time          	`bson:"updated_at"`
}

func (r Repository) SaveSports (sports *[]bookmaker.SportApiModel) ([]interface{}, error) {
	collection := r.db.Collection(SportsCollection)

	var sportsModelList []interface{}

	for _, v := range *sports {
		var sportModel SportsModel

		sportModel.ID = primitive.NewObjectID()
		copier.Copy(&sportModel, &v)
		sportsModelList = append(sportsModelList, sportModel)
	}

	res, err := collection.InsertMany(context.TODO(), sportsModelList)
	if err != nil {
		return []interface{}{}, err
	}
	return res.InsertedIDs, nil
}
