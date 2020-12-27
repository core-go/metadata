package metadata

type Project struct {
	Root       string            `json:"root,omitempty" bson:"root,omitempty" dynamodbav:"root,omitempty" firestore:"root,omitempty"`
	Env        map[string]string `json:"env,omitempty" bson:"env,omitempty" dynamodbav:"env,omitempty" firestore:"env,omitempty"`
	Statics    []Template        `json:"statics,omitempty" bson:"statics,omitempty" dynamodbav:"statics,omitempty" firestore:"statics,omitempty"`
	Collection []string          `json:"collection,omitempty" bson:"collection,omitempty" dynamodbav:"collection,omitempty" firestore:"collection,omitempty"`
	Arrays     []Template        `json:"arrays,omitempty" bson:"arrays,omitempty" dynamodbav:"arrays,omitempty" firestore:"arrays,omitempty"`
	Entities   []Template        `json:"entities,omitempty" bson:"entities,omitempty" dynamodbav:"entities,omitempty" firestore:"entities,omitempty"`
	TypesFile  string            `json:"typesFile,omitempty" bson:"typesFile,omitempty" dynamodbav:"typesFile,omitempty" firestore:"typesFile,omitempty"`
	Types      map[string]string `json:"types,omitempty" bson:"types,omitempty" dynamodbav:"types,omitempty" firestore:"types,omitempty"`
	ModelsFile string            `json:"modelsFile,omitempty" bson:"modelsFile,omitempty" dynamodbav:"modelsFile,omitempty" firestore:"modelsFile,omitempty"`
	Models     []Model           `json:"models,omitempty" bson:"models,omitempty" dynamodbav:"models,omitempty" firestore:"models,omitempty"`
}
