package request

var allowedRequestTypes = []string{"GET", "POST", "PATCH", "HEAD", "PUT", "DELETE"}
var allowedBodyFormats = []string{"json", "form", "raw"}

type Request struct {
	BaseUrl     string "base_url"
	Path        string
	BodyFormat  string "body_format"
	Method      string
	Url         string
	Body        map[string]string // This should be an interface to respect yaml types!
	Headers     map[string]string
	QueryString map[string]string "query_str"
}

func New() *Request {
	return new(Request)
}

func (request Request) IsValid() (bool, error) {
	err := newValidationError()
	isValid := false

	if request.Url == "" {
		err.addExplanation("url", "cannot be blank")
	}
	if request.Method == "" {
		err.addExplanation("method", "cannot be blank")
	}
	if request.BodyFormat == "" {
		err.addExplanation("body_format", "cannot be blank")
	}
	if !isStringInArray(request.Method, allowedRequestTypes) {
		err.addExplanation("method", "is not allowed")
	}
	if !isStringInArray(request.BodyFormat, allowedBodyFormats) {
		err.addExplanation("body_format", "is not allowed")
	}

	if len(err.explanations) == 0 {
		isValid = true
	}

	return isValid, err
}

func isStringInArray(expectedValue string, array []string) bool {
	for _, value := range array {
		if value == expectedValue {
			return true
		}
	}
	return false
}
