package helper

type SysOptionResp struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}
