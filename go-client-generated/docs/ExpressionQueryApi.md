# {{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryExemplarsGET**](ExpressionQueryApi.md#QueryExemplarsGET) | **Get** /query_exemplars | Returns list of Exemplars
[**QueryExemplarsPOST**](ExpressionQueryApi.md#QueryExemplarsPOST) | **Post** /query_exemplars | Returns list of Exemplars
[**QueryGET**](ExpressionQueryApi.md#QueryGET) | **Get** /query | Evaluates instant query
[**QueryPOST**](ExpressionQueryApi.md#QueryPOST) | **Post** /query | Evaluates instant query
[**QueryRangeGET**](ExpressionQueryApi.md#QueryRangeGET) | **Get** /query_range | Evaluates query over range of time.
[**QueryRangePOST**](ExpressionQueryApi.md#QueryRangePOST) | **Post** /query_range | Evaluates query over range of time.

# **QueryExemplarsGET**
> ResponseQueryExemplars QueryExemplarsGET(ctx, query, optional)
Returns list of Exemplars

This is <b>experimental</b> and might change in the future. The following endpoint returns a list of exemplars for a valid PromQL query for a specific time range 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| Prometheus expression query string.  Example: &#x60;&#x60;&#x60;?query&#x3D;test_exemplar_metric_total&#x60;&#x60;&#x60;  | 
 **optional** | ***ExpressionQueryApiQueryExemplarsGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExpressionQueryApiQueryExemplarsGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp.  Example: &#x60;&#x60;&#x60;&amp;start&#x3D;2020-09-14T15:22:25.479Z&#x60;&#x60;&#x60;  | 
 **end** | **optional.String**| End timestamp.  Example: &#x60;&#x60;&#x60;&amp;end&#x3D;020-09-14T15:23:25.479Z&#x60;&#x60;&#x60;  | 

### Return type

[**ResponseQueryExemplars**](responseQuery_exemplars.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryExemplarsPOST**
> ResponseQueryExemplars QueryExemplarsPOST(ctx, query, optional)
Returns list of Exemplars

This is <b>experimental</b> and might change in the future. The following endpoint returns a list of exemplars for a valid PromQL query for a specific time range 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| Prometheus expression query string.  Example: &#x60;&#x60;&#x60;?query&#x3D;test_exemplar_metric_total&#x60;&#x60;&#x60;  | 
 **optional** | ***ExpressionQueryApiQueryExemplarsPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExpressionQueryApiQueryExemplarsPOSTOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp.  Example: &#x60;&#x60;&#x60;&amp;start&#x3D;2020-09-14T15:22:25.479Z&#x60;&#x60;&#x60;  | 
 **end** | **optional.String**| End timestamp.  Example: &#x60;&#x60;&#x60;&amp;end&#x3D;020-09-14T15:23:25.479Z&#x60;&#x60;&#x60;  | 

### Return type

[**ResponseQueryExemplars**](responseQuery_exemplars.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryGET**
> QueryData QueryGET(ctx, query, optional)
Evaluates instant query

The following endpoint evaluates an instant query at a single point in time  You can URL-encode these parameters directly in the request body by using the ```POST``` method and ```Content-Type: application/x-www-form-urlencoded``` header. This is useful when specifying a large query that may breach server-side URL character limits.  The data section of the query result has the following format ``` {   \"resultType\": \"matrix\" | \"vector\" | \"scalar\" | \"string\",   \"result\": <value> } ``` ```<value>``` refers to the query result data, which has varying formats depending on the ```resultType```. See the [expression query result formats](https://prometheus.io/docs/prometheus/latest/querying/api/#expression-query-result-formats). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| Prometheus expression query string.  Example: &#x60;&#x60;&#x60;?query&#x3D;up&#x60;&#x60;&#x60;  | 
 **optional** | ***ExpressionQueryApiQueryGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExpressionQueryApiQueryGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.String**| Evaluation timestamp. Optional.  The current server time is used if the &#x60;&#x60;&#x60;time&#x60;&#x60;&#x60; parameter is omitted.  Example: &#x60;&#x60;&#x60;?metric&#x3D;http_requests_total&#x60;&#x60;&#x60;  | 
 **timeout** | **optional.String**| Evaluation timeout. Optional. Defaults to and is capped by the value of the &#x60;&#x60;&#x60;-query.timeout&#x60;&#x60;&#x60; flag.  Example: &#x60;&#x60;&#x60;?metric&#x3D;http_requests_total&#x60;&#x60;&#x60;  | 

### Return type

[**QueryData**](queryData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryPOST**
> QueryData QueryPOST(ctx, query, optional)
Evaluates instant query

The following endpoint evaluates an instant query at a single point in time  You can URL-encode these parameters directly in the request body by using the ```POST``` method and ```Content-Type: application/x-www-form-urlencoded``` header. This is useful when specifying a large query that may breach server-side URL character limits.  The data section of the query result has the following format ``` {   \"resultType\": \"matrix\" | \"vector\" | \"scalar\" | \"string\",   \"result\": <value> } ``` ```<value>``` refers to the query result data, which has varying formats depending on the ```resultType```. See the [expression query result formats](https://prometheus.io/docs/prometheus/latest/querying/api/#expression-query-result-formats). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| Prometheus expression query string.  Example: &#x60;&#x60;&#x60;?query&#x3D;up&#x60;&#x60;&#x60;  | 
 **optional** | ***ExpressionQueryApiQueryPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExpressionQueryApiQueryPOSTOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **optional.String**| Evaluation timestamp. Optional.  The current server time is used if the &#x60;&#x60;&#x60;time&#x60;&#x60;&#x60; parameter is omitted.  Example: &#x60;&#x60;&#x60;?metric&#x3D;http_requests_total&#x60;&#x60;&#x60;  | 
 **timeout** | **optional.String**| Evaluation timeout. Optional. Defaults to and is capped by the value of the &#x60;&#x60;&#x60;-query.timeout&#x60;&#x60;&#x60; flag.  Example: &#x60;&#x60;&#x60;?metric&#x3D;http_requests_total&#x60;&#x60;&#x60;  | 

### Return type

[**QueryData**](queryData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryRangeGET**
> ResponseQueryRange QueryRangeGET(ctx, query, optional)
Evaluates query over range of time.

The following endpoint evaluates an expression query over a range of time  You can URL-encode these parameters directly in the request body by using the ```POST``` method and ```Content-Type: application/x-www-form-urlencoded``` header. This is useful when specifying a large query that may breach server-side URL character limits.  The data section of the query result has the following format ``` {   \"resultType\": \"matrix\",   \"result\": <value> } ``` For the format of the ```<value>``` placeholder, see the [range-vector result format](https://prometheus.io/docs/prometheus/latest/querying/api/#range-vectors). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| Prometheus expression query string.  Example: &#x60;&#x60;&#x60;?query&#x3D;up&#x60;&#x60;&#x60;  | 
 **optional** | ***ExpressionQueryApiQueryRangeGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExpressionQueryApiQueryRangeGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp.  Example: &#x60;&#x60;&#x60;&amp;start&#x3D;2015-07-01T20:10:30.781Z&#x60;&#x60;&#x60;  | 
 **end** | **optional.String**| End timestamp.  Example: &#x60;&#x60;&#x60;&amp;end&#x3D;2015-07-01T20:11:00.781Z&#x60;&#x60;&#x60;  | 
 **step** | **optional.String**| Query resolution step width in &#x60;&#x60;&#x60;duration&#x60;&#x60;&#x60; format or float number of seconds.  Example: &#x60;&#x60;&#x60;&amp;step&#x3D;15s&#x60;&#x60;&#x60;  | 
 **timeout** | **optional.String**| Evaluation timeout. Optional. Defaults to and is capped by the value of the &#x60;&#x60;&#x60;-query.timeout&#x60;&#x60;&#x60; flag.  Example: &#x60;&#x60;&#x60;?metric&#x3D;http_requests_total&#x60;&#x60;&#x60;  | 

### Return type

[**ResponseQueryRange**](responseQuery_range.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryRangePOST**
> ResponseQueryRange QueryRangePOST(ctx, query, optional)
Evaluates query over range of time.

The following endpoint evaluates an expression query over a range of time  You can URL-encode these parameters directly in the request body by using the ```POST``` method and ```Content-Type: application/x-www-form-urlencoded``` header. This is useful when specifying a large query that may breach server-side URL character limits.  The data section of the query result has the following format ``` {   \"resultType\": \"matrix\",   \"result\": <value> } ``` For the format of the ```<value>``` placeholder, see the [range-vector result format](https://prometheus.io/docs/prometheus/latest/querying/api/#range-vectors). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| Prometheus expression query string.  Example: &#x60;&#x60;&#x60;?query&#x3D;up&#x60;&#x60;&#x60;  | 
 **optional** | ***ExpressionQueryApiQueryRangePOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExpressionQueryApiQueryRangePOSTOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp.  Example: &#x60;&#x60;&#x60;&amp;start&#x3D;2015-07-01T20:10:30.781Z&#x60;&#x60;&#x60;  | 
 **end** | **optional.String**| End timestamp.  Example: &#x60;&#x60;&#x60;&amp;end&#x3D;2015-07-01T20:11:00.781Z&#x60;&#x60;&#x60;  | 
 **step** | **optional.String**| Query resolution step width in &#x60;&#x60;&#x60;duration&#x60;&#x60;&#x60; format or float number of seconds.  Example: &#x60;&#x60;&#x60;&amp;step&#x3D;15s&#x60;&#x60;&#x60;  | 
 **timeout** | **optional.String**| Evaluation timeout. Optional. Defaults to and is capped by the value of the &#x60;&#x60;&#x60;-query.timeout&#x60;&#x60;&#x60; flag.  Example: &#x60;&#x60;&#x60;?metric&#x3D;http_requests_total&#x60;&#x60;&#x60;  | 

### Return type

[**ResponseQueryRange**](responseQuery_range.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

