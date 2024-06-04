package pingone

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"golang.org/x/oauth2"
)

type Config struct {
	AccessToken          *string
	accessTokenObject    *oauth2.Token
	APIHostnameOverride  *string
	AuthHostnameOverride *string
	ClientID             *string
	ClientSecret         *string
	EnvironmentID        *string
	ProxyURL             *string
	Region               string
	UserAgentOverride    *string
	UserAgentSuffix      *string
	validated            bool
}

var p1ResourceIDRegexp = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
var isHostname = regexp.MustCompile(`^(?:[\w-]+\.)+[a-z]{2,}(?:\/[\w-]+)*(?:\/[\w.-]+)?\/?(?:\?.*)?$`)

func (c *Config) validateAccessToken() error {
	if !checkForValue(c.AccessToken) {
		if v := envVar("PINGONE_API_ACCESS_TOKEN"); v != "" {
			c.AccessToken = &v
		}
	}

	return nil
}

func (c *Config) validateAPIHostnameOverride() error {
	if !checkForValue(c.APIHostnameOverride) {
		if v := envVar("PINGONE_API_SERVICE_HOSTNAME"); v != "" {
			c.APIHostnameOverride = &v
		}
	}

	if checkForValue(c.APIHostnameOverride) {
		if !isHostname.MatchString(*c.APIHostnameOverride) {
			return fmt.Errorf("Invalid parameter format.  Expected hostname format, got: %s", *c.APIHostnameOverride)
		}
	}

	return nil
}

func (c *Config) validateAuthHostnameOverride() error {
	if !checkForValue(c.AuthHostnameOverride) {
		if v := envVar("PINGONE_AUTH_SERVICE_HOSTNAME"); v != "" {
			c.AuthHostnameOverride = &v
		}
	}

	if checkForValue(c.AuthHostnameOverride) {
		if !isHostname.MatchString(*c.AuthHostnameOverride) {
			return fmt.Errorf("Invalid parameter format.  Expected hostname format, got: %s", *c.AuthHostnameOverride)
		}
	}

	return nil
}

func (c *Config) validateClientID() error {
	if !checkForValue(c.ClientID) {
		if v := envVar("PINGONE_CLIENT_ID"); v != "" {
			c.ClientID = &v
		}
	}

	if checkForValue(c.ClientID) {
		if !p1ResourceIDRegexp.MatchString(*c.ClientID) {
			return fmt.Errorf("Invalid parameter format.  Expected PingOne resource ID, got: %s", *c.ClientID)
		}
	}

	return nil
}

func (c *Config) validateClientSecret() error {
	if !checkForValue(c.ClientSecret) {
		if v := envVar("PINGONE_CLIENT_SECRET"); v != "" {
			c.ClientSecret = &v
		}
	}

	return nil
}

func (c *Config) validateEnvironmentID() error {
	if !checkForValue(c.EnvironmentID) {
		if v := envVar("PINGONE_ENVIRONMENT_ID"); v != "" {
			c.EnvironmentID = &v
		}
	}

	if checkForValue(c.EnvironmentID) {
		if !p1ResourceIDRegexp.MatchString(*c.EnvironmentID) {
			return fmt.Errorf("Invalid parameter format.  Expected PingOne resource ID, got: %s", *c.EnvironmentID)
		}
	}

	return nil
}

func (c *Config) validateRegion() error {
	if !checkForValue(c.Region) {
		if v := envVar("PINGONE_REGION"); v != "" {
			c.Region = v
		}
	}

	if !checkForValue(c.Region) {
		return fmt.Errorf("Must provide the region parameter.")
	} else {
		if !slices.Contains(model.RegionsAvailableList(), c.Region) {
			return fmt.Errorf("Invalid region value.  The region parameter is case sensitive and must be one of the following values: %s", strings.Join(model.RegionsAvailableList(), ", "))
		}
	}

	return nil
}

func (c *Config) validateProxyURL() error {
	// Defer validation to the net/http library
	return nil
}

func (c *Config) Validate() error {

	if c.validated {
		return nil
	}

	if err := c.validateAccessToken(); err != nil {
		return err
	}

	if err := c.validateAPIHostnameOverride(); err != nil {
		return err
	}

	if err := c.validateAuthHostnameOverride(); err != nil {
		return err
	}

	if err := c.validateClientID(); err != nil {
		return err
	}

	if err := c.validateClientSecret(); err != nil {
		return err
	}

	if err := c.validateEnvironmentID(); err != nil {
		return err
	}

	if err := c.validateRegion(); err != nil {
		return err
	}

	if err := c.validateProxyURL(); err != nil {
		return err
	}

	err := fmt.Errorf("Must provide the region parameter and either client ID, client secret and environment ID, or an api access token.")

	if (!checkForValue(c.ClientID) ||
		!checkForValue(c.ClientSecret) ||
		!checkForValue(c.EnvironmentID)) &&
		!checkForValue(c.AccessToken) {
		return err
	}

	// Conflicting attrs check
	if (checkForValue(c.ClientID) || checkForValue(c.ClientSecret) || checkForValue(c.EnvironmentID)) && checkForValue(c.AccessToken) {
		return fmt.Errorf("API access token cannot be set with client ID, client secret or environment ID")
	}

	// Service overrides
	if servicesOverridden := checkForValue(c.APIHostnameOverride) || checkForValue(c.AuthHostnameOverride); servicesOverridden {

		if servicesOverridden && (!checkForValue(c.APIHostnameOverride) || !checkForValue(c.AuthHostnameOverride)) {
			return fmt.Errorf("Required service endpoints not configured.  When overriding service endpoints, both auth (e.g. auth.pingone.com) and api service (e.g. api.pingone.com) endpoints must be set.")
		}
	}

	c.validated = true

	return nil
}

func envVar(name string) string {
	return strings.TrimSpace(os.Getenv(name))
}

func checkForValue(o any) bool {

	switch v := o.(type) {
	case *string:
		if v == nil || *v == "" {
			return false
		}
		return true
	case string:
		if v == "" {
			return false
		}
		return true
	}

	return false
}
