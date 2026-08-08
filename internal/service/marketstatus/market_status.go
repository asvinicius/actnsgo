package marketstatus

import (
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
