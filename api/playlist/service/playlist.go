package service

import (
	"errors"
	"github.com/andregumieri/music-for-all/playlist/models"
	"github.com/andregumieri/music-for-all/playlist/repository"
	"strings"
)

func CretePlaylist(playlist models.Playlist) (*models.Playlist, error) {
	if strings.TrimSpace(playlist.Artist) == "" {
		return nil, errors.New("artist cannot be null")
	}

	playlistResponse, err := repository.CretePlayList(playlist)
	if err != nil {
		return nil, err
	}
	return playlistResponse, nil
}

func GetPlaylist(id string) (*models.Playlist, error) {
	playlistResponse, err := repository.GetPlaylist(id)
	if err != nil {
		return nil, err
	}

	return playlistResponse, nil
}