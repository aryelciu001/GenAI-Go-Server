package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/vertexai/genai"

	"ex.com/basicws/internal/config"
)

type VertexAIService struct {
	GenAIClient *genai.Client
	ModelName   string
}

func MustInitializeVertexAIService(cfg *config.Config) *VertexAIService {
	client, err := genai.NewClient(context.Background(), cfg.ProjectID, cfg.Location)
	if err != nil {
		panic("failed to initialize Vertex AI client")
	}
	return &VertexAIService{
		GenAIClient: client,
		ModelName:   cfg.ModelName,
	}
}

func (s *VertexAIService) GenerateText(ctx context.Context, prompt string) (string, error) {
	go func() {
		log.Println("Gemini Call")
	}()

	geminiClient := s.GenAIClient.GenerativeModel(s.ModelName)
	_prompt := genai.Text(prompt)

	resp, err := geminiClient.GenerateContent(ctx, _prompt)
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("no content is returned")
	}

	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}

func (s *VertexAIService) PickOneRandomPoem(ctx context.Context) (string, error) {
	wg := sync.WaitGroup{}
	chanSize := 2
	resChan := make(chan string, chanSize)

	for i := 0; i < chanSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := s.GenerateText(ctx, "Generate a poem with less than 20 words")

			if err != nil {
				resChan <- ""
			} else {
				resChan <- res
			}
		}()
	}

	wg.Wait()
	close(resChan)

	var generatedPoems string
	count := 1
	for res := range resChan {
		generatedPoems = generatedPoems + fmt.Sprintf("Poem %d: %v\n", count, res)
	}

	prompt := fmt.Sprintf(
		`Given a list of prompt: %v
		Pick one poem that is the best.`, generatedPoems)

	res, err := s.GenerateText(ctx, prompt)
	return res, err
}
