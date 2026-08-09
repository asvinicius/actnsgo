package marketstatus

import (
	"time"

	"github.com/asvinicius/actnsgo/internal/client"
	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
)

type MarketStatusService struct {
	marketStatusRepository *repository.MarketStatusRepository
	statusClient           *client.Client
}

func NewMarketStatusService(marketStatusRepository *repository.MarketStatusRepository, statusClient *client.Client) *MarketStatusService {
	return &MarketStatusService{
		marketStatusRepository: marketStatusRepository,
		statusClient:           statusClient,
	}
}

func (mss *MarketStatusService) GetStatus() (*model.MarketStatus, error) {
	return mss.marketStatusRepository.GetStatus()
}

func (mss *MarketStatusService) UpdateStatus(ms model.MarketStatus) error {
	return mss.marketStatusRepository.UpdateStatus(ms)
}

func (mss *MarketStatusService) HasMarketClosed() (bool, error) {

	marketstatus, err := mss.GetStatus()

	if err != nil {
		return false, err
	}

	apistatus, err := mss.statusClient.GetStatus()

	if err != nil {
		return false, err
	}

	if marketstatus.MsCurrentRound != apistatus.CurrentRound {
		return false, nil
	}

	if marketstatus.MsStatus == apistatus.MarketStatus {
		return false, nil
	}

	return true, nil
}

func (mss *MarketStatusService) SyncStatus() error {

	current, err := mss.GetStatus()

	if err != nil {
		return err
	}

	apistatus, err := mss.statusClient.GetStatus()

	if err != nil {
		return err
	}

	updated := model.MarketStatus{
		MsID:                current.MsID,
		MsStatus:            apistatus.MarketStatus,
		MsCurrentRound:      apistatus.CurrentRound,
		MsCurrentMonth:      current.MsCurrentMonth,
		MsCurrentSeason:     apistatus.Season,
		MsShutdownDay:       apistatus.Shutdown.Day,
		MsShutdownMonth:     apistatus.Shutdown.Month,
		MsShutdownYear:      apistatus.Shutdown.Year,
		MsShutdownHour:      apistatus.Shutdown.Hour,
		MsShutdownMinute:    apistatus.Shutdown.Minute,
		MsShutdownTimestamp: time.Unix(apistatus.Shutdown.Timestamp, 0),
		MsPostRound:         apistatus.PostRound,
		MsNewMonth:          apistatus.NewMonth,
		MsRunning:           apistatus.Running,
	}

	return mss.UpdateStatus(updated)
}
