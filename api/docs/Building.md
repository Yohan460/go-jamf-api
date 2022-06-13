# Building

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**StreetAddress1** | Pointer to **NullableString** |  | [optional] 
**StreetAddress2** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**StateProvince** | Pointer to **NullableString** |  | [optional] 
**ZipPostalCode** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBuilding

`func NewBuilding(name string, ) *Building`

NewBuilding instantiates a new Building object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildingWithDefaults

`func NewBuildingWithDefaults() *Building`

NewBuildingWithDefaults instantiates a new Building object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Building) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Building) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Building) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Building) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Building) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Building) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Building) SetName(v string)`

SetName sets Name field to given value.


### GetStreetAddress1

`func (o *Building) GetStreetAddress1() string`

GetStreetAddress1 returns the StreetAddress1 field if non-nil, zero value otherwise.

### GetStreetAddress1Ok

`func (o *Building) GetStreetAddress1Ok() (*string, bool)`

GetStreetAddress1Ok returns a tuple with the StreetAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress1

`func (o *Building) SetStreetAddress1(v string)`

SetStreetAddress1 sets StreetAddress1 field to given value.

### HasStreetAddress1

`func (o *Building) HasStreetAddress1() bool`

HasStreetAddress1 returns a boolean if a field has been set.

### SetStreetAddress1Nil

`func (o *Building) SetStreetAddress1Nil(b bool)`

 SetStreetAddress1Nil sets the value for StreetAddress1 to be an explicit nil

### UnsetStreetAddress1
`func (o *Building) UnsetStreetAddress1()`

UnsetStreetAddress1 ensures that no value is present for StreetAddress1, not even an explicit nil
### GetStreetAddress2

`func (o *Building) GetStreetAddress2() string`

GetStreetAddress2 returns the StreetAddress2 field if non-nil, zero value otherwise.

### GetStreetAddress2Ok

`func (o *Building) GetStreetAddress2Ok() (*string, bool)`

GetStreetAddress2Ok returns a tuple with the StreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress2

`func (o *Building) SetStreetAddress2(v string)`

SetStreetAddress2 sets StreetAddress2 field to given value.

### HasStreetAddress2

`func (o *Building) HasStreetAddress2() bool`

HasStreetAddress2 returns a boolean if a field has been set.

### SetStreetAddress2Nil

`func (o *Building) SetStreetAddress2Nil(b bool)`

 SetStreetAddress2Nil sets the value for StreetAddress2 to be an explicit nil

### UnsetStreetAddress2
`func (o *Building) UnsetStreetAddress2()`

UnsetStreetAddress2 ensures that no value is present for StreetAddress2, not even an explicit nil
### GetCity

`func (o *Building) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Building) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Building) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Building) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Building) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Building) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetStateProvince

`func (o *Building) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *Building) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *Building) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *Building) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### SetStateProvinceNil

`func (o *Building) SetStateProvinceNil(b bool)`

 SetStateProvinceNil sets the value for StateProvince to be an explicit nil

### UnsetStateProvince
`func (o *Building) UnsetStateProvince()`

UnsetStateProvince ensures that no value is present for StateProvince, not even an explicit nil
### GetZipPostalCode

`func (o *Building) GetZipPostalCode() string`

GetZipPostalCode returns the ZipPostalCode field if non-nil, zero value otherwise.

### GetZipPostalCodeOk

`func (o *Building) GetZipPostalCodeOk() (*string, bool)`

GetZipPostalCodeOk returns a tuple with the ZipPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipPostalCode

`func (o *Building) SetZipPostalCode(v string)`

SetZipPostalCode sets ZipPostalCode field to given value.

### HasZipPostalCode

`func (o *Building) HasZipPostalCode() bool`

HasZipPostalCode returns a boolean if a field has been set.

### SetZipPostalCodeNil

`func (o *Building) SetZipPostalCodeNil(b bool)`

 SetZipPostalCodeNil sets the value for ZipPostalCode to be an explicit nil

### UnsetZipPostalCode
`func (o *Building) UnsetZipPostalCode()`

UnsetZipPostalCode ensures that no value is present for ZipPostalCode, not even an explicit nil
### GetCountry

`func (o *Building) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Building) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Building) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Building) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Building) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Building) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


