package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Author      string             `bson:"author" json:"author"`
	Text        string             `bson:"text" json:"text"`
	Category    []string           `bson:"category" json:"category"`
	PublishedAt time.Time          `bson:"publishedAt" json:"published_at"`
	UpdateddAt  time.Time          `bson:"upatedAt" json:"updated_at"`
}

func NewArticle(id primitive.ObjectID, title, author, text string, category []string) *Article {

	return &Article{
		ID:          id,
		Title:       title,
		Author:      author,
		Text:        text,
		Category:    category,
		PublishedAt: time.Now(),
	}

}
