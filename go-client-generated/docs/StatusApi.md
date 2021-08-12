# {{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServeBuildInfo**](StatusApi.md#ServeBuildInfo) | **Get** /status/buildinfo | Returns build information
[**ServeConfig**](StatusApi.md#ServeConfig) | **Get** /status/config | Returns configuration file
[**ServeFlags**](StatusApi.md#ServeFlags) | **Get** /status/flags | Returns flag values
[**ServeRuntimeInfo**](StatusApi.md#ServeRuntimeInfo) | **Get** /status/runtimeinfo | Returns runtime info
[**ServeTSDBStatus**](StatusApi.md#ServeTSDBStatus) | **Get** /status/tsdb | Returns statistics about TSBD
[**ServeWALReplayStatus**](StatusApi.md#ServeWALReplayStatus) | **Get** /status/walreplay | Returns info about WAL replay.

# **ServeBuildInfo**
> PrometheusVersion ServeBuildInfo(ctx, )
Returns build information

The following endpoint returns various build information properties about the Prometheus server  All values are of the result type ```string```.  --- **NOTE:** The exact returned build properties may change without notice between Prometheus versions.  ---  New in v2.14 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PrometheusVersion**](PrometheusVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServeConfig**
> PrometheusConfig ServeConfig(ctx, )
Returns configuration file

The following endpoint returns currently loaded configuration file  The config is returned as dumped YAML file. Due to limitation of the YAML library, YAML comments are not included. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PrometheusConfig**](prometheusConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServeFlags**
> ServeFlags(ctx, )
Returns flag values

The following endpoint returns flag values that Prometheus was configured with  All values are of the result type ```string```.  New in v2.2 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServeRuntimeInfo**
> RuntimeInfo ServeRuntimeInfo(ctx, )
Returns runtime info

The following endpoint returns various runtime information properties about the Prometheus server  The returned values are of different types, depending on the nature of the runtime property  --- **NOTE:** The exact returned runtime properties may change without notice between Prometheus versions.  ---  New in v2.14 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RuntimeInfo**](RuntimeInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServeTSDBStatus**
> TsdbStatus ServeTSDBStatus(ctx, )
Returns statistics about TSBD

The following endpoint returns various cardinality statistics about the Prometheus TSDB  Response Data ---  **headStats:** This provides the following data about the head block of the TSDB: >**numSeries:** The number of series. **chunkCount:** The number of chunks. **minTime:** The current minimum timestamp in milliseconds. **maxTime:** The current maximum timestamp in milliseconds.  **seriesCountByMetricName:** This will provide a list of metrics names and their series count. **labelValueCountByLabelName:** This will provide a list of the label names and their value count. **memoryInBytesByLabelName:** This will provide a list of the label names and memory used in bytes. Memory usage is calculated by adding the length of all values for a given label name. **seriesCountByLabelPair:** This will provide a list of label value pairs and their series count. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TsdbStatus**](tsdbStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServeWALReplayStatus**
> WalReplayStatus ServeWALReplayStatus(ctx, )
Returns info about WAL replay.

The following endpoint returns information about the WAL replay  Response Data ---  **read:** The number of segments replayed so far.  **total:** The total number segments needed to be replayed.  **progress:** The progress of the replay (0 - 100%).  **state:** The state of the replay.  **Possible states:**    - **waiting:** Waiting for the replay to start.    - **in progress:** The replay is in progress.    - **done:** The replay has finished.    --- **NOTE:** This endpoint is available before the server has been marked ready and is updated in real time to facilitate monitoring the progress of the WAL replay.  ---  New in v2.28 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WalReplayStatus**](walReplayStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

