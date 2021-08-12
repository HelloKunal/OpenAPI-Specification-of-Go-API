# {{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertManagersGET**](DefaultApi.md#AlertManagersGET) | **Get** /alertmanagers | Returns current alertmanager discovery
[**AlertsGET**](DefaultApi.md#AlertsGET) | **Get** /alerts | Returns active alerts
[**MetricMetadataGET**](DefaultApi.md#MetricMetadataGET) | **Get** /metadata | Returns metric metadata
[**RulesGET**](DefaultApi.md#RulesGET) | **Get** /rules | Returns currently loaded rules
[**TargetMetadataGET**](DefaultApi.md#TargetMetadataGET) | **Get** /targets/metadata | Returns target metadata
[**TargetsGET**](DefaultApi.md#TargetsGET) | **Get** /targets | Returns current target discovery.

# **AlertManagersGET**
> AlertmanagerDiscovery AlertManagersGET(ctx, )
Returns current alertmanager discovery

Returns an overview of the current state of the Prometheus alertmanager discovery  Both the active and dropped Alertmanagers are part of the response. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AlertmanagerDiscovery**](AlertmanagerDiscovery.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertsGET**
> Alert AlertsGET(ctx, )
Returns active alerts

The /alerts endpoint returns a list of all active alerts.  As the /alerts endpoint is fairly new, it does not have the same stability guarantees as the overarching API v1. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Alert**](Alert.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MetricMetadataGET**
> map[string]Metadata MetricMetadataGET(ctx, limit, optional)
Returns metric metadata

It returns metadata about metrics currently scrapped from targets. However, it does not provide any target information. This is considered experimental and might change in the future.  The data section of the query result consists of an object where each key is a metric name and each value is a list of unique metadata objects, as exposed for that metric name across all targets. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **float64**| Maximum number of metrics to return.  Example: &#x60;&#x60;&#x60;?limit&#x3D;2&#x60;&#x60;&#x60;  | 
 **optional** | ***DefaultApiMetricMetadataGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiMetricMetadataGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metric** | **optional.String**| A metric name to filter metadata for. All metric metadata is retrieved if left empty.  Example: &#x60;&#x60;&#x60;?metric&#x3D;http_requests_total&#x60;&#x60;&#x60;  | 

### Return type

[**map[string]Metadata**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RulesGET**
> RuleDiscovery RulesGET(ctx, optional)
Returns currently loaded rules

The ```/rules``` API endpoint returns a list of alerting and recording rules that are currently loaded. In addition it returns the currently active alerts fired by the Prometheus instance of each alerting rule.  As the ```/rules``` endpoint is fairly new, it does not have the same stability guarantees as the overarching API v1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiRulesGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRulesGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Return only the alerting rules (e.g. &#x60;&#x60;&#x60;type&#x3D;alert&#x60;&#x60;&#x60;) or the recording rules (e.g. &#x60;&#x60;&#x60;type&#x3D;record&#x60;&#x60;&#x60;). When the parameter is absent or empty, no filtering is done.  | 

### Return type

[**RuleDiscovery**](RuleDiscovery.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetMetadataGET**
> []MetricMetadata TargetMetadataGET(ctx, optional)
Returns target metadata

The following endpoint returns metadata about metrics currently scraped from targets. This is experimental and might change in the future.  The ```data``` section of the query result consists of a list of objects that contain metric metadata and the target label set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiTargetMetadataGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTargetMetadataGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **matchTarget** | **optional.String**| Label selectors that match targets by their label sets. All targets are selected if left empty.  Example: &#x60;&#x60;&#x60;match_target&#x3D;{job&#x3D;\&quot;prometheus\&quot;}&#x60;&#x60;&#x60;  | 
 **metric** | **optional.String**| A metric name to retrieve metadata for. All metric metadata is retrieved if left empty.  Example: &#x60;&#x60;&#x60;metric&#x3D;go_goroutines&#x60;&#x60;&#x60;  | 
 **limit** | **optional.Float64**| Maximum number of targets to match.  Example: &#x60;&#x60;&#x60;limit&#x3D;2&#x60;&#x60;&#x60;  | 

### Return type

[**[]MetricMetadata**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsGET**
> TargetDiscovery TargetsGET(ctx, optional)
Returns current target discovery.

Both the active and dropped targets are part of the response by default. ```labels``` represents the label set after relabelling has occurred. ```discoveredLabels``` represent the unmodified labels retrieved during service discovery before relabelling has occurred. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiTargetsGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTargetsGETOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| The &#x60;&#x60;&#x60;state&#x60;&#x60;&#x60; query parameter allows the caller to filter by active or dropped targets, (e.g., &#x60;&#x60;&#x60;state&#x3D;active&#x60;&#x60;&#x60;, &#x60;&#x60;&#x60;state&#x3D;dropped&#x60;&#x60;&#x60;, &#x60;&#x60;&#x60;state&#x3D;any&#x60;&#x60;&#x60;).  | 

### Return type

[**TargetDiscovery**](TargetDiscovery.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

