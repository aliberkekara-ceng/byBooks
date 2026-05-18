package models

type URLRequest struct {
	URL       string `json:"url" binding:"required"`
	Operation string `json:"operation" binding:"required,oneof=redirection canonical all"`
}

type URLResponse struct {
	ProcessedURL string `json:"processed_url"`
}
