package http

import (
	"net/http"
	"testing"
	"time"

	"github.com/cloudevents/sdk-go"
)

func TestClientConversion_v1(t *testing.T) {
	now := time.Now()

	testCases := ConversionTestCases{
		"Conversion v1.0": {
			now:       now,
			convertFn: UnitTestConvert,
			data:      map[string]string{"hello": "unittest"},
			want: &cloudevents.Event{
				Context: cloudevents.EventContextV1{
					ID:     "321-CBA",
					Type:   "io.cloudevents.conversion.http.post",
					Source: *cloudevents.ParseURIRef("github.com/cloudevents/test/http/conversion"),
				}.AsV1(),
				Data: map[string]string{"hello": "unittest"},
			},
			asSent: &TapValidation{
				Method: "POST",
				URI:    "/",
				Header: map[string][]string{
					"content-type": {"application/json"},
				},
				Body:          `{"hello":"unittest"}`,
				ContentLength: 20,
			},
			asRecv: &TapValidation{
				Header:        http.Header{},
				Status:        "202 Accepted",
				ContentLength: 0,
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			ClientConversion(t, tc)
		})
	}
}

func TestClientConversion_nil(t *testing.T) {
	now := time.Now()

	testCases := ConversionTestCases{
		"Conversion NOP": {
			now:       now,
			convertFn: UnitTestConvertNilNil,
			data:      map[string]string{"hello": "unittest"},
			want:      nil,
			asSent: &TapValidation{
				Method: "POST",
				URI:    "/",
				Header: map[string][]string{
					"content-type": {"application/json"},
				},
				Body:          `{"hello":"unittest"}`,
				ContentLength: 20,
			},
			asRecv: &TapValidation{
				Header:        http.Header{},
				Status:        "204 No Content",
				ContentLength: 0,
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			ClientConversion(t, tc)
		})
	}
}
