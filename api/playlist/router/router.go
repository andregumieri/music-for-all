package router

import (
	"github.com/andregumieri/music-for-all/playlist/controllers"
	"github.com/andregumieri/music-for-all/responses"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func Router (request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.Path {
	case "/music-for-all/playlists":
		{
			if request.HTTPMethod == http.MethodPost {
				return controllers.HandleCreatePlaylist(request), nil
			}
			if request.HTTPMethod == http.MethodGet {
				return controllers.HandleGetPlaylist(request), nil
			}
		}
	default:
		return responses.NotFound(), nil
	}

	return responses.MethodNotAllowed(request.HTTPMethod), nil
}