package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoryAPI interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	CreateNewCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryWithTasks(w http.ResponseWriter, r *http.Request)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryService}
}

func (c *categoryAPI) GetCategory(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("id").(string)

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	id, _ := strconv.Atoi(userID)

	getCategory, err := c.categoryService.GetCategories(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getCategory)
	// TODO: answer here
}

func (c *categoryAPI) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	log.Println("category: ", category)

	if category.Type == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	userID := r.Context().Value("id").(string)
	id, _ := strconv.Atoi(userID)

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var categoryPointer = &entity.Category{
		Type: category.Type,
	}

	storeCategory, err := c.categoryService.StoreCategory(r.Context(), categoryPointer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	response := map[string]interface{}{
		"user_id":     id,
		"category_id": storeCategory.ID,
		"message":     "success create new category",
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	// TODO: answer here
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("id").(string)
	id, _ := strconv.Atoi(userID)

	categoryID := r.URL.Query().Get("category_id")
	category_id, _ := strconv.Atoi(categoryID)

	deleteCategory := c.categoryService.DeleteCategory(r.Context(), id)

	if deleteCategory != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	response := map[string]interface{}{
		"user_id":     id,
		"category_id": category_id,
		"message":     "success delete category",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	// TODO: answer here
}

func (c *categoryAPI) GetCategoryWithTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategoriesWithTasks(r.Context(), int(idLogin))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

}
