package api

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Yustinia/gopaper"
	"github.com/Yustinia/kaiten-wall/internal/config"
)

var ErrNoWallpapers = errors.New("no wallpapers found")

func buildSearchParams(settings *config.ConfigModel) gopaper.SearchParams {
	w := settings.Wallhaven

	params := gopaper.NewSearch()
	params.AtLeast = w.AtLeast
	params.Categories = w.Categories
	params.KeySearch = w.Query
	params.Order = w.Order
	params.Purity = w.Purity
	params.Ratios = w.Ratios
	params.Resolution = w.Resolution
	params.Seed = w.Seed
	params.Sorting = w.Sorting

	return params
}

func FetchWallpapers(settings *config.ConfigModel) (gopaper.SearchResponse, error) {
	client := gopaper.NewClientWithKey(settings.ClientParams.APIKey)

	params := buildSearchParams(settings)
	result, err := client.Search(params)
	if err != nil {
		return gopaper.SearchResponse{}, fmt.Errorf("something went wrong: %w", err)
	}

	return result, nil
}

func SelectRandomWall(result gopaper.SearchResponse) (string, error) {
	wallCount := len(result.Wallpapers)
	if wallCount == 0 {
		return "", ErrNoWallpapers
	}

	randIndex := rand.Intn(wallCount)
	selectedWall := result.Wallpapers[randIndex]

	return selectedWall.Path, nil
}
