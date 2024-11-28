package util

import "errors"

func BaseUrl() (string, error) {
	env := LoadEnv()

	if env.Environment == "development" {
		return env.DevUrl, nil
	}

	if env.Environment == "production" {
		return env.ProdUrl, nil
	}

	if env.Environment == "staging" {
		return "youre in staging", nil // TODO: add staging url
	}

	return "", errors.New("environment not found")
}
