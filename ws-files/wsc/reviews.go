package wsc

import (
	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/gen"
)

func GetReviews() (*axios.Response, error) {
	return client.Get("/reviews")
}

func GetReview() (*axios.Response, error) {
	return client.Get("/reviews/1")
}

func CreateReview() (*axios.Response, error) {
	return client.Post("/reviews", map[string]any{
		"plan_category": "subscription",
		"description":   gen.Sentence(),
		"rating":        gen.Int(5),
	})
}

func UpdateReview() (*axios.Response, error) {
	return client.Patch("/reviews/9", map[string]any{
		"plan_category": "design",
		"description":   gen.Sentence(),
		"rating":        gen.Int(5),
	})
}

func Delete() (*axios.Response, error) {
	return client.Delete("/reviews/10")
}
