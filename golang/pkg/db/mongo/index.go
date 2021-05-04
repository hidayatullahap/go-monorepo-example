package mongo

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnsureIndexes(db *mongo.Database) {
	t := true
	opt := &options.IndexOptions{
		Background: &t,
	}

	EnsureIndex(db, CollectionUsers, bson.M{"username": 1}, opt)
	EnsureIndex(db, CollectionUserToken, bson.M{"token": 1}, opt)
	EnsureIndex(db, CollectionWatchlist, bson.M{"omdb_id": 1}, opt)

}

// EnsureIndex is for create index if not exist
func EnsureIndex(db *mongo.Database, collectionName string, keys bson.M, opt *options.IndexOptions) {
	var keyIndex []string
	for k := range keys {
		keyIndex = append(keyIndex, k)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Collection(collectionName)

	indexes := collection.Indexes()
	cursor, err := indexes.List(ctx)
	if err != nil {
		log.Panicf("index list error %v", err)
	}

	if cursor != nil {
		for cursor.Next(ctx) {
			var index []primitive.E
			errCursor := cursor.Decode(&index)
			if errCursor != nil {
				log.Panicf("index list error %v", errCursor)
			}

			// skip creating index if key field already exist
			keyIsExist := keyFieldIndexIsExist(index, keyIndex)
			if keyIsExist {
				return
			}
		}

		mod := mongo.IndexModel{
			Keys:    keys,
			Options: opt,
		}

		opts := options.CreateIndexes().SetMaxTime(5 * time.Second)
		_, err = collection.Indexes().CreateOne(ctx, mod, opts)
		if err != nil {
			log.Panicf("ensure index error %v", err)
		}
	}
}

func keyFieldIndexIsExist(index []primitive.E, keyIndex []string) bool {
	sort.Strings(keyIndex)

	for _, e := range index {
		if e.Key == "key" {
			values, ok := e.Value.([]primitive.E)
			if ok {
				var keyFields []string
				for _, value := range values {
					keyFields = append(keyFields, value.Key)
				}

				sort.Strings(keyFields)
				equal := funk.Equal(keyIndex, keyFields)
				if equal {
					return true
				}
			}
		}
	}

	return false
}
