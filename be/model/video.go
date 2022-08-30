package model

type Video struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Duration  string `json:"duration"`
	Names     string `json:"names"`
	Tags      string `json:"tags"`
	Index     int
}
