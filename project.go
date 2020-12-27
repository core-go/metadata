package metadata

type Project struct {
	Root       string            `json:"root,omitempty" gorm:"column:root" bson:"root,omitempty" dynamodbav:"root,omitempty" firestore:"root,omitempty"`
	Env        map[string]string `json:"env,omitempty" gorm:"column:env" bson:"env,omitempty" dynamodbav:"env,omitempty" firestore:"env,omitempty"`
	Statics    []Template        `json:"statics,omitempty" gorm:"column:statics" bson:"statics,omitempty" dynamodbav:"statics,omitempty" firestore:"statics,omitempty"`
	Collection []string          `json:"collection,omitempty" gorm:"column:collection" bson:"collection,omitempty" dynamodbav:"collection,omitempty" firestore:"collection,omitempty"`
	Arrays     []Template        `json:"arrays,omitempty" gorm:"column:arrays" bson:"arrays,omitempty" dynamodbav:"arrays,omitempty" firestore:"arrays,omitempty"`
	Entities   []Template        `json:"entities,omitempty" gorm:"column:entities" bson:"entities,omitempty" dynamodbav:"entities,omitempty" firestore:"entities,omitempty"`
	TypesFile  string            `json:"typesFile,omitempty" gorm:"column:typesfile" bson:"typesFile,omitempty" dynamodbav:"typesFile,omitempty" firestore:"typesFile,omitempty"`
	Types      map[string]string `json:"types,omitempty" gorm:"column:types" bson:"types,omitempty" dynamodbav:"types,omitempty" firestore:"types,omitempty"`
	ModelsFile string            `json:"modelsFile,omitempty" gorm:"column:modelsfile" bson:"modelsFile,omitempty" dynamodbav:"modelsFile,omitempty" firestore:"modelsFile,omitempty"`
	Models     []Model           `json:"models,omitempty" gorm:"column:models" bson:"models,omitempty" dynamodbav:"models,omitempty" firestore:"models,omitempty"`
}
