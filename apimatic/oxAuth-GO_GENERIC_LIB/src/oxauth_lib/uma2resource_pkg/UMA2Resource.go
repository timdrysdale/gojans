/*
 * oxauth_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package uma2resource_pkg

import "oxauth_lib/configuration_pkg"
import "oxauth_lib/models_pkg"

/*
 * Interface for the UMA2RESOURCE_IMPL
 */
type UMA2RESOURCE interface {
    PostHostRsrcResourceSet (string, *models_pkg.UmaResource) (*models_pkg.UmaResourceResponse, error)

    PutHostRsrcResourceSetRsid (string, string, *models_pkg.UmaResource) (*models_pkg.UmaResourceResponse, error)

    GetHostRsrcResourceSet (string, *string) ([]string, error)

    DeleteHostRsrcResourceSet (string, string) (error)

    GetHostRsrcResourceSetRsid (string, string) (*models_pkg.UmaResourceWithId, error)
}

/*
 * Factory for the UMA2RESOURCE interaface returning UMA2RESOURCE_IMPL
 */
func NewUMA2RESOURCE(config configuration_pkg.CONFIGURATION) *UMA2RESOURCE_IMPL {
    client := new(UMA2RESOURCE_IMPL)
    client.config = config
    return client
}