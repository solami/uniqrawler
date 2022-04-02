package sheet

import (
	"context"
	"log"

	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SpreadSheets struct {
	sheetService *sheets.Service
	driveService *drive.Service
}

func NewSpreadSheets(credentialPath string) (*SpreadSheets, error) {
	ctx := context.Background()
	sheetService, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialPath))
	if err != nil {
		log.Fatalf("Unable to create sheet service: %v", err)
	}
	driveSerice, err := drive.NewService(ctx, option.WithCredentialsFile(credentialPath))
	if err != nil {
		log.Fatalf("Unable to create drive service: %v", err)
	}
	return &SpreadSheets{
		sheetService: sheetService,
		driveService: driveSerice,
	}, nil
}

func (s *SpreadSheets) Create(filename string, folderId string) (string, error) {
	f := &drive.File{
		Title:    filename,
		MimeType: "application/vnd.google-apps.spreadsheet",
		Parents: []*drive.ParentReference{
			{Id: folderId},
		},
	}
	file, err := s.driveService.Files.Insert(f).Do()
	if err != nil {
		return "", err
	}
	return file.Id, nil
}

func (s *SpreadSheets) Append(id string, values [][]interface{}) error {
	_, err := s.sheetService.Spreadsheets.Values.Append(id, "Sheet1", &sheets.ValueRange{
		Values: values,
	}).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		return err
	}
	return nil
}
