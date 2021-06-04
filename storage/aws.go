package storage

type AWS struct {
	AccessKeyID     string `json:"-"`
	SecretAccessKey string `json:"-"`
	SessionToken    string `json:"-"`
	Region          string `json:"region,omitempty"`
}
