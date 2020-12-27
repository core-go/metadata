package metadata

type Link struct {
	Column string `json:"column,omitempty" gorm:"column:column" bson:"column,omitempty" dynamodbav:"column,omitempty" firestore:"column,omitempty"`
	To     string `json:"to,omitempty" gorm:"column:to" bson:"to,omitempty" dynamodbav:"to,omitempty" firestore:"to,omitempty"`
}
