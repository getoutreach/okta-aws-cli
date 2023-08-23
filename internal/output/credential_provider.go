/*
 * Copyright (c) 2023-Present, Okta, Inc.
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

package output

import (
	"encoding/json"
	"fmt"

	"github.com/okta/okta-aws-cli/internal/aws"
	"github.com/okta/okta-aws-cli/internal/config"
	"github.com/pkg/errors"
)

// CredentialProvider defines how to output AWS credentials in the
// format used by the AWS credentials provider feature.
type CredentialProvider struct {
	Version         uint   `json:"Version"`
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"Expiration"`
}

// NewCredentialProvider Creates a new
func NewCredentialProvider(expiry string) *CredentialProvider {
	return &CredentialProvider{
		Version:    1,
		Expiration: expiry,
	}
}

// Output Satisfies the Outputter interface and outputs credential
// provider JSON to STDOUT.
func (e *CredentialProvider) Output(c *config.Config, ac *aws.Credential) error {
	e.AccessKeyID = ac.AccessKeyID
	e.SecretAccessKey = ac.SecretAccessKey
	e.SessionToken = ac.SessionToken
	serialized, err := json.Marshal(e)
	if err != nil {
		return errors.Wrap(err, "could not serialize credential provider")
	}
	_, err = fmt.Print(string(serialized))
	return err
}
