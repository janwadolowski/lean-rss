package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Image         string             `bson:"image,omitempty"`
	Title         string             `bson:"title,omitempty"`
	Body          string             `bson:"body,omitempty"`
	Author        string             `bson:"author,omitempty"`
	Language      string             `bson:"language,omitempty"`
	Tags          []string           `bson:"tags,omitempty"`
	URL           string             `bson:"url,omitempty"`
	PublishedTime time.Time          `bson:"postedTimestamp,omitempty"`
	ScrappedTime  time.Time          `bson:"scrappedTimestamp,omitempty"`
}

//TableName return name of database table
func (a *Article) TableName() string {
	return "article"
}
