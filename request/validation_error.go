package request

import "fmt"

type validationError struct {
	explanations map[string][]string
}

func newValidationError() validationError {
	expl := make(map[string][]string)
	err := validationError{}
	err.explanations = expl
	return err
}

func (err *validationError) addExplanation(fieldName string, explanation string) {
	err.explanations[fieldName] = append(err.explanations[fieldName], explanation)
}

func (err validationError) Error() string {
	var errString string = "Invalid request:\n"
	for field, errors := range err.explanations {
		for _, err := range errors {
			errString += fmt.Sprintf("\t%s %s\n", field, err)
		}
	}
	return errString
}
