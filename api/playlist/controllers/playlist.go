package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andregumieri/music-for-all/tree/develop/api/models"
	"github.com/andregumieri/music-for-all/tree/develop/api/responses"
	"github.com/andregumieri/music-for-all/tree/develop/api/service"
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

func HandleCreatePlaylist(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var playlist models.Playlist

	if err := json.Unmarshal([]byte(request.Body), &playlist); err != nil {
		fmt.Println(fmt.Sprintf("Error parsing playlist body. Error: %s", err.Error()))
		return responses.BadRequest(err)
	}

	playlistResponse, err := service.CretePlaylist(playlist)

	if err != nil {
		if strings.Contains(err.Error(), "artist") {
			return responses.BadRequest(err)
		}
		fmt.Println(fmt.Sprintf("Error creating playlist. Error: %s", err.Error()))
		return responses.InternalServerError(err)
	}

	js, err := json.Marshal(playlistResponse)
	return responses.Created(string(js))
}

func HandleGetPlaylist(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	id := request.PathParameters["id"]
	if strings.TrimSpace(id) == "" {
		return responses.BadRequest(errors.New("id parameter cannot be null"))
	}

	playlistResponse, err := service.GetPlaylist(id)

	if err != nil {
		if strings.Contains(err.Error(), "playlist not found") {
			return responses.NotFound()
		}
		fmt.Println(fmt.Sprintf("Error creating playlist. Error: %s", err.Error()))
		return responses.InternalServerError(err)
	}

	js, err := json.Marshal(playlistResponse)
	return responses.Ok(string(js))
}
