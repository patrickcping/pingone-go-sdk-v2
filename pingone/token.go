package pingone

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	EnvironmentID  string   `json:"env,omitempty"`
	OrganizationID string   `json:"org,omitempty"`
	ClientID       string   `json:"client_id,omitempty"`
	Audience       []string `json:"aud,omitempty"`
	ExpiresAt      int64    `json:"exp,omitempty"`
	Id             string   `json:"jti,omitempty"`
	IssuedAt       int64    `json:"iat,omitempty"`
	Issuer         string   `json:"iss,omitempty"`
	NotBefore      int64    `json:"nbf,omitempty"`
	Subject        string   `json:"sub,omitempty"`
	jwt.Claims
}

func tokenKeyFunction() jwt.Keyfunc {
	return nil
}

func (c TokenClaims) GetEnvironmentID() string {
	return c.EnvironmentID
}

func (c TokenClaims) GetOrganizationID() string {
	return c.OrganizationID
}

func (c TokenClaims) GetClientID() string {
	return c.ClientID
}

func (c TokenClaims) GetAudience() []string {
	return c.Audience
}

func (c TokenClaims) GetExpiresAt() int64 {
	return c.ExpiresAt
}

func (c TokenClaims) GetId() string {
	return c.Id
}

func (c TokenClaims) GetIssuedAt() int64 {
	return c.IssuedAt
}

func (c TokenClaims) GetIssuer() string {
	return c.Issuer
}

func (c TokenClaims) GetNotBefore() int64 {
	return c.NotBefore
}

func (c TokenClaims) GetSubject() string {
	return c.Subject
}
