package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/category"
	"github.com/hktrib/RetailGo/internal/ent/item"
)

/*
CatDelete Brief:

	________________________________________

-> Deletes category given category id specified in Query Parameters (Ex: ?category=10)

Response:

-> On success, returns 200.
-> If no category supplied, 400.
-> If the specified category is not found, 404.

*/

func (srv *Server) CatDelete(w http.ResponseWriter, r *http.Request) {

	// Read category parameter
	catId, err := strconv.Atoi(r.URL.Query().Get("category"))
	if err != nil {

		log.Debug().Err(err).Msg("CatDelete: unable to parse category from URL")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Delete the category with this id in the database.
	err = srv.DBClient.
		Category.DeleteOneID(catId).
		Exec(r.Context())

	if ent.IsNotFound(err) {
		log.Debug().Err(err).Msg("CatDelete: unable to find category in database")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		log.Debug().Err(err).Msg("CatDelete: server failed to delete category from database")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

/*
CatCreate Brief:

	________________________________________

-> Creates category given store_id specified in query parameters (ex: ?store_id=10), and at least one of the category's name and its attached photo in the request body.

Response:

-> On success, returns 201.
-> If no store_id in URL or no category supplied (requires at least photo or name) in Body, 400.

*/

func (srv *Server) CatCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Read store_id parameter
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if err != nil {
		log.Debug().Err(err).Msg("CatCreate: no store id provided")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Load the message parameters (Name, Photo, Quantity, Category)
	req_body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Debug().Err(err).Msg("CatCreate: server failed to read message body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Parse requested category.
	req := ent.Category{}
	body_parse_err := json.Unmarshal(req_body, &req)

	if body_parse_err != nil {
		log.Debug().Err(err).Msg("CatCreate: server failed to Unmarshal to request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Create new category with provided name, photo and attached to the provided store id.
	cat, err := srv.DBClient.Category.Create().SetName(req.Name).SetPhoto(req.Photo).SetStoreID(store_id).Save(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("CatCreate: server failed to create category")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Response as JSON with the id of the newly created category.
	responseBody, _ := json.Marshal(map[string]interface{}{
		"id": cat.ID,
	})

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}

/*
CatItemRead Brief:

-> Retrieves items belonging to a specified category by category id in Query Parameters (Ex: ?category_id=15)

Response:

-> On success, returns a 200, with a list of items in body.
-> If no category id or an invalid category id is provided, returns 400.

*/

func (srv *Server) CatItemRead(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	// Read category id parameter
	category_id, err := strconv.Atoi(chi.URLParam(r, "category_id"))
	if err != nil {
		log.Debug().Err(err).Msg("CatItemRead: Invalid category ID")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Query for the specified category.
	targetCategory, err := srv.DBClient.Category.Query().Where(category.ID(category_id)).Only(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("CatItemRead: Unable to Query for Category")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := make(map[string][]*ent.Item)

	// get items for the target category
	items, err := srv.DBClient.Item.Query().Where(item.HasCategoryWith(category.ID(targetCategory.ID))).All(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("CatItemRead: unable to query for all items in category")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response[targetCategory.Name] = append(response[targetCategory.Name], items...)

	// send json response
	responseBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

/*
CatItemAdd Brief:

-> Adds items (provided as ids) from the request body to the category specified by category id from the URL parameters (Example: /?category_id=10).
Note: Adds edges, DOES NOT SET CATEGORY NAME.

Response:

-> On success, returns 200.
-> If an invalid category id is provided, return 400.
-> If the category doesn't exist, return 404.

*/

func (srv *Server) CatItemAdd(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Read category id parameter.
	cat_id, err := strconv.Atoi(chi.URLParam(r, "category_id"))

	if err != nil {
		log.Debug().Err(err).Msg("CatItemAdd: Invalid category id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Load the message parameters (List of Item ids)
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Debug().Err(err).Msg("CatItemAdd: server failed to read message body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var itemIDs []int

	// Load all supplied item ids
	body_parse_err := json.Unmarshal(reqBody, &itemIDs)

	if body_parse_err != nil {
		log.Debug().Err(body_parse_err).Msg("CatItemAdd: server failed to read message body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Add Item ids in batch to the category in the database, by linking the category and these items.
	err = srv.DBClient.Category.UpdateOneID(cat_id).AddItemIDs(itemIDs...).Exec(ctx)
	if ent.IsNotFound(err) {
		log.Debug().Err(body_parse_err).Msg("CatItemAdd: server query of cat_id database")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		log.Debug().Err(body_parse_err).Msg("CatItemAdd: server query of cat_id database")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

/*
CatItemList Brief:

-> Provides list of all categories for the store specified by store_id url parameter (Example: /category/10)

Response:

-> On Success, returns 200, with list of categories ({"id": _, "name": _}) in the message body.
-> If invalid store id, return 400.

*/

func (srv *Server) CatItemList(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	// Read store id parameter.
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		fmt.Println("CatItemList: Invalid store id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Query all categories for this store.
	categories, err := srv.DBClient.Category.Query().Where(category.StoreID(store_id)).All(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("CatItemList: Unable to Query for Category Item list")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// JSON response with list of categories, {"id": _, "name": _}.

	response := make([]map[string]interface{}, 0)

	for _, cat := range categories {
		mp := make(map[string]interface{}, 1)
		mp["id"] = cat.ID
		mp["name"] = cat.Name
		response = append(response, mp)

	}
	responseBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
