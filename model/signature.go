package models

// Signature represents a signed test
type Signature struct {
    ID         int64     `json:"id"`
    UserID     int64     `json:"user_id"`
    Questions  []Question `json:"questions"`
    Answers    []Answer   `json:"answers"`
    Timestamp  string    `json:"timestamp"`
    Signature  string    `json:"signature"` // Hashed representation of the signature
}
