package marketstatus

import (
	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
)

type MarketStatusService struct {
	marketStatusRepository *repository.MarketStatusRepository
}

func NewMarketStatusService(marketStatusRepository *repository.MarketStatusRepository) *MarketStatusService {
	return &MarketStatusService{
		marketStatusRepository: marketStatusRepository,
	}
}

func (mss *MarketStatusService) GetStatus() (*model.MarketStatus, error) {
	return mss.marketStatusRepository.GetStatus()
}

func (mss *MarketStatusService) UpdateStatus(ms model.MarketStatus) error {
	return mss.marketStatusRepository.UpdateStatus(ms)
}
