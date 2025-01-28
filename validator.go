package forms

type Validator interface {
	Validate() (bool, []interface{})
}

type RuleValidator struct {
}
