package models

type Playlist struct {
	ID     string   `json:"id,omitempty"`
	Artist string   `json:"artist"`
	Albums  []Album `json:"albums,omitempty"`
}