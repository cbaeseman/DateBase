package dateapp

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/models"
    "github.com/pocketbase/pocketbase/apis"
)

// Profile represents a user's profile in the date app
type Profile struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Age       int       `json:"age"`
    Gender    string    `json:"gender"`
    Interests string    `json:"interests"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

// DateAppClient wraps the PocketBase client
type DateAppClient struct {
    pb *pocketbase.PocketBase
}

// NewDateAppClient initializes a new PocketBase client
func NewDateAppClient(baseURL string) (*DateAppClient, error) {
    pb := pocketbase.New(baseURL)

    return &DateAppClient{pb: pb}, nil
}

// CreateProfile creates a new profile record
func (c *DateAppClient) CreateProfile(profile Profile) (*Profile, error) {
    record := &models.Record{}
    record.Set("name", profile.Name)
    record.Set("age", profile.Age)
    record.Set("gender", profile.Gender)
    record.Set("interests", profile.Interests)

    err := c.pb.Client.Collection("profiles").Create(record)
    if err != nil {
        return nil, fmt.Errorf("failed to create profile: %w", err)
    }

    profile.ID = record.GetID()
    profile.CreatedAt = record.GetCreatedAt()
    profile.UpdatedAt = record.GetUpdatedAt()

    return &profile, nil
}

// GetProfile retrieves a profile by ID
func (c *DateAppClient) GetProfile(id string) (*Profile, error) {
    record, err := c.pb.Client.Collection("profiles").Find(id)
    if err != nil {
        return nil, fmt.Errorf("failed to get profile: %w", err)
    }

    profile := Profile{
        ID:        record.GetID(),
        Name:      record.GetString("name"),
        Age:       record.GetInt("age"),
        Gender:    record.GetString("gender"),
        Interests: record.GetString("interests"),
        CreatedAt: record.GetCreatedAt(),
        UpdatedAt: record.GetUpdatedAt(),
    }

    return &profile, nil
}

// UpdateProfile updates an existing profile
func (c *DateAppClient) UpdateProfile(profile Profile) (*Profile, error) {
    record, err := c.pb.Client.Collection("profiles").Find(profile.ID)
    if err != nil {
        return nil, fmt.Errorf("failed to find profile: %w", err)
    }

    record.Set("name", profile.Name)
    record.Set("age", profile.Age)
    record.Set("gender", profile.Gender)
    record.Set("interests", profile.Interests)

    err = c.pb.Client.Collection("profiles").Update(record)
    if err != nil {
        return nil, fmt.Errorf("failed to update profile: %w", err)
    }

    profile.UpdatedAt = record.GetUpdatedAt()

    return &profile, nil
}

// DeleteProfile deletes a profile by ID
func (c *DateAppClient) DeleteProfile(id string) error {
    err := c.pb.Client.Collection("profiles").Delete(id)
    if err != nil {
        return fmt.Errorf("failed to delete profile: %w", err)
    }

    return nil
}

