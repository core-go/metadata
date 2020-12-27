package metadata

type Template struct {
	Name string `json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	File string `json:"file,omitempty" gorm:"column:file" bson:"file,omitempty" dynamodbav:"file,omitempty" firestore:"file,omitempty"`
}