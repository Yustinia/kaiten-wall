package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yustinia/gopaper"
	"github.com/Yustinia/kaiten-wall/internal/api"
	"github.com/Yustinia/kaiten-wall/internal/config"
	"github.com/Yustinia/kaiten-wall/internal/daemon"
	"github.com/Yustinia/kaiten-wall/internal/defaults"
	"github.com/Yustinia/kaiten-wall/internal/download"
	"github.com/Yustinia/kaiten-wall/internal/theming"
	"github.com/spf13/cobra"
)

var (
	queryFlag      string
	categoriesFlag string
	purityFlag     string
	sortingFlag    string
	orderFlag      string
	topRangeFlag   string
	atLeastFlag    string
	resolutionFlag string
	ratioFlag      string
	seedFlag       string
	fetchFlag      string
)

var rootCmd = &cobra.Command{
	Use:   "kaiten",
	Short: "A simple random wallpaper switcher for wayland that fetches from wallhaven and applies it",
	Run: func(cdm *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("home not found: %v", err)
		}

		configDir := filepath.Join(homeDir, ".config", "kaiten-wall")
		configPath := filepath.Join(homeDir, ".config", "kaiten-wall", "config.toml")

		if err = defaults.FirstLaunch(configDir, configPath); err != nil {
			if errors.Is(err, defaults.ErrFirstLaunch) {
				log.Println(err)
				os.Exit(0)
			}
			log.Fatal(err)
		}

		settings, err := config.ParseSettings(configPath)
		if err != nil {
			log.Fatal(err)
		}

		applyFlagOverrides(&settings)

		start := time.Now()
		result, err := api.FetchWallpapers(&settings)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("fetched %d wallpapers in %s\n", len(result), time.Since(start).Round(time.Millisecond))

		scanner := bufio.NewScanner(os.Stdin)

	WallpaperLoop:
		for {
			err = wallpaperApplication(result, &settings)
			if err != nil {
				log.Fatalln(err)
			}

			for {
				log.Print("select another wallpaper? (Y/N)")

				if !scanner.Scan() {
					if err := scanner.Err(); err != nil {
						log.Fatalf("error reading input %v\n", err)
					}
					break WallpaperLoop
				}

				userInput := strings.TrimSpace(strings.ToLower(scanner.Text()))

				switch userInput {
				case "y", "yes":
					continue WallpaperLoop
				case "n", "no":
					break WallpaperLoop
				default:
					log.Printf("%q is invalid, please try again", userInput)
				}
			}
		}
	},
}

func wallpaperApplication(result []gopaper.Wallpaper, settings *config.ConfigModel) error {
	selectedWall, err := api.SelectRandomWall(result)
	if err != nil {
		return err
	}
	log.Printf("selected wallpaper: %s\n", selectedWall)

	start := time.Now()
	wallLocation, err := download.DownloadWall(selectedWall, settings.General.DefaultPath)
	if err != nil {
		return err
	}
	log.Printf("downloaded wallpaper to %s in %s\n", wallLocation, time.Since(start).Round(time.Millisecond))

	start = time.Now()
	switch settings.General.UseDaemon {
	case "awww":
		err = daemon.RunAwww(wallLocation, &settings.Awww)
	default:
		return fmt.Errorf("unknown daemon: %q", settings.General.UseDaemon)
	}
	if err != nil {
		return err
	}
	log.Printf("applied wallpaper using %s in %s\n", settings.General.UseDaemon, time.Since(start).Round(time.Millisecond))

	if settings.General.UseThemer != "" {
		start = time.Now()

		switch settings.General.UseThemer {
		case "matugen":
			err = theming.ApplyMatugen(wallLocation, &settings.Matugen)
		case "wallust":
			err = theming.ApplyWallust(wallLocation, &settings.Wallust)
		default:
			return fmt.Errorf("unknown themer: %q", settings.General.UseThemer)
		}
		if err != nil {
			return err
		}

		log.Printf("applied color schemes from %s in %s\n", settings.General.UseThemer, time.Since(start).Round(time.Millisecond))
	}

	return nil
}

func applyFlagOverrides(settings *config.ConfigModel) {
	if queryFlag != "" {
		settings.Wallhaven.Query = queryFlag
	}
	if categoriesFlag != "" {
		settings.Wallhaven.Categories = categoriesFlag
	}
	if purityFlag != "" {
		settings.Wallhaven.Purity = purityFlag
	}
	if sortingFlag != "" {
		settings.Wallhaven.Sorting = sortingFlag
	}
	if orderFlag != "" {
		settings.Wallhaven.Order = orderFlag
	}
	if topRangeFlag != "" {
		settings.Wallhaven.TopRange = topRangeFlag
	}
	if atLeastFlag != "" {
		settings.Wallhaven.AtLeast = atLeastFlag
	}
	if resolutionFlag != "" {
		settings.Wallhaven.Resolution = resolutionFlag
	}
	if ratioFlag != "" {
		settings.Wallhaven.Ratios = ratioFlag
	}
	if seedFlag != "" {
		settings.Wallhaven.Seed = seedFlag
	}
	if fetchFlag != "" {
		settings.Wallhaven.Fetch = fetchFlag
	}
}

func init() {
	rootCmd.Flags().StringVar(&queryFlag, "query", "", "filter by search terms")
	rootCmd.Flags().StringVar(&categoriesFlag, "categories", "", "categorize wallpapers: General, Anime, People (e.g. 100, 110, 111)")
	rootCmd.Flags().StringVar(&purityFlag, "purity", "", "set content rating: SFW, Sketchy, NSFW (e.g. 100, 110, 111)")
	rootCmd.Flags().StringVar(&sortingFlag, "sorting", "", "sort wallpaper fetches (e.g. date_added, relevance, random, views, favorites, toplist)")
	rootCmd.Flags().StringVar(&orderFlag, "order", "", "set order (e.g. asc, desc)")
	rootCmd.Flags().StringVar(&topRangeFlag, "toprange", "", "filter popular wallpapers (e.g. 1d, 3d, 1w, 1M, 3M, 6M, 1y)")
	rootCmd.Flags().StringVar(&atLeastFlag, "atleast", "", "set minimum resolution for wallpapers: accepts one resolution (e.g. 1920x1080)")
	rootCmd.Flags().StringVar(&resolutionFlag, "resolution", "", "set exact wallpaper resolutions: accepts more than one value separated by comma (e.g. 1920x1080,2560x1440...)")
	rootCmd.Flags().StringVar(&ratioFlag, "ratio", "", "set exact wallpaper ratios: accepts more than one value separated by comma (e.g. 16x9,21x9 or landscape)")
	rootCmd.Flags().StringVar(&seedFlag, "seed", "", "set seed to randomize wallpaper fetches")
	rootCmd.Flags().StringVar(&fetchFlag, "fetch", "", "set fetch mode for wallpapers: accepts a key pair value separated by colon (e.g. count:24)")
}
