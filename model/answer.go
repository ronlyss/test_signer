package models

// Question represents a test question
type Question struct {
    ID      int64  `json:"id"`
    Text    string `json:"text"`
}
