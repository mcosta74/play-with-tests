package utils

func Hello(name string) string {
	if name == "" {
		return "Who are you?"
	}
	return "Hello " + name
}
