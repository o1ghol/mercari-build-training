package app

import (
	"context"
	"errors"
	"encoding/json"
	"fmt"
	"os"
	// STEP 5-1: uncomment this line
	// _ "github.com/mattn/go-sqlite3"
)

var errImageNotFound = errors.New("image not found")

type Item struct {
	ID       int    `db:"id" json:"-"`
	Name     string `db:"name" json:"name"`
	Category string `db:"category" json:"category"`
}

// Please run `go generate ./...` to generate the mock implementation
// ItemRepository is an interface to manage items.
//
//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -package=${GOPACKAGE} -destination=./mock_$GOFILE
type ItemRepository interface {
	Insert(ctx context.Context, item *Item) error
}

// itemRepository is an implementation of ItemRepository
type itemRepository struct {
	// fileName is the path to the JSON file storing items.
	fileName string
}

// NewItemRepository creates a new itemRepository.
func NewItemRepository() ItemRepository {
	return &itemRepository{fileName: "items.json"}
}

// Insert inserts an item into the repository.
func (i *itemRepository) Insert(ctx context.Context, item *Item) error {
	// 既存のアイテムを読み込む
	existingItems := []Item{}
	data, err := os.ReadFile(i.fileName)
	if err == nil {
		if err := json.Unmarshal(data, &existingItems); err != nil {
			var singleItem Item
			if err := json.Unmarshal(data, &singleItem); err == nil {
				existingItems = append(existingItems, singleItem) // JSONが単一オブジェクトだった場合の対応
			} else {
				return fmt.Errorf("failed to parse JSON: %w", err)
			}
		}
	}

	// アイテムをリストに追加
	existingItems = append(existingItems, *item)

	// JSON に変換
	data, err = json.MarshalIndent(existingItems, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// `items.json` に書き込む
	err = os.WriteFile(i.fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}


// StoreImage stores an image and returns an error if any.
// This package doesn't have a related interface for simplicity.
func StoreImage(fileName string, image []byte) error {
	// STEP 4-4: add an implementation to store an image

	return nil
}