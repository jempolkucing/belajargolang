package services

import (
	"belajargolang/models"
	"belajargolang/repositories"
	"time"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReport(startDateStr, endDateStr string) (*models.DailyReport, error) {
	var start, end time.Time
	var err error

	layout := "2006-01-02"
	now := time.Now()

	if startDateStr == "" {
		startDateStr = now.Format(layout)
	}

	start, err = time.Parse(layout, startDateStr)
	if err != nil {
		return nil, err
	}

	if endDateStr == "" {
		end = start
	} else {
		end, err = time.Parse(layout, endDateStr)
		if err != nil {
			return nil, err
		}
	}

	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local)
	end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999999999, time.Local)

	return s.repo.GetReport(start, end)
}
