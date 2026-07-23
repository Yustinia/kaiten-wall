package api

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/Yustinia/gopaper"
	"github.com/Yustinia/kaiten-wall/internal/config"
)

type fetchMode struct {
	mode  string
	start int
	end   int
}

var ErrNoWallpapers = errors.New("no wallpapers found")

var ErrModeFetch = errors.New("failed to parse fetch modes")
var ErrFetchParts = errors.New("invalid value for fetch")

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
	params.Sorting = w.Sorting

	switch params.Sorting {
	case "toplist":
		params.TopRange = w.TopRange
	case "random":
		params.Seed = w.Seed
	}

	return params
}

func parseFetchMode(fetch string) (fetchMode, error) {
	parts := strings.SplitN(fetch, ":", 2)

	if len(parts) != 2 {
		return fetchMode{}, ErrFetchParts
	}

	mode := parts[0]
	value := parts[1]

	switch mode {
	case "page":
		startPage, err := strconv.Atoi(value)
		if err != nil {
			return fetchMode{}, fmt.Errorf("failed to parse page: %w", err)
		}
		return fetchMode{mode: mode, start: startPage}, nil

	case "pages":
		page := strings.SplitN(value, "-", 2)
		startPage, err := strconv.Atoi(page[0])
		if err != nil {
			return fetchMode{}, fmt.Errorf("failed to parse page: %w", err)
		}
		endPage, err := strconv.Atoi(page[1])
		if err != nil {
			return fetchMode{}, fmt.Errorf("failed to parse page: %w", err)
		}
		return fetchMode{mode: mode, start: startPage, end: endPage}, nil

	case "count":
		wallCount, err := strconv.Atoi(value)
		if err != nil {
			return fetchMode{}, fmt.Errorf("failed to parse wall count: %w", err)
		}
		return fetchMode{mode: mode, start: wallCount}, nil

	default:
		return fetchMode{}, ErrModeFetch
	}
}

func FetchWallpapers(settings *config.ConfigModel) ([]gopaper.Wallpaper, error) {
	client := gopaper.NewClient(settings.ClientParams.APIKey)

	params := buildSearchParams(settings)

	if settings.Wallhaven.Fetch == "" {
		result, err := client.Search(params)
		if err != nil {
			return nil, err
		}

		return result.Wallpapers, nil
	}

	modeFetch, err := parseFetchMode(settings.Wallhaven.Fetch)
	if err != nil {
		return nil, err
	}

	switch modeFetch.mode {
	case "page":
		result, err := client.FetchPage(&params, modeFetch.start)

		if err != nil {
			return nil, err
		}
		return result, nil

	case "pages":
		result, err := client.FetchPages(&params, modeFetch.start, modeFetch.end)

		if err != nil {
			return nil, err
		}
		return result, nil

	case "count":
		result, err := client.FetchWallpaperCount(&params, modeFetch.start)

		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, ErrModeFetch
	}
}

func SelectRandomWall(wallSlice []gopaper.Wallpaper) (string, error) {
	wallCount := len(wallSlice)
	if wallCount == 0 {
		return "", ErrNoWallpapers
	}

	randIndex := rand.Intn(wallCount)
	selectedWall := wallSlice[randIndex]

	return selectedWall.Path, nil
}
