package repository

import (
	"errors"
	"fmt"
	"github.com/andregumieri/music-for-all/playlist/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
	"os"
)

var region = os.Getenv("AWS_REGION")
var tableName = os.Getenv("PLAYLIST_TABLE_NAME")
var db *dynamodb.DynamoDB

func init() {
	db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(region))
}

func CretePlayList(playlist models.Playlist) (*models.Playlist, error) {
	playlist.ID = uuid.NewV4().String()

	dynamoPlaylist, err := dynamodbattribute.MarshalMap(playlist)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: &tableName,
		Item: dynamoPlaylist,
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func GetPlaylist(id string) (*models.Playlist, error) {
	// Prepares the input to retrieve the item with the given ID
	input := &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	// Retrieves the item
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, errors.New("playlist not found")
	}

	var playlist = new(models.Playlist)
	err = dynamodbattribute.UnmarshalMap(result.Item, &playlist)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}