package data

import (
	"fmt"
	"io"
	"os"

	api "plants/internal/pkg/pb/data/v1"
)

func (h *Handler) StreamExportDBV1(
	req *api.ExportDBV1Request,
	stream api.DataAPI_ExportDBV1Server,
) error {
	ctx := stream.Context()

	// Экспортируем базу данных в файл
	filePath := "exported_db.json"
	jsonData, err := h.storage.ExportDB(ctx)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка сохранения файла: %w", err)
	}

	// Открываем файл для стриминга
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	// Читаем файл по частям и отправляем
	buffer := make([]byte, 1024) // Размер буфера
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("ошибка чтения файла: %w", err)
		}

		// Отправляем часть данных клиенту
		if err := stream.Send(&api.ExportDBV1Response{
			Db: buffer[:n],
		}); err != nil {
			return fmt.Errorf("ошибка отправки данных: %w", err)
		}
	}

	return nil
}
