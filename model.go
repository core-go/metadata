package metadata

import "strings"

type Model struct {
	Name      string          `json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Source    string          `json:"source,omitempty" gorm:"column:source" bson:"source,omitempty" dynamodbav:"source,omitempty" firestore:"source,omitempty"`
	Alias     []TypeAlias     `json:"alias,omitempty" gorm:"column:alias" bson:"alias,omitempty" dynamodbav:"alias,omitempty" firestore:"alias,omitempty"`
	Ones      []Relationship  `json:"ones,omitempty" gorm:"column:ones" bson:"ones,omitempty" dynamodbav:"ones,omitempty" firestore:"ones,omitempty"`
	Models    []Relationship  `json:"models,omitempty" gorm:"column:models" bson:"models,omitempty" dynamodbav:"models,omitempty" firestore:"models,omitempty"`
	Arrays    []Relationship  `json:"arrays,omitempty" gorm:"column:arrays" bson:"arrays,omitempty" dynamodbav:"arrays,omitempty" firestore:"arrays,omitempty"`
	Fields    []Field         `json:"fields,omitempty" gorm:"column:fields" bson:"fields,omitempty" dynamodbav:"fields,omitempty" firestore:"fields,omitempty"`
	WriteFile strings.Builder `json:"-" bson:"-" dynamodbav:"-" firestore:"-"`
}
