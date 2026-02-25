//go:build !beta

package pingone

import "context"

// BetaAPIClients holds API clients for beta features.
// In non-beta builds this struct is empty.
type BetaAPIClients struct{}

func (c *Config) betaAPIClients(_ context.Context) (BetaAPIClients, error) {
	return BetaAPIClients{}, nil
}
