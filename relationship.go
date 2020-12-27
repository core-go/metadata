package metadata

type Relationship struct {
	Ref    string `json:"table,omitempty" gorm:"column:table" bson:"table,omitempty" dynamodbav:"table,omitempty" firestore:"table,omitempty"`
	Fields []Link `json:"fields,omitempty" gorm:"column:root" bson:"fields,omitempty" dynamodbav:"fields,omitempty" firestore:"fields,omitempty"`
}
