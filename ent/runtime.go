// Code generated by entc, DO NOT EDIT.

package ent

import (
	"aio/ent/barkaddress"
	"aio/ent/extensionclient"
	"aio/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	barkaddressFields := schema.BarkAddress{}.Fields()
	_ = barkaddressFields
	// barkaddressDescName is the schema descriptor for name field.
	barkaddressDescName := barkaddressFields[0].Descriptor()
	// barkaddress.NameValidator is a validator for the "name" field. It is called by the builders before save.
	barkaddress.NameValidator = barkaddressDescName.Validators[0].(func(string) error)
	// barkaddressDescTarget is the schema descriptor for target field.
	barkaddressDescTarget := barkaddressFields[1].Descriptor()
	// barkaddress.TargetValidator is a validator for the "target" field. It is called by the builders before save.
	barkaddress.TargetValidator = barkaddressDescTarget.Validators[0].(func(string) error)
	extensionclientFields := schema.ExtensionClient{}.Fields()
	_ = extensionclientFields
	// extensionclientDescName is the schema descriptor for name field.
	extensionclientDescName := extensionclientFields[0].Descriptor()
	// extensionclient.NameValidator is a validator for the "name" field. It is called by the builders before save.
	extensionclient.NameValidator = extensionclientDescName.Validators[0].(func(string) error)
	// extensionclientDescExtensionID is the schema descriptor for extension_id field.
	extensionclientDescExtensionID := extensionclientFields[1].Descriptor()
	// extensionclient.ExtensionIDValidator is a validator for the "extension_id" field. It is called by the builders before save.
	extensionclient.ExtensionIDValidator = extensionclientDescExtensionID.Validators[0].(func(string) error)
}
