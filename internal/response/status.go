package response

type StatusResponse struct {
	CurrentRound int      `json:"rodada_atual"`
	MarketStatus int      `json:"status_mercado"`
	Season       int      `json:"temporada"`
	Shutdown     Shutdown `json:"fechamento"`
	PostRound    bool     `json:"mercado_pos_rodada"`
	NewMonth     bool     `json:"novo_mes_ranking"`
	Running      bool     `json:"bola_rolando"`
}

type Shutdown struct {
	Day       int   `json:"dia"`
	Month     int   `json:"mes"`
	Year      int   `json:"ano"`
	Hour      int   `json:"hora"`
	Minute    int   `json:"minuto"`
	Timestamp int64 `json:"timestamp"` // Unix timestamp, precisa converter pra time.Time
}
