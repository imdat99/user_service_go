// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/pkg/database/ent/notificationsetting"
	"app/pkg/database/ent/predicate"
	"app/pkg/database/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationSettingUpdate is the builder for updating NotificationSetting entities.
type NotificationSettingUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationSettingMutation
}

// Where appends a list predicates to the NotificationSettingUpdate builder.
func (nsu *NotificationSettingUpdate) Where(ps ...predicate.NotificationSetting) *NotificationSettingUpdate {
	nsu.mutation.Where(ps...)
	return nsu
}

// SetUserID sets the "user_id" field.
func (nsu *NotificationSettingUpdate) SetUserID(s string) *NotificationSettingUpdate {
	nsu.mutation.SetUserID(s)
	return nsu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillableUserID(s *string) *NotificationSettingUpdate {
	if s != nil {
		nsu.SetUserID(*s)
	}
	return nsu
}

// ClearUserID clears the value of the "user_id" field.
func (nsu *NotificationSettingUpdate) ClearUserID() *NotificationSettingUpdate {
	nsu.mutation.ClearUserID()
	return nsu
}

// SetEmailNotifications sets the "email_notifications" field.
func (nsu *NotificationSettingUpdate) SetEmailNotifications(b bool) *NotificationSettingUpdate {
	nsu.mutation.SetEmailNotifications(b)
	return nsu
}

// SetNillableEmailNotifications sets the "email_notifications" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillableEmailNotifications(b *bool) *NotificationSettingUpdate {
	if b != nil {
		nsu.SetEmailNotifications(*b)
	}
	return nsu
}

// ClearEmailNotifications clears the value of the "email_notifications" field.
func (nsu *NotificationSettingUpdate) ClearEmailNotifications() *NotificationSettingUpdate {
	nsu.mutation.ClearEmailNotifications()
	return nsu
}

// SetSmsNotifications sets the "sms_notifications" field.
func (nsu *NotificationSettingUpdate) SetSmsNotifications(b bool) *NotificationSettingUpdate {
	nsu.mutation.SetSmsNotifications(b)
	return nsu
}

// SetNillableSmsNotifications sets the "sms_notifications" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillableSmsNotifications(b *bool) *NotificationSettingUpdate {
	if b != nil {
		nsu.SetSmsNotifications(*b)
	}
	return nsu
}

// ClearSmsNotifications clears the value of the "sms_notifications" field.
func (nsu *NotificationSettingUpdate) ClearSmsNotifications() *NotificationSettingUpdate {
	nsu.mutation.ClearSmsNotifications()
	return nsu
}

// SetPushNotifications sets the "push_notifications" field.
func (nsu *NotificationSettingUpdate) SetPushNotifications(b bool) *NotificationSettingUpdate {
	nsu.mutation.SetPushNotifications(b)
	return nsu
}

// SetNillablePushNotifications sets the "push_notifications" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillablePushNotifications(b *bool) *NotificationSettingUpdate {
	if b != nil {
		nsu.SetPushNotifications(*b)
	}
	return nsu
}

// ClearPushNotifications clears the value of the "push_notifications" field.
func (nsu *NotificationSettingUpdate) ClearPushNotifications() *NotificationSettingUpdate {
	nsu.mutation.ClearPushNotifications()
	return nsu
}

// SetMarketingEmails sets the "marketing_emails" field.
func (nsu *NotificationSettingUpdate) SetMarketingEmails(b bool) *NotificationSettingUpdate {
	nsu.mutation.SetMarketingEmails(b)
	return nsu
}

// SetNillableMarketingEmails sets the "marketing_emails" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillableMarketingEmails(b *bool) *NotificationSettingUpdate {
	if b != nil {
		nsu.SetMarketingEmails(*b)
	}
	return nsu
}

// ClearMarketingEmails clears the value of the "marketing_emails" field.
func (nsu *NotificationSettingUpdate) ClearMarketingEmails() *NotificationSettingUpdate {
	nsu.mutation.ClearMarketingEmails()
	return nsu
}

// SetSecurityAlerts sets the "security_alerts" field.
func (nsu *NotificationSettingUpdate) SetSecurityAlerts(b bool) *NotificationSettingUpdate {
	nsu.mutation.SetSecurityAlerts(b)
	return nsu
}

// SetNillableSecurityAlerts sets the "security_alerts" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillableSecurityAlerts(b *bool) *NotificationSettingUpdate {
	if b != nil {
		nsu.SetSecurityAlerts(*b)
	}
	return nsu
}

// ClearSecurityAlerts clears the value of the "security_alerts" field.
func (nsu *NotificationSettingUpdate) ClearSecurityAlerts() *NotificationSettingUpdate {
	nsu.mutation.ClearSecurityAlerts()
	return nsu
}

// SetLoginAlerts sets the "login_alerts" field.
func (nsu *NotificationSettingUpdate) SetLoginAlerts(b bool) *NotificationSettingUpdate {
	nsu.mutation.SetLoginAlerts(b)
	return nsu
}

// SetNillableLoginAlerts sets the "login_alerts" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillableLoginAlerts(b *bool) *NotificationSettingUpdate {
	if b != nil {
		nsu.SetLoginAlerts(*b)
	}
	return nsu
}

// ClearLoginAlerts clears the value of the "login_alerts" field.
func (nsu *NotificationSettingUpdate) ClearLoginAlerts() *NotificationSettingUpdate {
	nsu.mutation.ClearLoginAlerts()
	return nsu
}

// SetProfileUpdates sets the "profile_updates" field.
func (nsu *NotificationSettingUpdate) SetProfileUpdates(b bool) *NotificationSettingUpdate {
	nsu.mutation.SetProfileUpdates(b)
	return nsu
}

// SetNillableProfileUpdates sets the "profile_updates" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillableProfileUpdates(b *bool) *NotificationSettingUpdate {
	if b != nil {
		nsu.SetProfileUpdates(*b)
	}
	return nsu
}

// ClearProfileUpdates clears the value of the "profile_updates" field.
func (nsu *NotificationSettingUpdate) ClearProfileUpdates() *NotificationSettingUpdate {
	nsu.mutation.ClearProfileUpdates()
	return nsu
}

// SetPaymentNotifications sets the "payment_notifications" field.
func (nsu *NotificationSettingUpdate) SetPaymentNotifications(b bool) *NotificationSettingUpdate {
	nsu.mutation.SetPaymentNotifications(b)
	return nsu
}

// SetNillablePaymentNotifications sets the "payment_notifications" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillablePaymentNotifications(b *bool) *NotificationSettingUpdate {
	if b != nil {
		nsu.SetPaymentNotifications(*b)
	}
	return nsu
}

// ClearPaymentNotifications clears the value of the "payment_notifications" field.
func (nsu *NotificationSettingUpdate) ClearPaymentNotifications() *NotificationSettingUpdate {
	nsu.mutation.ClearPaymentNotifications()
	return nsu
}

// SetCreatedAt sets the "created_at" field.
func (nsu *NotificationSettingUpdate) SetCreatedAt(t time.Time) *NotificationSettingUpdate {
	nsu.mutation.SetCreatedAt(t)
	return nsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillableCreatedAt(t *time.Time) *NotificationSettingUpdate {
	if t != nil {
		nsu.SetCreatedAt(*t)
	}
	return nsu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (nsu *NotificationSettingUpdate) ClearCreatedAt() *NotificationSettingUpdate {
	nsu.mutation.ClearCreatedAt()
	return nsu
}

// SetUpdatedAt sets the "updated_at" field.
func (nsu *NotificationSettingUpdate) SetUpdatedAt(t time.Time) *NotificationSettingUpdate {
	nsu.mutation.SetUpdatedAt(t)
	return nsu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nsu *NotificationSettingUpdate) SetNillableUpdatedAt(t *time.Time) *NotificationSettingUpdate {
	if t != nil {
		nsu.SetUpdatedAt(*t)
	}
	return nsu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nsu *NotificationSettingUpdate) ClearUpdatedAt() *NotificationSettingUpdate {
	nsu.mutation.ClearUpdatedAt()
	return nsu
}

// SetUser sets the "user" edge to the User entity.
func (nsu *NotificationSettingUpdate) SetUser(u *User) *NotificationSettingUpdate {
	return nsu.SetUserID(u.ID)
}

// Mutation returns the NotificationSettingMutation object of the builder.
func (nsu *NotificationSettingUpdate) Mutation() *NotificationSettingMutation {
	return nsu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (nsu *NotificationSettingUpdate) ClearUser() *NotificationSettingUpdate {
	nsu.mutation.ClearUser()
	return nsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nsu *NotificationSettingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nsu.sqlSave, nsu.mutation, nsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nsu *NotificationSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := nsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nsu *NotificationSettingUpdate) Exec(ctx context.Context) error {
	_, err := nsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsu *NotificationSettingUpdate) ExecX(ctx context.Context) {
	if err := nsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nsu *NotificationSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(notificationsetting.Table, notificationsetting.Columns, sqlgraph.NewFieldSpec(notificationsetting.FieldID, field.TypeString))
	if ps := nsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nsu.mutation.EmailNotifications(); ok {
		_spec.SetField(notificationsetting.FieldEmailNotifications, field.TypeBool, value)
	}
	if nsu.mutation.EmailNotificationsCleared() {
		_spec.ClearField(notificationsetting.FieldEmailNotifications, field.TypeBool)
	}
	if value, ok := nsu.mutation.SmsNotifications(); ok {
		_spec.SetField(notificationsetting.FieldSmsNotifications, field.TypeBool, value)
	}
	if nsu.mutation.SmsNotificationsCleared() {
		_spec.ClearField(notificationsetting.FieldSmsNotifications, field.TypeBool)
	}
	if value, ok := nsu.mutation.PushNotifications(); ok {
		_spec.SetField(notificationsetting.FieldPushNotifications, field.TypeBool, value)
	}
	if nsu.mutation.PushNotificationsCleared() {
		_spec.ClearField(notificationsetting.FieldPushNotifications, field.TypeBool)
	}
	if value, ok := nsu.mutation.MarketingEmails(); ok {
		_spec.SetField(notificationsetting.FieldMarketingEmails, field.TypeBool, value)
	}
	if nsu.mutation.MarketingEmailsCleared() {
		_spec.ClearField(notificationsetting.FieldMarketingEmails, field.TypeBool)
	}
	if value, ok := nsu.mutation.SecurityAlerts(); ok {
		_spec.SetField(notificationsetting.FieldSecurityAlerts, field.TypeBool, value)
	}
	if nsu.mutation.SecurityAlertsCleared() {
		_spec.ClearField(notificationsetting.FieldSecurityAlerts, field.TypeBool)
	}
	if value, ok := nsu.mutation.LoginAlerts(); ok {
		_spec.SetField(notificationsetting.FieldLoginAlerts, field.TypeBool, value)
	}
	if nsu.mutation.LoginAlertsCleared() {
		_spec.ClearField(notificationsetting.FieldLoginAlerts, field.TypeBool)
	}
	if value, ok := nsu.mutation.ProfileUpdates(); ok {
		_spec.SetField(notificationsetting.FieldProfileUpdates, field.TypeBool, value)
	}
	if nsu.mutation.ProfileUpdatesCleared() {
		_spec.ClearField(notificationsetting.FieldProfileUpdates, field.TypeBool)
	}
	if value, ok := nsu.mutation.PaymentNotifications(); ok {
		_spec.SetField(notificationsetting.FieldPaymentNotifications, field.TypeBool, value)
	}
	if nsu.mutation.PaymentNotificationsCleared() {
		_spec.ClearField(notificationsetting.FieldPaymentNotifications, field.TypeBool)
	}
	if value, ok := nsu.mutation.CreatedAt(); ok {
		_spec.SetField(notificationsetting.FieldCreatedAt, field.TypeTime, value)
	}
	if nsu.mutation.CreatedAtCleared() {
		_spec.ClearField(notificationsetting.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := nsu.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationsetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if nsu.mutation.UpdatedAtCleared() {
		_spec.ClearField(notificationsetting.FieldUpdatedAt, field.TypeTime)
	}
	if nsu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationsetting.UserTable,
			Columns: []string{notificationsetting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationsetting.UserTable,
			Columns: []string{notificationsetting.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, nsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nsu.mutation.done = true
	return n, nil
}

// NotificationSettingUpdateOne is the builder for updating a single NotificationSetting entity.
type NotificationSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationSettingMutation
}

// SetUserID sets the "user_id" field.
func (nsuo *NotificationSettingUpdateOne) SetUserID(s string) *NotificationSettingUpdateOne {
	nsuo.mutation.SetUserID(s)
	return nsuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillableUserID(s *string) *NotificationSettingUpdateOne {
	if s != nil {
		nsuo.SetUserID(*s)
	}
	return nsuo
}

// ClearUserID clears the value of the "user_id" field.
func (nsuo *NotificationSettingUpdateOne) ClearUserID() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearUserID()
	return nsuo
}

// SetEmailNotifications sets the "email_notifications" field.
func (nsuo *NotificationSettingUpdateOne) SetEmailNotifications(b bool) *NotificationSettingUpdateOne {
	nsuo.mutation.SetEmailNotifications(b)
	return nsuo
}

// SetNillableEmailNotifications sets the "email_notifications" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillableEmailNotifications(b *bool) *NotificationSettingUpdateOne {
	if b != nil {
		nsuo.SetEmailNotifications(*b)
	}
	return nsuo
}

// ClearEmailNotifications clears the value of the "email_notifications" field.
func (nsuo *NotificationSettingUpdateOne) ClearEmailNotifications() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearEmailNotifications()
	return nsuo
}

// SetSmsNotifications sets the "sms_notifications" field.
func (nsuo *NotificationSettingUpdateOne) SetSmsNotifications(b bool) *NotificationSettingUpdateOne {
	nsuo.mutation.SetSmsNotifications(b)
	return nsuo
}

// SetNillableSmsNotifications sets the "sms_notifications" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillableSmsNotifications(b *bool) *NotificationSettingUpdateOne {
	if b != nil {
		nsuo.SetSmsNotifications(*b)
	}
	return nsuo
}

// ClearSmsNotifications clears the value of the "sms_notifications" field.
func (nsuo *NotificationSettingUpdateOne) ClearSmsNotifications() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearSmsNotifications()
	return nsuo
}

// SetPushNotifications sets the "push_notifications" field.
func (nsuo *NotificationSettingUpdateOne) SetPushNotifications(b bool) *NotificationSettingUpdateOne {
	nsuo.mutation.SetPushNotifications(b)
	return nsuo
}

// SetNillablePushNotifications sets the "push_notifications" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillablePushNotifications(b *bool) *NotificationSettingUpdateOne {
	if b != nil {
		nsuo.SetPushNotifications(*b)
	}
	return nsuo
}

// ClearPushNotifications clears the value of the "push_notifications" field.
func (nsuo *NotificationSettingUpdateOne) ClearPushNotifications() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearPushNotifications()
	return nsuo
}

// SetMarketingEmails sets the "marketing_emails" field.
func (nsuo *NotificationSettingUpdateOne) SetMarketingEmails(b bool) *NotificationSettingUpdateOne {
	nsuo.mutation.SetMarketingEmails(b)
	return nsuo
}

// SetNillableMarketingEmails sets the "marketing_emails" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillableMarketingEmails(b *bool) *NotificationSettingUpdateOne {
	if b != nil {
		nsuo.SetMarketingEmails(*b)
	}
	return nsuo
}

// ClearMarketingEmails clears the value of the "marketing_emails" field.
func (nsuo *NotificationSettingUpdateOne) ClearMarketingEmails() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearMarketingEmails()
	return nsuo
}

// SetSecurityAlerts sets the "security_alerts" field.
func (nsuo *NotificationSettingUpdateOne) SetSecurityAlerts(b bool) *NotificationSettingUpdateOne {
	nsuo.mutation.SetSecurityAlerts(b)
	return nsuo
}

// SetNillableSecurityAlerts sets the "security_alerts" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillableSecurityAlerts(b *bool) *NotificationSettingUpdateOne {
	if b != nil {
		nsuo.SetSecurityAlerts(*b)
	}
	return nsuo
}

// ClearSecurityAlerts clears the value of the "security_alerts" field.
func (nsuo *NotificationSettingUpdateOne) ClearSecurityAlerts() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearSecurityAlerts()
	return nsuo
}

// SetLoginAlerts sets the "login_alerts" field.
func (nsuo *NotificationSettingUpdateOne) SetLoginAlerts(b bool) *NotificationSettingUpdateOne {
	nsuo.mutation.SetLoginAlerts(b)
	return nsuo
}

// SetNillableLoginAlerts sets the "login_alerts" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillableLoginAlerts(b *bool) *NotificationSettingUpdateOne {
	if b != nil {
		nsuo.SetLoginAlerts(*b)
	}
	return nsuo
}

// ClearLoginAlerts clears the value of the "login_alerts" field.
func (nsuo *NotificationSettingUpdateOne) ClearLoginAlerts() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearLoginAlerts()
	return nsuo
}

// SetProfileUpdates sets the "profile_updates" field.
func (nsuo *NotificationSettingUpdateOne) SetProfileUpdates(b bool) *NotificationSettingUpdateOne {
	nsuo.mutation.SetProfileUpdates(b)
	return nsuo
}

// SetNillableProfileUpdates sets the "profile_updates" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillableProfileUpdates(b *bool) *NotificationSettingUpdateOne {
	if b != nil {
		nsuo.SetProfileUpdates(*b)
	}
	return nsuo
}

// ClearProfileUpdates clears the value of the "profile_updates" field.
func (nsuo *NotificationSettingUpdateOne) ClearProfileUpdates() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearProfileUpdates()
	return nsuo
}

// SetPaymentNotifications sets the "payment_notifications" field.
func (nsuo *NotificationSettingUpdateOne) SetPaymentNotifications(b bool) *NotificationSettingUpdateOne {
	nsuo.mutation.SetPaymentNotifications(b)
	return nsuo
}

// SetNillablePaymentNotifications sets the "payment_notifications" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillablePaymentNotifications(b *bool) *NotificationSettingUpdateOne {
	if b != nil {
		nsuo.SetPaymentNotifications(*b)
	}
	return nsuo
}

// ClearPaymentNotifications clears the value of the "payment_notifications" field.
func (nsuo *NotificationSettingUpdateOne) ClearPaymentNotifications() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearPaymentNotifications()
	return nsuo
}

// SetCreatedAt sets the "created_at" field.
func (nsuo *NotificationSettingUpdateOne) SetCreatedAt(t time.Time) *NotificationSettingUpdateOne {
	nsuo.mutation.SetCreatedAt(t)
	return nsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillableCreatedAt(t *time.Time) *NotificationSettingUpdateOne {
	if t != nil {
		nsuo.SetCreatedAt(*t)
	}
	return nsuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (nsuo *NotificationSettingUpdateOne) ClearCreatedAt() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearCreatedAt()
	return nsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nsuo *NotificationSettingUpdateOne) SetUpdatedAt(t time.Time) *NotificationSettingUpdateOne {
	nsuo.mutation.SetUpdatedAt(t)
	return nsuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nsuo *NotificationSettingUpdateOne) SetNillableUpdatedAt(t *time.Time) *NotificationSettingUpdateOne {
	if t != nil {
		nsuo.SetUpdatedAt(*t)
	}
	return nsuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nsuo *NotificationSettingUpdateOne) ClearUpdatedAt() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearUpdatedAt()
	return nsuo
}

// SetUser sets the "user" edge to the User entity.
func (nsuo *NotificationSettingUpdateOne) SetUser(u *User) *NotificationSettingUpdateOne {
	return nsuo.SetUserID(u.ID)
}

// Mutation returns the NotificationSettingMutation object of the builder.
func (nsuo *NotificationSettingUpdateOne) Mutation() *NotificationSettingMutation {
	return nsuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (nsuo *NotificationSettingUpdateOne) ClearUser() *NotificationSettingUpdateOne {
	nsuo.mutation.ClearUser()
	return nsuo
}

// Where appends a list predicates to the NotificationSettingUpdate builder.
func (nsuo *NotificationSettingUpdateOne) Where(ps ...predicate.NotificationSetting) *NotificationSettingUpdateOne {
	nsuo.mutation.Where(ps...)
	return nsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nsuo *NotificationSettingUpdateOne) Select(field string, fields ...string) *NotificationSettingUpdateOne {
	nsuo.fields = append([]string{field}, fields...)
	return nsuo
}

// Save executes the query and returns the updated NotificationSetting entity.
func (nsuo *NotificationSettingUpdateOne) Save(ctx context.Context) (*NotificationSetting, error) {
	return withHooks(ctx, nsuo.sqlSave, nsuo.mutation, nsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nsuo *NotificationSettingUpdateOne) SaveX(ctx context.Context) *NotificationSetting {
	node, err := nsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nsuo *NotificationSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := nsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsuo *NotificationSettingUpdateOne) ExecX(ctx context.Context) {
	if err := nsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nsuo *NotificationSettingUpdateOne) sqlSave(ctx context.Context) (_node *NotificationSetting, err error) {
	_spec := sqlgraph.NewUpdateSpec(notificationsetting.Table, notificationsetting.Columns, sqlgraph.NewFieldSpec(notificationsetting.FieldID, field.TypeString))
	id, ok := nsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NotificationSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationsetting.FieldID)
		for _, f := range fields {
			if !notificationsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notificationsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nsuo.mutation.EmailNotifications(); ok {
		_spec.SetField(notificationsetting.FieldEmailNotifications, field.TypeBool, value)
	}
	if nsuo.mutation.EmailNotificationsCleared() {
		_spec.ClearField(notificationsetting.FieldEmailNotifications, field.TypeBool)
	}
	if value, ok := nsuo.mutation.SmsNotifications(); ok {
		_spec.SetField(notificationsetting.FieldSmsNotifications, field.TypeBool, value)
	}
	if nsuo.mutation.SmsNotificationsCleared() {
		_spec.ClearField(notificationsetting.FieldSmsNotifications, field.TypeBool)
	}
	if value, ok := nsuo.mutation.PushNotifications(); ok {
		_spec.SetField(notificationsetting.FieldPushNotifications, field.TypeBool, value)
	}
	if nsuo.mutation.PushNotificationsCleared() {
		_spec.ClearField(notificationsetting.FieldPushNotifications, field.TypeBool)
	}
	if value, ok := nsuo.mutation.MarketingEmails(); ok {
		_spec.SetField(notificationsetting.FieldMarketingEmails, field.TypeBool, value)
	}
	if nsuo.mutation.MarketingEmailsCleared() {
		_spec.ClearField(notificationsetting.FieldMarketingEmails, field.TypeBool)
	}
	if value, ok := nsuo.mutation.SecurityAlerts(); ok {
		_spec.SetField(notificationsetting.FieldSecurityAlerts, field.TypeBool, value)
	}
	if nsuo.mutation.SecurityAlertsCleared() {
		_spec.ClearField(notificationsetting.FieldSecurityAlerts, field.TypeBool)
	}
	if value, ok := nsuo.mutation.LoginAlerts(); ok {
		_spec.SetField(notificationsetting.FieldLoginAlerts, field.TypeBool, value)
	}
	if nsuo.mutation.LoginAlertsCleared() {
		_spec.ClearField(notificationsetting.FieldLoginAlerts, field.TypeBool)
	}
	if value, ok := nsuo.mutation.ProfileUpdates(); ok {
		_spec.SetField(notificationsetting.FieldProfileUpdates, field.TypeBool, value)
	}
	if nsuo.mutation.ProfileUpdatesCleared() {
		_spec.ClearField(notificationsetting.FieldProfileUpdates, field.TypeBool)
	}
	if value, ok := nsuo.mutation.PaymentNotifications(); ok {
		_spec.SetField(notificationsetting.FieldPaymentNotifications, field.TypeBool, value)
	}
	if nsuo.mutation.PaymentNotificationsCleared() {
		_spec.ClearField(notificationsetting.FieldPaymentNotifications, field.TypeBool)
	}
	if value, ok := nsuo.mutation.CreatedAt(); ok {
		_spec.SetField(notificationsetting.FieldCreatedAt, field.TypeTime, value)
	}
	if nsuo.mutation.CreatedAtCleared() {
		_spec.ClearField(notificationsetting.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := nsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationsetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if nsuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(notificationsetting.FieldUpdatedAt, field.TypeTime)
	}
	if nsuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationsetting.UserTable,
			Columns: []string{notificationsetting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationsetting.UserTable,
			Columns: []string{notificationsetting.UserColumn},
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
	_node = &NotificationSetting{config: nsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nsuo.mutation.done = true
	return _node, nil
}
