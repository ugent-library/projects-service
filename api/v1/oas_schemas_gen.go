// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"time"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Merged schema.
// Ref: #/components/schemas/AddProject
type AddProject struct {
	ID              OptString                   `json:"id"`
	Created         OptDateTime                 `json:"created"`
	Modified        OptDateTime                 `json:"modified"`
	Type            string                      `json:"type"`
	Identifier      []AddProjectIdentifierItem  `json:"identifier"`
	IsFundedBy      OptAddProjectIsFundedBy     `json:"isFundedBy"`
	HasAcronym      []string                    `json:"hasAcronym"`
	Name            []AddProjectNameItem        `json:"name"`
	Description     []AddProjectDescriptionItem `json:"description"`
	FoundingDate    OptString                   `json:"foundingDate"`
	DissolutionDate OptString                   `json:"dissolutionDate"`
}

// GetID returns the value of ID.
func (s *AddProject) GetID() OptString {
	return s.ID
}

// GetCreated returns the value of Created.
func (s *AddProject) GetCreated() OptDateTime {
	return s.Created
}

// GetModified returns the value of Modified.
func (s *AddProject) GetModified() OptDateTime {
	return s.Modified
}

// GetType returns the value of Type.
func (s *AddProject) GetType() string {
	return s.Type
}

// GetIdentifier returns the value of Identifier.
func (s *AddProject) GetIdentifier() []AddProjectIdentifierItem {
	return s.Identifier
}

// GetIsFundedBy returns the value of IsFundedBy.
func (s *AddProject) GetIsFundedBy() OptAddProjectIsFundedBy {
	return s.IsFundedBy
}

// GetHasAcronym returns the value of HasAcronym.
func (s *AddProject) GetHasAcronym() []string {
	return s.HasAcronym
}

// GetName returns the value of Name.
func (s *AddProject) GetName() []AddProjectNameItem {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *AddProject) GetDescription() []AddProjectDescriptionItem {
	return s.Description
}

// GetFoundingDate returns the value of FoundingDate.
func (s *AddProject) GetFoundingDate() OptString {
	return s.FoundingDate
}

// GetDissolutionDate returns the value of DissolutionDate.
func (s *AddProject) GetDissolutionDate() OptString {
	return s.DissolutionDate
}

// SetID sets the value of ID.
func (s *AddProject) SetID(val OptString) {
	s.ID = val
}

// SetCreated sets the value of Created.
func (s *AddProject) SetCreated(val OptDateTime) {
	s.Created = val
}

// SetModified sets the value of Modified.
func (s *AddProject) SetModified(val OptDateTime) {
	s.Modified = val
}

// SetType sets the value of Type.
func (s *AddProject) SetType(val string) {
	s.Type = val
}

// SetIdentifier sets the value of Identifier.
func (s *AddProject) SetIdentifier(val []AddProjectIdentifierItem) {
	s.Identifier = val
}

// SetIsFundedBy sets the value of IsFundedBy.
func (s *AddProject) SetIsFundedBy(val OptAddProjectIsFundedBy) {
	s.IsFundedBy = val
}

// SetHasAcronym sets the value of HasAcronym.
func (s *AddProject) SetHasAcronym(val []string) {
	s.HasAcronym = val
}

// SetName sets the value of Name.
func (s *AddProject) SetName(val []AddProjectNameItem) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *AddProject) SetDescription(val []AddProjectDescriptionItem) {
	s.Description = val
}

// SetFoundingDate sets the value of FoundingDate.
func (s *AddProject) SetFoundingDate(val OptString) {
	s.FoundingDate = val
}

// SetDissolutionDate sets the value of DissolutionDate.
func (s *AddProject) SetDissolutionDate(val OptString) {
	s.DissolutionDate = val
}

type AddProjectDescriptionItem struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// GetLanguage returns the value of Language.
func (s *AddProjectDescriptionItem) GetLanguage() string {
	return s.Language
}

// GetValue returns the value of Value.
func (s *AddProjectDescriptionItem) GetValue() string {
	return s.Value
}

// SetLanguage sets the value of Language.
func (s *AddProjectDescriptionItem) SetLanguage(val string) {
	s.Language = val
}

// SetValue sets the value of Value.
func (s *AddProjectDescriptionItem) SetValue(val string) {
	s.Value = val
}

type AddProjectIdentifierItem struct {
	Type       string `json:"type"`
	PropertyID string `json:"propertyID"`
	Value      string `json:"value"`
}

// GetType returns the value of Type.
func (s *AddProjectIdentifierItem) GetType() string {
	return s.Type
}

// GetPropertyID returns the value of PropertyID.
func (s *AddProjectIdentifierItem) GetPropertyID() string {
	return s.PropertyID
}

// GetValue returns the value of Value.
func (s *AddProjectIdentifierItem) GetValue() string {
	return s.Value
}

// SetType sets the value of Type.
func (s *AddProjectIdentifierItem) SetType(val string) {
	s.Type = val
}

// SetPropertyID sets the value of PropertyID.
func (s *AddProjectIdentifierItem) SetPropertyID(val string) {
	s.PropertyID = val
}

// SetValue sets the value of Value.
func (s *AddProjectIdentifierItem) SetValue(val string) {
	s.Value = val
}

type AddProjectIsFundedBy struct {
	Type          string                                `json:"type"`
	HasCallNumber OptString                             `json:"hasCallNumber"`
	IsAwardedBy   OptNilAddProjectIsFundedByIsAwardedBy `json:"isAwardedBy"`
}

// GetType returns the value of Type.
func (s *AddProjectIsFundedBy) GetType() string {
	return s.Type
}

// GetHasCallNumber returns the value of HasCallNumber.
func (s *AddProjectIsFundedBy) GetHasCallNumber() OptString {
	return s.HasCallNumber
}

// GetIsAwardedBy returns the value of IsAwardedBy.
func (s *AddProjectIsFundedBy) GetIsAwardedBy() OptNilAddProjectIsFundedByIsAwardedBy {
	return s.IsAwardedBy
}

// SetType sets the value of Type.
func (s *AddProjectIsFundedBy) SetType(val string) {
	s.Type = val
}

// SetHasCallNumber sets the value of HasCallNumber.
func (s *AddProjectIsFundedBy) SetHasCallNumber(val OptString) {
	s.HasCallNumber = val
}

// SetIsAwardedBy sets the value of IsAwardedBy.
func (s *AddProjectIsFundedBy) SetIsAwardedBy(val OptNilAddProjectIsFundedByIsAwardedBy) {
	s.IsAwardedBy = val
}

type AddProjectIsFundedByIsAwardedBy struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// GetType returns the value of Type.
func (s *AddProjectIsFundedByIsAwardedBy) GetType() string {
	return s.Type
}

// GetName returns the value of Name.
func (s *AddProjectIsFundedByIsAwardedBy) GetName() string {
	return s.Name
}

// SetType sets the value of Type.
func (s *AddProjectIsFundedByIsAwardedBy) SetType(val string) {
	s.Type = val
}

// SetName sets the value of Name.
func (s *AddProjectIsFundedByIsAwardedBy) SetName(val string) {
	s.Name = val
}

type AddProjectNameItem struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// GetLanguage returns the value of Language.
func (s *AddProjectNameItem) GetLanguage() string {
	return s.Language
}

// GetValue returns the value of Value.
func (s *AddProjectNameItem) GetValue() string {
	return s.Value
}

// SetLanguage sets the value of Language.
func (s *AddProjectNameItem) SetLanguage(val string) {
	s.Language = val
}

// SetValue sets the value of Value.
func (s *AddProjectNameItem) SetValue(val string) {
	s.Value = val
}

// AddProjectOK is response for AddProject operation.
type AddProjectOK struct{}

type ApiKey struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *ApiKey) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *ApiKey) SetAPIKey(val string) {
	s.APIKey = val
}

// DeleteProjectOK is response for DeleteProject operation.
type DeleteProjectOK struct{}

// Ref: #/components/schemas/DeleteProjectRequest
type DeleteProjectRequest struct {
	ID string `json:"id"`
}

// GetID returns the value of ID.
func (s *DeleteProjectRequest) GetID() string {
	return s.ID
}

// SetID sets the value of ID.
func (s *DeleteProjectRequest) SetID(val string) {
	s.ID = val
}

// Ref: #/components/schemas/Error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

func (*ErrorStatusCode) getProjectRes() {}

// Merged schema.
// Ref: #/components/schemas/GetProject
type GetProject struct {
	ID              string                      `json:"id"`
	Created         time.Time                   `json:"created"`
	Modified        time.Time                   `json:"modified"`
	Type            string                      `json:"type"`
	Identifier      []GetProjectIdentifierItem  `json:"identifier"`
	IsFundedBy      OptGetProjectIsFundedBy     `json:"isFundedBy"`
	HasAcronym      []string                    `json:"hasAcronym"`
	Name            []GetProjectNameItem        `json:"name"`
	Description     []GetProjectDescriptionItem `json:"description"`
	FoundingDate    OptString                   `json:"foundingDate"`
	DissolutionDate OptString                   `json:"dissolutionDate"`
}

// GetID returns the value of ID.
func (s *GetProject) GetID() string {
	return s.ID
}

// GetCreated returns the value of Created.
func (s *GetProject) GetCreated() time.Time {
	return s.Created
}

// GetModified returns the value of Modified.
func (s *GetProject) GetModified() time.Time {
	return s.Modified
}

// GetType returns the value of Type.
func (s *GetProject) GetType() string {
	return s.Type
}

// GetIdentifier returns the value of Identifier.
func (s *GetProject) GetIdentifier() []GetProjectIdentifierItem {
	return s.Identifier
}

// GetIsFundedBy returns the value of IsFundedBy.
func (s *GetProject) GetIsFundedBy() OptGetProjectIsFundedBy {
	return s.IsFundedBy
}

// GetHasAcronym returns the value of HasAcronym.
func (s *GetProject) GetHasAcronym() []string {
	return s.HasAcronym
}

// GetName returns the value of Name.
func (s *GetProject) GetName() []GetProjectNameItem {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *GetProject) GetDescription() []GetProjectDescriptionItem {
	return s.Description
}

// GetFoundingDate returns the value of FoundingDate.
func (s *GetProject) GetFoundingDate() OptString {
	return s.FoundingDate
}

// GetDissolutionDate returns the value of DissolutionDate.
func (s *GetProject) GetDissolutionDate() OptString {
	return s.DissolutionDate
}

// SetID sets the value of ID.
func (s *GetProject) SetID(val string) {
	s.ID = val
}

// SetCreated sets the value of Created.
func (s *GetProject) SetCreated(val time.Time) {
	s.Created = val
}

// SetModified sets the value of Modified.
func (s *GetProject) SetModified(val time.Time) {
	s.Modified = val
}

// SetType sets the value of Type.
func (s *GetProject) SetType(val string) {
	s.Type = val
}

// SetIdentifier sets the value of Identifier.
func (s *GetProject) SetIdentifier(val []GetProjectIdentifierItem) {
	s.Identifier = val
}

// SetIsFundedBy sets the value of IsFundedBy.
func (s *GetProject) SetIsFundedBy(val OptGetProjectIsFundedBy) {
	s.IsFundedBy = val
}

// SetHasAcronym sets the value of HasAcronym.
func (s *GetProject) SetHasAcronym(val []string) {
	s.HasAcronym = val
}

// SetName sets the value of Name.
func (s *GetProject) SetName(val []GetProjectNameItem) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *GetProject) SetDescription(val []GetProjectDescriptionItem) {
	s.Description = val
}

// SetFoundingDate sets the value of FoundingDate.
func (s *GetProject) SetFoundingDate(val OptString) {
	s.FoundingDate = val
}

// SetDissolutionDate sets the value of DissolutionDate.
func (s *GetProject) SetDissolutionDate(val OptString) {
	s.DissolutionDate = val
}

func (*GetProject) getProjectRes() {}

type GetProjectDescriptionItem struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// GetLanguage returns the value of Language.
func (s *GetProjectDescriptionItem) GetLanguage() string {
	return s.Language
}

// GetValue returns the value of Value.
func (s *GetProjectDescriptionItem) GetValue() string {
	return s.Value
}

// SetLanguage sets the value of Language.
func (s *GetProjectDescriptionItem) SetLanguage(val string) {
	s.Language = val
}

// SetValue sets the value of Value.
func (s *GetProjectDescriptionItem) SetValue(val string) {
	s.Value = val
}

type GetProjectIdentifierItem struct {
	Type       string `json:"type"`
	PropertyID string `json:"propertyID"`
	Value      string `json:"value"`
}

// GetType returns the value of Type.
func (s *GetProjectIdentifierItem) GetType() string {
	return s.Type
}

// GetPropertyID returns the value of PropertyID.
func (s *GetProjectIdentifierItem) GetPropertyID() string {
	return s.PropertyID
}

// GetValue returns the value of Value.
func (s *GetProjectIdentifierItem) GetValue() string {
	return s.Value
}

// SetType sets the value of Type.
func (s *GetProjectIdentifierItem) SetType(val string) {
	s.Type = val
}

// SetPropertyID sets the value of PropertyID.
func (s *GetProjectIdentifierItem) SetPropertyID(val string) {
	s.PropertyID = val
}

// SetValue sets the value of Value.
func (s *GetProjectIdentifierItem) SetValue(val string) {
	s.Value = val
}

type GetProjectIsFundedBy struct {
	Type          string                                `json:"type"`
	HasCallNumber OptString                             `json:"hasCallNumber"`
	IsAwardedBy   OptNilGetProjectIsFundedByIsAwardedBy `json:"isAwardedBy"`
}

// GetType returns the value of Type.
func (s *GetProjectIsFundedBy) GetType() string {
	return s.Type
}

// GetHasCallNumber returns the value of HasCallNumber.
func (s *GetProjectIsFundedBy) GetHasCallNumber() OptString {
	return s.HasCallNumber
}

// GetIsAwardedBy returns the value of IsAwardedBy.
func (s *GetProjectIsFundedBy) GetIsAwardedBy() OptNilGetProjectIsFundedByIsAwardedBy {
	return s.IsAwardedBy
}

// SetType sets the value of Type.
func (s *GetProjectIsFundedBy) SetType(val string) {
	s.Type = val
}

// SetHasCallNumber sets the value of HasCallNumber.
func (s *GetProjectIsFundedBy) SetHasCallNumber(val OptString) {
	s.HasCallNumber = val
}

// SetIsAwardedBy sets the value of IsAwardedBy.
func (s *GetProjectIsFundedBy) SetIsAwardedBy(val OptNilGetProjectIsFundedByIsAwardedBy) {
	s.IsAwardedBy = val
}

type GetProjectIsFundedByIsAwardedBy struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// GetType returns the value of Type.
func (s *GetProjectIsFundedByIsAwardedBy) GetType() string {
	return s.Type
}

// GetName returns the value of Name.
func (s *GetProjectIsFundedByIsAwardedBy) GetName() string {
	return s.Name
}

// SetType sets the value of Type.
func (s *GetProjectIsFundedByIsAwardedBy) SetType(val string) {
	s.Type = val
}

// SetName sets the value of Name.
func (s *GetProjectIsFundedByIsAwardedBy) SetName(val string) {
	s.Name = val
}

type GetProjectNameItem struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// GetLanguage returns the value of Language.
func (s *GetProjectNameItem) GetLanguage() string {
	return s.Language
}

// GetValue returns the value of Value.
func (s *GetProjectNameItem) GetValue() string {
	return s.Value
}

// SetLanguage sets the value of Language.
func (s *GetProjectNameItem) SetLanguage(val string) {
	s.Language = val
}

// SetValue sets the value of Value.
func (s *GetProjectNameItem) SetValue(val string) {
	s.Value = val
}

// Ref: #/components/schemas/GetProjectRequest
type GetProjectRequest struct {
	ID string `json:"id"`
}

// GetID returns the value of ID.
func (s *GetProjectRequest) GetID() string {
	return s.ID
}

// SetID sets the value of ID.
func (s *GetProjectRequest) SetID(val string) {
	s.ID = val
}

// NewOptAddProjectIsFundedBy returns new OptAddProjectIsFundedBy with value set to v.
func NewOptAddProjectIsFundedBy(v AddProjectIsFundedBy) OptAddProjectIsFundedBy {
	return OptAddProjectIsFundedBy{
		Value: v,
		Set:   true,
	}
}

// OptAddProjectIsFundedBy is optional AddProjectIsFundedBy.
type OptAddProjectIsFundedBy struct {
	Value AddProjectIsFundedBy
	Set   bool
}

// IsSet returns true if OptAddProjectIsFundedBy was set.
func (o OptAddProjectIsFundedBy) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptAddProjectIsFundedBy) Reset() {
	var v AddProjectIsFundedBy
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptAddProjectIsFundedBy) SetTo(v AddProjectIsFundedBy) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptAddProjectIsFundedBy) Get() (v AddProjectIsFundedBy, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptAddProjectIsFundedBy) Or(d AddProjectIsFundedBy) AddProjectIsFundedBy {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptGetProjectIsFundedBy returns new OptGetProjectIsFundedBy with value set to v.
func NewOptGetProjectIsFundedBy(v GetProjectIsFundedBy) OptGetProjectIsFundedBy {
	return OptGetProjectIsFundedBy{
		Value: v,
		Set:   true,
	}
}

// OptGetProjectIsFundedBy is optional GetProjectIsFundedBy.
type OptGetProjectIsFundedBy struct {
	Value GetProjectIsFundedBy
	Set   bool
}

// IsSet returns true if OptGetProjectIsFundedBy was set.
func (o OptGetProjectIsFundedBy) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptGetProjectIsFundedBy) Reset() {
	var v GetProjectIsFundedBy
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptGetProjectIsFundedBy) SetTo(v GetProjectIsFundedBy) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptGetProjectIsFundedBy) Get() (v GetProjectIsFundedBy, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptGetProjectIsFundedBy) Or(d GetProjectIsFundedBy) GetProjectIsFundedBy {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilAddProjectIsFundedByIsAwardedBy returns new OptNilAddProjectIsFundedByIsAwardedBy with value set to v.
func NewOptNilAddProjectIsFundedByIsAwardedBy(v AddProjectIsFundedByIsAwardedBy) OptNilAddProjectIsFundedByIsAwardedBy {
	return OptNilAddProjectIsFundedByIsAwardedBy{
		Value: v,
		Set:   true,
	}
}

// OptNilAddProjectIsFundedByIsAwardedBy is optional nullable AddProjectIsFundedByIsAwardedBy.
type OptNilAddProjectIsFundedByIsAwardedBy struct {
	Value AddProjectIsFundedByIsAwardedBy
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilAddProjectIsFundedByIsAwardedBy was set.
func (o OptNilAddProjectIsFundedByIsAwardedBy) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilAddProjectIsFundedByIsAwardedBy) Reset() {
	var v AddProjectIsFundedByIsAwardedBy
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilAddProjectIsFundedByIsAwardedBy) SetTo(v AddProjectIsFundedByIsAwardedBy) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o OptNilAddProjectIsFundedByIsAwardedBy) IsNull() bool { return o.Null }

// SetNull sets value to null.
func (o *OptNilAddProjectIsFundedByIsAwardedBy) SetToNull() {
	o.Set = true
	o.Null = true
	var v AddProjectIsFundedByIsAwardedBy
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilAddProjectIsFundedByIsAwardedBy) Get() (v AddProjectIsFundedByIsAwardedBy, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilAddProjectIsFundedByIsAwardedBy) Or(d AddProjectIsFundedByIsAwardedBy) AddProjectIsFundedByIsAwardedBy {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilGetProjectIsFundedByIsAwardedBy returns new OptNilGetProjectIsFundedByIsAwardedBy with value set to v.
func NewOptNilGetProjectIsFundedByIsAwardedBy(v GetProjectIsFundedByIsAwardedBy) OptNilGetProjectIsFundedByIsAwardedBy {
	return OptNilGetProjectIsFundedByIsAwardedBy{
		Value: v,
		Set:   true,
	}
}

// OptNilGetProjectIsFundedByIsAwardedBy is optional nullable GetProjectIsFundedByIsAwardedBy.
type OptNilGetProjectIsFundedByIsAwardedBy struct {
	Value GetProjectIsFundedByIsAwardedBy
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilGetProjectIsFundedByIsAwardedBy was set.
func (o OptNilGetProjectIsFundedByIsAwardedBy) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilGetProjectIsFundedByIsAwardedBy) Reset() {
	var v GetProjectIsFundedByIsAwardedBy
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilGetProjectIsFundedByIsAwardedBy) SetTo(v GetProjectIsFundedByIsAwardedBy) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o OptNilGetProjectIsFundedByIsAwardedBy) IsNull() bool { return o.Null }

// SetNull sets value to null.
func (o *OptNilGetProjectIsFundedByIsAwardedBy) SetToNull() {
	o.Set = true
	o.Null = true
	var v GetProjectIsFundedByIsAwardedBy
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilGetProjectIsFundedByIsAwardedBy) Get() (v GetProjectIsFundedByIsAwardedBy, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilGetProjectIsFundedByIsAwardedBy) Or(d GetProjectIsFundedByIsAwardedBy) GetProjectIsFundedByIsAwardedBy {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/SuggestProjectsRequest
type SuggestProjectsRequest struct {
	Query string `json:"query"`
}

// GetQuery returns the value of Query.
func (s *SuggestProjectsRequest) GetQuery() string {
	return s.Query
}

// SetQuery sets the value of Query.
func (s *SuggestProjectsRequest) SetQuery(val string) {
	s.Query = val
}

// Ref: #/components/schemas/SuggestProjectsResponse
type SuggestProjectsResponse struct {
	Cursor OptString    `json:"cursor"`
	Data   []GetProject `json:"data"`
}

// GetCursor returns the value of Cursor.
func (s *SuggestProjectsResponse) GetCursor() OptString {
	return s.Cursor
}

// GetData returns the value of Data.
func (s *SuggestProjectsResponse) GetData() []GetProject {
	return s.Data
}

// SetCursor sets the value of Cursor.
func (s *SuggestProjectsResponse) SetCursor(val OptString) {
	s.Cursor = val
}

// SetData sets the value of Data.
func (s *SuggestProjectsResponse) SetData(val []GetProject) {
	s.Data = val
}
