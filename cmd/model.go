package cmd

type KeyVal struct {
	Key   string `validate:"min=1,max=50"`
	Value string `validate:"min=1"`
}
