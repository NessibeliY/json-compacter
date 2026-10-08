package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const (
	inputPath  = "input.json"
	outputPath = "output.json"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}
}

func run() error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	compacted, err := compact(data)
	if err != nil {
		return fmt.Errorf("%s: %w", inputPath, err)
	}

	return os.WriteFile(outputPath, compacted, 0o644)
}

// compact убирает из JSON все незначащие пробелы, сохраняя порядок ключей
// и исходную запись чисел и строк.
func compact(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := json.Compact(&buf, data); err != nil {
		return nil, fmt.Errorf("некорректный JSON: %w", err)
	}
	return buf.Bytes(), nil
}