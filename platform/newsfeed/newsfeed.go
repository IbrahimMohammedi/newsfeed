package newsfeed

// Item sturcts : store the news feed items.
type Item struct {
	Title string `json:"title"`
	Post  string `json :"post"`
}

// Repo struct : Creates a list of Items from Item.
type Repo struct {
	// remark : when we create repo it doesn't instantiate any items inside of that list
	Items []Item
}

// New func: returns a pointer to Repo.
func New() *Repo {
	return &Repo{
		// we have to manually declare the items inside the constructer func, otherwise we get null because
		// when we create repo it doesn't instantiate any items inside of that list
		Items: []Item{},
	}
}

// Add func: Adds an item to Items, recieves a pointer to Repo.
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// GetAll func: returns the value of Items, recieves a pointer to Repo.
func (r *Repo) GetAll() []Item {
	return r.Items
}
