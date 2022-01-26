package errors

type ApiRequestErrors struct {
	Errors []ValidationErrors `json:"errors"`
}

type ValidationErrors struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

func SingleError(field string, description string) *ApiRequestErrors {

	var items = []ValidationErrors{
		{field, description},
	}

	return &ApiRequestErrors{items}

}
