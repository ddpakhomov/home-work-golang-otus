package models

type Term struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Term        string `json:"term" bson:"term,omitempty"`
	Name        string `json:"name", bson:"name"`
	Description string `json:"description" bson:"description,omitempty"`
	Theme       string `json:"theme" bson:"theme,omitempty"`
}
