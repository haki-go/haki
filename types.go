package haki

type HakiError struct {
	Message    string
	StatusCode int
}

type HakiMap map[string]any
type HakiArray []any
