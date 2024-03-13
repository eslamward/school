package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleDatabase struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func NewAticleDatabase(collection *mongo.Collection, ctx context.Context) *ArticleDatabase {

	return &ArticleDatabase{
		Collection: collection,
		Ctx:        ctx,
	}
}

type UserDatabase struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func NewUserDatabase(collection *mongo.Collection, ctx context.Context) *UserDatabase {

	return &UserDatabase{
		Collection: collection,
		Ctx:        ctx,
	}
}
