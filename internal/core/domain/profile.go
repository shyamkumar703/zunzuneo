package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"zunzuneo/internal/dependencies"

	"github.com/google/uuid"
)

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

func CreateRandomPersonality() Personality {
	return Personality{
		Openness:          int8(rand.Intn(11) - 5), // -5 to 5
		Conscientiousness: int8(rand.Intn(11) - 5), // -5 to 5
		Extraversion:      int8(rand.Intn(11) - 5), // -5 to 5
		Agreeableness:     int8(rand.Intn(11) - 5), // -5 to 5
		Neuroticism:       int8(rand.Intn(11) - 5), // -5 to 5
	}
}

type Profile struct {
	ID          string      `json:"id"`
	Handle      string      `json:"handle"`
	DisplayName string      `json:"displayName"`
	Bio         string      `json:"bio"`
	Personality Personality `json:"personality"`
	Interests   []Interest  `json:"interests"`
	JoinedAt    time.Time   `json:"joinedAt"`
	IsAI        bool        `json:"isAI"`
	IsVerified  bool        `json:"isVerified"`
}

func CreateRandomProfile(ctx *context.Context, prompt string) (*Profile, error) {
	personality := CreateRandomPersonality()
	id := uuid.New().String()
	type GeneratedProfileResponse struct {
		Handle      string     `json:"handle"`
		DisplayName string     `json:"displayName"`
		Bio         string     `json:"bio"`
		Interests   []Interest `json:"interests"`
	}
	formattedPrompt := fmt.Sprintf(
		`
		You are generating a Twitter profile for a user who lives in a world just like the one you know, but
		with some fictional elements. Here are the fictional elements: %s. Here are the user's big 5 personality
		traits: %+v. You need to generate:
		1) the user's display name. Generate a name for this user, this can include both a first name and a last name, just
		a first name, just a last name, even a first name, middle name, and a last name. Try to make the name as interesting and as
		realistic as possible. THIS DOES NOT HAVE TO BE THE USER'S REAL NAME. Think of this more so as an extension of the display name. If the user's personality
		traits indicate that they may use their real name as their display name on Twitter however, make it their real name. You DO NOT
		have to worry about capitalization rules that would normally apply to names, you can mix in all lowercase and all uppercase names.
		2) the user's interests. An interest needs to be a short, specific topic that a random Twitter user may be interested in. For example, instead
		of a user being interested in basketball, try to pinpoint a particular team, like the New York Knicks or the Golden State Warriors. In conjuction with
		generating an interest, generate a corresponding interest level; a floating point value between -1 and 1 that indicates the user's attitude towards the
		interest. A low value (closer to -1) will indicate a negative interest in the topic (the user is interested insofar as they do not like that topic), and a high value
		(closer to 1) will indicate a positive interest in the topic (the user likes and enjoys the topic). Allow these to build on top of each other - for example,
		if a user has a high positive interest in the San Francisco Giants, they should have a significant negative interest in the Los Angeles Dodgers, as those two teams
		are rivals. Generate at least 40 of these topics, and make them as specific as possible. Some of them should be political, some of them should be hobby-related, some of
		them should be job-related. Do not generate something like "music festivals", instead generate "Coachella". Prefer specific proper nouns over overarching topics.
		3) the user's Twitter handle. Try to keep this short, include numbers and punctuation, lean into anything that could
		be potentially funny, avoid any direct references to the big 5 personality traits but use them to inform your generation, incorporate the user's interests, incorporate
		the user's display name if necessary.
		4) the user's Twitter bio. Try to keep this to roughly 2-4 sentences, reference the user's interests in a way that would align with their personality traits.
		This needs to be outputted in the JSON format below:
		{
			handle: string, // warriorsfan22
			displayName: string // frosty
			bio: string // die hard warriors fan. FUCK THE KINGS
			interests: Interest[] // [{interest: 'Golden State Warriors', interestLevel: 0.94}]
		}

		Here is the JSON format for interests:
		{
			interest: string, // Golden State Warriors
			interestLevel: float32 // 0.94
		}
		`,
		prompt,
		personality,
	)
	response, err := dependencies.RequestLLM(formattedPrompt, ctx)
	if err != nil {
		return nil, err
	}
	trimmedString := strings.TrimPrefix(*response, "```json")
	trimmedString = strings.TrimSuffix(trimmedString, "```")
	var profileResponse GeneratedProfileResponse
	jsonErr := json.Unmarshal([]byte(trimmedString), &profileResponse)
	if jsonErr != nil {
		return nil, jsonErr
	}
	profile := &Profile{
		ID:          id,
		Handle:      profileResponse.Handle,
		DisplayName: profileResponse.DisplayName,
		Bio:         profileResponse.Bio,
		Personality: personality,
		Interests:   profileResponse.Interests,
		JoinedAt:    time.Now(),
		IsAI:        true,
		IsVerified:  false,
	}
	return profile, nil
}

type Interest struct {
	Interest string `json:"interest"`
	// a value between -1 and 1
	// a negative value indicates NEGATIVE interest,
	// almost derision. a positive value indicates
	// POSITIVE interest.
	InterestLevel float32 `json:"interestLevel"`
}
