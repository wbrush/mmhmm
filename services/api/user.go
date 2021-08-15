package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/wbrush/go-common/errorhandler"
	"github.com/wbrush/go-common/httphelper"
	"github.com/wbrush/mmhmm/models"
)

// swagger:operation POST /v1/users Users CreateUser
// ---
// summary: Create new user
// description: returns new user
// parameters:
// - name: body
//   in: body
//   description: User object that needs to be added
//   schema:
//       $ref: "#/definitions/UserStruct"
//   required: true
// consumes:
//   - application/json
// produces:
//   - application/json
// responses:
//   200:
//     description: "OK"
//     schema:
//       $ref: "#/definitions/UserStruct"
func (api *API) CreateUser(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested CreateUser")

	var newUser models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		logrus.Warnf("wrong input data provided: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, "body"))
		return
	}

	err = newUser.Validate()
	if err != nil {
		logrus.Warnf("User data is invalid: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, "body"))
		return
	}

	isDuplicate, err := api.dao.CreateUser(&newUser)
	if err != nil {
		logrus.Errorf("User creation error: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}
	if isDuplicate {
		logrus.Errorf("User already exist: %d", newUser.Id)
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrAlreadyExists, strconv.FormatInt(int64(newUser.Id), 10)))
		return
	}

	httphelper.Json(w, newUser)
	logrus.Debug("finished CreateUser()")
	return
}

// swagger:operation GET /v1/users/{id} Users getUser
// ---
// summary: Get User by given id
// description: returns user
// parameters:
// - name: id
//   in: path
//   description: Numeric ID of the user to get
//   type: number
//   required: true
// produces:
//   - application/json
// responses:
//   200:
//     description: "OK"
//     schema:
//       $ref: "#/definitions/UserStruct"
func (api *API) GetUser(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested GetUser()")

	ids := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		logrus.Warnf("wrong Note id %s provided: %s", ids, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, ids))
		return
	}

	user, found, err := api.dao.GetUserById(int(id))
	if err != nil {
		logrus.Errorf("user id %d select error: %s", id, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}
	if !found || user == nil {
		logrus.Warnf("No such user by id %d was found", id)
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrNotFound, ids))
		return
	}

	httphelper.Json(w, user)
	logrus.Debug("finished GetUser()")
	return
}

// swagger:operation GET /v1/users Users listUsers
// ---
// description: returns users list
// produces:
//   - application/json
// summary: Return list of users
// parameters:
// - in: query
//   name: orderBy
//   description: on which field results must to be ordered
//   schema:
//     type: string
// - in: query
//   name: filters
//   description: 'filters on fields, which can be provided in format described here:
//     https://godoc.org/github.com/go-pg/pg/urlvalues#Filter'
//   schema:
//     type: object
// responses:
//   '200':
//     description: OK
//     schema:
//       type: array
//       items:
//        $ref: "#/definitions/UserStruct"
func (api *API) ListUsers(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested ListUsers()")

	Users, err := api.dao.ListUsers(r.URL.Query())
	if err != nil {
		logrus.Errorf("ListUsers() list select error: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}

	httphelper.Json(w, Users)
	logrus.Debug("finished ListUsers()")
	return
}

// swagger:operation PUT /v1/users/{id} Users updateUser
// ---
// summary: Update User by given id
// description: returns user
// parameters:
// - name: id
//   in: path
//   description: Numeric ID of the user to update
//   type: number
//   required: true
// - name: body
//   in: body
//   description: User object that needs to be updated
//   schema:
//      $ref: "#/definitions/UserStruct"
//   required: true
// produces:
//   - application/json
// responses:
//   200:
//     description: "OK"
//     schema:
//        $ref: "#/definitions/UserStruct"
func (api *API) UpdateUser(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested UpdateUser")

	ids := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		logrus.Warnf("wrong User id %s: %s", ids, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, ids))
		return
	}

	user, found, err := api.dao.GetUserById(int(id))
	if err != nil {
		logrus.Errorf("user id %d select error: %s", id, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}
	if !found || user == nil {
		logrus.Warnf("No such user by id %d was found", id)
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrNotFound, ids))
		return
	}

	var updUser models.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updUser)
	if err != nil {
		logrus.Warnf("wrong input data provided: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, "body"))
		return
	}

	isFoundUpd := false
	if updUser.FirstName != "" && updUser.FirstName != user.FirstName {
		user.FirstName = updUser.FirstName
		isFoundUpd = true
	}
	if updUser.LastName != "" && updUser.LastName != user.LastName {
		user.LastName = updUser.LastName
		isFoundUpd = true
	}
	if err != nil {
		logrus.Warnf("cannot fill updateable fields: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}
	if !isFoundUpd {
		logrus.Infof("nothing to update for User %d", user.Id)
		httphelper.Json(w, user)
		return
	}

	err = api.dao.UpdateUserById(user)
	if err != nil {
		logrus.Errorf("User with id %d update error: %s", id, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}

	httphelper.Json(w, user)
	logrus.Debug("finished UpdateUser()")
	return
}

// swagger:operation DELETE /v1/users/{id} Users deleteUser
// ---
// summary: Delete user by given id
// description: returns id
// parameters:
//   - name: id
//     in: path
//     description: Numeric ID of the User to delete
//     type: number
//     required: true
// produces:
//   - application/json
// responses:
//   200:
//     description: "id"
//     schema:
//       type: integer
func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested DeleteUser()")

	ids := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		logrus.Warnf("wrong User id %s provided: %s", ids, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, ids))
		return
	}

	found, err := api.dao.DeleteUserById(int(id))
	if err != nil {
		logrus.Errorf("User id %d delete error: %s", id, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}
	if !found {
		logrus.Warnf("No such User by id %d was found", id)
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrNotFound, ids))
		return
	}

	httphelper.Json(w, map[string]interface{}{"id": id})
	logrus.Debug("finished DeleteUser()")
	return
}
