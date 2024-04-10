package services

import (
	"7solutionstest3/internal/models"
	"log/slog"
	"strings"

	"golang.org/x/net/context"
)

type BeefAPICaller interface {
	FetchBeefData(ctx context.Context) (*string, error)
}

type BeefSummaryService struct {
	beefAPICaller BeefAPICaller
}

func NewBeefSummaryService(beefAPICaller BeefAPICaller) *BeefSummaryService {
	return &BeefSummaryService{beefAPICaller: beefAPICaller}
}

func (s *BeefSummaryService) BeefSummary(ctx context.Context) (*models.BeefResponse, error) {
	beefData, err := s.beefAPICaller.FetchBeefData(ctx)
	if err != nil {
		slog.Error(err.Error(), err)
		return nil, err
	}

	beefWords := sanitizeWords(*beefData)
	beefCount := countWords(beefWords)

	return &models.BeefResponse{Beef: beefCount}, nil
}

func sanitizeWords(text string) []string {
	cleanText := strings.ReplaceAll(text, ",", "")
	cleanText = strings.ReplaceAll(cleanText, ".", "")
	cleanText = strings.ReplaceAll(cleanText, "\n", "")

	arrInput := strings.Fields(cleanText)

	return arrInput
}

func countWords(words []string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[strings.ToLower(word)]++
	}
	return wordCount
}
