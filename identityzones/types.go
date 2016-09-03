package identityzones

type IdentityZone struct {
	ID           string
	Subdomain    string
	Name         string
	Version      int
	Description  string
	Created      int
	LastModified int `json:"last_modified"`
	Config       *Config
}

type Config struct {
	TokenPolicy *TokenPolicy
	SamlConfig  *SamlConfig
	//IDP Discovery should be set to true if you have configured more than one identity provider for UAA. The discovery relies on email domain being set for each additional provider
	IdpDiscoveryEnabled bool
	Prompts             []*Prompt
	Links               *Links
}

type Prompt struct {
	Name string
	Type string
	Text string
}

type Links struct {
	Logout      *Logout
	SelfService *SelfService
}

type Logout struct {
	RedirectURL           string
	RedirectParameterName string
	//Whether or not to allow the redirect parameter on logout
	DisableRedirectParameter bool
	Whitelist                []string
}

type SelfService struct {
	//Whether or not users are allowed to sign up or reset their passwords via the UI
	SelfServiceLinksEnabled bool
	//Where users are directed upon clicking the account creation link
	Signup string
	//Where users are directed upon clicking the password reset link
	Passwd string
}

type TokenPolicy struct {
	//The ID for the key that is being used to sign tokens
	ActiveKeyID string
	//Time in seconds between when a access token is issued and when it expires. Defaults to global accessTokenValidit
	AccessTokenValidity int
	//Time in seconds between when a refresh token is issued and when it expires. Defaults to global refreshTokenValidity
	RefreshTokenValidity int
}

type SamlConfig struct {
	//If true, the SAML provider will sign all assertions
	AssertionSigned bool
	//Exposed SAML metadata property. If true, all assertions received by the SAML provider must be signed. Defaults to true.
	WantAssertionSigned bool
	//Exposed SAML metadata property. If true, the service provider will sign all outgoing authentication requests. Defaults to true.
	RequestSigned bool
	//If true, the authentication request from the partner service provider must be signed.
	WantAuthnRequestSigned bool
	//The lifetime of a SAML assertion in seconds. Defaults to 600.
	AssertionTimeToLiveSeconds int
	//Exposed SAML metadata property. The certificate used to sign all communications.
	Certificate string
	//Exposed SAML metadata property. The SAML provider’s private key.
	PrivateKey string
	//Exposed SAML metadata property. The SAML provider’s private key password. Reserved for future use.
	PrivateKeyPassword string
}
