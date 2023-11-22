package mapping

type BootNotificationResponse struct {
	CurrentTime string `json:"currentTime"`
	Interval    int    `json:"interval"`
	Status      string `json:"status"`
}
