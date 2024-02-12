// ent/schema/invitation.go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// Invitation holds the schema definition for the Invitation entity.
type Invitation struct {
    ent.Schema
}

// Fields of the Invitation.
func (Invitation) Fields() []ent.Field {
    return []ent.Field{
        field.String("email").NotEmpty().Comment("The email address to which the invitation was sent"),
        field.String("token").Unique().NotEmpty().Comment("The unique token associated with the invitation"),
        field.Bool("accepted").Default(false).Comment("Whether the invitation has been accepted"),
        field.Time("created_at").DefaultNow().Immutable().Comment("The time at which the invitation was created"),
        field.Time("expires_at").Optional().Nillable().Comment("The time at which the invitation expires"),
        // ... other fields like 'invited_by', 'accepted_at', etc. can be added here
    }
}

// Edges of the Invitation.
func (Invitation) Edges() []ent.Edge {
    return []ent.Edge{
        // This creates a foreign key linking the 'invitations' table to the 'groups' table
        edge.From("group", Group.Type).
            Ref("invitations").
            Unique().
            Required().
            Comment("The group that the user is being invited to"),
    }
}
