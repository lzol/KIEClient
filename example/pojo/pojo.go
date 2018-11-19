package pojo

type Address struct {
	City     string `json:"city"`
	Detail   string `json:"detail"`
	Distract string `json:"distract"`
	Province string `json:"province"`
}

type ApplyInfo struct {
	Age           int64       `json:"age"`
	FamilyAddress Address     `json:"familyAddress"`
	Name          interface{} `json:"name"`
}
