/*
 * oxauth_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package umascope_pkg


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
type UMASCOPE_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Provides UMA Scope description by scope id.
 * @param    string        id     parameter: Required
 * @return	Returns the *models_pkg.UmaScopeDescription response from the API call
 */
func (me *UMASCOPE_IMPL) GetUmaScopes (
            id string) (*models_pkg.UmaScopeDescription, error) {
    //the endpoint path uri
    _pathUrl := "/uma/scopes"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

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
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("Invalid parameters provided to endpoint.", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Invalid parameters provided to endpoint.", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.UmaScopeDescription = &models_pkg.UmaScopeDescription{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

