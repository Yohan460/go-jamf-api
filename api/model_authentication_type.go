/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AuthenticationType the model 'AuthenticationType'
type AuthenticationType string

// List of AuthenticationType
const (
	AUTHENTICATIONTYPE_JSS AuthenticationType = "JSS"
	AUTHENTICATIONTYPE_LDAP AuthenticationType = "LDAP"
	AUTHENTICATIONTYPE_SAML AuthenticationType = "SAML"
	AUTHENTICATIONTYPE_INVITE AuthenticationType = "INVITE"
	AUTHENTICATIONTYPE_NATIVE_APP_API_INTEGRATION AuthenticationType = "NATIVE_APP_API_INTEGRATION"
	AUTHENTICATIONTYPE_DEVICE_SIGNATURE AuthenticationType = "DEVICE_SIGNATURE"
	AUTHENTICATIONTYPE_CLOUD_CONNECTOR AuthenticationType = "CLOUD_CONNECTOR"
	AUTHENTICATIONTYPE_SYSTEM_ACCOUNT AuthenticationType = "SYSTEM_ACCOUNT"
	AUTHENTICATIONTYPE_USER_ENROLLMENT AuthenticationType = "USER_ENROLLMENT"
	AUTHENTICATIONTYPE_CLIENT_CREDENTIALS AuthenticationType = "CLIENT_CREDENTIALS"
	AUTHENTICATIONTYPE_M2_M AuthenticationType = "M2M"
)

// All allowed values of AuthenticationType enum
var AllowedAuthenticationTypeEnumValues = []AuthenticationType{
	"JSS",
	"LDAP",
	"SAML",
	"INVITE",
	"NATIVE_APP_API_INTEGRATION",
	"DEVICE_SIGNATURE",
	"CLOUD_CONNECTOR",
	"SYSTEM_ACCOUNT",
	"USER_ENROLLMENT",
	"CLIENT_CREDENTIALS",
	"M2M",
}

func (v *AuthenticationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticationType(value)
	for _, existing := range AllowedAuthenticationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticationType", value)
}

// NewAuthenticationTypeFromValue returns a pointer to a valid AuthenticationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticationTypeFromValue(v string) (*AuthenticationType, error) {
	ev := AuthenticationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticationType: valid values are %v", v, AllowedAuthenticationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticationType) IsValid() bool {
	for _, existing := range AllowedAuthenticationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthenticationType value
func (v AuthenticationType) Ptr() *AuthenticationType {
	return &v
}

type NullableAuthenticationType struct {
	value *AuthenticationType
	isSet bool
}

func (v NullableAuthenticationType) Get() *AuthenticationType {
	return v.value
}

func (v *NullableAuthenticationType) Set(val *AuthenticationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationType(val *AuthenticationType) *NullableAuthenticationType {
	return &NullableAuthenticationType{value: val, isSet: true}
}

func (v NullableAuthenticationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

