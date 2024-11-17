package service

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/vertexai/genai"

	"ex.com/basicws/internal/constants"
)

type VertexAIService struct {
	GenAIClient *genai.Client
}

func InitializeVertexAIService() *VertexAIService {
	client, err := genai.NewClient(context.Background(), constants.PROJECT_ID, constants.LOCATION)
	if err != nil {
		panic("failed to initialize Vertex AI client")
	}
	return &VertexAIService{
		GenAIClient: client,
	}
}

func (s *VertexAIService) GenerateText(ctx context.Context) (string, error) {
	geminiClient := s.GenAIClient.GenerativeModel(constants.MODEL_NAME)
	prompt := genai.Text("Generate random poem minimum of 1000 words")

	resp, err := geminiClient.GenerateContent(ctx, prompt)
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("no content is returned")
	}

	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}
