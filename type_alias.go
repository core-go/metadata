package metadata

type TypeAlias struct {
	Name string `json:"name,omitempty" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty"`
}
