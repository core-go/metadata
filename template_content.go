package metadata

type TemplateContent struct {
	Name    string `json:"name,omitempty" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty" dynamodbav:"content,omitempty" firestore:"content,omitempty"`
}
