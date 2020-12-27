package metadata

type Link struct {
	Column string `json:"column,omitempty" bson:"column,omitempty" dynamodbav:"column,omitempty" firestore:"column,omitempty"`
	To     string `json:"to,omitempty" bson:"to,omitempty" dynamodbav:"to,omitempty" firestore:"to,omitempty"`
}
