package token

import (
	"strings"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

// PasetoMaker is a struct that implements Maker interface
type PasetoMaker struct {
	symmetricKey paseto.V4SymmetricKey
	key          []byte
}

func NewPasetoMaker(key string) Maker {
	return &PasetoMaker{paseto.NewV4SymmetricKey(), []byte(key)}
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	// create paseto token
	token := paseto.NewToken()
	// Create uuid for token id
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", nil, err
	}
	// add data to the token.
	if err = token.Set("id", tokenID.String()); err != nil {
		return "", nil, err
	}
	if err = token.Set("username", username); err != nil {
		return "", nil, err
	}
	issuedAt := time.Now()
	token.SetIssuedAt(issuedAt)
	expiredAt := time.Now().Add(duration)
	token.SetExpiration(expiredAt)

	return token.V4Encrypt(maker.symmetricKey, maker.key),
		&Payload{
			ID:        tokenID,
			Username:  username,
			Role:      role,
			IssuedAt:  issuedAt,
			ExpiredAt: expiredAt,
		},
		nil
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	// construct payload from token
	payload, err := maker.GetPayloadFromToken(token)
	if err != nil {
		return nil, ErrInvalidToken
	}
	return payload, nil

}

// GetPayloadFromToken todo refactor this function
func (maker *PasetoMaker) GetPayloadFromToken(token string) (*Payload, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())
	parsedToken, err := parser.ParseV4Local(maker.symmetricKey, token, maker.key)
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	id, err := parsedToken.GetString("id")
	if err != nil {
		return nil, ErrInvalidToken
	}
	username, err := parsedToken.GetString("username")
	if err != nil {
		return nil, ErrInvalidToken
	}
	issuedAt, err := parsedToken.GetIssuedAt()
	if err != nil {
		return nil, ErrInvalidToken
	}
	expiredAt, err := parsedToken.GetExpiration()
	if err != nil {
		return nil, ErrInvalidToken
	}

	return &Payload{
		ID:        uuid.MustParse(id),
		Username:  username,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}, nil
}

var _ Maker = (*PasetoMaker)(nil)
