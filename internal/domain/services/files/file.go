package files

import (
	"encoding/json"
	"golang.org/x/net/context"
	"log/slog"
	"os"
	"project-auction/internal/domain/services"
	"strconv"
)

type Data struct {
	OwnerID  int
	ItemID   int
	ItemName string
	Points   float64
}

func CreateAndWrite(ctx context.Context, itmProvider services.ItemService, ownerID, itemID int, points float64, log *slog.Logger) (string, error) {
	item, err := itmProvider.GetItemByID(ctx, itemID)
	if err != nil {
		log.Error("error getting item", slog.String("error", err.Error()))
		return "", err
	}

	data := Data{
		OwnerID:  item.OwnerID,
		ItemID:   item.ID,
		ItemName: item.Name,
		Points:   points,
	}

	fileName := strconv.Itoa(ownerID) + "_bids_" + strconv.Itoa(itemID) + ".json"

	file, err := os.Create(fileName)
	if err != nil {
		log.Error("error creating file", slog.String("error", err.Error()))
		return "", err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Error("error marshalling to json", slog.String("error", err.Error()))
		return "", err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Error("error writing to file", slog.String("fileName", fileName), slog.String("error", err.Error()))
		return "", err
	}

	log.Info("JSON data successfully written", slog.String("filename", fileName))

	return fileName, nil
}
