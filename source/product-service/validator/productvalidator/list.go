package productvalidator

import (
	"clean-code-structure/param/productparam"
	"clean-code-structure/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateListRequest(req productparam.ListRequest) (map[string]string, error) {
	const op = "messagevalidator.ValidateListRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.PrePage, validation.Required),
		validation.Field(&req.Page, validation.Required),
	); err != nil {
		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).WithErr(err)
	}

	return nil, nil
}
