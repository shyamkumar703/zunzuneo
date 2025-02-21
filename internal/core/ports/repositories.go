package ports

import (
	"context"
	"zunzuneo/internal/core/domain"
)

type ProfileRepository interface {
	GenerateProfile(ctx context.Context, profile domain.Profile) (bool, error)
	GetByID(ctx context.Context, id string) (*domain.Profile, error)
	Update(ctx context.Context, profile domain.Profile) (*domain.Profile, error)
	// TODO: some sort of search feature? refer to ProfileFilter in planning/
}

type TweetRepository interface {
	Create(ctx context.Context, tweet *domain.Tweet) (bool, error)
	GetByID(ctx context.Context, id string) (*domain.Tweet, error)
	ListByProfile(ctx context.Context, profileID string) ([]*domain.Tweet, error)
	ListReplies(ctx context.Context, tweetID string) ([]*domain.Tweet, error)
}
