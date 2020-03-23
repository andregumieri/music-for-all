package router

import (
	"fmt"
	"github.com/andregumieri/music-for-all/playlist/controllers"
	"github.com/andregumieri/music-for-all/responses"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const path = "/music-for-all/playlists"
func Router (request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(fmt.Sprintf("Requested path: %s", request.Path))
	fmt.Println(fmt.Sprintf("Requested method: %s", request.HTTPMethod))
	hashPathId, _ := regexp.MatchString(fmt.Sprintf("%s/.+", path), request.Path)
	fmt.Println(fmt.Sprintf("Has id in path? %s", strconv.FormatBool(hashPathId)))

	if strings.Contains(request.Path, path) {
		switch request.HTTPMethod {
		case http.MethodPost:
			return controllers.HandleCreatePlaylist(request), nil
		case http.MethodGet:
			return controllers.HandleGetPlaylist(request), nil
		default:
			return responses.MethodNotAllowed(request.HTTPMethod), nil
		}
	}

	return responses.NotFound(), nil
}