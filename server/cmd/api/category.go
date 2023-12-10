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

External Package Calls:

	DBClient.Category.DeleteOneID
*/
func (srv *Server) CatDelete(w http.ResponseWriter, r *http.Request) {

	catId, err := strconv.Atoi(r.URL.Query().Get("category"))
	if err != nil {

		log.Debug().Err(err).Msg("CatDelete: unable to parse category from URL")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

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

-> Creates category given category id specified in Query Parameters (Ex: ?category=10)

External Package Calls:

	DBClient.Category.Create()
*/
func (srv *Server) CatCreate(w http.ResponseWriter, r *http.Request) {
	// name, photo, quantity, store_id, category

	ctx := r.Context()
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	// Load store_id first

	// Load the message parameters (Name, Photo, Quantity, Category)
	req_body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Debug().Err(err).Msg("CatCreate: server failed to read message body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req := ent.Category{}
	body_parse_err := json.Unmarshal(req_body, &req)

	if body_parse_err != nil {
		log.Debug().Err(err).Msg("CatCreate: server failed to Unmarshal to request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cat, err := srv.DBClient.Category.Create().SetName(req.Name).SetPhoto(req.Photo).SetStoreID(store_id).Save(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("CatCreate: server failed to create category")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(map[string]interface{}{
		"id": cat.ID,
	})

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}

/*
CatItemRead Brief:

-> Retrieves items belonging to a specified category by category ID from Query Parameters (Ex: ?category_id=15)

External Package Calls:
- DBClient.Category.Query()
- srv.DBClient.Item.Query()
*/
func (srv *Server) CatItemRead(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	category_id, err := strconv.Atoi(chi.URLParam(r, "category_id"))
	if err != nil {
		log.Debug().Err(err).Msg("CatItemRead: Invalid Store ID")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	targetCategory, err := srv.DBClient.Category.Query().Where(category.ID(category_id)).Only(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("CatItemRead: Unable to Query for Category")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// send json response
	response := make(map[string][]*ent.Item)

	// get items for each targetCategory
	items, err := srv.DBClient.Item.Query().Where(item.HasCategoryWith(category.ID(targetCategory.ID))).All(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("CatItemRead: unable to query for all items in category")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response[targetCategory.Name] = append(response[targetCategory.Name], items...)

	responseBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

/*
CatItemAdd Brief:

-> Adds items to a category specified by category ID from the URL parameters.

External Package Calls:
- srv.DBClient.Category.UpdateOneID()
*/
func (srv *Server) CatItemAdd(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	cat_id, err := strconv.Atoi(chi.URLParam(r, "category_id"))

	if err != nil {
		log.Debug().Err(err).Msg("CatItemAdd: Invalid category id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Load the message parameters (Name, Photo, Quantity, Category)
	req_body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Debug().Err(err).Msg("CatItemAdd: server failed to read message body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var itemIDs []int

	body_parse_err := json.Unmarshal(req_body, &itemIDs)

	if body_parse_err != nil {
		log.Debug().Err(body_parse_err).Msg("CatItemAdd: server failed to read message body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.Category.UpdateOneID(cat_id).AddItemIDs(itemIDs...).Exec(ctx)
	if ent.IsNotFound(err) {
		log.Debug().Err(body_parse_err).Msg("CatItemAdd: server query of cat_id database")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
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

func (srv *Server) CatItemList(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		fmt.Println("CatItemList: Invalid store id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	categories, err := srv.DBClient.Category.Query().Where(category.StoreID(store_id)).All(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("CatItemList: Unable to Query for Category Item list")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
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
