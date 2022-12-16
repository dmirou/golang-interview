package concat

import (
	"bytes"
	"strings"
)

func slowConcat(elems []string) string {
	var res string

	for _, elem := range elems {
		res += elem
	}

	return res
}

func bufConcat(elems []string) string {
	var b bytes.Buffer

	for _, elem := range elems {
		b.WriteString(elem)
	}

	return b.String()
}

func builderConcat(elems []string) string {
	var b strings.Builder

	for _, elem := range elems {
		b.WriteString(elem)
	}

	return b.String()
}
