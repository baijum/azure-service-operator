// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

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

func Test_Redis_PatchSchedule_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_PatchSchedule_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_PatchSchedule_Spec_ARM, Redis_PatchSchedule_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_PatchSchedule_Spec_ARM runs a test to see if a specific instance of Redis_PatchSchedule_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_PatchSchedule_Spec_ARM(subject Redis_PatchSchedule_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_PatchSchedule_Spec_ARM
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

// Generator of Redis_PatchSchedule_Spec_ARM instances for property testing - lazily instantiated by
// Redis_PatchSchedule_Spec_ARMGenerator()
var redis_PatchSchedule_Spec_ARMGenerator gopter.Gen

// Redis_PatchSchedule_Spec_ARMGenerator returns a generator of Redis_PatchSchedule_Spec_ARM instances for property testing.
// We first initialize redis_PatchSchedule_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_PatchSchedule_Spec_ARMGenerator() gopter.Gen {
	if redis_PatchSchedule_Spec_ARMGenerator != nil {
		return redis_PatchSchedule_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_PatchSchedule_Spec_ARM(generators)
	redis_PatchSchedule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Redis_PatchSchedule_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_PatchSchedule_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForRedis_PatchSchedule_Spec_ARM(generators)
	redis_PatchSchedule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Redis_PatchSchedule_Spec_ARM{}), generators)

	return redis_PatchSchedule_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedis_PatchSchedule_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_PatchSchedule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedis_PatchSchedule_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_PatchSchedule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ScheduleEntries_ARMGenerator())
}

func Test_ScheduleEntries_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntries_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntries_ARM, ScheduleEntries_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntries_ARM runs a test to see if a specific instance of ScheduleEntries_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntries_ARM(subject ScheduleEntries_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntries_ARM
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

// Generator of ScheduleEntries_ARM instances for property testing - lazily instantiated by
// ScheduleEntries_ARMGenerator()
var scheduleEntries_ARMGenerator gopter.Gen

// ScheduleEntries_ARMGenerator returns a generator of ScheduleEntries_ARM instances for property testing.
func ScheduleEntries_ARMGenerator() gopter.Gen {
	if scheduleEntries_ARMGenerator != nil {
		return scheduleEntries_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduleEntries_ARM(generators)
	scheduleEntries_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduleEntries_ARM{}), generators)

	return scheduleEntries_ARMGenerator
}

// AddRelatedPropertyGeneratorsForScheduleEntries_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduleEntries_ARM(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntry_ARMGenerator())
}

func Test_ScheduleEntry_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntry_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntry_ARM, ScheduleEntry_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntry_ARM runs a test to see if a specific instance of ScheduleEntry_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntry_ARM(subject ScheduleEntry_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntry_ARM
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

// Generator of ScheduleEntry_ARM instances for property testing - lazily instantiated by ScheduleEntry_ARMGenerator()
var scheduleEntry_ARMGenerator gopter.Gen

// ScheduleEntry_ARMGenerator returns a generator of ScheduleEntry_ARM instances for property testing.
func ScheduleEntry_ARMGenerator() gopter.Gen {
	if scheduleEntry_ARMGenerator != nil {
		return scheduleEntry_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduleEntry_ARM(generators)
	scheduleEntry_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduleEntry_ARM{}), generators)

	return scheduleEntry_ARMGenerator
}

// AddIndependentPropertyGeneratorsForScheduleEntry_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduleEntry_ARM(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.OneConstOf(
		ScheduleEntry_DayOfWeek_Everyday,
		ScheduleEntry_DayOfWeek_Friday,
		ScheduleEntry_DayOfWeek_Monday,
		ScheduleEntry_DayOfWeek_Saturday,
		ScheduleEntry_DayOfWeek_Sunday,
		ScheduleEntry_DayOfWeek_Thursday,
		ScheduleEntry_DayOfWeek_Tuesday,
		ScheduleEntry_DayOfWeek_Wednesday,
		ScheduleEntry_DayOfWeek_Weekend))
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}