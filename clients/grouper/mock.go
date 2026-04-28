package grouper

import (
	"context"

	"github.com/cyverse-de/permissions/models"
)

// MockGrouperClient represents a mock Grouper client.
type MockGrouperClient struct {
	groups map[string][]*GroupInfo
}

// NewMockGrouperClient returns a new Mock grouper client.
func NewMockGrouperClient(groups map[string][]*GroupInfo) *MockGrouperClient {
	return &MockGrouperClient{groups: groups}
}

// IsGroupSource returns true if the given source ID refers to a group in Grouper.
func (gc *MockGrouperClient) IsGroupSource(sourceID models.SubjectSourceID) bool {
	return sourceID == groupSubjectSource
}

// GroupsForSubject returns a mock list of groups for a subject.
func (gc *MockGrouperClient) GroupsForSubject(_ context.Context, subjectID string) ([]*GroupInfo, error) {
	return gc.groups[subjectID], nil
}

// AddSourceIDToPermissions is a no-op for now.
func (gc *MockGrouperClient) AddSourceIDToPermissions(_ context.Context, _ []*models.Permission) error {
	return nil
}

// AddSourceIDToPermission is a no-op for now.
func (gc *MockGrouperClient) AddSourceIDToPermission(_ context.Context, _ *models.Permission) error {
	return nil
}

// ListGroupMembers is a no-op for now.
func (gc *MockGrouperClient) ListGroupMembers(
	_ context.Context,
	_ models.ExternalSubjectID,
) ([]*models.SubjectOut, error) {
	return nil, nil
}
