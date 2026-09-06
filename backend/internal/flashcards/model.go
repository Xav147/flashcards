package flashcards

type Deck struct {
	Name string `bson:"name" json:"name"`
	Size int    `bson:"size" json:"size"`
}

type Flashcard struct {
	Title   string   `bson:"title" json:"title"`
	Content string   `bson:"content" json:"content"`
	Tags    []string `bson:"tags" json:"tags"`
}
