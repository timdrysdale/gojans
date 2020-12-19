/*
 * oxauth_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package serverconfiguration_pkg

import "oxauth_lib/configuration_pkg"
import "oxauth_lib/models_pkg"

/*
 * Interface for the SERVERCONFIGURATION_IMPL
 */
type SERVERCONFIGURATION interface {
    CreateWellKnownGluuConfiguration () (*models_pkg.GluuConfigurationResponse, error)
}

/*
 * Factory for the SERVERCONFIGURATION interaface returning SERVERCONFIGURATION_IMPL
 */
func NewSERVERCONFIGURATION(config configuration_pkg.CONFIGURATION) *SERVERCONFIGURATION_IMPL {
    client := new(SERVERCONFIGURATION_IMPL)
    client.config = config
    return client
}