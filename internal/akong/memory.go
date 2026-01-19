package akong

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Memory struct {
	Date     string `json:"data"`
	Input    string `json:"input"`
	Response string `json:"response"`
}

func getFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".akong")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "memories.json"), nil
}

func SaveMemory(input, response string) error {
	newMem := Memory{
		Date:     time.Now().Format("2006-01-02 15:04:05"),
		Input:    input,
		Response: response,
	}
	filepath, err := getFilePath()
	if err != nil {
		return err
	}
	data, err := os.ReadFile(filepath)
	var memories []Memory
	if err == nil {
		json.Unmarshal(data, &memories)
	}

	memories = append(memories, newMem)

	updatedData, err := json.MarshalIndent(memories, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, updatedData, 0644)

}

func GetRandomMemory() (*Memory, error) {
	filepath, err := getFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filepath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var memories []Memory
	if err := json.Unmarshal(data, &memories); err != nil {
		return nil, err
	}
	if len(memories) == 0 {
		return nil, nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(memories))
	return &memories[index], nil
}
