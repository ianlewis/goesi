# \LocationApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdOnline**](LocationApi.md#GetCharactersCharacterIdOnline) | **Get** /characters/{character_id}/online/ | Get character online


# **GetCharactersCharacterIdOnline**
> GetCharactersCharacterIdOnlineOk GetCharactersCharacterIdOnline(ctx, characterId, optional)
Get character online

Checks if the character is currently online  ---  Alternate route: `/dev/characters/{character_id}/online/`   ---  This route is cached for up to 60 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetCharactersCharacterIdOnlineOk**](get_characters_character_id_online_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
