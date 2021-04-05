package db

import (
	"context"
	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

// The repetition of the model struct is to avoid tight coupling to the odds api structure
// If the odds api changes we want to avoid our own database breaking
type MatchOddsModel struct {
	H2h		[]float32	`bson:"h2h" json:"h2h"`
}

type SportOddSitesModel struct {
	SiteKey 	string			`bson:"site_key" json:"site_key"`
	SiteNice	string			`bson:"site_nice" json:"site_nice"`
	LastUpdate	int				`bson:"last_update" json:"last_update"`
	Odds		MatchOddsModel	`bson:"odds" json:"odds"`
	SitesCount	int				`bson:"sites_count" json:"sites_count"`
}

type OddsModel struct {
	ID				primitive.ObjectID 		`bson:"_id" json:"_id,omitempty"`
	ApiID			string					`bson:"apiId" json:"id,omitempty"`
	SportKey		string					`bson:"sport_key" json:"sport_key"`
	SportNice		string					`bson:"sport_nice" json:"sport_nice"`
	Teams			[]string				`bson:"teams" json:"teams"`
	CommenceTime	int						`bson:"commence_time" json:"commence_time"`
	HomeTeam		string					`bson:"home_team" json:"home_team"`
	Sites			[]SportOddSitesModel	`bson:"sites" json:"sites"`
	CreatedAt 		time.Time          		`bson:"created_at"`
	UpdatedAt 		time.Time          		`bson:"updated_at"`
}

func (r Repository) SaveFixtures (odds *[]bookmaker.SportOddsApiModel) ([]interface{}, error) {
	collection := r.db.Collection(FixturesCollection)

	var oddsModelList []interface{}

	for _, odd := range *odds {
		var oddsModel OddsModel

		oddsModel.ID = primitive.NewObjectID()

		copier.Copy(&oddsModel, &odd)
		oddsModelList = append(oddsModelList, oddsModel)
	}
	res, err := collection.InsertMany(context.TODO(), oddsModelList)
	if err != nil {
		return []interface{}{}, err
	}
	return res.InsertedIDs, nil
}
