package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	// Version can be set at build time with: go build -ldflags "-X github.com/benfaerber/stripey-horse/app.Version=1.0.0"
	Version = "dev"
)

type LabelConfig struct {
	LabelWidthMm  float64 `json:"labelWidthMm"`
	LabelHeightMm float64 `json:"labelHeightMm"`
	Dpmm          int     `json:"dpmm"`
	Rotation      int     `json:"rotation"`
}

func printHelpMenu() {
	fmt.Println("stripey-horse: a knock-off zebra renderer")
	fmt.Println("https://github.com/benfaerber/stripey-horse")
	fmt.Println("")

	pflag.Usage()
}

type Signature struct {
	App       string `json:"app"`
	Version   string `json:"version"`
	Signature string `json:"signature"`
}

func generateSignature() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("getting executable path: %w", err)
	}

	data, err := os.ReadFile(execPath)
	if err != nil {
		return "", fmt.Errorf("reading executable: %w", err)
	}

	hash := sha256.Sum256(data)
	hashStr := hex.EncodeToString(hash[:])

	sig := Signature{
		App:       "stripey-horse",
		Version:   Version,
		Signature: hashStr,
	}

	jsonBytes, err := json.Marshal(sig)
	if err != nil {
		return "", fmt.Errorf("marshaling signature: %w", err)
	}

	return string(jsonBytes), nil
}

func parseConfig() (LabelConfig, string, error) {
	configJSON := pflag.StringP("config", "c", "", "JSON configuration for label dimensions")
	outputFile := pflag.StringP("output", "o", "", "Output file path (defaults to stdout)")
	showSignature := pflag.Bool("signature", false, "Output binary signature for verification")
	pflag.Parse()

	if *showSignature {
		sig, err := generateSignature()
		if err != nil {
			return LabelConfig{}, "", fmt.Errorf("generating signature: %w", err)
		}
		fmt.Println(sig)
		os.Exit(0)
	}

	if *configJSON == "" {
		printHelpMenu()
		return LabelConfig{}, "", fmt.Errorf("config is required")
	}

	var config LabelConfig
	if err := json.Unmarshal([]byte(*configJSON), &config); err != nil {
		return LabelConfig{}, "", fmt.Errorf("parsing config JSON: %w", err)
	}

	return config, *outputFile, nil
}
