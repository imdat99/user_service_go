// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"app/pkg/database/ent/activitylog"
	"app/pkg/database/ent/apikey"
	"app/pkg/database/ent/notificationsetting"
	"app/pkg/database/ent/paymentmethod"
	"app/pkg/database/ent/privacysetting"
	"app/pkg/database/ent/schema"
	"app/pkg/database/ent/transaction"
	"app/pkg/database/ent/user"
	"app/pkg/database/ent/user2fa"
	"app/pkg/database/ent/userprofile"
	"app/pkg/database/ent/usersession"
	"app/pkg/database/ent/usertoken"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	activitylogHooks := schema.ActivityLog{}.Hooks()
	activitylog.Hooks[0] = activitylogHooks[0]
	apikeyHooks := schema.ApiKey{}.Hooks()
	apikey.Hooks[0] = apikeyHooks[0]
	notificationsettingHooks := schema.NotificationSetting{}.Hooks()
	notificationsetting.Hooks[0] = notificationsettingHooks[0]
	paymentmethodHooks := schema.PaymentMethod{}.Hooks()
	paymentmethod.Hooks[0] = paymentmethodHooks[0]
	privacysettingHooks := schema.PrivacySetting{}.Hooks()
	privacysetting.Hooks[0] = privacysettingHooks[0]
	transactionHooks := schema.Transaction{}.Hooks()
	transaction.Hooks[0] = transactionHooks[0]
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	user2faHooks := schema.User2fa{}.Hooks()
	user2fa.Hooks[0] = user2faHooks[0]
	userprofileHooks := schema.UserProfile{}.Hooks()
	userprofile.Hooks[0] = userprofileHooks[0]
	usersessionHooks := schema.UserSession{}.Hooks()
	usersession.Hooks[0] = usersessionHooks[0]
	usertokenHooks := schema.UserToken{}.Hooks()
	usertoken.Hooks[0] = usertokenHooks[0]
}

const (
	Version = "v0.14.4"                                         // Version of ent codegen.
	Sum     = "h1:/DhDraSLXIkBhyiVoJeSshr4ZYi7femzhj6/TckzZuI=" // Sum of ent codegen.
)
