# PingOne Administration GO SDK

The PingOne GO SDK provides a set of functions and stucts that help with interacting with the PingOne public cloud API.

The code is intended to be delivered as a sample, until an official GO SDK is released from Ping Identity.  As such, the code is highly likely to change significantly between releases. 

Code for each service is generated with the help of the [OpenAPI Generator](https://openapi-generator.tech/).

## Packages

The SDK provides a core package, and a package per PingOne service, each with their own directory off the root of the project:

* **agreementmanagement** for the PingOne end-user agreements managements service
* **authorize** for the PingOne Authorize service
* **credentials** for the PingOne Credentials service, part of PingOne Neo
* **management** for the PingOne platform common and SSO services
* **mfa** for the PingOne MFA service
* **risk** for the PingOne Protect service
* **verify** for the PingOne Verify service, part of PingOne Neo

## Getting Started

The client can be invoked using the following syntax:

```go

...

import (
	"context"

	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
)

...

config := &pingone.Config{
	ClientID:      clientId,
	ClientSecret:  clientSecret,
	EnvironmentID: environmentId,
	AccessToken:   accessToken,
	Region:        region,
}

client, err := config.APIClient(ctx)
if err != nil {
	return nil, err
}
```

The result is an object with clients initialised for each service:
* `client.AgreementManagementAPIClient`
* `client.AuthorizeAPIClient`
* `client.CredentialsAPIClient`
* `client.ManagementAPIClient`
* `client.MFAAPIClient`
* `client.RiskAPIClient`
* `client.VerifyAPIClient`

In the above, if an `AccessToken` is provided, this will be verified and used.  If the `AccessToken` is not provided, the SDK will retrieve an access token from the provided `ClientID`, `ClientSecret`, `EnvironmentID` and `Region` parameters.

The client SDK defaults to production hostnames, and the `Region` is used to add the relevant suffix to the hostname.  For example, `Europe` as a `Region` value with suffix the service hostname with `.eu`.  Hostnames can be overridden with the optional `APIHostnameOverride`, `AgreementMgmtHostnameOverride`, and `AuthHostnameOverride` parameters.

An API call can be made against the API objects, as in the following example to get all environments in a tenant:

```go
...

resp, r, err := client.ManagementAPIClient.EnvironmentsApi.ReadAllEnvironments(ctx).Execute()
if err != nil {
	return nil, err
}
...
```

## Contributing

Each package is generated from an underlying OpenAPI 3 specification.  Currently the OpenAPI 3 specification is stored in the `./<<module>>/generate/pingone-<<module>>.yml` file, although this will be subject to change in the future.

* [**agreementmanagement** OpenAPI 3 Specification file](./authorize/generate/pingone-authorize.yml)
* [**authorize** OpenAPI 3 Specification file](./authorize/generate/pingone-authorize.yml)
* [**credentials** OpenAPI 3 Specification file](./credentials/generate/pingone-credentials.yml)
* [**management** OpenAPI 3 Specification file](./management/generate/pingone-management.yml)
* [**mfa** OpenAPI 3 Specification file](./mfa/generate/pingone-mfa.yml)
* [**risk** OpenAPI 3 Specification file](./risk/generate/pingone-risk.yml)
* [**verify** OpenAPI 3 Specification file](./verify/generate/pingone-verify.yml)

Once this file has been updated, from the module directory itself run `make generate`.  This will generate the required `api_*.go` files, `model_*.go` files and associated documentation.

Before raising a Pull request, the resulting code can be checked using the `make devcheck` command.  This will build, lint and verify the code.