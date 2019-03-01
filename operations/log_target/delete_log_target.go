// Code generated by go-swagger; DO NOT EDIT.

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteLogTargetHandlerFunc turns a function with the right signature into a delete log target handler
type DeleteLogTargetHandlerFunc func(DeleteLogTargetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteLogTargetHandlerFunc) Handle(params DeleteLogTargetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteLogTargetHandler interface for that can handle valid delete log target params
type DeleteLogTargetHandler interface {
	Handle(DeleteLogTargetParams, interface{}) middleware.Responder
}

// NewDeleteLogTarget creates a new http.Handler for the delete log target operation
func NewDeleteLogTarget(ctx *middleware.Context, handler DeleteLogTargetHandler) *DeleteLogTarget {
	return &DeleteLogTarget{Context: ctx, Handler: handler}
}

/*DeleteLogTarget swagger:route DELETE /services/haproxy/configuration/log_targets/{id} LogTarget deleteLogTarget

Delete a Log Target

Deletes a Log Target configuration by it's ID from the specified parent.

*/
type DeleteLogTarget struct {
	Context *middleware.Context
	Handler DeleteLogTargetHandler
}

func (o *DeleteLogTarget) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteLogTargetParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}