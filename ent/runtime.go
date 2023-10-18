// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ugent-library/projects/ent/project"
	"github.com/ugent-library/projects/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescIdentifier is the schema descriptor for identifier field.
	projectDescIdentifier := projectFields[2].Descriptor()
	// project.DefaultIdentifier holds the default value on creation for the identifier field.
	project.DefaultIdentifier = projectDescIdentifier.Default.(schema.Identifier)
	// projectDescCreated is the schema descriptor for created field.
	projectDescCreated := projectFields[10].Descriptor()
	// project.DefaultCreated holds the default value on creation for the created field.
	project.DefaultCreated = projectDescCreated.Default.(func() time.Time)
	// projectDescModified is the schema descriptor for modified field.
	projectDescModified := projectFields[11].Descriptor()
	// project.DefaultModified holds the default value on creation for the modified field.
	project.DefaultModified = projectDescModified.Default.(func() time.Time)
	// project.UpdateDefaultModified holds the default value on update for the modified field.
	project.UpdateDefaultModified = projectDescModified.UpdateDefault.(func() time.Time)
	// projectDescID is the schema descriptor for id field.
	projectDescID := projectFields[0].Descriptor()
	// project.DefaultID holds the default value on creation for the id field.
	project.DefaultID = projectDescID.Default.(func() string)
}
