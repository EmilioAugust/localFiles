package repository

type File struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Size int64 `json:"size"`
}

type IPAddress struct {
	IPAddr string `json:"ip"`
}

type Device struct {
	DeviceID string `json:"device_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Online bool `json:"online"`
}

type FindDevice struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type FileToken struct {
	Name string `json:"name"`
	Token int `json:"token"`
}