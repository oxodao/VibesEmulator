package models

type Answer struct {
	ID          uint64 `json:"id"`
	Text        string `json:"text"`
	CommentText string `json:"commentText"`
}
