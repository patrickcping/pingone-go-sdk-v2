package model

import (
	"fmt"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
)

func TestRemarshalGenericOpenAPIErrorObj_Success(t *testing.T) {

	errorName := "418 I'm a teapot"
	p1ErrorId := "dummyerrorid"
	p1ErrorCode := "ACCESS_FAILED"
	p1ErrorMessage := "The request could not be completed. You do not have access to this resource."
	p1ErrorDetailsCode := "INSUFFICIENT_PERMISSIONS"
	p1ErrorDetailsMessage := "The actor attempting to perform the request is not authorized."

	moduleApiError := GenericOpenAPIError{
		body:  []byte(fmt.Sprintf(`{"id": "%s","code": "%s","message": "%s","details": [{"code": "INSUFFICIENT_PERMISSIONS","message": "The actor attempting to perform the request is not authorized."}]}`, p1ErrorId, p1ErrorCode, p1ErrorMessage)),
		error: errorName,
		model: &management.P1Error{
			Id:      &p1ErrorId,
			Code:    &p1ErrorCode,
			Message: &p1ErrorMessage,
			Details: []management.P1ErrorDetailsInner{
				{
					Code:    &p1ErrorDetailsCode,
					Message: &p1ErrorDetailsMessage,
				},
			},
		},
	}

	v, err := RemarshalGenericOpenAPIErrorObj(moduleApiError)

	if err != nil {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %v", err)
	}

	if v.error != errorName {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", v.error, errorName)
	}

	if c, ok := v.Model().(*P1Error).GetIdOk(); !ok || *c != p1ErrorId {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorId)
	}

	if c, ok := v.Model().(*P1Error).GetCodeOk(); !ok || *c != p1ErrorCode {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorCode)
	}

	if c, ok := v.Model().(*P1Error).GetMessageOk(); !ok || *c != p1ErrorMessage {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorMessage)
	}

	if c, ok := v.Model().(*P1Error).GetDetailsOk(); ok {

		if c, ok := c[0].GetCodeOk(); !ok || *c != p1ErrorDetailsCode {
			t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorDetailsCode)
		}

		if c, ok := c[0].GetMessageOk(); !ok || *c != p1ErrorDetailsMessage {
			t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorDetailsMessage)
		}
	} else {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj model missing")
	}
}

func TestRemarshalErrorObj_Success(t *testing.T) {

	p1ErrorId := "dummyerrorid"
	p1ErrorCode := "ACCESS_FAILED"
	p1ErrorMessage := "The request could not be completed. You do not have access to this resource."
	p1ErrorDetailsCode := "INSUFFICIENT_PERMISSIONS"
	p1ErrorDetailsMessage := "The actor attempting to perform the request is not authorized."

	moduleApiError := &management.P1Error{
		Id:      &p1ErrorId,
		Code:    &p1ErrorCode,
		Message: &p1ErrorMessage,
		Details: []management.P1ErrorDetailsInner{
			{
				Code:    &p1ErrorDetailsCode,
				Message: &p1ErrorDetailsMessage,
			},
		},
	}

	v, err := RemarshalErrorObj(moduleApiError)

	if err != nil {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %v", err)
	}

	if c, ok := v.GetIdOk(); !ok || *c != p1ErrorId {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorId)
	}

	if c, ok := v.GetCodeOk(); !ok || *c != p1ErrorCode {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorCode)
	}

	if c, ok := v.GetMessageOk(); !ok || *c != p1ErrorMessage {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorMessage)
	}

	if c, ok := v.GetDetailsOk(); ok {

		if c, ok := c[0].GetCodeOk(); !ok || *c != p1ErrorDetailsCode {
			t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorDetailsCode)
		}

		if c, ok := c[0].GetMessageOk(); !ok || *c != p1ErrorDetailsMessage {
			t.Fatalf("TestRemarshalGenericOpenAPIErrorObj resulted in %s, expected %v", *c, p1ErrorDetailsMessage)
		}
	} else {
		t.Fatalf("TestRemarshalGenericOpenAPIErrorObj model missing")
	}
}
