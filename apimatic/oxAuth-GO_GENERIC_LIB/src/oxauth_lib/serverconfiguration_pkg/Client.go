/*
 * oxauth_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package serverconfiguration_pkg


import(
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"oxauth_lib/apihelper_pkg"
	"oxauth_lib/configuration_pkg"
	"oxauth_lib/models_pkg"
)
/*
 * Client structure as interface implementation
 */
type SERVERCONFIGURATION_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Gets GluuServer configuration data that contains non-standard OpenID Connect discovery metadata.
 * @return	Returns the *models_pkg.GluuConfigurationResponse response from the API call
 */
func (me *SERVERCONFIGURATION_IMPL) CreateWellKnownGluuConfiguration () (*models_pkg.GluuConfigurationResponse, error) {
    //the endpoint path uri
    _pathUrl := "/.well-known/gluu-configuration"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.ENUM_DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, nil)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal error occured. Please check log file for details.", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GluuConfigurationResponse = &models_pkg.GluuConfigurationResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}
