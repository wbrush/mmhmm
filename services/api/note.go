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

var requestFunc = httphelper.MakeHTTPRequest

// swagger:operation POST /v1/notes notes createNote
// ---
// summary: Create new Note
// description: returns new Note
// parameters:
// - name: body
//   in: body
//   description: Note object that needs to be added
//   schema:
//     $ref: "#/definitions/NoteStruct"
//   required: true
// consumes:
//   - application/json
// produces:
//   - application/json
// responses:
//   200:
//     description: "OK"
//     schema:
//       $ref: "#/definitions/NoteStruct"
func (api *API) CreateNote(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested CreateNote()")

	var newNote models.Note
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newNote)
	if err != nil {
		logrus.Warnf("wrong input data provided: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, "body"))
		return
	}

	err = newNote.Validate()
	if err != nil {
		logrus.Warnf("Note data is invalid: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, "body"))
		return
	}

	author, found, err := api.dao.GetUserById(newNote.Author.Id)
	if err == nil && found && author != nil { // TODO: normally would check to make sure the names match??
		if newNote.Author.FirstName == "" {
			newNote.Author.FirstName = author.FirstName
		}
		if newNote.Author.LastName == "" {
			newNote.Author.LastName = author.LastName
		}
	}

	//  could pass in the author found flag but that can make it harder to change in the future. Don't think
	//  this will slow it down noticeably but might want to check performance as development moves forward
	isDuplicate, err := api.dao.CreateNote(&newNote)
	if err != nil {
		logrus.Errorf("Note creation error: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}
	if isDuplicate {
		logrus.Errorf("Note already exist: %d", newNote.Id)
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrAlreadyExists, strconv.FormatInt(int64(newNote.Id), 10)))
		return
	}

	httphelper.Json(w, newNote)
	logrus.Trace("finished CreateNote()")
	return
}

// swagger:operation GET /v1/notes/{id} notes getNote
// ---
// summary: Get note by given id
// description: returns note
// parameters:
// - name: id
//   in: path
//   description: Numeric ID of the note to get
//   type: number
//   required: true
// produces:
//   - application/json
// responses:
//   200:
//     description: "OK"
//     schema:
//       $ref: "#/definitions/NoteStruct"
func (api *API) GetNote(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested GetNote()")

	ids := mux.Vars(r)["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		logrus.Warnf("wrong note id %s provided: %s", ids, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, ids))
		return
	}

	note, found, err := api.dao.GetNoteById(id)
	if err != nil {
		logrus.Errorf("Note id %d select error: %s", id, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}
	if !found || note == nil {
		logrus.Warnf("No such Note by id %d was found", id)
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrNotFound, ids))
		return
	}

	httphelper.Json(w, note)
	logrus.Debug("finished GetNote()")
	return
}

// swagger:operation GET /v1/notes notes listNotes
// ---
// description: returns notes list
// produces:
//   - application/json
// summary: Return list of notes
// parameters:
// - in: query
//   name: authorId
//   description: authors user id to filter by
//   schema:
//     type: string
// responses:
//   '200':
//     description: OK
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/NoteStruct"
func (api *API) ListNotes(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested ListNotes()")

	notes, err := api.dao.ListNotes(r.URL.Query())
	if err != nil {
		logrus.Errorf("ListNotes list select error: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}

	httphelper.Json(w, notes)
	logrus.Trace("finished ListNotes()")
	return
}

// swagger:operation PUT /v1/notes/{id} notes updateNote
// ---
// summary: Update Note by given id
// description: returns Note
// parameters:
// - name: id
//   in: path
//   description: Numeric ID of the Note to update
//   type: number
//   required: true
// - name: body
//   in: body
//   description: Note update object that needs to be updated
//   schema:
//   $ref: "#/definitions/NoteStruct"
//   required: true
// produces:
//   - application/json
// responses:
//   200:
//     description: "OK"
//     schema:
//       $ref: "#/definitions/NoteStruct"
func (api *API) UpdateNote(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested UpdateNote()")

	ids := mux.Vars(r)["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		logrus.Warnf("Error: wrong id %s: %s", ids, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, ids))
		return
	}

	var updNote models.Note
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updNote)
	if err != nil {
		logrus.Warnf("wrong input data provided: %s", err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, "body"))
		return
	}

	if updNote.Id >= 0 {
		updNote.Id = id
	}

	err = api.dao.UpdateNoteById(&updNote)
	if err != nil {
		logrus.Errorf("Note with id %d update error: %s", id, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}

	note := &models.Note{}
	note, _, err = api.dao.GetNoteById(id)
	if err != nil {
		logrus.Errorf("Note with id %d read error: %s", id, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}

	httphelper.Json(w, note)
	logrus.Debug("finished UpdateNote()")
	return
}

// swagger:operation DELETE /v1/notes/{id} notes DeleteNote
// ---
// summary: Delete Template by given id
// description: returns id
// parameters:
//   - name: id
//     in: path
//     description: Numeric ID of the Template to delete
//     type: number
//     required: true
// produces:
//   - application/json
// responses:
//   200:
//     description: "id"
//     schema:
//       type: integer
func (api *API) DeleteNote(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("requested DeleteNote()")

	ids := mux.Vars(r)["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		logrus.Warnf("wrong note id %s provided: %s", ids, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrBadParam, ids))
		return
	}

	found, err := api.dao.DeleteNoteById(id)
	if err != nil {
		logrus.Errorf("note id %d delete error: %s", id, err.Error())
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrService))
		return
	}
	if !found {
		logrus.Warnf("No such note by id %d was found", id)
		httphelper.JsonError(w, errorhandler.NewError(errorhandler.ErrNotFound, ids))
		return
	}

	httphelper.Json(w, map[string]interface{}{"id": id})
	logrus.Debug("finished DeleteNote()")
	return
}
