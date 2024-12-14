package utils

import (
	"github.com/aejoy/go-pkg/consts"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateNanoID() (string, error) {
	return gonanoid.Generate(consts.NanoAlphabet, consts.NanoLength)
}
