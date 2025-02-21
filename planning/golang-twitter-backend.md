# Virtual Twitter Backend Architecture

## Project Structure
```
virtual-twitter/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── core/
│   │   ├── domain/
│   │   │   ├── profile.go
│   │   │   ├── tweet.go
│   │   │   └── event.go
│   │   └── ports/
│   │       ├── repositories.go
│   │       └── services.go
│   ├── handlers/
│   │   ├── event_handler.go
│   │   ├── profile_handler.go
│   │   └── tweet_handler.go
│   ├── repositories/
│   │   ├── supabase/
│   │   │   ├── profile_repository.go
│   │   │   └── tweet_repository.go
│   │   └── cache/
│   │       └── redis_repository.go
│   └── services/
│       ├── profile_service.go
│       ├── tweet_service.go
│       └── llm_service.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   └── utils/
│       ├── logger.go
│       └── validator.go
└── go.mod
```

## Core Domain Models

```go
// internal/core/domain/profile.go
package domain

type Personality struct {
    Openness        float64
    Conscientiousness float64
    Extraversion    float64
    Agreeableness   float64
    Neuroticism     float64
}

type Profile struct {
    ID           string
    Handle       string
    DisplayName  string
    Bio          string
    Personality  Personality
    Interests    []string
    JoinedAt     time.Time
    IsVerified   bool
}

// internal/core/domain/tweet.go
type Tweet struct {
    ID          string
    ProfileID   string
    Content     string
    MediaURLs   []string
    ReplyToID   *string
    QuoteID     *string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Metadata    TweetMetadata
}

type TweetMetadata struct {
    Sentiment    float64
    Topics      []string
    IsGenerated bool
}
```

## Interface Definitions

```go
// internal/core/ports/repositories.go
package ports

type ProfileRepository interface {
    Create(ctx context.Context, profile *domain.Profile) error
    GetByID(ctx context.Context, id string) (*domain.Profile, error)
    Update(ctx context.Context, profile *domain.Profile) error
    List(ctx context.Context, filters ProfileFilters) ([]*domain.Profile, error)
}

type TweetRepository interface {
    Create(ctx context.Context, tweet *domain.Tweet) error
    GetByID(ctx context.Context, id string) (*domain.Tweet, error)
    ListByProfile(ctx context.Context, profileID string) ([]*domain.Tweet, error)
    ListReplies(ctx context.Context, tweetID string) ([]*domain.Tweet, error)
}

// internal/core/ports/services.go
type ProfileService interface {
    GenerateProfile(ctx context.Context) (*domain.Profile, error)
    GetRelevantProfiles(ctx context.Context, event domain.Event) ([]*domain.Profile, error)
}

type TweetService interface {
    GenerateTweet(ctx context.Context, profile *domain.Profile, event domain.Event) (*domain.Tweet, error)
    GenerateReply(ctx context.Context, profile *domain.Profile, parentTweet *domain.Tweet) (*domain.Tweet, error)
}
```

## Implementation Examples

```go
// internal/repositories/supabase/profile_repository.go
package supabase

import (
    "github.com/your/project/internal/core/domain"
    "github.com/your/project/internal/core/ports"
)

type profileRepository struct {
    db *database.Supabase
}

func NewProfileRepository(db *database.Supabase) ports.ProfileRepository {
    return &profileRepository{db: db}
}

func (r *profileRepository) Create(ctx context.Context, profile *domain.Profile) error {
    // Implementation
}

// internal/services/profile_service.go
package services

type profileService struct {
    repo    ports.ProfileRepository
    llm     ports.LLMService
    cache   ports.CacheRepository
}

func NewProfileService(
    repo ports.ProfileRepository,
    llm ports.LLMService,
    cache ports.CacheRepository,
) ports.ProfileService {
    return &profileService{
        repo: repo,
        llm: llm,
        cache: cache,
    }
}

func (s *profileService) GenerateProfile(ctx context.Context) (*domain.Profile, error) {
    // Implementation
}

// cmd/server/main.go
func main() {
    // Initialize configuration
    cfg := config.Load()

    // Setup dependencies
    db := supabase.NewClient(cfg.SupabaseURL, cfg.SupabaseKey)
    cache := redis.NewClient(cfg.RedisURL)
    llm := llm.NewClient(cfg.LLMConfig)

    // Initialize repositories
    profileRepo := supabase.NewProfileRepository(db)
    tweetRepo := supabase.NewTweetRepository(db)
    cacheRepo := redis.NewCacheRepository(cache)

    // Initialize services
    profileSvc := services.NewProfileService(profileRepo, llm, cacheRepo)
    tweetSvc := services.NewTweetService(tweetRepo, llm, cacheRepo)

    // Initialize handlers
    profileHandler := handlers.NewProfileHandler(profileSvc)
    tweetHandler := handlers.NewTweetHandler(tweetSvc)

    // Setup router
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)

    // Register routes
    router.Route("/api/v1", func(r chi.Router) {
        r.Mount("/profiles", profileHandler.Routes())
        r.Mount("/tweets", tweetHandler.Routes())
    })

    // Start server
    log.Fatal(http.ListenAndServe(":8080", router))
}
```

## Key Design Principles

1. **Clean Architecture**:
   - Core domain models are independent of external concerns
   - Dependencies point inward
   - Interfaces define boundaries between layers

2. **Dependency Injection**:
   - Services receive their dependencies through constructors
   - Makes testing and swapping implementations easier

3. **Repository Pattern**:
   - Abstracts data storage details
   - Makes it easy to switch between Supabase, cache, or other storage

4. **Interface Segregation**:
   - Small, focused interfaces
   - Each component only depends on what it needs

5. **Concurrent Processing**:
   - Use contexts for cancellation and timeouts
   - Goroutines for parallel processing where appropriate

## Important Considerations for Virtual Twitter

1. **Caching Strategy**:
```go
type CacheRepository interface {
    GetGeneratedTweets(ctx context.Context, eventType string) ([]*domain.Tweet, error)
    CacheGeneratedTweets(ctx context.Context, eventType string, tweets []*domain.Tweet) error
    InvalidateCache(ctx context.Context, pattern string) error
}
```

2. **Event Processing**:
```go
type EventProcessor interface {
    ProcessEvent(ctx context.Context, event domain.Event) error
    BatchProcessEvents(ctx context.Context, events []domain.Event) error
}
```

3. **Rate Limiting**:
```go
type RateLimiter interface {
    AllowRequest(ctx context.Context, key string) bool
    GetQuotaRemaining(ctx context.Context, key string) (int64, error)
}
```
