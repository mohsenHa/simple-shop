package productvalidator

import (
	"clean-code-structure/param/productparam"
	"clean-code-structure/pkg/errmsg"
	"clean-code-structure/pkg/richerror"
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateGetRequest(req productparam.GetRequest) (map[string]string, error) {
	const op = "messagevalidator.ValidateGetRequest"

	if err := validation.ValidateStruct(&req, validation.Field(&req.ProductId,
		validation.Required,
		validation.By(v.isProductIdValid)),
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

func (v Validator) isProductIdValid(value interface{}) error {
	productId := value.(int)

	if isExist, err := v.productRepo.IsProductExist(context.Background(), productId); err != nil || !isExist {
		if err != nil {
			return err
		}

		if !isExist {
			return fmt.Errorf(errmsg.ErrorMsgInvalidInput)
		}
	}

	return nil
}
