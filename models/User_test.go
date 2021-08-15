package models

import (
	"testing"
)

func TestUser_String(t *testing.T) {

	type request struct {
		user User
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
				user: User{
					Id:        1,
					FirstName: "Test",
					LastName:  "User",
				},
			},
			want: response{
				name: "User<1, Test User>",
			},
		},
		{
			name: "missing FirstName",
			fields: request{
				user: User{
					Id:        99999999,
					FirstName: "",
					LastName:  "User",
				},
			},
			want: response{
				name: "User<99999999,  User>",
			},
		},
		{
			name: "missing FirstName and LastName",
			fields: request{
				user: User{
					Id:        99999999,
					FirstName: "",
					LastName:  "",
				},
			},
			want: response{
				name: "User<99999999,  >",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := tt.fields.user.String()
			if response != tt.want.name {
				t.Errorf("Response *(%s) is not equal to expected (%s)", response, tt.want.name)
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {

	type request struct {
		user User
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
				user: User{
					Id:        1,
					FirstName: "Test1",
					LastName:  "User",
				},
			},
			want: response{
				response: nil,
			},
		},
		{
			name: "invalid id",
			fields: request{
				user: User{
					Id:        -1,
					FirstName: "Test1",
					LastName:  "User",
				},
			},
			want: response{
				response: nil,
			},
		},
		{
			name: "no first name",
			fields: request{
				user: User{
					Id:        1,
					FirstName: "",
					LastName:  "User",
				},
			},
			want: response{
				response: nil,
			},
		},
		{
			name: "no last name",
			fields: request{
				user: User{
					Id:        1,
					FirstName: "Test1",
					LastName:  "",
				},
			},
			want: response{
				response: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := tt.fields.user.Validate()
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
