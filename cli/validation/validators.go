package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var ErrorFieldIsNotPointer = fmt.Errorf("validators.Field: First argument must be a pointer to a struct field!")

// Validator represents a validation rule.
type Validator struct {
	Tags string
	Err  string
}

// initialize creates new singleton validator if its value is nil.
func initialize() {
	if validate != nil {
		return
	}

	validate = validator.New()
	validate.RegisterTagNameFunc(fieldName)
}

// validate validates the provided value against the validator.
// It returns encountered validation errors and boolean indicating
// whether to skip further validation of a variable or not.
func (v *Validator) validate(value interface{}) (ValidationErrors, bool) {
	initialize()

	if v.Tags == "omitempty" {
		return nil, isEmpty(value)
	}

	errs := validate.Var(value, v.Tags)
	es := toValidationErrors(errs)

	if len(v.Err) > 0 {
		for i := range es {
			es[i].Err = v.Err
		}
	}

	return es, false
}

// Error overrides the default error of the validator and returns modified validator.
func (v Validator) Error(err string) Validator {
	v.Err = err
	return v
}

// Tags returns a new validator with the given tags. It is a generic validator that
// allows use of any validation rule from 'github.com/go-playground/validator' library.
func Tags(tags string) Validator {
	return Validator{
		Tags: tags,
	}
}

// OmitEmpty validator prevents further validation of a variable, if variable is empty.
func OmitEmpty() Validator {
	return Validator{
		Tags: "omitempty",
	}
}

// Required validator verifies that value is provided.
func Required() Validator {
	return Validator{
		Tags: "required",
		Err:  "Property '{.Namespace}' is required.",
	}
}

// Min validator verifies that the field value is greater then or equal to the given value.
// In case of slices, arrays and maps, the length is verified.
func Min(value int) Validator {
	tag := fmt.Sprintf("min=%d", value)

	return Validator{
		Tags: tag,
		Err:  "Minimum value for property '{.Namespace}' is {.Param} (actual: {.Value}).",
	}
}

// Max validator verifies that the field value is less then or equal to the given value.
// In case of slices, arrays and maps, the length is verified.
func Max(value int) Validator {
	tag := fmt.Sprintf("max=%d", value)

	return Validator{
		Tags: tag,
		Err:  "Maximum value for property '{.Namespace}' is {.Param} (actual: {.Value}).",
	}
}

// Len validator verifies that the field length equals to the given value.
func Len(value int) Validator {
	tag := fmt.Sprintf("len=%d", value)

	return Validator{
		Tags: tag,
		Err:  "Length of '{.Namespace}' must be {.Param} (actual: {.Value}).",
	}
}

// MinLen validator verifies that the field value is greater then or equal to the given value.
// In case of slices, arrays and maps, the length is verified.
func MinLen(value int) Validator {
	return Min(value).Error("Minimum length of '{.StructField}' is {.Param} (actual: {.Value})")
}

// MaxLen validator verifies that the field value is less then or equal to the given value.
// In case of slices, arrays and maps, the length is verified.
func MaxLen(value int) Validator {
	return Max(value).Error("Maximum length of '{.StructField}' is {.Param} (actual: {.Value})")
}
