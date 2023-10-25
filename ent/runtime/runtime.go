// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/yawnak/fuel-record-crud/ent/car"
	"github.com/yawnak/fuel-record-crud/ent/fuelrecord"
	"github.com/yawnak/fuel-record-crud/ent/odometerrecord"
	"github.com/yawnak/fuel-record-crud/ent/schema"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescMake is the schema descriptor for make field.
	carDescMake := carFields[1].Descriptor()
	// car.MakeValidator is a validator for the "make" field. It is called by the builders before save.
	car.MakeValidator = carDescMake.Validators[0].(func(string) error)
	// carDescModel is the schema descriptor for model field.
	carDescModel := carFields[2].Descriptor()
	// car.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	car.ModelValidator = carDescModel.Validators[0].(func(string) error)
	fuelrecord.Policy = privacy.NewPolicies(schema.FuelRecord{})
	fuelrecord.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := fuelrecord.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	fuelrecordFields := schema.FuelRecord{}.Fields()
	_ = fuelrecordFields
	// fuelrecordDescCurrentFuelLiters is the schema descriptor for current_fuel_liters field.
	fuelrecordDescCurrentFuelLiters := fuelrecordFields[1].Descriptor()
	// fuelrecord.CurrentFuelLitersValidator is a validator for the "current_fuel_liters" field. It is called by the builders before save.
	fuelrecord.CurrentFuelLitersValidator = fuelrecordDescCurrentFuelLiters.Validators[0].(func(float64) error)
	odometerrecordHooks := schema.OdometerRecord{}.Hooks()
	odometerrecord.Hooks[0] = odometerrecordHooks[0]
	odometerrecordFields := schema.OdometerRecord{}.Fields()
	_ = odometerrecordFields
	// odometerrecordDescCurrentFuelLiters is the schema descriptor for current_fuel_liters field.
	odometerrecordDescCurrentFuelLiters := odometerrecordFields[1].Descriptor()
	// odometerrecord.CurrentFuelLitersValidator is a validator for the "current_fuel_liters" field. It is called by the builders before save.
	odometerrecord.CurrentFuelLitersValidator = odometerrecordDescCurrentFuelLiters.Validators[0].(func(float64) error)
	// odometerrecordDescDifference is the schema descriptor for difference field.
	odometerrecordDescDifference := odometerrecordFields[2].Descriptor()
	// odometerrecord.DifferenceValidator is a validator for the "difference" field. It is called by the builders before save.
	odometerrecord.DifferenceValidator = odometerrecordDescDifference.Validators[0].(func(float64) error)
}

const (
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
