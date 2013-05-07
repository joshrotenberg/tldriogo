package tldrio

type Category struct {
	Name string
	Slug string
}

// Categories returns a slice of TL;DR categories.
func (t *TldrIo) Categories() (*[]Category, error) {

	var categories []Category
	callApi(t, "GET", "categories", "", nil, &categories)
	return &categories, nil
}
