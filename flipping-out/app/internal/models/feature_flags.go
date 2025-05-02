package models

type FeatureFlag struct {
	ID          string             `bson:"_id,omitempty" json:"id,omitempty"`
	Flag        string             `bson:"flag" json:"flag"`
	Variations  VariationSet       `bson:"variations" json:"variations"`
	DefaultRule FeatureFlagRuleSet `bson:"defaultRule" json:"defaultRule"`
}

type VariationSet struct {
	DefaultVar bool `bson:"default_var" json:"default_var"`
	FalseVar   bool `bson:"false_var" json:"false_var"`
	TrueVar    bool `bson:"true_var" json:"true_var"`
}

type FeatureFlagRuleSet struct {
	Percentage map[string]int `bson:"percentage" json:"percentage"`
}
