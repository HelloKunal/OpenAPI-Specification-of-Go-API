# {{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LabelNamesGET**](QueryingMetadataApi.md#LabelNamesGET) | **Get** /labels | Returns label names
[**LabelNamesPOST**](QueryingMetadataApi.md#LabelNamesPOST) | **Post** /labels | Returns label names
[**LabelValuesGET**](QueryingMetadataApi.md#LabelValuesGET) | **Get** /label/{label_name}/values | Returns label values
[**SeriesDELETE**](QueryingMetadataApi.md#SeriesDELETE) | **Delete** /series | Returns time series
[**SeriesGET**](QueryingMetadataApi.md#SeriesGET) | **Get** /series | Returns time series
[**SeriesPOST**](QueryingMetadataApi.md#SeriesPOST) | **Post** /series | Returns time series

# **LabelNamesGET**
> []string LabelNamesGET(ctx, optional)
Returns label names

The following endpoint returns a list of label names  The ```data``` section of the JSON response is a list of string label names.  --- **NOTE:** These API endpoints may return metadata for series for which there is no sample within the selected time range, and/or for series whose samples have been marked as deleted via the deletion API endpoint. The exact extent of additionally returned series metadata is an implementation detail that may change in the future.  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryingMetadataApiLabelNamesGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryingMetadataApiLabelNamesGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.String**| Start timestamp. Optional.  | 
 **end** | **optional.String**| End timestamp. Optional.  | 
 **match** | **optional.String**| Repeated series selector argument that selects the series from which to read the label values. Optional.  | 

### Return type

[**[]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelNamesPOST**
> []string LabelNamesPOST(ctx, optional)
Returns label names

The following endpoint returns a list of label names  The ```data``` section of the JSON response is a list of string label names. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryingMetadataApiLabelNamesPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryingMetadataApiLabelNamesPOSTOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.String**| Start timestamp. Optional.  | 
 **end** | **optional.String**| End timestamp. Optional.  | 
 **match** | **optional.String**| Repeated series selector argument that selects the series from which to read the label values. Optional.  | 

### Return type

[**[]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelValuesGET**
> []string LabelValuesGET(ctx, labelName, optional)
Returns label values

The following endpoint returns a list of label values for a provided label name  The ```data``` section of the JSON response is a list of string label values.  --- **NOTE:** These API endpoints may return metadata for series for which there is no sample within the selected time range, and/or for series whose samples have been marked as deleted via the deletion API endpoint. The exact extent of additionally returned series metadata is an implementation detail that may change in the future.  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelName** | **string**| Label name  Example: &#x60;&#x60;&#x60;/label/job/values&#x60;&#x60;&#x60;  | 
 **optional** | ***QueryingMetadataApiLabelValuesGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryingMetadataApiLabelValuesGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp. Optional.  | 
 **end** | **optional.String**| End timestamp. Optional.  | 
 **match** | **optional.String**| Repeated series selector argument that selects the series from which to read the label values. Optional.  | 

### Return type

[**[]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SeriesDELETE**
> []QueryDataResultMetric SeriesDELETE(ctx, match, optional)
Returns time series

The following endpoint returns the list of time series that match a certain label set.  You can URL-encode these parameters directly in the request body by using the ```POST``` method and ```Content-Type: application/x-www-form-urlencoded``` header. This is useful when specifying a large or dynamic number of series selectors that may breach server-side URL character limits.  The ```data``` section of the query result consists of a list of objects that contain the label name/value pairs which identify each series.  --- **NOTE:** These API endpoints may return metadata for series for which there is no sample within the selected time range, and/or for series whose samples have been marked as deleted via the deletion API endpoint. The exact extent of additionally returned series metadata is an implementation detail that may change in the future.  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **match** | **string**| Repeated series selector argument that selects the series to return. At least one &#x60;&#x60;&#x60;match[]&#x60;&#x60;&#x60; argument must be provided.  Example: &#x60;&#x60;&#x60;?&#x27; --data-urlencode &#x27;match[]&#x3D;up&#x27;&#x60;&#x60;&#x60;  | 
 **optional** | ***QueryingMetadataApiSeriesDELETEOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryingMetadataApiSeriesDELETEOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp. Optional.  | 
 **end** | **optional.String**| End timestamp. Optional.  | 

### Return type

[**[]QueryDataResultMetric**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SeriesGET**
> []QueryDataResultMetric SeriesGET(ctx, match, optional)
Returns time series

The following endpoint returns the list of time series that match a certain label set.  You can URL-encode these parameters directly in the request body by using the ```POST``` method and ```Content-Type: application/x-www-form-urlencoded``` header. This is useful when specifying a large or dynamic number of series selectors that may breach server-side URL character limits.  The ```data``` section of the query result consists of a list of objects that contain the label name/value pairs which identify each series. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **match** | **string**| Repeated series selector argument that selects the series to return. At least one &#x60;&#x60;&#x60;match[]&#x60;&#x60;&#x60; argument must be provided.  Example: &#x60;&#x60;&#x60;?&#x27; --data-urlencode &#x27;match[]&#x3D;up&#x27;&#x60;&#x60;&#x60;  | 
 **optional** | ***QueryingMetadataApiSeriesGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryingMetadataApiSeriesGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp. Optional.  | 
 **end** | **optional.String**| End timestamp. Optional.  | 

### Return type

[**[]QueryDataResultMetric**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SeriesPOST**
> []QueryDataResultMetric SeriesPOST(ctx, match, optional)
Returns time series

The following endpoint returns the list of time series that match a certain label set.  You can URL-encode these parameters directly in the request body by using the ```POST``` method and ```Content-Type: application/x-www-form-urlencoded``` header. This is useful when specifying a large or dynamic number of series selectors that may breach server-side URL character limits.  The ```data``` section of the query result consists of a list of objects that contain the label name/value pairs which identify each series. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **match** | **string**| Repeated series selector argument that selects the series to return. At least one &#x60;&#x60;&#x60;match[]&#x60;&#x60;&#x60; argument must be provided.  Example: &#x60;&#x60;&#x60;?&#x27; --data-urlencode &#x27;match[]&#x3D;up&#x27;&#x60;&#x60;&#x60;  | 
 **optional** | ***QueryingMetadataApiSeriesPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryingMetadataApiSeriesPOSTOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp. Optional.  | 
 **end** | **optional.String**| End timestamp. Optional.  | 

### Return type

[**[]QueryDataResultMetric**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

