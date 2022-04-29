package model

type Session struct {
	Socket              string // 终端套接字
	AccessType          string
	PreviousSatelliteId string
	SessionKey          []byte
	ExpirationDate      int64
	StartAt             int64
}

type SessionRecord struct {
	MacAddr        string `json:"macAddr"`
	SessionKey     string `json:"sessionKey"`
	ExpirationDate int64  `json:"expirationDate"`
}
