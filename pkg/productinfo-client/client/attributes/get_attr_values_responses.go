// Code generated by go-swagger; DO NOT EDIT.

package attributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/banzaicloud/productinfo/pkg/productinfo-client/models"
)

// GetAttrValuesReader is a Reader for the GetAttrValues structure.
type GetAttrValuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAttrValuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAttrValuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAttrValuesOK creates a GetAttrValuesOK with default headers values
func NewGetAttrValuesOK() *GetAttrValuesOK {
	return &GetAttrValuesOK{}
}

/*GetAttrValuesOK handles this case with default header values.

AttributeResponse
*/
type GetAttrValuesOK struct {
	Payload *models.AttributeResponse
}

func (o *GetAttrValuesOK) Error() string {
	return fmt.Sprintf("[GET /providers/{provider}/services/{service}/regions/{region}/products/{attribute}][%d] getAttrValuesOK  %+v", 200, o.Payload)
}

func (o *GetAttrValuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttributeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
