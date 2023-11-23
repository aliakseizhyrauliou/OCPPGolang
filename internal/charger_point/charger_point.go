package charger_point

import "time"

type ChargerPoint struct {
	ID            int       `json:"id"`
	WebID         string    `json:"web_id"`
	Name          string    `json:"name"`
	LastHeartbeat time.Time `json:"last_heartbeat"`
	Location      string    `json:"location"`
	OCPPVersion   string    `json:"ocpp_version"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsDeleted     bool      `json:"is_deleted"`
}
