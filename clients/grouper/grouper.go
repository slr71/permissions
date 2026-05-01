package grouper

import (
	"context"
	"database/sql"

	"github.com/cyverse-de/dbutil"
	"github.com/cyverse-de/permissions/logger"
	"github.com/cyverse-de/permissions/models"
	"go.opentelemetry.io/otel"

	"github.com/lib/pq"
)

const otelName = "github.com/cyverse-de/permissions/clients/grouper"
const groupSubjectSource = models.SubjectSourceID("g:gsa")

// GroupInfo represents information about a group in Grouper.
type GroupInfo struct {
	ID   string
	Name string
}

// Grouper is the interface implemented by a Grouper client instance.
type Grouper interface {
	IsGroupSource(sourceID models.SubjectSourceID) bool
	GroupsForSubject(context.Context, string) ([]*GroupInfo, error)
	AddSourceIDToPermissions(context.Context, []*models.Permission) error
	AddSourceIDToPermission(context.Context, *models.Permission) error
	ListUsersInGroup(context.Context, models.ExternalSubjectID) ([]*models.SubjectOut, error)
}

// Client represents a Grouper client instance.
//
// Note: the grouper client is intended to be a read-only client. Explicit transactions are not
// used here for that reason.
type Client struct {
	db     *sql.DB
	prefix string
}

// NewGrouperClient creates and returns a new Grouper client for the given database URI and group name prefix.
func NewGrouperClient(dbURI, prefix string) (*Client, error) {
	connector, err := dbutil.NewDefaultConnector("1m")
	if err != nil {
		return nil, err
	}

	db, err := connector.Connect("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	return &Client{
		db:     db,
		prefix: prefix,
	}, nil
}

// IsGroupSource returns true if the given source ID refers to a group in Grouper.
func (gc *Client) IsGroupSource(sourceID models.SubjectSourceID) bool {
	return sourceID == groupSubjectSource
}

// GroupsForSubject returns the list of groups that the subject with the given ID belongs to.
func (gc *Client) GroupsForSubject(ctx context.Context, subjectID string) ([]*GroupInfo, error) {
	ctx, span := otel.Tracer(otelName).Start(ctx, "GroupsForSubject")
	defer span.End()

	// Query the database.
	query := `SELECT group_id, group_name FROM grouper_memberships_v
            WHERE subject_id = $1 AND group_name LIKE $2 AND list_name = 'members'`
	rows, err := gc.db.QueryContext(ctx, query, subjectID, gc.prefix+"%")
	if err != nil {
		return nil, err
	}

	// Extract the groups from the database.
	groups := make([]*GroupInfo, 0)
	for rows.Next() {
		var group GroupInfo
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, &group)
	}

	return groups, nil
}

// AddSourceIDToPermissions adds the subject source IDs to a slice of Permission objects.
func (gc *Client) AddSourceIDToPermissions(ctx context.Context, permissions []*models.Permission) error {
	ctx, span := otel.Tracer(otelName).Start(ctx, "AddSourceIDToPermissions")
	defer span.End()

	// Get a list of subject identifiers.
	subjectIDs := make([]string, 0)
	for _, permission := range permissions {
		subjectIDs = append(subjectIDs, string(*permission.Subject.SubjectID))
	}

	// Query the database.
	query := `SELECT subject_id, subject_source FROM grouper_members
            WHERE subject_id = ANY($1)`
	rows, err := gc.db.QueryContext(ctx, query, pq.Array(subjectIDs))
	if err != nil {
		return err
	}
	defer logger.LogClose(rows, "grouper members rows")

	// Build a map from subject ID to source ID.
	m := make(map[string]string)
	for rows.Next() {
		var subjectID, sourceID string
		if err := rows.Scan(&subjectID, &sourceID); err != nil {
			return err
		}
		m[subjectID] = sourceID
	}

	// Add the subject IDs to the permission objects.
	for _, permission := range permissions {
		sourceID := models.SubjectSourceID(m[string(*permission.Subject.SubjectID)])
		permission.Subject.SubjectSourceID = &sourceID
	}

	return nil
}

// AddSourceIDToPermission adds the subject source ID to a permission object.
func (gc *Client) AddSourceIDToPermission(ctx context.Context, permission *models.Permission) error {
	return gc.AddSourceIDToPermissions(ctx, []*models.Permission{permission})
}

// ListUsersInGroup returns the list of users in the group with the given ID. Nested groups are supported, but only the
// users in the group or any nested groups are returned.
func (gc *Client) ListUsersInGroup(
	ctx context.Context,
	groupID models.ExternalSubjectID,
) ([]*models.SubjectOut, error) {
	ctx, span := otel.Tracer(otelName).Start(ctx, "ListUsersInGroup")
	defer span.End()

	// Query the database.
	query := `SELECT subject_id, subject_source FROM grouper_memberships_v
            WHERE group_id = $1 AND list_name = 'members' AND subject_source != $2`
	rows, err := gc.db.QueryContext(ctx, query, string(groupID), string(groupSubjectSource))
	if err != nil {
		return nil, err
	}
	defer logger.LogClose(rows, "grouper memberships rows")

	// Build the list of subjects.
	members := make([]*models.SubjectOut, 0)
	for rows.Next() {
		var id models.ExternalSubjectID
		var source models.SubjectSourceID
		if err := rows.Scan(&id, &source); err != nil {
			return nil, err
		}
		members = append(members, &models.SubjectOut{SubjectID: &id, SubjectSourceID: &source})
	}

	return members, nil
}
