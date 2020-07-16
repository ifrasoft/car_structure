package car_structure

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Axis struct {
	AxisID      int64  `json:"axisID"`
	Left        []Tire `json:"left"`
	Right       []Tire `json:"right"`
	SpareWheel  []Tire `json:"spareWheels"`
	ImageBase64 string `json:"image"`
}

type Wheel struct {
	Tires []Tire `json:"tires"`
}

type PolicyStatus struct {
	TireDepth struct {
		Status string `json:"status"`
	} `json:"tireDepth"`
	PSI struct {
		Status string `json:"status"`
	} `json:"psi"`
}

type Tire struct {
	IsEmpty          bool         `json:"isEmpty"`
	TireID           int64        `json:"tireID"`
	Position         string       `json:"position"`
	TireSerialNumber string       `json:"tireSerialNumber"`
	TireDepth        float64      `json:"tireDepth"`
	TirePressure     float64      `json:"tirePressure"`
	Turnable         bool         `json:"turnable"`
	PolicyStatus     PolicyStatus `json:"policyStatus"`
}

type Summary struct {
	AxisQTY  int    `json:"axisQTY"`
	WheelQTY int    `json:"wheelQTY"`
	Axles    []Axis `json:"axis"`
}

type carStructure struct {
	TextIntput       string
	TireInformations []*TireInformation
	Policies         []*Policy
	HeaderImage      image.Image
	BodyImage        image.Image
	FooterImage      image.Image
	CarType          string
}

type TireInformation struct {
	TireID              int64   `json:"tireID"`
	PositionCode        string  `json:"position"`
	TireSerialNumber    string  `json:"tireSerialNumber"`
	TireDepthMaximum    float64 `json:"tireDepthMaximum"`
	TireDepthMinimum    float64 `json:"tireDepthMinimum"`
	TirePressureMaximum float64 `json:"tirePressureMaximum"`
	TirePressureMinimum float64 `json:"tirePressureMinimum"`
}

func NewCarStructureConvertor(input, carType string, tireInformations []*TireInformation) *carStructure {
	return &carStructure{
		TextIntput:       input,
		TireInformations: tireInformations,
		CarType:          carType,
	}
}

func (cs *carStructure) ApplyPolicies(policies []*Policy) {
	cs.Policies = policies
}

func (cs *carStructure) InjectImageCarType(HeaderCar, BodyCar, FooterCar image.Image) {
	cs.HeaderImage = HeaderCar
	cs.BodyImage = BodyCar
	cs.FooterImage = FooterCar
}

func (sm *Summary) Sort() {

	for index, axle := range sm.Axles {
		sort.SliceStable(axle.Left, func(i, j int) bool {
			_, _, WheelNoi := extectCode(axle.Left[i].Position)
			_, _, WheelNoj := extectCode(axle.Left[j].Position)
			return WheelNoi < WheelNoj
		})

		sort.SliceStable(axle.Right, func(i, j int) bool {
			_, _, WheelNoi := extectCode(axle.Right[i].Position)
			_, _, WheelNoj := extectCode(axle.Right[j].Position)
			return WheelNoi > WheelNoj
		})

		sm.Axles[index].Left = axle.Left
		sm.Axles[index].Right = axle.Right
	}

}

func (cs *carStructure) Turnable(axisNo int) bool {
	axisLists := strings.Split(cs.TextIntput, "-")
	for no, axis := range axisLists {
		if axisNo == no+1 {
			str := axis[1:2]
			if str == "S" {
				return true
			}
		}

	}
	return false
}

func (cs *carStructure) GetAxisQTY() int {
	return len(strings.Split(cs.TextIntput, "-"))
}

func (cs *carStructure) GetWheelQTY() int {

	count := 0
	axisLists := strings.Split(cs.TextIntput, "-")
	for _, axis := range axisLists {
		numberStr := axis[0:1]
		numberInt, _ := strconv.Atoi(numberStr)
		count = count + numberInt

	}

	return count
}

func extectCode(positionCode string) (side string, axisNo, WheelNo int) {

	data := strings.Split(positionCode, "-")
	axisNo, _ = strconv.Atoi(data[0])
	WheelNo, _ = strconv.Atoi(data[1][1:2])

	side = "L"
	if strings.Contains(data[1], "R") {
		side = "R"
	}
	return
}

func (cs *carStructure) GetJsonResult() (Summary, error) {

	var summary Summary
	summary.AxisQTY = cs.GetAxisQTY()
	summary.WheelQTY = cs.GetWheelQTY()

	codes := strings.Split(cs.TextIntput, "-")

	for i := summary.AxisQTY; i >= 0; i-- {

		//Add initial struct
		summary.Axles = append(summary.Axles, Axis{
			AxisID: int64(i), ImageBase64: cs.GenerateImage(i, summary.AxisQTY),
		})

		///เพิ่ม default ล้อ กรณียังไม่มี
		if i != 0 {

			re := regexp.MustCompile("[0-9]+")
			wheelQTY, _ := strconv.Atoi(re.FindAllString(codes[i-1], -1)[0])

			for a := 0; a < len(summary.Axles); a++ {
				if summary.Axles[a].AxisID == int64(i) {
					for x := 1; x <= wheelQTY/2; x++ {
						var tire = Tire{
							IsEmpty:  true,
							Position: fmt.Sprintf("%d-%s%d", i, "R", x),
						}
						summary.Axles[a].Right = append(summary.Axles[a].Right, tire)
					}
					break
				}
			}
			for a := 0; a < len(summary.Axles); a++ {
				if summary.Axles[a].AxisID == int64(i) {
					for x := 1; x <= wheelQTY/2; x++ {
						var tire = Tire{
							IsEmpty:  true,
							Position: fmt.Sprintf("%d-%s%d", i, "L", x),
						}
						summary.Axles[a].Left = append(summary.Axles[a].Left, tire)
					}
					break
				}
			}
		}

		for _, tireInformation := range cs.TireInformations {

			side, axisNo, _ := extectCode(tireInformation.PositionCode)

			var policyStatus PolicyStatus

			var tire = Tire{
				TireID:           tireInformation.TireID,
				Position:         tireInformation.PositionCode,
				TireSerialNumber: tireInformation.TireSerialNumber,
				TireDepth:        tireInformation.TireDepthMinimum,
				TirePressure:     tireInformation.TirePressureMinimum,
				Turnable:         cs.Turnable(axisNo),
				PolicyStatus:     policyStatus,
			}

			if axisNo == 0 { //ยางอะไหล่
				for a := 0; a < len(summary.Axles); a++ {
					if summary.Axles[a].AxisID == 0 {
						summary.Axles[a].SpareWheel = append(summary.Axles[a].SpareWheel, tire)
						break
					}
				}
			}

			if axisNo == i && i != 0 { //Exclude axis 0

				//เพิ่ม Policy
				for _, policy := range cs.Policies {
					if policy.AxlesNo == axisNo {
						if tire.TireDepth <= policy.CriticalTireDepth {
							tire.PolicyStatus.TireDepth.Status = "critical"
						} else if tire.TireDepth <= policy.WarningTireDepth {
							tire.PolicyStatus.TireDepth.Status = "warning"
						} else {
							tire.PolicyStatus.TireDepth.Status = "good"
						}

						diff := Abs(tire.TirePressure-policy.StandardTirePressure) / policy.StandardTirePressure * 100

						if diff >= policy.CriticalTirePressure {
							tire.PolicyStatus.PSI.Status = "critical"
						} else if diff >= policy.WarningTirePressure {
							tire.PolicyStatus.PSI.Status = "warning"
						} else {
							tire.PolicyStatus.PSI.Status = "good"
						}
					}
				}

				if side == "R" { //ด้านขวา
					for a := 0; a < len(summary.Axles); a++ {
						if summary.Axles[a].AxisID == int64(i) {
							for index, r := range summary.Axles[a].Right {
								if r.Position == tire.Position {
									summary.Axles[a].Right[index] = tire
									break
								}

							}
							break
						}
					}
				} else { //ด้านซ้าย
					for a := 0; a < len(summary.Axles); a++ {
						if summary.Axles[a].AxisID == int64(i) {
							for index, l := range summary.Axles[a].Left {
								if l.Position == tire.Position {
									summary.Axles[a].Left[index] = tire
									break
								}

							}
							break
						}
					}
				}

			}

		}

	}

	//เรียงข้อมูล
	summary.Sort()

	//แปลง struct to json format
	b, _ := json.Marshal(summary)

	fmt.Println(string(b))

	return summary, nil
}

func (cs *carStructure) GenerateImage(axisQTY, numAxis int) string {

	if axisQTY != 0 {
		var croppedImg image.Image
		if axisQTY == numAxis {
			croppedImg = cs.FooterImage
		} else if axisQTY == 1 {
			croppedImg = cs.HeaderImage
		} else {
			croppedImg = cs.BodyImage
		}

		buf := new(bytes.Buffer)
		err := png.Encode(buf, croppedImg)
		if err != nil {
			fmt.Println(err)
		}
		imageBit := buf.Bytes()
		return base64.StdEncoding.EncodeToString([]byte(imageBit))
	}

	return ""

}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
