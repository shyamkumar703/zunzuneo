package domain

import "time"

// big 5 personality traits.
// values range from 0-40, with higher
// values indicating a stronger predominance
// of the trait
type Personality struct {
	Openness          int8
	Conscientiousness int8
	Extraversion      int8
	Agreeableness     int8
	Neuroticism       int8
}

type Profile struct {
	ID          string
	Handle      string
	DisplayName string
	Bio         string
	Personality Personality
	Interests   []Interest
	JoinedAt    time.Time
	IsAI        bool
	IsVerified  bool
}

type Interest struct {
	interest string
	// a value between -1 and 1
	// a negative value indicates NEGATIVE interest,
	// almost derision. a positive value indicates
	// POSITIVE interest.
	interestLevel float32
}
