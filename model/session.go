package model

type Session struct {
	SessionKey     []byte
	ExpirationDate int64
	StartAt        int64
}

type SessionRecord struct {
	MacAddr        string `json:"macAddr"`
	SessionKey     string `json:"sessionKey"`
	ExpirationDate int64  `json:"expirationDate"`
}
