package domain

import "time"

type Tweet struct {
	ID        string
	ProfileID string
	Content   string
	MediaURL  *string
	ReplyToID *string
	QuoteID   *string
	Metadata  TweetMetadata
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TweetMetadata struct {
	MainTopic string
	// a value between -1 and 1.
	// a lower value indicates negative sentiment
	// towards the `MainTopic`, and a higher value
	// indicates a positive sentiment towards the
	// `MainTopic`.
	Sentiment        float32
	TangentialTopics []string
}
