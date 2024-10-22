package pdc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

const (
	PathLogin    = "/cmc/login_system.html"
	PathWorkInfo = "/cmc/getWorkInfo.html"
	Protocol     = "41"
)

type Session struct {
	JSessionId   string
	BaseUrl      string
	SerialNumber string
	Protocol     string
	WorkInfo     WorkInfo
}

type WorkInfo struct {
	SerialNo                   string  `json:"serialNo"`
	GridFrequency1             float64 `json:"gridFrequency"`
	GridFrequency2             float64 `json:"gridFrequency2"`
	GridVoltage1               float64 `json:"gridVoltage"`
	GridVoltage2               float64 `json:"gridVoltage2"`
	PvInputVoltage1            float64 `json:"pvInputVoltage1"`
	PvInputVoltage2            float64 `json:"pvInputVoltage2"`
	PvInputCurrent1            float64 `json:"pvInputCurrent1"`
	PvInputCurrent2            float64 `json:"pvInputCurrent2"`
	TotalPvInputPower          float64 `json:"totalPvInputPower"`
	AcOutputVoltage1           float64 `json:"acOutputVoltage"`
	AcOutputVoltage2           float64 `json:"acOutputVoltage2"`
	AcOutputFrequency1         float64 `json:"acOutputFrequency"`
	AcOutputFrequency2         float64 `json:"acOutputFrequency2"`
	AcOutputApparentPower1     float64 `json:"acOutputApparentPower"`
	AcOutputApparentPower2     float64 `json:"acOutputApparentPower2"`
	AcOutputActivePower1       float64 `json:"acOutputActivePower"`
	AcOutputActivePower2       float64 `json:"acOutputActivePower2"`
	OutputLoadPercent1         float64 `json:"outputLoadPercent"`
	OutputLoadPercent2         float64 `json:"outputLoadPercent2"`
	TotalOutputLoadPercent     float64 `json:"totalOutputLoadPercent"`
	BatVoltage                 float64 `json:"batteryVoltage"`
	BatCapacity                float64 `json:"batteryCapacity"`
	BatChgCurrent              float64 `json:"batteryChgCurrent"`
	TotalBatChgCurrent         float64 `json:"totalChargingCurrent"`
	BatDischgCurrent           float64 `json:"batteryDischgCurrent"`
	TotalAcOutputApparentPower float64 `json:"totalAcOutputApparentPower"`
	TotalAcOutputActivePower   float64 `json:"totalAcOutputActivePower"`
	ChargeSource               string  `json:"chargeSource"`
	LoadSource                 string  `json:"loadSource"`
	WorkMode                   string  `json:"workMode"`
	MachineType                string  `json:"machineType"`
	HasLoad1                   bool    `json:"hasLoad"`
	HasLoad2                   bool    `json:"hasLoad2"`
	ACchargeOn1                bool    `json:"ACchargeOn"`
	ACchargeOn2                bool    `json:"ACchargeOn2"`
	ChargeOn                   bool    `json:"chargeOn"`
	SCCchargeOn1               bool    `json:"SCCchargeOn"`
	SCCchargeOn2               bool    `json:"SCCchargeOn2"`
	LineLoss1                  bool    `json:"lineLoss"`
	LineLoss2                  bool    `json:"lineLoss2"`
	OverLoad                   bool    `json:"overLoad"`
	Timestr                    string  `json:"timestr"`
	DataID                     float64 `json:"dataID"`
	Time                       struct {
		Date           int   `json:"date"`
		Hours          int   `json:"hours"`
		Seconds        int   `json:"seconds"`
		Month          int   `json:"month"`
		TimezoneOffset int   `json:"timezoneOffset"`
		Year           int   `json:"year"`
		Minutes        int   `json:"minutes"`
		Time           int64 `json:"time"`
		Day            int   `json:"day"`
	} `json:"time"`
}

// Returns a new session.
func NewSession(baseUrl, serialNumber string) *Session {
	return &Session{
		BaseUrl:      baseUrl,
		SerialNumber: serialNumber,
	}
}

// Retrieves a JSESSIONID using the provided username and password
// and stores it in the session.
func (s *Session) Login(username, password string) error {
	data := url.Values{}
	data.Add("username", username)
	data.Add("password", password)

	res, err := postRequestForm(s.BaseUrl, PathLogin, "", data)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	for _, c := range res.Cookies() {
		if c.Name == "JSESSIONID" {
			s.JSessionId = c.Value
			return nil
		}
	}

	return ErrLoginFailed
}

// Retrieves the current work info for the given session.
func (s *Session) GetWorkInfo() error {
	path := fmt.Sprintf("%v?serialNo=%v&protocol=%v", PathWorkInfo, s.SerialNumber, Protocol)

	res, err := postRequest(s.BaseUrl, path, s.JSessionId)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &s.WorkInfo)
	if err != nil {
		return err
	}

	return nil
}
