package metadata

type Field struct {
	Name   string `json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Source string `json:"source,omitempty" gorm:"column:source" bson:"source,omitempty" dynamodbav:"source,omitempty" firestore:"source,omitempty"`
	Type   string `json:"type,omitempty" gorm:"column:type" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty"`
	Id     bool   `json:"id,omitempty" gorm:"column:id" bson:"id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty"`
}
