// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/pkg/database/ent/predicate"
	"app/pkg/database/ent/user"
	"app/pkg/database/ent/userprofile"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserProfileUpdate is the builder for updating UserProfile entities.
type UserProfileUpdate struct {
	config
	hooks    []Hook
	mutation *UserProfileMutation
}

// Where appends a list predicates to the UserProfileUpdate builder.
func (upu *UserProfileUpdate) Where(ps ...predicate.UserProfile) *UserProfileUpdate {
	upu.mutation.Where(ps...)
	return upu
}

// SetUserID sets the "user_id" field.
func (upu *UserProfileUpdate) SetUserID(s string) *UserProfileUpdate {
	upu.mutation.SetUserID(s)
	return upu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableUserID(s *string) *UserProfileUpdate {
	if s != nil {
		upu.SetUserID(*s)
	}
	return upu
}

// ClearUserID clears the value of the "user_id" field.
func (upu *UserProfileUpdate) ClearUserID() *UserProfileUpdate {
	upu.mutation.ClearUserID()
	return upu
}

// SetBio sets the "bio" field.
func (upu *UserProfileUpdate) SetBio(s string) *UserProfileUpdate {
	upu.mutation.SetBio(s)
	return upu
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableBio(s *string) *UserProfileUpdate {
	if s != nil {
		upu.SetBio(*s)
	}
	return upu
}

// ClearBio clears the value of the "bio" field.
func (upu *UserProfileUpdate) ClearBio() *UserProfileUpdate {
	upu.mutation.ClearBio()
	return upu
}

// SetWebsite sets the "website" field.
func (upu *UserProfileUpdate) SetWebsite(s string) *UserProfileUpdate {
	upu.mutation.SetWebsite(s)
	return upu
}

// SetNillableWebsite sets the "website" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableWebsite(s *string) *UserProfileUpdate {
	if s != nil {
		upu.SetWebsite(*s)
	}
	return upu
}

// ClearWebsite clears the value of the "website" field.
func (upu *UserProfileUpdate) ClearWebsite() *UserProfileUpdate {
	upu.mutation.ClearWebsite()
	return upu
}

// SetLocation sets the "location" field.
func (upu *UserProfileUpdate) SetLocation(s string) *UserProfileUpdate {
	upu.mutation.SetLocation(s)
	return upu
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableLocation(s *string) *UserProfileUpdate {
	if s != nil {
		upu.SetLocation(*s)
	}
	return upu
}

// ClearLocation clears the value of the "location" field.
func (upu *UserProfileUpdate) ClearLocation() *UserProfileUpdate {
	upu.mutation.ClearLocation()
	return upu
}

// SetTimezone sets the "timezone" field.
func (upu *UserProfileUpdate) SetTimezone(s string) *UserProfileUpdate {
	upu.mutation.SetTimezone(s)
	return upu
}

// SetNillableTimezone sets the "timezone" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableTimezone(s *string) *UserProfileUpdate {
	if s != nil {
		upu.SetTimezone(*s)
	}
	return upu
}

// ClearTimezone clears the value of the "timezone" field.
func (upu *UserProfileUpdate) ClearTimezone() *UserProfileUpdate {
	upu.mutation.ClearTimezone()
	return upu
}

// SetLanguage sets the "language" field.
func (upu *UserProfileUpdate) SetLanguage(s string) *UserProfileUpdate {
	upu.mutation.SetLanguage(s)
	return upu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableLanguage(s *string) *UserProfileUpdate {
	if s != nil {
		upu.SetLanguage(*s)
	}
	return upu
}

// ClearLanguage clears the value of the "language" field.
func (upu *UserProfileUpdate) ClearLanguage() *UserProfileUpdate {
	upu.mutation.ClearLanguage()
	return upu
}

// SetGender sets the "gender" field.
func (upu *UserProfileUpdate) SetGender(u userprofile.Gender) *UserProfileUpdate {
	upu.mutation.SetGender(u)
	return upu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableGender(u *userprofile.Gender) *UserProfileUpdate {
	if u != nil {
		upu.SetGender(*u)
	}
	return upu
}

// ClearGender clears the value of the "gender" field.
func (upu *UserProfileUpdate) ClearGender() *UserProfileUpdate {
	upu.mutation.ClearGender()
	return upu
}

// SetProfileVisibility sets the "profile_visibility" field.
func (upu *UserProfileUpdate) SetProfileVisibility(uv userprofile.ProfileVisibility) *UserProfileUpdate {
	upu.mutation.SetProfileVisibility(uv)
	return upu
}

// SetNillableProfileVisibility sets the "profile_visibility" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableProfileVisibility(uv *userprofile.ProfileVisibility) *UserProfileUpdate {
	if uv != nil {
		upu.SetProfileVisibility(*uv)
	}
	return upu
}

// ClearProfileVisibility clears the value of the "profile_visibility" field.
func (upu *UserProfileUpdate) ClearProfileVisibility() *UserProfileUpdate {
	upu.mutation.ClearProfileVisibility()
	return upu
}

// SetCreatedAt sets the "created_at" field.
func (upu *UserProfileUpdate) SetCreatedAt(t time.Time) *UserProfileUpdate {
	upu.mutation.SetCreatedAt(t)
	return upu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableCreatedAt(t *time.Time) *UserProfileUpdate {
	if t != nil {
		upu.SetCreatedAt(*t)
	}
	return upu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (upu *UserProfileUpdate) ClearCreatedAt() *UserProfileUpdate {
	upu.mutation.ClearCreatedAt()
	return upu
}

// SetUpdatedAt sets the "updated_at" field.
func (upu *UserProfileUpdate) SetUpdatedAt(t time.Time) *UserProfileUpdate {
	upu.mutation.SetUpdatedAt(t)
	return upu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableUpdatedAt(t *time.Time) *UserProfileUpdate {
	if t != nil {
		upu.SetUpdatedAt(*t)
	}
	return upu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (upu *UserProfileUpdate) ClearUpdatedAt() *UserProfileUpdate {
	upu.mutation.ClearUpdatedAt()
	return upu
}

// SetUser sets the "user" edge to the User entity.
func (upu *UserProfileUpdate) SetUser(u *User) *UserProfileUpdate {
	return upu.SetUserID(u.ID)
}

// Mutation returns the UserProfileMutation object of the builder.
func (upu *UserProfileUpdate) Mutation() *UserProfileMutation {
	return upu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (upu *UserProfileUpdate) ClearUser() *UserProfileUpdate {
	upu.mutation.ClearUser()
	return upu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upu *UserProfileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, upu.sqlSave, upu.mutation, upu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upu *UserProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := upu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upu *UserProfileUpdate) Exec(ctx context.Context) error {
	_, err := upu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upu *UserProfileUpdate) ExecX(ctx context.Context) {
	if err := upu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upu *UserProfileUpdate) check() error {
	if v, ok := upu.mutation.Gender(); ok {
		if err := userprofile.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "UserProfile.gender": %w`, err)}
		}
	}
	if v, ok := upu.mutation.ProfileVisibility(); ok {
		if err := userprofile.ProfileVisibilityValidator(v); err != nil {
			return &ValidationError{Name: "profile_visibility", err: fmt.Errorf(`ent: validator failed for field "UserProfile.profile_visibility": %w`, err)}
		}
	}
	return nil
}

func (upu *UserProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := upu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userprofile.Table, userprofile.Columns, sqlgraph.NewFieldSpec(userprofile.FieldID, field.TypeString))
	if ps := upu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upu.mutation.Bio(); ok {
		_spec.SetField(userprofile.FieldBio, field.TypeString, value)
	}
	if upu.mutation.BioCleared() {
		_spec.ClearField(userprofile.FieldBio, field.TypeString)
	}
	if value, ok := upu.mutation.Website(); ok {
		_spec.SetField(userprofile.FieldWebsite, field.TypeString, value)
	}
	if upu.mutation.WebsiteCleared() {
		_spec.ClearField(userprofile.FieldWebsite, field.TypeString)
	}
	if value, ok := upu.mutation.Location(); ok {
		_spec.SetField(userprofile.FieldLocation, field.TypeString, value)
	}
	if upu.mutation.LocationCleared() {
		_spec.ClearField(userprofile.FieldLocation, field.TypeString)
	}
	if value, ok := upu.mutation.Timezone(); ok {
		_spec.SetField(userprofile.FieldTimezone, field.TypeString, value)
	}
	if upu.mutation.TimezoneCleared() {
		_spec.ClearField(userprofile.FieldTimezone, field.TypeString)
	}
	if value, ok := upu.mutation.Language(); ok {
		_spec.SetField(userprofile.FieldLanguage, field.TypeString, value)
	}
	if upu.mutation.LanguageCleared() {
		_spec.ClearField(userprofile.FieldLanguage, field.TypeString)
	}
	if value, ok := upu.mutation.Gender(); ok {
		_spec.SetField(userprofile.FieldGender, field.TypeEnum, value)
	}
	if upu.mutation.GenderCleared() {
		_spec.ClearField(userprofile.FieldGender, field.TypeEnum)
	}
	if value, ok := upu.mutation.ProfileVisibility(); ok {
		_spec.SetField(userprofile.FieldProfileVisibility, field.TypeEnum, value)
	}
	if upu.mutation.ProfileVisibilityCleared() {
		_spec.ClearField(userprofile.FieldProfileVisibility, field.TypeEnum)
	}
	if value, ok := upu.mutation.CreatedAt(); ok {
		_spec.SetField(userprofile.FieldCreatedAt, field.TypeTime, value)
	}
	if upu.mutation.CreatedAtCleared() {
		_spec.ClearField(userprofile.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := upu.mutation.UpdatedAt(); ok {
		_spec.SetField(userprofile.FieldUpdatedAt, field.TypeTime, value)
	}
	if upu.mutation.UpdatedAtCleared() {
		_spec.ClearField(userprofile.FieldUpdatedAt, field.TypeTime)
	}
	if upu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.UserTable,
			Columns: []string{userprofile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.UserTable,
			Columns: []string{userprofile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, upu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	upu.mutation.done = true
	return n, nil
}

// UserProfileUpdateOne is the builder for updating a single UserProfile entity.
type UserProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserProfileMutation
}

// SetUserID sets the "user_id" field.
func (upuo *UserProfileUpdateOne) SetUserID(s string) *UserProfileUpdateOne {
	upuo.mutation.SetUserID(s)
	return upuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableUserID(s *string) *UserProfileUpdateOne {
	if s != nil {
		upuo.SetUserID(*s)
	}
	return upuo
}

// ClearUserID clears the value of the "user_id" field.
func (upuo *UserProfileUpdateOne) ClearUserID() *UserProfileUpdateOne {
	upuo.mutation.ClearUserID()
	return upuo
}

// SetBio sets the "bio" field.
func (upuo *UserProfileUpdateOne) SetBio(s string) *UserProfileUpdateOne {
	upuo.mutation.SetBio(s)
	return upuo
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableBio(s *string) *UserProfileUpdateOne {
	if s != nil {
		upuo.SetBio(*s)
	}
	return upuo
}

// ClearBio clears the value of the "bio" field.
func (upuo *UserProfileUpdateOne) ClearBio() *UserProfileUpdateOne {
	upuo.mutation.ClearBio()
	return upuo
}

// SetWebsite sets the "website" field.
func (upuo *UserProfileUpdateOne) SetWebsite(s string) *UserProfileUpdateOne {
	upuo.mutation.SetWebsite(s)
	return upuo
}

// SetNillableWebsite sets the "website" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableWebsite(s *string) *UserProfileUpdateOne {
	if s != nil {
		upuo.SetWebsite(*s)
	}
	return upuo
}

// ClearWebsite clears the value of the "website" field.
func (upuo *UserProfileUpdateOne) ClearWebsite() *UserProfileUpdateOne {
	upuo.mutation.ClearWebsite()
	return upuo
}

// SetLocation sets the "location" field.
func (upuo *UserProfileUpdateOne) SetLocation(s string) *UserProfileUpdateOne {
	upuo.mutation.SetLocation(s)
	return upuo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableLocation(s *string) *UserProfileUpdateOne {
	if s != nil {
		upuo.SetLocation(*s)
	}
	return upuo
}

// ClearLocation clears the value of the "location" field.
func (upuo *UserProfileUpdateOne) ClearLocation() *UserProfileUpdateOne {
	upuo.mutation.ClearLocation()
	return upuo
}

// SetTimezone sets the "timezone" field.
func (upuo *UserProfileUpdateOne) SetTimezone(s string) *UserProfileUpdateOne {
	upuo.mutation.SetTimezone(s)
	return upuo
}

// SetNillableTimezone sets the "timezone" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableTimezone(s *string) *UserProfileUpdateOne {
	if s != nil {
		upuo.SetTimezone(*s)
	}
	return upuo
}

// ClearTimezone clears the value of the "timezone" field.
func (upuo *UserProfileUpdateOne) ClearTimezone() *UserProfileUpdateOne {
	upuo.mutation.ClearTimezone()
	return upuo
}

// SetLanguage sets the "language" field.
func (upuo *UserProfileUpdateOne) SetLanguage(s string) *UserProfileUpdateOne {
	upuo.mutation.SetLanguage(s)
	return upuo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableLanguage(s *string) *UserProfileUpdateOne {
	if s != nil {
		upuo.SetLanguage(*s)
	}
	return upuo
}

// ClearLanguage clears the value of the "language" field.
func (upuo *UserProfileUpdateOne) ClearLanguage() *UserProfileUpdateOne {
	upuo.mutation.ClearLanguage()
	return upuo
}

// SetGender sets the "gender" field.
func (upuo *UserProfileUpdateOne) SetGender(u userprofile.Gender) *UserProfileUpdateOne {
	upuo.mutation.SetGender(u)
	return upuo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableGender(u *userprofile.Gender) *UserProfileUpdateOne {
	if u != nil {
		upuo.SetGender(*u)
	}
	return upuo
}

// ClearGender clears the value of the "gender" field.
func (upuo *UserProfileUpdateOne) ClearGender() *UserProfileUpdateOne {
	upuo.mutation.ClearGender()
	return upuo
}

// SetProfileVisibility sets the "profile_visibility" field.
func (upuo *UserProfileUpdateOne) SetProfileVisibility(uv userprofile.ProfileVisibility) *UserProfileUpdateOne {
	upuo.mutation.SetProfileVisibility(uv)
	return upuo
}

// SetNillableProfileVisibility sets the "profile_visibility" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableProfileVisibility(uv *userprofile.ProfileVisibility) *UserProfileUpdateOne {
	if uv != nil {
		upuo.SetProfileVisibility(*uv)
	}
	return upuo
}

// ClearProfileVisibility clears the value of the "profile_visibility" field.
func (upuo *UserProfileUpdateOne) ClearProfileVisibility() *UserProfileUpdateOne {
	upuo.mutation.ClearProfileVisibility()
	return upuo
}

// SetCreatedAt sets the "created_at" field.
func (upuo *UserProfileUpdateOne) SetCreatedAt(t time.Time) *UserProfileUpdateOne {
	upuo.mutation.SetCreatedAt(t)
	return upuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableCreatedAt(t *time.Time) *UserProfileUpdateOne {
	if t != nil {
		upuo.SetCreatedAt(*t)
	}
	return upuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (upuo *UserProfileUpdateOne) ClearCreatedAt() *UserProfileUpdateOne {
	upuo.mutation.ClearCreatedAt()
	return upuo
}

// SetUpdatedAt sets the "updated_at" field.
func (upuo *UserProfileUpdateOne) SetUpdatedAt(t time.Time) *UserProfileUpdateOne {
	upuo.mutation.SetUpdatedAt(t)
	return upuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserProfileUpdateOne {
	if t != nil {
		upuo.SetUpdatedAt(*t)
	}
	return upuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (upuo *UserProfileUpdateOne) ClearUpdatedAt() *UserProfileUpdateOne {
	upuo.mutation.ClearUpdatedAt()
	return upuo
}

// SetUser sets the "user" edge to the User entity.
func (upuo *UserProfileUpdateOne) SetUser(u *User) *UserProfileUpdateOne {
	return upuo.SetUserID(u.ID)
}

// Mutation returns the UserProfileMutation object of the builder.
func (upuo *UserProfileUpdateOne) Mutation() *UserProfileMutation {
	return upuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (upuo *UserProfileUpdateOne) ClearUser() *UserProfileUpdateOne {
	upuo.mutation.ClearUser()
	return upuo
}

// Where appends a list predicates to the UserProfileUpdate builder.
func (upuo *UserProfileUpdateOne) Where(ps ...predicate.UserProfile) *UserProfileUpdateOne {
	upuo.mutation.Where(ps...)
	return upuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upuo *UserProfileUpdateOne) Select(field string, fields ...string) *UserProfileUpdateOne {
	upuo.fields = append([]string{field}, fields...)
	return upuo
}

// Save executes the query and returns the updated UserProfile entity.
func (upuo *UserProfileUpdateOne) Save(ctx context.Context) (*UserProfile, error) {
	return withHooks(ctx, upuo.sqlSave, upuo.mutation, upuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upuo *UserProfileUpdateOne) SaveX(ctx context.Context) *UserProfile {
	node, err := upuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upuo *UserProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := upuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upuo *UserProfileUpdateOne) ExecX(ctx context.Context) {
	if err := upuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upuo *UserProfileUpdateOne) check() error {
	if v, ok := upuo.mutation.Gender(); ok {
		if err := userprofile.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "UserProfile.gender": %w`, err)}
		}
	}
	if v, ok := upuo.mutation.ProfileVisibility(); ok {
		if err := userprofile.ProfileVisibilityValidator(v); err != nil {
			return &ValidationError{Name: "profile_visibility", err: fmt.Errorf(`ent: validator failed for field "UserProfile.profile_visibility": %w`, err)}
		}
	}
	return nil
}

func (upuo *UserProfileUpdateOne) sqlSave(ctx context.Context) (_node *UserProfile, err error) {
	if err := upuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userprofile.Table, userprofile.Columns, sqlgraph.NewFieldSpec(userprofile.FieldID, field.TypeString))
	id, ok := upuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserProfile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := upuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userprofile.FieldID)
		for _, f := range fields {
			if !userprofile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upuo.mutation.Bio(); ok {
		_spec.SetField(userprofile.FieldBio, field.TypeString, value)
	}
	if upuo.mutation.BioCleared() {
		_spec.ClearField(userprofile.FieldBio, field.TypeString)
	}
	if value, ok := upuo.mutation.Website(); ok {
		_spec.SetField(userprofile.FieldWebsite, field.TypeString, value)
	}
	if upuo.mutation.WebsiteCleared() {
		_spec.ClearField(userprofile.FieldWebsite, field.TypeString)
	}
	if value, ok := upuo.mutation.Location(); ok {
		_spec.SetField(userprofile.FieldLocation, field.TypeString, value)
	}
	if upuo.mutation.LocationCleared() {
		_spec.ClearField(userprofile.FieldLocation, field.TypeString)
	}
	if value, ok := upuo.mutation.Timezone(); ok {
		_spec.SetField(userprofile.FieldTimezone, field.TypeString, value)
	}
	if upuo.mutation.TimezoneCleared() {
		_spec.ClearField(userprofile.FieldTimezone, field.TypeString)
	}
	if value, ok := upuo.mutation.Language(); ok {
		_spec.SetField(userprofile.FieldLanguage, field.TypeString, value)
	}
	if upuo.mutation.LanguageCleared() {
		_spec.ClearField(userprofile.FieldLanguage, field.TypeString)
	}
	if value, ok := upuo.mutation.Gender(); ok {
		_spec.SetField(userprofile.FieldGender, field.TypeEnum, value)
	}
	if upuo.mutation.GenderCleared() {
		_spec.ClearField(userprofile.FieldGender, field.TypeEnum)
	}
	if value, ok := upuo.mutation.ProfileVisibility(); ok {
		_spec.SetField(userprofile.FieldProfileVisibility, field.TypeEnum, value)
	}
	if upuo.mutation.ProfileVisibilityCleared() {
		_spec.ClearField(userprofile.FieldProfileVisibility, field.TypeEnum)
	}
	if value, ok := upuo.mutation.CreatedAt(); ok {
		_spec.SetField(userprofile.FieldCreatedAt, field.TypeTime, value)
	}
	if upuo.mutation.CreatedAtCleared() {
		_spec.ClearField(userprofile.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := upuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userprofile.FieldUpdatedAt, field.TypeTime, value)
	}
	if upuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(userprofile.FieldUpdatedAt, field.TypeTime)
	}
	if upuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.UserTable,
			Columns: []string{userprofile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.UserTable,
			Columns: []string{userprofile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserProfile{config: upuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	upuo.mutation.done = true
	return _node, nil
}
