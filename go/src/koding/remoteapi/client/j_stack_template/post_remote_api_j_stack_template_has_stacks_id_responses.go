package j_stack_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJStackTemplateHasStacksIDReader is a Reader for the PostRemoteAPIJStackTemplateHasStacksID structure.
type PostRemoteAPIJStackTemplateHasStacksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJStackTemplateHasStacksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJStackTemplateHasStacksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJStackTemplateHasStacksIDOK creates a PostRemoteAPIJStackTemplateHasStacksIDOK with default headers values
func NewPostRemoteAPIJStackTemplateHasStacksIDOK() *PostRemoteAPIJStackTemplateHasStacksIDOK {
	return &PostRemoteAPIJStackTemplateHasStacksIDOK{}
}

/*PostRemoteAPIJStackTemplateHasStacksIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJStackTemplateHasStacksIDOK struct {
	Payload PostRemoteAPIJStackTemplateHasStacksIDOKBody
}

func (o *PostRemoteAPIJStackTemplateHasStacksIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.hasStacks/{id}][%d] postRemoteApiJStackTemplateHasStacksIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJStackTemplateHasStacksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJStackTemplateHasStacksIDOKBody post remote API j stack template has stacks ID o k body
swagger:model PostRemoteAPIJStackTemplateHasStacksIDOKBody
*/
type PostRemoteAPIJStackTemplateHasStacksIDOKBody struct {
	models.JStackTemplate

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJStackTemplateHasStacksIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJStackTemplateHasStacksIDOKBodyAO0 models.JStackTemplate
	if err := swag.ReadJSON(raw, &postRemoteAPIJStackTemplateHasStacksIDOKBodyAO0); err != nil {
		return err
	}
	o.JStackTemplate = postRemoteAPIJStackTemplateHasStacksIDOKBodyAO0

	var postRemoteAPIJStackTemplateHasStacksIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJStackTemplateHasStacksIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJStackTemplateHasStacksIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJStackTemplateHasStacksIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJStackTemplateHasStacksIDOKBodyAO0, err := swag.WriteJSON(o.JStackTemplate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJStackTemplateHasStacksIDOKBodyAO0)

	postRemoteAPIJStackTemplateHasStacksIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJStackTemplateHasStacksIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j stack template has stacks ID o k body
func (o *PostRemoteAPIJStackTemplateHasStacksIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JStackTemplate.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
