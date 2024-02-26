/*
 * Copyright (c) 2022-Present, Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const (
	// Version app version
	Version = "1.2.2"

	// AWSCredentialsFormat format const
	AWSCredentialsFormat = "aws-credentials"
	// EnvVarFormat format const
	EnvVarFormat = "env-var"

	// AWSAcctFedAppIDFlag cli flag const
	AWSAcctFedAppIDFlag = "aws-acct-fed-app-id"
	// AWSCredentialsFlag cli flag const
	AWSCredentialsFlag = "aws-credentials"
	// AWSIAMIdPFlag cli flag const
	AWSIAMIdPFlag = "aws-iam-idp"
	// AWSIAMRoleFlag cli flag const
	AWSIAMRoleFlag = "aws-iam-role"
	// DebugFlag cli flag const
	DebugFlag = "debug"
	// DebugAPICallsFlag cli flag const
	DebugAPICallsFlag = "debug-api-calls"
	// DebugConfigFlag cli flag const
	DebugConfigFlag = "debug-config"
	// FormatFlag cli flag const
	FormatFlag = "format"
	// OIDCClientIDFlag cli flag const
	OIDCClientIDFlag = "oidc-client-id"
	// OpenBrowserFlag cli flag const
	OpenBrowserFlag = "open-browser"
	// OrgDomainFlag cli flag const
	OrgDomainFlag = "org-domain"
	// ProfileFlag cli flag const
	ProfileFlag = "profile"
	// QRCodeFlag cli flag const
	QRCodeFlag = "qr-code"
	// SessionDurationFlag cli flag const
	SessionDurationFlag = "session-duration"
	// WriteAWSCredentialsFlag cli flag const
	WriteAWSCredentialsFlag = "write-aws-credentials"
	// LegacyAWSVariablesFlag cli flag const
	LegacyAWSVariablesFlag = "legacy-aws-variables"
	// ExpiryAWSVariablesFlag cli flag const
	ExpiryAWSVariablesFlag = "expiry-aws-variables"
	// CacheAccessTokenFlag cli flag const
	CacheAccessTokenFlag = "cache-access-token"

	// AWSCredentialsEnvVar env var const
	AWSCredentialsEnvVar = "OKTA_AWSCLI_AWS_CREDENTIALS"
	// AWSIAMIdPEnvVar env var const
	AWSIAMIdPEnvVar = "OKTA_AWSCLI_IAM_IDP"
	// AWSIAMRoleEnvVar env var const
	AWSIAMRoleEnvVar = "OKTA_AWSCLI_IAM_ROLE"
	// AWSSessionDurationEnvVar env var const
	AWSSessionDurationEnvVar = "OKTA_AWSCLI_SESSION_DURATION"
	// FormatEnvVar env var const
	FormatEnvVar = "OKTA_AWSCLI_FORMAT"
	// OktaOIDCClientIDEnvVar env var const
	OktaOIDCClientIDEnvVar = "OKTA_OIDC_CLIENT_ID"
	// OktaOrgDomainEnvVar env var const
	OktaOrgDomainEnvVar = "OKTA_ORG_DOMAIN"
	// OktaAWSAccountFederationAppIDEnvVar env var const
	OktaAWSAccountFederationAppIDEnvVar = "OKTA_AWS_ACCOUNT_FEDERATION_APP_ID"
	// OpenBrowserEnvVar env var const
	OpenBrowserEnvVar = "OKTA_AWSCLI_OPEN_BROWSER"
	// ProfileEnvVar env var const
	ProfileEnvVar = "OKTA_AWSCLI_PROFILE"
	// QRCodeEnvVar env var const
	QRCodeEnvVar = "OKTA_AWSCLI_QR_CODE"
	// WriteAWSCredentialsEnvVar env var const
	WriteAWSCredentialsEnvVar = "OKTA_AWSCLI_WRITE_AWS_CREDENTIALS"
	// DebugEnvVar env var const
	DebugEnvVar = "OKTA_AWSCLI_DEBUG"
	// DebugAPICallsEnvVar env var const
	DebugAPICallsEnvVar = "OKTA_AWSCLI_DEBUG_API_CALLS"
	// DebugConfigEnvVar env var const
	DebugConfigEnvVar = "OKTA_AWSCLI_DEBUG_CONFIG"
	// LegacyAWSVariablesEnvVar env var const
	LegacyAWSVariablesEnvVar = "OKTA_AWSCLI_LEGACY_AWS_VARIABLES"
	// ExpiryAWSVariablesEnvVar env var const
	ExpiryAWSVariablesEnvVar = "OKTA_AWSCLI_EXPIRY_AWS_VARIABLES"
	// CacheAccessTokenEnvVar env var const
	CacheAccessTokenEnvVar = "OKTA_AWSCLI_CACHE_ACCESS_TOKEN"

	// CannotBeBlankErrMsg error message const
	CannotBeBlankErrMsg = "cannot be blank"
	// OrgDomainMsg error message const
	OrgDomainMsg = "Org Domain"

	// DotOkta string const
	DotOkta = ".okta"
	// OktaYaml string const
	OktaYaml = "okta.yaml"
)

// Config A config object for the CLI
type Config struct {
	orgDomain           string
	oidcAppID           string
	fedAppID            string
	awsIAMIdP           string
	awsIAMRole          string
	awsSessionDuration  int64
	format              string
	profile             string
	qrCode              bool
	awsCredentials      string
	writeAWSCredentials bool
	openBrowser         bool
	debug               bool
	debugAPICalls       bool
	debugConfig         bool
	legacyAWSVariables  bool
	expiryAWSVariables  bool
	cacheAccessToken    bool
	httpClient          *http.Client
}

// OktaYamlConfig represents config settings from $HOME/.okta/okta.yaml
type OktaYamlConfig struct {
	AWSCLI struct {
		IDPS  map[string]string `yaml:"idps"`
		ROLES map[string]string `yaml:"roles"`
	} `yaml:"awscli"`
}

// Attributes config construction
type Attributes struct {
	OrgDomain           string
	OIDCAppID           string
	FedAppID            string
	AWSIAMIdP           string
	AWSIAMRole          string
	AWSSessionDuration  int64
	Format              string
	Profile             string
	QRCode              bool
	AWSCredentials      string
	WriteAWSCredentials bool
	OpenBrowser         bool
	Debug               bool
	DebugAPICalls       bool
	DebugConfig         bool
	LegacyAWSVariables  bool
	ExpiryAWSVariables  bool
	CacheAccessToken    bool
}

// CreateConfig Creates a new config gathering values in this order of precedence:
//  1. CLI flags
//  2. ENV variables
//  3. .env file
func CreateConfig() (*Config, error) {
	cfgAttrs, err := readConfig()
	if err != nil {
		return nil, err
	}
	return NewConfig(cfgAttrs)
}

// NewConfig create config from attributes
func NewConfig(attrs Attributes) (*Config, error) {
	var err error
	cfg := &Config{
		fedAppID:            attrs.FedAppID,
		awsIAMIdP:           attrs.AWSIAMIdP,
		awsIAMRole:          attrs.AWSIAMRole,
		format:              attrs.Format,
		profile:             attrs.Profile,
		qrCode:              attrs.QRCode,
		awsCredentials:      attrs.AWSCredentials,
		writeAWSCredentials: attrs.WriteAWSCredentials,
		openBrowser:         attrs.OpenBrowser,
		debug:               attrs.Debug,
		debugAPICalls:       attrs.DebugAPICalls,
		debugConfig:         attrs.DebugConfig,
		legacyAWSVariables:  attrs.LegacyAWSVariables,
		expiryAWSVariables:  attrs.ExpiryAWSVariables,
		cacheAccessToken:    attrs.CacheAccessToken,
	}
	if attrs.DebugConfig {
		return cfg, nil
	}
	err = cfg.SetOrgDomain(attrs.OrgDomain)
	if err != nil {
		return nil, err
	}
	err = cfg.SetOIDCAppID(attrs.OIDCAppID)
	if err != nil {
		return nil, err
	}
	err = cfg.SetAWSSessionDuration(attrs.AWSSessionDuration)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: newConfigTransport(cfg.DebugAPICalls()),
		Timeout:   time.Second * time.Duration(60),
	}
	err = cfg.SetHTTPClient(client)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func readConfig() (Attributes, error) {
	attrs := Attributes{
		AWSCredentials:      viper.GetString(AWSCredentialsFlag),
		AWSIAMIdP:           viper.GetString(AWSIAMIdPFlag),
		AWSIAMRole:          viper.GetString(AWSIAMRoleFlag),
		AWSSessionDuration:  viper.GetInt64(SessionDurationFlag),
		Debug:               viper.GetBool(DebugFlag),
		DebugAPICalls:       viper.GetBool(DebugAPICallsFlag),
		DebugConfig:         viper.GetBool(DebugConfigFlag),
		FedAppID:            viper.GetString(AWSAcctFedAppIDFlag),
		Format:              viper.GetString(FormatFlag),
		LegacyAWSVariables:  viper.GetBool(LegacyAWSVariablesFlag),
		ExpiryAWSVariables:  viper.GetBool(ExpiryAWSVariablesFlag),
		CacheAccessToken:    viper.GetBool(CacheAccessTokenFlag),
		OIDCAppID:           viper.GetString(OIDCClientIDFlag),
		OpenBrowser:         viper.GetBool(OpenBrowserFlag),
		OrgDomain:           viper.GetString(OrgDomainFlag),
		Profile:             viper.GetString(ProfileFlag),
		QRCode:              viper.GetBool(QRCodeFlag),
		WriteAWSCredentials: viper.GetBool(WriteAWSCredentialsFlag),
	}
	if attrs.Format == "" {
		attrs.Format = EnvVarFormat
	}

	// mimic AWS CLI behavior, if profile value is not set by flag check
	// the ENV VAR, else set to "default"
	if attrs.Profile == "" {
		attrs.Profile = viper.GetString(downCase(ProfileEnvVar))
	}
	if attrs.Profile == "" {
		attrs.Profile = "default"
	}

	// Viper binds ENV VARs to a lower snake version, set the configs with them
	// if they haven't already been set by cli flag binding.
	if attrs.OrgDomain == "" {
		attrs.OrgDomain = viper.GetString(downCase(OktaOrgDomainEnvVar))
	}
	if attrs.OIDCAppID == "" {
		attrs.OIDCAppID = viper.GetString(downCase(OktaOIDCClientIDEnvVar))
	}
	if attrs.FedAppID == "" {
		attrs.FedAppID = viper.GetString(downCase(OktaAWSAccountFederationAppIDEnvVar))
	}
	if attrs.AWSIAMIdP == "" {
		attrs.AWSIAMIdP = viper.GetString(downCase(AWSIAMIdPEnvVar))
	}
	if attrs.AWSIAMRole == "" {
		attrs.AWSIAMRole = viper.GetString(downCase(AWSIAMRoleEnvVar))
	}
	if !attrs.QRCode {
		attrs.QRCode = viper.GetBool(downCase(QRCodeEnvVar))
	}

	// if session duration is 0, inspect the ENV VAR for a value, else set
	// a default of 3600
	if attrs.AWSSessionDuration == 0 {
		attrs.AWSSessionDuration = viper.GetInt64(downCase(AWSSessionDurationEnvVar))
	}
	if attrs.AWSSessionDuration == 0 {
		attrs.AWSSessionDuration = 3600
	}

	// correct org domain if it's in admin form
	orgDomain := strings.Replace(attrs.OrgDomain, "-admin", "", -1)
	if orgDomain != attrs.OrgDomain {
		fmt.Printf("WARNING: proactively correcting org domain %q to non-admin form %q.\n\n", attrs.OrgDomain, orgDomain)
		attrs.OrgDomain = orgDomain
	}
	if strings.HasPrefix(attrs.OrgDomain, "http") {
		u, err := url.Parse(attrs.OrgDomain)
		// try to help correct org domain value if parsing occurs correctly,
		// else let the CLI error out else where
		if err == nil {
			orgDomain = u.Hostname()
			fmt.Printf("WARNING: proactively correcting URL format org domain %q value to hostname only form %q.\n\n", attrs.OrgDomain, orgDomain)
			attrs.OrgDomain = orgDomain
		}
	}
	if strings.HasSuffix(attrs.OrgDomain, "/") {
		orgDomain = string([]byte(attrs.OrgDomain)[0 : len(attrs.OrgDomain)-1])
		// try to help correct malformed org domain value
		fmt.Printf("WARNING: proactively correcting malformed org domain %q value to hostname only form %q.\n\n", attrs.OrgDomain, orgDomain)
		attrs.OrgDomain = orgDomain
	}

	// There is always a default aws credentials path set in root.go's init
	// function so overwrite the config value if the operator is attempting to
	// set it by ENV VAR value.
	if viper.GetString(downCase(AWSCredentialsEnvVar)) != "" {
		attrs.AWSCredentials = viper.GetString(downCase(AWSCredentialsEnvVar))
	}
	if !attrs.WriteAWSCredentials {
		attrs.WriteAWSCredentials = viper.GetBool(downCase(WriteAWSCredentialsEnvVar))
	}
	if attrs.WriteAWSCredentials {
		// writing aws creds option implies "aws-credentials" format
		attrs.Format = AWSCredentialsFormat
	}
	if !attrs.OpenBrowser {
		attrs.OpenBrowser = viper.GetBool(downCase(OpenBrowserEnvVar))
	}
	if !attrs.Debug {
		attrs.Debug = viper.GetBool(downCase(DebugEnvVar))
	}
	if !attrs.DebugAPICalls {
		attrs.DebugAPICalls = viper.GetBool(downCase(DebugAPICallsEnvVar))
	}
	if !attrs.LegacyAWSVariables {
		attrs.LegacyAWSVariables = viper.GetBool(downCase(LegacyAWSVariablesEnvVar))
	}
	if !attrs.ExpiryAWSVariables {
		attrs.ExpiryAWSVariables = viper.GetBool(downCase(ExpiryAWSVariablesEnvVar))
	}
	if !attrs.CacheAccessToken {
		attrs.CacheAccessToken = viper.GetBool(downCase(CacheAccessTokenEnvVar))
	}
	return attrs, nil
}

// downCase ToLower all alpha chars e.g. HELLO_WORLD -> hello_world
func downCase(s string) string {
	return strings.ToLower(s)
}

// OrgDomain --
func (c *Config) OrgDomain() string {
	return c.orgDomain
}

// SetOrgDomain --
func (c *Config) SetOrgDomain(domain string) error {
	if domain == "" {
		return NewValidationError(OrgDomainMsg, CannotBeBlankErrMsg)
	}
	c.orgDomain = domain
	return nil
}

// OIDCAppID --
func (c *Config) OIDCAppID() string {
	return c.oidcAppID
}

// SetOIDCAppID --
func (c *Config) SetOIDCAppID(appID string) error {
	if appID == "" {
		return NewValidationError("OIDC App ID", CannotBeBlankErrMsg)
	}
	c.oidcAppID = appID
	return nil
}

// FedAppID --
func (c *Config) FedAppID() string {
	return c.fedAppID
}

// SetFedAppID --
func (c *Config) SetFedAppID(appID string) error {
	c.fedAppID = appID
	return nil
}

// AWSIAMIdP --
func (c *Config) AWSIAMIdP() string {
	return c.awsIAMIdP
}

// SetAWSIAMIdP --
func (c *Config) SetAWSIAMIdP(idp string) error {
	c.awsIAMIdP = idp
	return nil
}

// AWSIAMRole --
func (c *Config) AWSIAMRole() string {
	return c.awsIAMRole
}

// SetAWSIAMRole --
func (c *Config) SetAWSIAMRole(role string) error {
	c.awsIAMRole = role
	return nil
}

// AWSSessionDuration --
func (c *Config) AWSSessionDuration() int64 {
	return c.awsSessionDuration
}

// SetAWSSessionDuration --
func (c *Config) SetAWSSessionDuration(duration int64) error {
	if duration < 60 || duration > 43200 {
		return NewValidationError("AWS Session Duration", "must be between 60 and 43200")
	}
	c.awsSessionDuration = duration
	return nil
}

// Format --
func (c *Config) Format() string {
	return c.format
}

// SetFormat --
func (c *Config) SetFormat(format string) error {
	c.format = format
	return nil
}

// Profile --
func (c *Config) Profile() string {
	return c.profile
}

// SetProfile --
func (c *Config) SetProfile(profile string) error {
	c.profile = profile
	return nil
}

// QRCode --
func (c *Config) QRCode() bool {
	return c.qrCode
}

// SetQRCode --
func (c *Config) SetQRCode(qrCode bool) error {
	c.qrCode = qrCode
	return nil
}

// AWSCredentials --
func (c *Config) AWSCredentials() string {
	return c.awsCredentials
}

// SetAWSCredentials --
func (c *Config) SetAWSCredentials(credentials string) error {
	c.awsCredentials = credentials
	return nil
}

// WriteAWSCredentials --
func (c *Config) WriteAWSCredentials() bool {
	return c.writeAWSCredentials
}

// SetWriteAWSCredentials --
func (c *Config) SetWriteAWSCredentials(writeCredentials bool) error {
	c.writeAWSCredentials = writeCredentials
	return nil
}

// OpenBrowser --
func (c *Config) OpenBrowser() bool {
	return c.openBrowser
}

// SetOpenBrowser --
func (c *Config) SetOpenBrowser(openBrowser bool) error {
	c.openBrowser = openBrowser
	return nil
}

// Debug --
func (c *Config) Debug() bool {
	return c.debug
}

// SetDebug --
func (c *Config) SetDebug(debug bool) error {
	c.debug = debug
	return nil
}

// DebugAPICalls --
func (c *Config) DebugAPICalls() bool {
	return c.debugAPICalls
}

// SetDebugAPICalls --
func (c *Config) SetDebugAPICalls(debugAPICalls bool) error {
	c.debugAPICalls = debugAPICalls
	return nil
}

// DebugConfig --
func (c *Config) DebugConfig() bool {
	return c.debugConfig
}

// SetDebugConfig --
func (c *Config) SetDebugConfig(debugConfig bool) error {
	c.debugConfig = debugConfig
	return nil
}

// LegacyAWSVariables --
func (c *Config) LegacyAWSVariables() bool {
	return c.legacyAWSVariables
}

// SetLegacyAWSVariables --
func (c *Config) SetLegacyAWSVariables(legacyAWSVariables bool) error {
	c.legacyAWSVariables = legacyAWSVariables
	return nil
}

// ExpiryAWSVariables --
func (c *Config) ExpiryAWSVariables() bool {
	return c.expiryAWSVariables
}

// SetExpiryAWSVariables --
func (c *Config) SetExpiryAWSVariables(expiryAWSVariables bool) error {
	c.expiryAWSVariables = expiryAWSVariables
	return nil
}

// CacheAccessToken --
func (c *Config) CacheAccessToken() bool {
	return c.cacheAccessToken
}

// SetCacheAccessToken --
func (c *Config) SetCacheAccessToken(cacheAccessToken bool) error {
	c.cacheAccessToken = cacheAccessToken
	return nil
}

// HTTPClient --
func (c *Config) HTTPClient() *http.Client {
	return c.httpClient
}

// SetHTTPClient --
func (c *Config) SetHTTPClient(client *http.Client) error {
	c.httpClient = client
	return nil
}

// OktaConfig returns an Okta YAML Config object representation of $HOME/.okta/okta.yaml
func (c *Config) OktaConfig() (config *OktaYamlConfig, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	configPath := filepath.Join(homeDir, DotOkta, OktaYaml)
	yamlConfig, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	conf := OktaYamlConfig{}
	err = yaml.Unmarshal(yamlConfig, &conf)
	if err != nil {
		return
	}
	config = &conf

	return
}

// RunConfigChecks runs a series of checks on the okta.yaml config file
func (c *Config) RunConfigChecks() (err error) {
	exampleYaml := `
---
awscli:
  idps:
    "arn:aws:iam::123456789012:saml-provider/company-okta-idp": "Data Production"
    "arn:aws:iam::012345678901:saml-provider/company-okta-idp": "Data Development"
  roles:
    "arn:aws:iam::123456789012:role/admin": "Prod Admin"
    "arn:aws:iam::123456789012:role/operator": "Prod Ops"
    "arn:aws:iam::012345678901:role/admin": "Dev Admin"
    "arn:aws:iam::012345678901:role/operator": "Dev Ops"
	`
	fmt.Fprintf(os.Stderr, "Given this YAML as an example template of okta.yaml for reference:\n%s\n", exampleYaml)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: can't find user home directory $HOME\n")
		fmt.Fprintf(os.Stderr, "         see https://pkg.go.dev/os#UserHomeDir\n")
		return
	}
	fmt.Fprintf(os.Stderr, "found home directory %q\n", homeDir)

	configPath := filepath.Join(homeDir, DotOkta, OktaYaml)
	yamlConfig, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: can't read okta config %q\n", configPath)
		return
	}
	fmt.Fprintf(os.Stderr, "okta.yaml is readable %q\n", configPath)

	conf := map[string]any{}
	err = yaml.Unmarshal(yamlConfig, &conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml is invalid yaml format\n")
		return
	}
	fmt.Fprintf(os.Stderr, "okta.yaml is valid yaml\n")

	awscli, ok := conf["awscli"]
	if !ok {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml missing \"awscli\" section\n")
		return
	}
	fmt.Fprintf(os.Stderr, "okta.yaml has root \"awscli\" section\n")

	if awscli == nil {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml \"awscli\" section has no values\n")
		return
	}
	_awscli, ok := awscli.(map[any]any)
	if !ok {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml \"awscli\" is not a map of values\n")
	}
	idps, ok := _awscli["idps"]
	if !ok {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml missing \"awscli.idps\" section\n")
		return
	}
	if idps == nil {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml \"awscli.idps\" section has no values\n")
		return
	}

	// map[interface {}]interface {}
	_idps, ok := idps.(map[any]any)
	if !ok {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml \"awscli.idps\" section is not a map of ARN string key to friendly string label values\n")
		return
	}
	if len(_idps) == 0 {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml \"awscli.idps\" section is an empty map of ARN string key to friendly string label values\n")
		return
	}

	for k, v := range _idps {
		if _, ok := k.(string); !ok {
			fmt.Fprintf(os.Stderr, "okta.yaml \"awscli.idps\" value of ARN key \"%v\" is not a string\n", k)
			return
		}
		if _, ok := v.(string); !ok {
			fmt.Fprintf(os.Stderr, "okta.yaml \"awscli.idps\" ARN key %q's friendly label value \"%v\" is not a string\n", k, v)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "okta.yaml \"awscli.idps\" section is a map of %d ARN string keys to friendly string label values\n", len(_idps))

	roles, ok := _awscli["roles"]
	if !ok {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml missing \"awscli.roles\" section\n")
		return
	}
	if roles == nil {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml \"awscli.roles\" section has no values\n")
		return
	}

	_roles, ok := roles.(map[any]any)
	if !ok {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml \"awscli.roles\" section is not a map of ARN string key to friendly string label values\n")
		return
	}
	if len(_roles) == 0 {
		fmt.Fprintf(os.Stderr, "WARNING: okta.yaml \"awscli.roles\" section is an empty map of ARN string key to friendly string label values\n")
		return
	}

	for k, v := range _roles {
		if _, ok := k.(string); !ok {
			fmt.Fprintf(os.Stderr, "okta.yaml \"awscli.roles\" value of ARN key \"%v\" is not a string\n", k)
			return
		}
		if _, ok := v.(string); !ok {
			fmt.Fprintf(os.Stderr, "okta.yaml \"awscli.roles\" ARN key %q's friendly label value \"%v\" is not a string\n", k, v)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "okta.yaml \"awscli.roles\" section is a map of %d ARN string keys to friendly string label values\n", len(_roles))

	fmt.Fprintf(os.Stderr, "okta.yaml is OK\n")
	return nil
}
