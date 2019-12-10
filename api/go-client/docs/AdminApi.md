# \AdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminLogin**](AdminApi.md#AdminLogin) | **Post** /admin/login | Admin login
[**DeleteAccountByAdmin**](AdminApi.md#DeleteAccountByAdmin) | **Delete** /admin/api/account/ | Delete account
[**GetAccountsByAdmin**](AdminApi.md#GetAccountsByAdmin) | **Get** /admin/api/account/ | Get accounts
[**GetCompanysByAdmin**](AdminApi.md#GetCompanysByAdmin) | **Get** /admin/api/company/ | Get companys
[**SetCompany**](AdminApi.md#SetCompany) | **Post** /admin/api/company/ | Set Company


# **AdminLogin**
> AdminLogin(ctx, body)
Admin login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Admin**](Admin.md)| Created user object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountByAdmin**
> DeleteAccountByAdmin(ctx, accountID)
Delete account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountID** | **int32**| Delete account | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountsByAdmin**
> Account GetAccountsByAdmin(ctx, )
Get accounts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompanysByAdmin**
> Company GetCompanysByAdmin(ctx, )
Get companys

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Company**](Company.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetCompany**
> SetCompany(ctx, body)
Set Company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Company**](Company.md)| Set Company | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

