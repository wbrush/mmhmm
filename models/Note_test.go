package models

import (
	"errors"
	"testing"
)

func TestNote_String(t *testing.T) {

	type request struct {
		note Note
	}
	type response struct {
		name string
	}

	tests := []struct {
		name   string
		fields request
		want   response
	}{
		{
			name: "happy path",
			fields: request{
				note: Note{
					Id: 1,
				},
			},
			want: response{
				name: "Note<1>",
			},
		},
		{
			name: "happy path - large number",
			fields: request{
				note: Note{
					Id: 9999999,
				},
			},
			want: response{
				name: "Note<9999999>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := tt.fields.note.String()
			if response != tt.want.name {
				t.Errorf("Response *(%s) is not equal to expected (%s)", response, tt.want.name)
			}
		})
	}
}

func TestNote_Validate(t *testing.T) {

	type request struct {
		note Note
	}
	type response struct {
		response error
	}

	tests := []struct {
		name   string
		fields request
		want   response
	}{
		{
			name: "happy path",
			fields: request{
				note: Note{
					Id: 1,
					Author: &User{
						Id:        1,
						FirstName: "Test1",
						LastName:  "User",
					},
				},
			},
			want: response{
				response: nil,
			},
		},
		{
			name: "no author",
			fields: request{
				note: Note{
					Id: 1,
				},
			},
			want: response{
				response: errors.New(NoteValidationAuthorMissingError),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := tt.fields.note.Validate()
			if response == nil && tt.want.response == nil {
				//  do nothing
			} else if (response == nil && tt.want.response != nil) || (response != nil && tt.want.response == nil) {
				t.Errorf("Response is not equal to expected")
			} else {
				if response.Error() != tt.want.response.Error() {
					t.Errorf("Response *(%s) is not equal to expected (%s)", response.Error(), tt.want.response.Error())
				}
			}
		})
	}
}
