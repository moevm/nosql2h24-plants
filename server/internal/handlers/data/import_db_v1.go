package data

import (
	"fmt"
	"io"
	"os"

	api "plants/internal/pkg/pb/data/v1"
)

func (h *Handler) StreamImportDBV1(
	stream api.DataAPI_ImportDBV1Server,
) error {
	ctx := stream.Context()

	// Создаём временный файл для записи полученных данных
	filePath := "imported_db.json"
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer file.Close()

	// Читаем части данных из стрима
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("ошибка получения данных: %w", err)
		}

		// Записываем данные в файл
		if _, err := file.Write(req.Db); err != nil {
			return fmt.Errorf("ошибка записи в файл: %w", err)
		}
	}

	// Читаем данные из файла и импортируем в базу данных
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	if err := h.storage.ImportDB(ctx, jsonData); err != nil {
		return fmt.Errorf("ошибка импорта данных: %w", err)
	}

	return nil
}
