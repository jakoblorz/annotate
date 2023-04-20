package main

import (
	"encoding/json"
	"github.com/jensteichert/colt"
	"log"
	"time"
)

type RepositorySettings struct {
	Host string `json:"host" bson:"host" default:"localhost"`
	Port int    `json:"port" bson:"port" default:"5432"`
}

type RecentUpdateInfo struct {
	LastUpdatedAt int64 `json:"last_updated_at" bson:"last_updated_at"`
}

func (r RecentUpdateInfo) Default() *RecentUpdateInfo {
	r.LastUpdatedAt = time.Now().Unix()
	return &r
}

type Repository struct {
	colt.DocWithTimestamps `bson:",inline"`
	Name                   string `json:"name" bson:"name"`

	// Settings uses the Default generic type (which is based on annotate.D), which allows to define a default value
	// for the type itself. In this example, we define a default value for the RepositorySettings type, which is a
	// struct with two fields. This is useful, because we don't have to set the default value for every field of the
	// struct.
	Settings Default[RepositorySettings] `json:"settings" bson:"settings"`

	// RecentUpdateInfo uses the DefaultX generic type (which is based on annotate.DX), which allows to define a
	// default value for the type itself. In this example, we define a default value for the RecentUpdateInfo type, which is the current
	// timestamp. This is useful, because we don't have to define a default value for every field of the struct, but
	// can define a default value for the struct itself.
	RecentUpdateInfo DefaultX[RecentUpdateInfo] `json:"recent_update_info" bson:"recent_update_info"`
}

func main() {
	db := colt.Database{}
	db.Connect("mongodb://...", "...")

	repositories := colt.GetCollection[*Repository](&db, "repositories")
	repository, _ := repositories.Insert(&Repository{
		Name: "annotate",
	})

	data, _ := json.Marshal(repository)
	log.Println(string(data)) // {"_id":"6442ce210ff68df56822dbb8","created_at":"2023-04-21T00:55:45.716764+02:00","name":"annotate","settings":{"host":"localhost","port":5432},"recent_update_info":{"last_updated_at":1682031345}}
}
