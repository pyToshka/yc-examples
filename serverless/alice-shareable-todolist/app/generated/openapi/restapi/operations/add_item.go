// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddItemHandlerFunc turns a function with the right signature into a add item handler
type AddItemHandlerFunc func(AddItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddItemHandlerFunc) Handle(params AddItemParams) middleware.Responder {
	return fn(params)
}

// AddItemHandler interface for that can handle valid add item params
type AddItemHandler interface {
	Handle(AddItemParams) middleware.Responder
}

// NewAddItem creates a new http.Handler for the add item operation
func NewAddItem(ctx *middleware.Context, handler AddItemHandler) *AddItem {
	return &AddItem{Context: ctx, Handler: handler}
}

/* AddItem swagger:route POST /api/todo-lists/{list_id}/items addItem

AddItem add item API

*/
type AddItem struct {
	Context *middleware.Context
	Handler AddItemHandler
}

func (o *AddItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddItemParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddItemBody add item body
//
// swagger:model AddItemBody
type AddItemBody struct {

	// text
	// Required: true
	Text *string `json:"text"`
}

// Validate validates this add item body
func (o *AddItemBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddItemBody) validateText(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"text", "body", o.Text); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add item body based on context it is used
func (o *AddItemBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddItemBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddItemBody) UnmarshalBinary(b []byte) error {
	var res AddItemBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
