package utils

import "math/rand"

// Generate 8 char random slug

func RandomSlug() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	slug := make([]byte, 8)
	for i := range slug {
		slug[i] = charset[rand.Intn(len(charset))]
	}
	return string(slug)
}
