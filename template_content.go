package metadata

type TemplateContent struct {
	Name    string `json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Content string `json:"content,omitempty" gorm:"column:content" bson:"content,omitempty" dynamodbav:"content,omitempty" firestore:"content,omitempty"`
}
