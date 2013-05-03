package tldrio

type Category struct {
	Name string `json:"name"`
}

func (t *Tldr) Categories() { (categories []Category, error) {

	var r interface{}
	callApi(t, "GET", "categories", "", &r)
}
