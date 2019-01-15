// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/bitmexwrap/bitmex/models"
)

// ChatGetConnectedReader is a Reader for the ChatGetConnected structure.
type ChatGetConnectedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChatGetConnectedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChatGetConnectedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewChatGetConnectedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewChatGetConnectedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewChatGetConnectedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChatGetConnectedOK creates a ChatGetConnectedOK with default headers values
func NewChatGetConnectedOK() *ChatGetConnectedOK {
	return &ChatGetConnectedOK{}
}

/*ChatGetConnectedOK handles this case with default header values.

Request was successful
*/
type ChatGetConnectedOK struct {
	Payload *models.ConnectedUsers
}

func (o *ChatGetConnectedOK) Error() string {
	return fmt.Sprintf("[GET /chat/connected][%d] chatGetConnectedOK  %+v", 200, *o.Payload)
}

func (o *ChatGetConnectedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectedUsers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChatGetConnectedBadRequest creates a ChatGetConnectedBadRequest with default headers values
func NewChatGetConnectedBadRequest() *ChatGetConnectedBadRequest {
	return &ChatGetConnectedBadRequest{}
}

/*ChatGetConnectedBadRequest handles this case with default header values.

Parameter Error
*/
type ChatGetConnectedBadRequest struct {
	Payload *models.Error
}

func (o *ChatGetConnectedBadRequest) Error() string {
	return fmt.Sprintf("[GET /chat/connected][%d] chatGetConnectedBadRequest  %+v", 400, *o.Payload)
}

func (o *ChatGetConnectedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChatGetConnectedUnauthorized creates a ChatGetConnectedUnauthorized with default headers values
func NewChatGetConnectedUnauthorized() *ChatGetConnectedUnauthorized {
	return &ChatGetConnectedUnauthorized{}
}

/*ChatGetConnectedUnauthorized handles this case with default header values.

Unauthorized
*/
type ChatGetConnectedUnauthorized struct {
	Payload *models.Error
}

func (o *ChatGetConnectedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /chat/connected][%d] chatGetConnectedUnauthorized  %+v", 401, o.Payload)
}

func (o *ChatGetConnectedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChatGetConnectedNotFound creates a ChatGetConnectedNotFound with default headers values
func NewChatGetConnectedNotFound() *ChatGetConnectedNotFound {
	return &ChatGetConnectedNotFound{}
}

/*ChatGetConnectedNotFound handles this case with default header values.

Not Found
*/
type ChatGetConnectedNotFound struct {
	Payload *models.Error
}

func (o *ChatGetConnectedNotFound) Error() string {
	return fmt.Sprintf("[GET /chat/connected][%d] chatGetConnectedNotFound  %+v", 404, o.Payload)
}

func (o *ChatGetConnectedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}