package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"trabalho-01-batalha-naval/config"
	"trabalho-01-batalha-naval/domain/entity"
)

type JwtService struct {
	key      []byte
	duration time.Duration
	renewDue time.Duration
}

func NewJwtService(cfg config.Config) *JwtService {
	return &JwtService{
		key:      []byte(cfg.JwtSecret),
		duration: time.Duration(cfg.JwtDurationInMinutes) * time.Minute,
		renewDue: time.Duration(cfg.JwtRenewDueInMinutes) * time.Minute,
	}
}

func (s *JwtService) GenerateAccessToken(user *entity.User) (*Token, error) {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(s.duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.UUID,
		"exp": expiresAt.Unix(),
		"iat": issuedAt.Unix(),
		"goa": Access,
	})

	signed, err := token.SignedString(s.key)

	if err != nil {
		return nil, err
	}

	return &Token{
		Content:   signed,
		Uid:       user.UUID,
		Type:      Bearer,
		Goal:      Access,
		ExpiresAt: expiresAt,
		IssuedAt:  issuedAt,
	}, nil
}

func (s *JwtService) GenerateRefreshToken(user *entity.User) (*Token, error) {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(s.renewDue)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.UUID,
		"exp": expiresAt.Unix(),
		"iat": issuedAt.Unix(),
		"goa": Refresh,
	})

	signed, err := token.SignedString(s.key)

	if err != nil {
		return nil, err
	}

	return &Token{
		Content:   signed,
		Uid:       user.UUID,
		Type:      Bearer,
		Goal:      Refresh,
		ExpiresAt: expiresAt,
		IssuedAt:  issuedAt,
	}, nil
}

func (s *JwtService) GenerateBroadcastToken(user *entity.User) (*Token, error) {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(s.duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.UUID,
		"exp": expiresAt.Unix(),
		"iat": issuedAt.Unix(),
		"goa": Broadcast,
	})

	signed, err := token.SignedString(s.key)

	if err != nil {
		return nil, err
	}

	return &Token{
		Content:   signed,
		Uid:       user.UUID,
		Type:      Bearer,
		Goal:      Broadcast,
		ExpiresAt: expiresAt,
		IssuedAt:  issuedAt,
	}, nil
}

func (s *JwtService) ParseTokenFromContent(content string) (*Token, error) {
	rawToken, err := jwt.Parse(content, func(token *jwt.Token) (interface{}, error) {
		return s.key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := rawToken.Claims.(jwt.MapClaims)

	if !ok || !rawToken.Valid {
		return nil, InvalidTokenError
	}

	uid, ok := claims["uid"].(string)

	if !ok {
		return nil, InvalidTokenError
	}

	exp, ok := claims["exp"].(float64)

	if !ok {
		return nil, InvalidTokenError
	}

	iat, ok := claims["iat"].(float64)

	if !ok {
		return nil, InvalidTokenError
	}

	goa, ok := claims["goa"].(string)

	if !ok {
		return nil, InvalidTokenError
	}

	goal, err := ParseGoal(goa)

	if err != nil {
		return nil, err
	}

	return &Token{
		Content:   content,
		Uid:       uid,
		Type:      Bearer,
		Goal:      goal,
		ExpiresAt: time.Unix(int64(exp), 0),
		IssuedAt:  time.Unix(int64(iat), 0),
	}, nil
}
