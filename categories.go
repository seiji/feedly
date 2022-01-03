package feedly

type APICategories struct {
	api *apiV3
}

type Category struct {
	Id    string `json:"id"`
	Label string `json:"label"`
}

func (a *APICategories) Get() ([]Category, *Response, error) {
	rel := "categories"

	req, err := a.api.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	categories := new([]Category)

	res, err := a.api.Do(req, categories)
	if err != nil {
		return nil, res, err
	}

	return *categories, res, nil
}
