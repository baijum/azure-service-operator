// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Rule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Rule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRule_STATUS_ARM, Rule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRule_STATUS_ARM runs a test to see if a specific instance of Rule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRule_STATUS_ARM(subject Rule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Rule_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Rule_STATUS_ARM instances for property testing - lazily instantiated by Rule_STATUS_ARMGenerator()
var rule_STATUS_ARMGenerator gopter.Gen

// Rule_STATUS_ARMGenerator returns a generator of Rule_STATUS_ARM instances for property testing.
// We first initialize rule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Rule_STATUS_ARMGenerator() gopter.Gen {
	if rule_STATUS_ARMGenerator != nil {
		return rule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRule_STATUS_ARM(generators)
	rule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Rule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForRule_STATUS_ARM(generators)
	rule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Rule_STATUS_ARM{}), generators)

	return rule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Ruleproperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_Ruleproperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Ruleproperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRuleproperties_STATUS_ARM, Ruleproperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRuleproperties_STATUS_ARM runs a test to see if a specific instance of Ruleproperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRuleproperties_STATUS_ARM(subject Ruleproperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Ruleproperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Ruleproperties_STATUS_ARM instances for property testing - lazily instantiated by
// Ruleproperties_STATUS_ARMGenerator()
var ruleproperties_STATUS_ARMGenerator gopter.Gen

// Ruleproperties_STATUS_ARMGenerator returns a generator of Ruleproperties_STATUS_ARM instances for property testing.
// We first initialize ruleproperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Ruleproperties_STATUS_ARMGenerator() gopter.Gen {
	if ruleproperties_STATUS_ARMGenerator != nil {
		return ruleproperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleproperties_STATUS_ARM(generators)
	ruleproperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Ruleproperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleproperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForRuleproperties_STATUS_ARM(generators)
	ruleproperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Ruleproperties_STATUS_ARM{}), generators)

	return ruleproperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRuleproperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRuleproperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["FilterType"] = gen.PtrOf(gen.OneConstOf(FilterType_STATUS_CorrelationFilter, FilterType_STATUS_SqlFilter))
}

// AddRelatedPropertyGeneratorsForRuleproperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRuleproperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(Action_STATUS_ARMGenerator())
	gens["CorrelationFilter"] = gen.PtrOf(CorrelationFilter_STATUS_ARMGenerator())
	gens["SqlFilter"] = gen.PtrOf(SqlFilter_STATUS_ARMGenerator())
}

func Test_SystemData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS_ARM, SystemData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS_ARM runs a test to see if a specific instance of SystemData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS_ARM(subject SystemData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SystemData_STATUS_ARM instances for property testing - lazily instantiated by
// SystemData_STATUS_ARMGenerator()
var systemData_STATUS_ARMGenerator gopter.Gen

// SystemData_STATUS_ARMGenerator returns a generator of SystemData_STATUS_ARM instances for property testing.
func SystemData_STATUS_ARMGenerator() gopter.Gen {
	if systemData_STATUS_ARMGenerator != nil {
		return systemData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(generators)
	systemData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS_ARM{}), generators)

	return systemData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}

func Test_Action_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Action_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAction_STATUS_ARM, Action_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAction_STATUS_ARM runs a test to see if a specific instance of Action_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAction_STATUS_ARM(subject Action_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Action_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Action_STATUS_ARM instances for property testing - lazily instantiated by Action_STATUS_ARMGenerator()
var action_STATUS_ARMGenerator gopter.Gen

// Action_STATUS_ARMGenerator returns a generator of Action_STATUS_ARM instances for property testing.
func Action_STATUS_ARMGenerator() gopter.Gen {
	if action_STATUS_ARMGenerator != nil {
		return action_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAction_STATUS_ARM(generators)
	action_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Action_STATUS_ARM{}), generators)

	return action_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAction_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAction_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CompatibilityLevel"] = gen.PtrOf(gen.Int())
	gens["RequiresPreprocessing"] = gen.PtrOf(gen.Bool())
	gens["SqlExpression"] = gen.PtrOf(gen.AlphaString())
}

func Test_CorrelationFilter_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorrelationFilter_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorrelationFilter_STATUS_ARM, CorrelationFilter_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorrelationFilter_STATUS_ARM runs a test to see if a specific instance of CorrelationFilter_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorrelationFilter_STATUS_ARM(subject CorrelationFilter_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorrelationFilter_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of CorrelationFilter_STATUS_ARM instances for property testing - lazily instantiated by
// CorrelationFilter_STATUS_ARMGenerator()
var correlationFilter_STATUS_ARMGenerator gopter.Gen

// CorrelationFilter_STATUS_ARMGenerator returns a generator of CorrelationFilter_STATUS_ARM instances for property testing.
func CorrelationFilter_STATUS_ARMGenerator() gopter.Gen {
	if correlationFilter_STATUS_ARMGenerator != nil {
		return correlationFilter_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorrelationFilter_STATUS_ARM(generators)
	correlationFilter_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(CorrelationFilter_STATUS_ARM{}), generators)

	return correlationFilter_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCorrelationFilter_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorrelationFilter_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ContentType"] = gen.PtrOf(gen.AlphaString())
	gens["CorrelationId"] = gen.PtrOf(gen.AlphaString())
	gens["Label"] = gen.PtrOf(gen.AlphaString())
	gens["MessageId"] = gen.PtrOf(gen.AlphaString())
	gens["Properties"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["ReplyTo"] = gen.PtrOf(gen.AlphaString())
	gens["ReplyToSessionId"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresPreprocessing"] = gen.PtrOf(gen.Bool())
	gens["SessionId"] = gen.PtrOf(gen.AlphaString())
	gens["To"] = gen.PtrOf(gen.AlphaString())
}

func Test_SqlFilter_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlFilter_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlFilter_STATUS_ARM, SqlFilter_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlFilter_STATUS_ARM runs a test to see if a specific instance of SqlFilter_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlFilter_STATUS_ARM(subject SqlFilter_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlFilter_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlFilter_STATUS_ARM instances for property testing - lazily instantiated by
// SqlFilter_STATUS_ARMGenerator()
var sqlFilter_STATUS_ARMGenerator gopter.Gen

// SqlFilter_STATUS_ARMGenerator returns a generator of SqlFilter_STATUS_ARM instances for property testing.
func SqlFilter_STATUS_ARMGenerator() gopter.Gen {
	if sqlFilter_STATUS_ARMGenerator != nil {
		return sqlFilter_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlFilter_STATUS_ARM(generators)
	sqlFilter_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlFilter_STATUS_ARM{}), generators)

	return sqlFilter_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlFilter_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlFilter_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CompatibilityLevel"] = gen.PtrOf(gen.Int())
	gens["RequiresPreprocessing"] = gen.PtrOf(gen.Bool())
	gens["SqlExpression"] = gen.PtrOf(gen.AlphaString())
}
