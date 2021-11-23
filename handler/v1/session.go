package v1

import (
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mazrean/todoList/common"
	"github.com/mazrean/todoList/domain"
	"github.com/mazrean/todoList/domain/values"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Session struct {
	key    string
	secret string
	store  sessions.Store
}

func NewSession(key common.SessionKey, secret common.SessionSecret) *Session {
	store := cookie.NewStore([]byte(secret))

	/*
		gorilla/sessionsの内部で使われているgobが
		time.Time,uuid.UUIDのエンコード・デコードをできるようにRegisterする
	*/
	gob.Register(time.Time{})
	gob.Register(uuid.UUID{})

	return &Session{
		key:    string(key),
		secret: string(secret),
		store:  store,
	}
}

func (s *Session) Use(r *gin.Engine) {
	r.Use(sessions.Sessions("session", s.store))
}

func (s *Session) getSession(c *gin.Context) sessions.Session {
	session := sessions.Default(c)

	return session
}

func (s *Session) save(c *gin.Context, session sessions.Session) error {
	err := session.Save()
	if err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}

func (s *Session) revoke(session sessions.Session) {
	session.Clear()
}

var (
	ErrNoValue     = errors.New("no value")
	ErrValueBroken = errors.New("value broken")
)

const (
	userIDSessionKey             = "userID"
	userNameSessionKey           = "userName"
	userHashedPasswordSessionKey = "userHashedPassword"
)

func (s *Session) setUser(session sessions.Session, user *domain.User) {
	session.Set(userIDSessionKey, uuid.UUID(user.GetID()))
	session.Set(userNameSessionKey, string(user.GetName()))
	session.Set(userHashedPasswordSessionKey, []byte(user.GetHashedPassword()))
}

func (s *Session) getUser(session sessions.Session) (*domain.User, error) {
	iUserID := session.Get(userIDSessionKey)
	if iUserID == nil {
		return nil, ErrNoValue
	}

	userID, ok := session.Get(userIDSessionKey).(uuid.UUID)
	if !ok {
		return nil, ErrValueBroken
	}

	iUserName := session.Get(userNameSessionKey)
	if iUserName == nil {
		return nil, ErrNoValue
	}

	userName, ok := iUserName.(string)
	if !ok {
		return nil, ErrValueBroken
	}

	iUserHashedPassword := session.Get(userHashedPasswordSessionKey)
	if iUserHashedPassword == nil {
		return nil, ErrNoValue
	}

	userHashedPassword, ok := iUserHashedPassword.([]byte)
	if !ok {
		return nil, ErrValueBroken
	}

	return domain.NewUser(
		values.NewUserIDFromUUID(userID),
		values.NewUserName(userName),
		values.NewUserHashedPassword(userHashedPassword),
	), nil
}
