/*
 * oxauth_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package fido2_pkg

import "oxauth_lib/configuration_pkg"
import "oxauth_lib/models_pkg"

/*
 * Interface for the FIDO2_IMPL
 */
type FIDO2 interface {
    GetFido2Configuration () (*models_pkg.FIDO2Configuration, error)

    CreateOptions (*models_pkg.AssertionOptions) (*models_pkg.AssertionOptionsResponse, error)

    CreateResult (*models_pkg.AssertionOptions1) (error)

    CreateAttestationOptions (*models_pkg.AttestationOptions) (*models_pkg.CredentialCreationOptions, error)

    CreateAttestationResult (*models_pkg.AttestationOptions1) (*models_pkg.ErrorResponse, error)
}

/*
 * Factory for the FIDO2 interaface returning FIDO2_IMPL
 */
func NewFIDO2(config configuration_pkg.CONFIGURATION) *FIDO2_IMPL {
    client := new(FIDO2_IMPL)
    client.config = config
    return client
}