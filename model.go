package metadata

import "strings"

type Model struct {
	Name      string          `json:"name,omitempty" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Source    string          `json:"source,omitempty" bson:"source,omitempty" dynamodbav:"source,omitempty" firestore:"source,omitempty"`
	Alias     []TypeAlias     `json:"alias,omitempty" bson:"alias,omitempty" dynamodbav:"alias,omitempty" firestore:"alias,omitempty"`
	Ones      []Relationship  `json:"ones,omitempty" bson:"ones,omitempty" dynamodbav:"ones,omitempty" firestore:"ones,omitempty"`
	Models    []Relationship  `json:"models,omitempty" bson:"models,omitempty" dynamodbav:"models,omitempty" firestore:"models,omitempty"`
	Arrays    []Relationship  `json:"arrays,omitempty" bson:"arrays,omitempty" dynamodbav:"arrays,omitempty" firestore:"arrays,omitempty"`
	Fields    []Field         `json:"fields,omitempty" bson:"fields,omitempty" dynamodbav:"fields,omitempty" firestore:"fields,omitempty"`
	WriteFile strings.Builder `json:"-" bson:"-" dynamodbav:"-" firestore:"-"`
}
