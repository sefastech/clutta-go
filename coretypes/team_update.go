package coretypes

type TeamUpdate struct {
	UUID     string `json:"uuid"`
	TeamName string `json:"teamName"`
	TeamUUID string `json:"teamUuid"`
	OwnerID  string `json:"ownerId"`
}
