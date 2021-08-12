# {{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanTombstonesPOST**](TSDBAdminAPIApi.md#CleanTombstonesPOST) | **Post** /admin/tsdb/clean_tombstones | Removes deleted data
[**CleanTombstonesPUT**](TSDBAdminAPIApi.md#CleanTombstonesPUT) | **Put** /admin/tsdb/clean_tombstones | Removes deleted data
[**DeleteSeriesPOST**](TSDBAdminAPIApi.md#DeleteSeriesPOST) | **Post** /admin/tsdb/delete_series | Deletes selected data
[**DeleteSeriesPUT**](TSDBAdminAPIApi.md#DeleteSeriesPUT) | **Put** /admin/tsdb/delete_series | Deletes selected data
[**SnapshotPOST**](TSDBAdminAPIApi.md#SnapshotPOST) | **Post** /admin/tsdb/snapshot | Creates Snapshot of current data
[**SnapshotPUT**](TSDBAdminAPIApi.md#SnapshotPUT) | **Put** /admin/tsdb/snapshot | Creates Snapshot of current data

# **CleanTombstonesPOST**
> CleanTombstonesPOST(ctx, )
Removes deleted data

CleanTombstones removes the deleted data from disk and cleans up the existing tombstones. This can be used after deleting series to free up space.  <i>New in v2.1 and supports PUT from v2.9</i> 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CleanTombstonesPUT**
> CleanTombstonesPUT(ctx, )
Removes deleted data

CleanTombstones removes the deleted data from disk and cleans up the existing tombstones. This can be used after deleting series to free up space.  <i>New in v2.1 and supports PUT from v2.9</i> 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSeriesPOST**
> DeleteSeriesPOST(ctx, match, optional)
Deletes selected data

DeleteSeries deletes data for a selection of series in a time range. The actual data still exists on disk and is cleaned up in future compactions or can be explicitly cleaned up by hitting the [Clean Tombstones](https://prometheus.io/docs/prometheus/latest/querying/api/#clean-tombstones) endpoint.  --- **NOTE:** This endpoint marks samples from series as deleted, but will not necessarily prevent associated series metadata from still being returned in metadata queries for the affected time range (even after cleaning tombstones). The exact extent of metadata deletion is an implementation detail that may change in the future.  ---  New in v2.1 and supports PUT from v2.9 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **match** | **string**| Repeated label matcher argument that selects the series to delete. At least one match[] argument must be provided.  Example: &#x60;&#x60;&#x60;?match[]&#x3D;up&amp;match[]&#x3D;process_start_time_seconds{job&#x3D;\&quot;prometheus\&quot;}&#x27;&#x60;&#x60;&#x60;  | 
 **optional** | ***TSDBAdminAPIApiDeleteSeriesPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TSDBAdminAPIApiDeleteSeriesPOSTOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp. Optional and defaults to minimum possible time. | 
 **end** | **optional.String**| End timestamp. Optional and defaults to maximum possible time.  Not mentioning both start and end times would clear all the data for the matched series in the database.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSeriesPUT**
> DeleteSeriesPUT(ctx, match, optional)
Deletes selected data

DeleteSeries deletes data for a selection of series in a time range. The actual data still exists on disk and is cleaned up in future compactions or can be explicitly cleaned up by hitting the [Clean Tombstones](https://prometheus.io/docs/prometheus/latest/querying/api/#clean-tombstones) endpoint.  New in v2.1 and supports PUT from v2.9 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **match** | **string**| Repeated label matcher argument that selects the series to delete. At least one match[] argument must be provided.  Example: &#x60;&#x60;&#x60;?match[]&#x3D;up&amp;match[]&#x3D;process_start_time_seconds{job&#x3D;\&quot;prometheus\&quot;}&#x27;&#x60;&#x60;&#x60;  | 
 **optional** | ***TSDBAdminAPIApiDeleteSeriesPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TSDBAdminAPIApiDeleteSeriesPUTOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| Start timestamp. Optional and defaults to minimum possible time. | 
 **end** | **optional.String**| End timestamp. Optional and defaults to maximum possible time.  Not mentioning both start and end times would clear all the data for the matched series in the database.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotPOST**
> ResponseSnapshot SnapshotPOST(ctx, optional)
Creates Snapshot of current data

Snapshot creates a snapshot of all current data into ```snapshots/<datetime>-<rand>``` under the TSDB's data directory and returns the directory as response. It will optionally skip snapshotting data that is only present in the head block, and which has not yet been compacted to disk.  New in v2.1 and supports PUT from v2.9 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TSDBAdminAPIApiSnapshotPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TSDBAdminAPIApiSnapshotPOSTOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipHead** | **optional.Bool**| Skip data present in the head block. Optional.  | 

### Return type

[**ResponseSnapshot**](responseSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotPUT**
> ResponseSnapshot SnapshotPUT(ctx, optional)
Creates Snapshot of current data

Snapshot creates a snapshot of all current data into ```snapshots/<datetime>-<rand>``` under the TSDB's data directory and returns the directory as response. It will optionally skip snapshotting data that is only present in the head block, and which has not yet been compacted to disk.  New in v2.1 and supports PUT from v2.9 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TSDBAdminAPIApiSnapshotPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TSDBAdminAPIApiSnapshotPUTOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipHead** | **optional.Bool**| Skip data present in the head block. Optional.  | 

### Return type

[**ResponseSnapshot**](responseSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

