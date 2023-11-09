// Code generated by "terraform-provider-keyhub-generator"; DO NOT EDIT.
// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

//lint:ignore U1000 Ignore unused functions in generated code
package provider

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"golang.org/x/exp/slices"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	rsschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	keyhub "github.com/topicuskeyhub/sdk-go"
	keyhubaccount "github.com/topicuskeyhub/sdk-go/account"
	keyhubcertificate "github.com/topicuskeyhub/sdk-go/certificate"
	keyhubclient "github.com/topicuskeyhub/sdk-go/client"
	keyhubdirectory "github.com/topicuskeyhub/sdk-go/directory"
	keyhubgroup "github.com/topicuskeyhub/sdk-go/group"
	keyhubgroupclassification "github.com/topicuskeyhub/sdk-go/groupclassification"
	keyhubmodels "github.com/topicuskeyhub/sdk-go/models"
	keyhuborganizationalunit "github.com/topicuskeyhub/sdk-go/organizationalunit"
	keyhubserviceaccount "github.com/topicuskeyhub/sdk-go/serviceaccount"
	keyhubsystem "github.com/topicuskeyhub/sdk-go/system"
	keyhubvaultrecord "github.com/topicuskeyhub/sdk-go/vaultrecord"
)

type contextKey int

const (
	keyHubClientKey contextKey = iota
)

func sliceToTF[T any](elemType attr.Type, vals []T, toValue func(T, *diag.Diagnostics) attr.Value) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	ret := make([]attr.Value, 0, len(vals))
	for _, curVal := range vals {
		ret = append(ret, toValue(curVal, &diags))
	}
	return types.ListValue(elemType, ret)
}

func tfToSlice[T any](val basetypes.ListValue, toValue func(attr.Value, *diag.Diagnostics) T) ([]T, diag.Diagnostics) {
	var diags diag.Diagnostics
	vals := val.Elements()
	ret := make([]T, 0, len(vals))
	for _, curVal := range vals {
		ret = append(ret, toValue(curVal, &diags))
	}
	return ret, diags
}

func mapToTF[T any](elemType attr.Type, vals map[string]T, toValue func(T, *diag.Diagnostics) attr.Value) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	ret := make(map[string]attr.Value)
	for name, val := range vals {
		ret[name] = toValue(val, &diags)
	}
	return types.MapValue(elemType, ret)
}

func tfToMap[T serialization.AdditionalDataHolder](val basetypes.MapValue, toValue func(attr.Value, *diag.Diagnostics) any, ret T) (T, diag.Diagnostics) {
	var diags diag.Diagnostics
	vals := val.Elements()
	retMap := make(map[string]any)
	for name, val := range vals {
		retMap[name] = toValue(val, &diags)
	}
	ret.SetAdditionalData(retMap)
	return ret, diags
}

func int32PToInt64P(in *int32) *int64 {
	if in == nil {
		return nil
	}
	ret := int64(*in)
	return &ret
}

func int64PToInt32P(in *int64) *int32 {
	if in == nil {
		return nil
	}
	ret := int32(*in)
	return &ret
}

func stringerToTF[T fmt.Stringer](val *T) attr.Value {
	if val == nil {
		return types.StringNull()
	}
	return types.StringValue((*val).String())
}

func timeToTF(val time.Time) attr.Value {
	ret, _ := val.MarshalText()
	return types.StringValue(string(ret))
}

func timePointerToTF(val *time.Time) attr.Value {
	if val == nil {
		return types.StringNull()
	}
	ret, _ := val.MarshalText()
	return types.StringValue(string(ret))
}

func tfToTime(val basetypes.StringValue) (time.Time, diag.Diagnostics) {
	var diags diag.Diagnostics
	parsed, err := time.Parse(time.RFC3339, val.ValueString())
	if err != nil {
		diags.AddError("Conversion error", fmt.Sprintf("Unable to parse time: %s", err))
	}
	return parsed, diags
}

func tfToTimePointer(val basetypes.StringValue) (*time.Time, diag.Diagnostics) {
	if val.IsNull() || val.IsUnknown() {
		return nil, diag.Diagnostics{}
	}
	parsed, diags := tfToTime(val)
	return &parsed, diags
}

func withUuidToTF(val interface{ GetUuid() *string }) attr.Value {
	if val == nil {
		return types.StringNull()
	}
	return types.StringPointerValue(val.GetUuid())
}

func toItemsList(ctx context.Context, val attr.Value) basetypes.ObjectValue {
	attrType := map[string]attr.Type{"items": val.Type(ctx)}
	if val.IsNull() || val.IsUnknown() {
		return types.ObjectNull(attrType)
	}
	return types.ObjectValueMust(attrType, map[string]attr.Value{"items": val})
}

func getItemsAttr(val basetypes.ObjectValue, attrType attr.Type) attr.Value {
	if val.IsNull() || val.IsUnknown() {
		return types.ListNull(attrType.(basetypes.ListType).ElementType())
	}
	return val.Attributes()["items"]
}

func parsePointer[T any](val basetypes.StringValue, parser func(string) (T, error)) (*T, diag.Diagnostics) {
	if val.IsNull() || val.IsUnknown() {
		return nil, diag.Diagnostics{}
	}
	parsed, diags := parse(val, parser)
	return &parsed, diags
}

func parsePointer2[T any](val basetypes.StringValue, parser func(string) (*T, error)) (*T, diag.Diagnostics) {
	if val.IsNull() || val.IsUnknown() {
		return nil, diag.Diagnostics{}
	}
	parsed, diags := parse(val, parser)
	return parsed, diags
}

func parse[T any](val basetypes.StringValue, parser func(string) (T, error)) (T, diag.Diagnostics) {
	var diags diag.Diagnostics
	parsed, err := parser(val.ValueString())
	if err != nil {
		diags.AddError("Conversion error", fmt.Sprintf("Unable to parse %s: %s", val.ValueString(), err))
	}
	return parsed, diags
}

func parseCastPointer[T any, Z any](val basetypes.StringValue, parser func(string) (Z, error), caster func(Z) T) (*T, diag.Diagnostics) {
	if val.IsNull() || val.IsUnknown() {
		return nil, diag.Diagnostics{}
	}
	parsed, diags := parseCast(val, parser, caster)
	return &parsed, diags
}

func parseCast[T any, Z any](val basetypes.StringValue, parser func(string) (Z, error), caster func(Z) T) (T, diag.Diagnostics) {
	parsed, diags := parse(val, parser)
	var ret T
	if diags.HasError() {
		return ret, diags
	}
	return caster(parsed), diags
}

func findFirst[T keyhubmodels.Linkableable](ctx context.Context, wrapper interface{ GetItems() []T }, name string, uuid *string, notFoundIsNil bool, err error) (T, diag.Diagnostics) {
	var diags diag.Diagnostics
	var noVal T
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read %s, got error: %s", name, errorReportToString(ctx, err)))
		return noVal, diags
	}
	if len(wrapper.GetItems()) == 0 {
		if !notFoundIsNil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to find %s with UUID %s", name, *uuid))
		}
		return noVal, diags
	}
	return wrapper.GetItems()[0], nil
}

func findGroupGroupPrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.GroupGroupPrimerable, diag.Diagnostics) {
	return findGroupGroupPrimerByUUIDOptionallyNil(ctx, uuid, false)
}

func findGroupGroupPrimerByUUIDOrNil(ctx context.Context, uuid *string) (keyhubmodels.GroupGroupPrimerable, diag.Diagnostics) {
	return findGroupGroupPrimerByUUIDOptionallyNil(ctx, uuid, true)
}

func findGroupGroupPrimerByUUIDOptionallyNil(ctx context.Context, uuid *string, notFoundIsNil bool) (keyhubmodels.GroupGroupPrimerable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Group().Get(ctx, &keyhubgroup.GroupRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubgroup.GroupRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.GroupGroupable](ctx, wrapper, "group", uuid, notFoundIsNil, err)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.GroupGroupPrimer{})); ok {
		ret := primer.(keyhubmodels.GroupGroupPrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type GroupGroupPrimer")
	return nil, diag
}

func findDirectoryAccountDirectoryPrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.DirectoryAccountDirectoryPrimerable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Directory().Get(ctx, &keyhubdirectory.DirectoryRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubdirectory.DirectoryRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.DirectoryAccountDirectoryable](ctx, wrapper, "directory", uuid, false, err)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.DirectoryAccountDirectoryPrimer{})); ok {
		ret := primer.(keyhubmodels.DirectoryAccountDirectoryPrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type DirectoryAccountDirectoryPrimer")
	return nil, diag
}

func findOrganizationOrganizationalUnitPrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.OrganizationOrganizationalUnitPrimerable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Organizationalunit().Get(ctx, &keyhuborganizationalunit.OrganizationalunitRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhuborganizationalunit.OrganizationalunitRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.OrganizationOrganizationalUnitable](ctx, wrapper, "organizational unit", uuid, false, err)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.OrganizationOrganizationalUnitPrimer{})); ok {
		ret := primer.(keyhubmodels.OrganizationOrganizationalUnitPrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type OrganizationOrganizationalUnitPrimer")
	return nil, diag
}

func findCertificateCertificatePrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.CertificateCertificatePrimerable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Certificate().Get(ctx, &keyhubcertificate.CertificateRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubcertificate.CertificateRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.CertificateCertificateable](ctx, wrapper, "certificate", uuid, false, err)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.CertificateCertificatePrimer{})); ok {
		ret := primer.(keyhubmodels.CertificateCertificatePrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type CertificateCertificatePrimer")
	return nil, diag
}

func findProvisioningProvisionedSystemPrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.ProvisioningProvisionedSystemPrimerable, diag.Diagnostics) {
	return findProvisioningProvisionedSystemPrimerByUUIDOptionallyNil(ctx, uuid, false)
}

func findProvisioningProvisionedSystemPrimerByUUIDOrNil(ctx context.Context, uuid *string) (keyhubmodels.ProvisioningProvisionedSystemPrimerable, diag.Diagnostics) {
	return findProvisioningProvisionedSystemPrimerByUUIDOptionallyNil(ctx, uuid, true)
}

func findProvisioningProvisionedSystemPrimerByUUIDOptionallyNil(ctx context.Context, uuid *string, notFoundIsNil bool) (keyhubmodels.ProvisioningProvisionedSystemPrimerable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.System().Get(ctx, &keyhubsystem.SystemRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubsystem.SystemRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.ProvisioningProvisionedSystemable](ctx, wrapper, "provisioned system", uuid, notFoundIsNil, err)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.ProvisioningProvisionedSystemPrimer{})); ok {
		ret := primer.(keyhubmodels.ProvisioningProvisionedSystemPrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type ProvisioningProvisionedSystemPrimer")
	return nil, diag
}

func findGroupGroupClassificationPrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.GroupGroupClassificationPrimerable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Groupclassification().Get(ctx, &keyhubgroupclassification.GroupclassificationRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubgroupclassification.GroupclassificationRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.GroupGroupClassificationable](ctx, wrapper, "group classification", uuid, false, err)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.GroupGroupClassificationPrimer{})); ok {
		ret := primer.(keyhubmodels.GroupGroupClassificationPrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type GroupGroupClassificationPrimer")
	return nil, diag
}

func findClientClientApplicationPrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.ClientClientApplicationPrimerable, diag.Diagnostics) {
	return findClientClientApplicationPrimerByUUIDOptionallyNil(ctx, uuid, false)
}

func findClientClientApplicationPrimerByUUIDOrNil(ctx context.Context, uuid *string) (keyhubmodels.ClientClientApplicationPrimerable, diag.Diagnostics) {
	return findClientClientApplicationPrimerByUUIDOptionallyNil(ctx, uuid, true)
}

func findClientClientApplicationPrimerByUUIDOptionallyNil(ctx context.Context, uuid *string, notFoundIsNil bool) (keyhubmodels.ClientClientApplicationPrimerable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Client().Get(ctx, &keyhubclient.ClientRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubclient.ClientRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.ClientClientApplicationable](ctx, wrapper, "client application", uuid, notFoundIsNil, err)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.ClientClientApplicationPrimer{})); ok {
		ret := primer.(keyhubmodels.ClientClientApplicationPrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type ClientClientApplicationPrimer")
	return nil, diag
}

func findClientOAuth2ClientByUUID(ctx context.Context, uuid *string) (keyhubmodels.ClientOAuth2Clientable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Client().Get(ctx, &keyhubclient.ClientRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubclient.ClientRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.ClientClientApplicationable](ctx, wrapper, "client application", uuid, false, err)
	if ret == nil {
		return nil, diag
	}
	if retSub, ok := ret.(*keyhubmodels.ClientOAuth2Client); ok {
		return retSub, diag
	}
	diag.AddError("Type error", "Result not of type ClientOAuth2Client")
	return nil, diag
}

func findClientLdapClientByUUID(ctx context.Context, uuid *string) (keyhubmodels.ClientLdapClientable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Client().Get(ctx, &keyhubclient.ClientRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubclient.ClientRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.ClientClientApplicationable](ctx, wrapper, "client application", uuid, false, err)
	if ret == nil {
		return nil, diag
	}
	if retSub, ok := ret.(*keyhubmodels.ClientLdapClient); ok {
		return retSub, diag
	}
	diag.AddError("Type error", "Result not of type ClientLdapClient")
	return nil, diag
}

func findAuthAccountPrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.AuthAccountPrimerable, diag.Diagnostics) {
	ret, diag := findAuthAccountByUUID(ctx, uuid)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.AuthAccountPrimer{})); ok {
		ret := primer.(keyhubmodels.AuthAccountPrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type AuthAccountPrimer")
	return nil, diag
}

func findAuthAccountByUUID(ctx context.Context, uuid *string) (keyhubmodels.AuthAccountable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Account().Get(ctx, &keyhubaccount.AccountRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubaccount.AccountRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	return findFirst[keyhubmodels.AuthAccountable](ctx, wrapper, "account", uuid, false, err)
}

func findServiceaccountServiceAccountPrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.ServiceaccountServiceAccountPrimerable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Serviceaccount().Get(ctx, &keyhubserviceaccount.ServiceaccountRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubserviceaccount.ServiceaccountRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	ret, diag := findFirst[keyhubmodels.ServiceaccountServiceAccountable](ctx, wrapper, "service account", uuid, false, err)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.ServiceaccountServiceAccountPrimer{})); ok {
		ret := primer.(keyhubmodels.ServiceaccountServiceAccountPrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type ServiceaccountServiceAccountPrimer")
	return nil, diag
}

func findVaultVaultRecordPrimerByUUID(ctx context.Context, uuid *string) (keyhubmodels.VaultVaultRecordPrimerable, diag.Diagnostics) {
	ret, diag := findVaultVaultRecordByUUID(ctx, uuid)
	if ret == nil {
		return ret, diag
	}
	if primer, ok := findSuperStruct(ret, reflect.TypeOf(keyhubmodels.VaultVaultRecordPrimer{})); ok {
		ret := primer.(keyhubmodels.VaultVaultRecordPrimer)
		return &ret, diag
	}
	diag.AddError("Type error", "Result not of type VaultVaultRecordPrimer")
	return nil, diag
}

func findVaultVaultRecordByUUID(ctx context.Context, uuid *string) (keyhubmodels.VaultVaultRecordable, diag.Diagnostics) {
	if uuid == nil || *uuid == "" {
		return nil, diag.Diagnostics{}
	}
	client := ctx.Value(keyHubClientKey).(*keyhub.KeyHubClient)
	wrapper, err := client.Vaultrecord().Get(ctx, &keyhubvaultrecord.VaultrecordRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubvaultrecord.VaultrecordRequestBuilderGetQueryParameters{
			Uuid: []string{*uuid},
		},
	})
	return findFirst[keyhubmodels.VaultVaultRecordable](ctx, wrapper, "vault record", uuid, false, err)
}

func errorReportToString(ctx context.Context, err error) string {
	report, ok := err.(keyhubmodels.ErrorReportable)
	if !ok {
		return err.Error()
	}
	var msg string
	if report.GetApplicationError() == nil {
		msg = fmt.Sprintf("Error %d from backend: %s", *report.GetCode(), stringPointerToString(report.GetMessage()))
	} else {
		msg = fmt.Sprintf("Error %d (%s) from backend: %s", *report.GetCode(), *report.GetApplicationError(), stringPointerToString(report.GetMessage()))
	}
	tflog.Info(ctx, msg)
	if report.GetStacktrace() != nil {
		tflog.Info(ctx, strings.Join(report.GetStacktrace(), "\n"))
	}
	return msg
}

func stringPointerToString(input *string) string {
	if input != nil {
		return *input
	}
	return ""
}

func getSelfLink(linksAttr basetypes.ListValue) restLinkDataRS {
	var links restLinkDataRS
	fillDataStructFromTFObjectRSRestLink(&links, linksAttr.Elements()[0].(basetypes.ObjectValue))
	return links
}

func resetListNestedAttributeFlags(schema rsschema.ListNestedAttribute) rsschema.ListNestedAttribute {
	schema.Optional = false
	schema.Computed = false
	schema.Required = false
	schema.PlanModifiers = nil
	return schema
}

func resetListAttributeFlags(schema rsschema.ListAttribute) rsschema.ListAttribute {
	schema.Optional = false
	schema.Computed = false
	schema.Required = false
	schema.PlanModifiers = nil
	return schema
}

func isHttpStatusCodeOk(ctx context.Context, status int32, err error, diags *diag.Diagnostics) bool {
	if err != nil {
		report, ok := err.(keyhubmodels.ErrorReportable)
		if !ok || *report.GetCode() != status {
			diags.AddError("Client Error", fmt.Sprintf("Unexpected status code: %s", errorReportToString(ctx, err)))
			return false
		}
	}
	return true
}

func setAttributeValue(ctx context.Context, tf basetypes.ObjectValue, key string, value attr.Value) basetypes.ObjectValue {
	obj := tf.Attributes()
	obj[key] = value
	return types.ObjectValueMust(tf.AttributeTypes(ctx), obj)
}

func collectAdditional(ctx context.Context, data any, additional types.List) []string {
	listValue, _ := additional.ToListValue(ctx)
	ret, _ := tfToSlice(listValue, func(val attr.Value, diags *diag.Diagnostics) string {
		return val.(basetypes.StringValue).ValueString()
	})
	reflectValue := reflect.ValueOf(data)
	reflectType := reflectValue.Type()
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		tkhoa := field.Tag.Get("tkhao")
		if tkhoa != "" {
			attr := reflectValue.Field(i).Interface().(attr.Value)
			if !attr.IsNull() && !attr.IsUnknown() && !slices.Contains(ret, tkhoa) {
				ret = append(ret, tkhoa)
			}
		}
	}
	return ret
}

func findSuperStruct(data any, targetType reflect.Type) (any, bool) {
	reflectValue := reflect.ValueOf(data)
	reflectType := reflectValue.Type()
	if reflectType.Kind() == reflect.Pointer {
		return findSuperStruct(reflectValue.Elem().Interface(), targetType)
	}
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		if field.Anonymous {
			fieldValue := reflectValue.Field(i).Interface()
			if field.Type == targetType {
				return fieldValue, true
			}
			if ret, ok := findSuperStruct(fieldValue, targetType); ok {
				return ret, true
			}
		}
	}
	return nil, false
}
