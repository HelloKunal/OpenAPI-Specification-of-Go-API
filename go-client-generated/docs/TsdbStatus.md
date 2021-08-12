# TsdbStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeadStats** | [***HeadStats**](HeadStats.md) |  | [optional] [default to null]
**LabelValueCountByLabelName** | [**[]Stat**](stat.md) | This will provide a list of the label names and their value count. | [optional] [default to null]
**MemoryInBytesByLabelName** | [**[]Stat**](stat.md) | This will provide a list of the label names and memory used in bytes. Memory usage is calculated by adding the length of all values for a given label name. | [optional] [default to null]
**SeriesCountByLabelValuePair** | [**[]Stat**](stat.md) | This will provide a list of label value pairs and their series count. | [optional] [default to null]
**SeriesCountByMetricName** | [**[]Stat**](stat.md) | This will provide a list of metrics names and their series count. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

