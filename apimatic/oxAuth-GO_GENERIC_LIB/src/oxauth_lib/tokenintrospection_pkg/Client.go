/*
 * oxauth_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package tokenintrospection_pkg


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
type TOKENINTROSPECTION_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * The Introspection OAuth 2 Endpoint for RPT.
 * @param    string         authorization       parameter: Required
 * @param    string         token               parameter: Required
 * @param    *string        tokenTypeHint       parameter: Optional
 * @return	Returns the *models_pkg.RptIntrospectionResponse response from the API call
 */
func (me *TOKENINTROSPECTION_IMPL) GetRptStatus (
            authorization string,
            token string,
            tokenTypeHint *string) (*models_pkg.RptIntrospectionResponse, error) {
    //the endpoint path uri
    _pathUrl := "/rpt/status"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.ENUM_DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "token" : token,
        "token_type_hint" : tokenTypeHint,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

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
        "Authorization" : apihelper_pkg.ToString(authorization, ""),
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
    if (_response.Code == 405) {
        err = apihelper_pkg.NewAPIError("Introspection of RPT is not allowed.", _response.Code, _response.RawBody)
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
    var retVal *models_pkg.RptIntrospectionResponse = &models_pkg.RptIntrospectionResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * The Introspection OAuth 2 Endpoint for RPT.
 * @param    string         authorization       parameter: Required
 * @param    string         token               parameter: Required
 * @param    *string        tokenTypeHint       parameter: Optional
 * @return	Returns the *models_pkg.RptIntrospectionResponse1 response from the API call
 */
func (me *TOKENINTROSPECTION_IMPL) PostRptStatus (
            authorization string,
            token string,
            tokenTypeHint *string) (*models_pkg.RptIntrospectionResponse1, error) {
    //the endpoint path uri
    _pathUrl := "/rpt/status"

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
        "Authorization" : apihelper_pkg.ToString(authorization, ""),
    }

    //form parameters
    parameters := map[string]interface{} {

        "token" : token,
        "token_type_hint" : tokenTypeHint,

    }


    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 405) {
        err = apihelper_pkg.NewAPIError("Introspection of RPT is not allowed.", _response.Code, _response.RawBody)
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
    var retVal *models_pkg.RptIntrospectionResponse1 = &models_pkg.RptIntrospectionResponse1{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}
