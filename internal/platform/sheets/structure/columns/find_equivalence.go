package columns

import "errors"

var DictEquivalences = map[string]string{
	"ingeniería" : "ingenieria",
}

func FindEquivalence(columnName string) (string, error) {
	eq, err := DictEquivalences[columnName]

	if !err{
		return "", errors.New("equivalence not found")
	}
	return eq, nil
}
