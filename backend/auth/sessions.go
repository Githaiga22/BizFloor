package auth

import (
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

type Session struct {
	ID           uuid.UUID
	UserID       uint
	UserName     string
	UserEmail    string
	IPAddress    string
	CreatedAt    time.Time
	ExpiresAt    time.Time
	LastActivity time.Time
}

type SessionStore struct {
	mu       sync.RWMutex
	sessions map[uuid.UUID]*Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[uuid.UUID]*Session),
	}
}

func (store *SessionStore) CreateSession(userID uint, userName, userEmail, ipAddress string) (*Session, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	sessionID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	session := &Session{
		ID:           sessionID,
		UserID:       userID,
		UserName:     userName,
		UserEmail:    userEmail,
		IPAddress:    ipAddress,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(24 * time.Hour),
		LastActivity: time.Now(),
	}

	store.sessions[sessionID] = session
	return session, nil
}

func (store *SessionStore) GetSession(sessionID uuid.UUID) (*Session, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	session, exists := store.sessions[sessionID]
	if !exists || time.Now().After(session.ExpiresAt) {
		if exists {
			delete(store.sessions, sessionID)
		}
		return nil, false
	}

	return session, true
}

func (store *SessionStore) DeleteSession(sessionID uuid.UUID) {
	store.mu.Lock()
	defer store.mu.Unlock()

	delete(store.sessions, sessionID)
}

func (store *SessionStore) ExtendSession(sessionID uuid.UUID) {
	store.mu.Lock()
	defer store.mu.Unlock()

	if session, exists := store.sessions[sessionID]; exists {
		session.ExpiresAt = time.Now().Add(24 * time.Hour)
		session.LastActivity = time.Now()
	}
}
