package schema

import (
	"app/pkg/database/ent"
	"app/pkg/database/ent/hook"
	"context"

	"github.com/google/uuid"
)

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (ActivityLog) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.ActivityLogFunc(func(ctx context.Context, m *ent.ActivityLogMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (ApiKey) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.ApiKeyFunc(func(ctx context.Context, m *ent.ApiKeyMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (NotificationSetting) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.NotificationSettingFunc(func(ctx context.Context, m *ent.NotificationSettingMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (PaymentMethod) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.PaymentMethodFunc(func(ctx context.Context, m *ent.PaymentMethodMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (PrivacySetting) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.PrivacySettingFunc(func(ctx context.Context, m *ent.PrivacySettingMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (Transaction) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.TransactionFunc(func(ctx context.Context, m *ent.TransactionMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (UserProfile) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.UserProfileFunc(func(ctx context.Context, m *ent.UserProfileMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (User2fa) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.User2faFunc(func(ctx context.Context, m *ent.User2faMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (UserSession) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.UserSessionFunc(func(ctx context.Context, m *ent.UserSessionMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
func (UserToken) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.UserTokenFunc(func(ctx context.Context, m *ent.UserTokenMutation) (ent.Value, error) {
				uuidValue := uuid.New().String()
				m.SetID(uuidValue)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
