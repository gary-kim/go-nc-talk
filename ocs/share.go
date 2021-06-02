package ocs

const (
	// ShareTypeRoom is OC.share.SHARE_TYPE_ROOM
	ShareTypeRoom = "10"
)

// ShareReturn is the response for a file share request
type ShareReturn struct {
	OCS struct {
		ocs
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"ocs"`
}
