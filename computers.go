package jamf

type Computer struct {
	ID           int    `json:"id,omitempty" xml:"id,omitempty"`
	Name         string `json:"name,omitempty" xml:"name,omitempty"`
	SerialNumber string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
}
