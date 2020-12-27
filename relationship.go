package metadata

type Relationship struct {
	Ref    string `json:"table,omitempty" bson:"table,omitempty" dynamodbav:"table,omitempty" firestore:"table,omitempty"`
	Fields []Link `json:"fields,omitempty" bson:"fields,omitempty" dynamodbav:"fields,omitempty" firestore:"fields,omitempty"`
}
