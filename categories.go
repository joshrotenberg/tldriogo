package tldrio

type Category struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Categories returns a slice of TL;DR categories.
func (t *TldrIo) Categories() (*[]Category, error) {

	var categories []Category
	callApi(t, "GET", "categories", "", &categories)
	return &categories, nil
}
