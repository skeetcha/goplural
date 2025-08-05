package main

import (
	"bytes"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))

	for i, t := range ts {
		result[i] = fn(t)
	}

	return result
}

func Filter[T any](ts []T, fn func(T) bool) []T {
	result := make([]T, 0)

	for _, t := range ts {
		if fn(t) {
			result = append(result, t)
		}
	}

	return result
}

func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	fdata, err := f.Stat()

	if err != nil {
		return nil, err
	}

	data := make([]byte, fdata.Size())

	defer f.Close()
	_, err = f.Read(data)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	img, _, err := image.Decode(reader)

	if err != nil {
		return nil, err
	}

	return img, nil
}
