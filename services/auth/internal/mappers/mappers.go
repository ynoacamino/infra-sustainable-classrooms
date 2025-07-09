package mappers

import (
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
	auth "github.com/ynoacamino/infrastructure/services/auth/gen/auth"
	authdb "github.com/ynoacamino/infrastructure/services/auth/gen/database"
)

func timestampToMillis(timestamp pgtype.Timestamptz) int64 {
	if timestamp.Valid {
		return timestamp.Time.UnixMilli()
	}
	return 0
}

func UserDBToAPI(user *authdb.User) *auth.User {
	if user == nil {
		return nil
	}

	apiUser := &auth.User{
		ID:         user.ID,
		Identifier: user.Identifier,
		IsVerified: user.IsVerified,
		CreatedAt:  timestampToMillis(user.CreatedAt),
	}

	if user.LastLogin.Valid {
		lastLogin := timestampToMillis(user.LastLogin)
		apiUser.LastLogin = &lastLogin
	}

	if len(user.Metadata) > 0 {
		var metadata map[string]string
		if err := json.Unmarshal(user.Metadata, &metadata); err == nil {
			apiUser.Metadata = metadata
		}
	}

	return apiUser
}

// SessionDBToAPI convierte un modelo de sesión de la base de datos al modelo de la API
func SessionDBToAPI(session *authdb.Session) *auth.Session {
	if session == nil {
		return nil
	}

	apiSession := &auth.Session{
		ID:        session.ID,
		UserID:    session.UserID,
		CreatedAt: timestampToMillis(session.CreatedAt),
		ExpiresAt: timestampToMillis(session.ExpiresAt),
		IsActive:  session.IsActive,
	}

	if session.LastAccessed.Valid {
		lastAccessed := timestampToMillis(session.LastAccessed)
		apiSession.LastAccessed = &lastAccessed
	}

	if session.UserAgent.Valid {
		apiSession.UserAgent = &session.UserAgent.String
	}

	if session.IpAddress != nil {
		ipAddr := session.IpAddress.String()
		apiSession.IPAddress = &ipAddr
	}

	if session.DeviceID.Valid {
		apiSession.DeviceID = &session.DeviceID.String
	}

	if session.Platform.Valid {
		apiSession.Platform = &session.Platform.String
	}

	return apiSession
}

func UsersDBToAPI(users []authdb.User) []*auth.User {
	apiUsers := make([]*auth.User, len(users))
	for i, user := range users {
		apiUsers[i] = UserDBToAPI(&user)
	}
	return apiUsers
}

func SessionsDBToAPI(sessions []authdb.Session) []*auth.Session {
	apiSessions := make([]*auth.Session, len(sessions))
	for i, session := range sessions {
		apiSessions[i] = SessionDBToAPI(&session)
	}
	return apiSessions
}
