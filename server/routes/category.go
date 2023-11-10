package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/ent/category"
	"github.com/hktrib/RetailGo/ent/item"
	"io"
	"net/http"
	"strconv"

	"github.com/hktrib/RetailGo/ent"
)

func (srv *Server) CatDelete(w http.ResponseWriter, r *http.Request) {

	catId, err := strconv.Atoi(r.URL.Query().Get("category"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.
		Category.DeleteOneID(catId).
		Exec(r.Context())

	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (srv *Server) CatCreate(w http.ResponseWriter, r *http.Request) {
	// name, photo, quantity, store_id, category

	ctx := r.Context()
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	// Load store_id first

	// Load the message parameters (Name, Photo, Quantity, Category)
	req_body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Server failed to read message body:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req := ent.Category{}

	body_parse_err := json.Unmarshal(req_body, &req)

	if body_parse_err != nil {
		fmt.Println("Unmarshalling failed:", body_parse_err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cat, create_err := srv.DBClient.Category.Create().SetName(req.Name).SetPhoto(req.Photo).SetStoreID(store_id).Save(ctx)

	if create_err != nil {
		fmt.Println("Create didn't work:", create_err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(map[string]interface{}{
		"id": cat.ID,
	})

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}

func (srv *Server) CatItemRead(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	category_id, err := strconv.Atoi(chi.URLParam(r, "category_id"))
	if err != nil {
		fmt.Println("Invalid store id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	targetCategory, err := srv.DBClient.Category.Query().Where(category.ID(category_id)).Only(ctx)
	// send json response
	response := make(map[string][]string)

	// get items for each targetCategory
	items, _ := srv.DBClient.Item.Query().Where(item.HasCategoryWith(category.ID(targetCategory.ID))).All(ctx)
	for _, item := range items {
		itemData, _ := json.Marshal(item)
		response[targetCategory.Name] = append(response[targetCategory.Name], string(itemData))
	}

	responseBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (srv *Server) CatItemAdd(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		fmt.Println("Invalid store id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	cat_id, err := strconv.Atoi(chi.URLParam(r, "category_id"))

	if err != nil {
		fmt.Println("Invalid category id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Load the message parameters (Name, Photo, Quantity, Category)
	req_body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Server failed to read message body:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req := ent.Item{}

	body_parse_err := json.Unmarshal(req_body, &req)

	if body_parse_err != nil {
		fmt.Println("Unmarshalling failed:", body_parse_err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	createdItem, err := srv.DBClient.Item.Create().SetName(req.Name).SetPhoto(req.Photo).SetQuantity(req.Quantity).SetStoreID(store_id).SetPrice(req.Price).AddCategoryIDs(cat_id).Save(ctx)

	if err != nil {
		fmt.Println("Create didn't work:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	srv.DBClient.Category.UpdateOneID(cat_id).AddItems(createdItem).ExecX(ctx)

	responseBody, _ := json.Marshal(map[string]interface{}{
		"id": createdItem.ID,
	})

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)

}
func (srv *Server) CatItemList(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		fmt.Println("Invalid store id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	categories, err := srv.DBClient.Category.Query().Where(category.StoreID(store_id)).All(ctx)
	// send json response
	response := make(map[string][]string)

	for _, cat := range categories {
		// get items for each category
		items, _ := srv.DBClient.Item.Query().Where(item.HasCategoryWith(category.ID(cat.ID))).All(ctx)
		for _, item := range items {
			itemData, _ := json.Marshal(item)
			response[cat.Name] = append(response[cat.Name], string(itemData))
		}

	}
	responseBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
