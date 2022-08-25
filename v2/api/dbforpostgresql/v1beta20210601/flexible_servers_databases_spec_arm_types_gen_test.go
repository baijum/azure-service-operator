// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

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

func Test_FlexibleServers_Databases_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Databases_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Databases_SpecARM, FlexibleServers_Databases_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Databases_SpecARM runs a test to see if a specific instance of FlexibleServers_Databases_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Databases_SpecARM(subject FlexibleServers_Databases_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Databases_SpecARM
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

// Generator of FlexibleServers_Databases_SpecARM instances for property testing - lazily instantiated by
// FlexibleServers_Databases_SpecARMGenerator()
var flexibleServers_Databases_SpecARMGenerator gopter.Gen

// FlexibleServers_Databases_SpecARMGenerator returns a generator of FlexibleServers_Databases_SpecARM instances for property testing.
// We first initialize flexibleServers_Databases_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServers_Databases_SpecARMGenerator() gopter.Gen {
	if flexibleServers_Databases_SpecARMGenerator != nil {
		return flexibleServers_Databases_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Databases_SpecARM(generators)
	flexibleServers_Databases_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Databases_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Databases_SpecARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServers_Databases_SpecARM(generators)
	flexibleServers_Databases_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Databases_SpecARM{}), generators)

	return flexibleServers_Databases_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Databases_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Databases_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServers_Databases_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServers_Databases_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabasePropertiesARMGenerator())
}

func Test_DatabasePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabasePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabasePropertiesARM, DatabasePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabasePropertiesARM runs a test to see if a specific instance of DatabasePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabasePropertiesARM(subject DatabasePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabasePropertiesARM
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

// Generator of DatabasePropertiesARM instances for property testing - lazily instantiated by
// DatabasePropertiesARMGenerator()
var databasePropertiesARMGenerator gopter.Gen

// DatabasePropertiesARMGenerator returns a generator of DatabasePropertiesARM instances for property testing.
func DatabasePropertiesARMGenerator() gopter.Gen {
	if databasePropertiesARMGenerator != nil {
		return databasePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabasePropertiesARM(generators)
	databasePropertiesARMGenerator = gen.Struct(reflect.TypeOf(DatabasePropertiesARM{}), generators)

	return databasePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabasePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabasePropertiesARM(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}