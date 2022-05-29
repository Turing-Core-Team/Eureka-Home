package columns

import "strings"

const (
	carreras     string = "carreras"
	tecnologos   string = "tecnologos"
	cursos       string = "cursos"
	diplomados   string = "diplomados"
	concursos    string = "concursos"
	eventos      string = "eventos"
	divulgacion  string = "divulgacion"
	financiacion string = "financiacion"
	ferias       string = "ferias"
)

var Index = map[string]string{
	carreras:   "A",
	tecnologos: "B",
	cursos:     "C",
	diplomados: "D",

	concursos:    "N",
	eventos:      "O",
	divulgacion:  "P",
	financiacion: "Q",
	ferias:       "R",
}

func GetRange(isFirstPartition bool, columnName string) (string, error) {
	columnName = strings.ToLower(columnName)
	if isFirstPartition {
		switch columnName {
		case carreras:
			return Index[carreras], nil
		case tecnologos:
			return Index[tecnologos], nil
		case cursos:
			return Index[cursos], nil
		case diplomados:
			return Index[diplomados], nil
		default:
			return FindEquivalence(isFirstPartition, columnName)
		}
	} else {
		switch columnName {
		case concursos:
			return Index[concursos], nil
		case eventos:
			return Index[eventos], nil
		case divulgacion:
			return Index[divulgacion], nil
		case financiacion:
			return Index[financiacion], nil
		case ferias:
			return Index[ferias], nil
		default:
			return FindEquivalence(isFirstPartition, columnName)
		}
	}
}
