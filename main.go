package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type SalahTimeRecord struct {
	Month     int         `json:"month"`
	Year      int         `json:"year"`
	Salatimes []Salatimes `json:"salatimes"`
}
type SalahTime struct {
	Hour int `json:"hour"`
	Min  int `json:"min"`
}
type Salatimes struct {
	Month   int
	Day     int
	Fajar18 int `json:"fajar18"`
	Fajar15 int `json:"fajar15"`
	Sunrise int `json:"sunrise"`
	//Ishraq    SalahTime `json:"ishraq"`
	//Chasht    SalahTime `json:"chasht"`
	//Zawal     SalahTime `json:"zawal"`
	Zuhar     int `json:"zuhar"`
	AsrShafai int `json:"asarshafai"`
	AsrHanfi  int `json:"asrhanfi"`
	Maghrib   int `json:"maghrib"`
	Isha15    int `json:"isha15"`
	Isha18    int `json:"isha18"`
}

type Salatimes2 struct {
	Fajar18 int `json:"fajar18"`
	Fajar15 int `json:"fajar15"`
	Sunrise int `json:"sunrise"`
	//Ishraq    SalahTime `json:"ishraq"`
	//Chasht    SalahTime `json:"chasht"`
	//Zawal     SalahTime `json:"zawal"`
	Zuhar     int `json:"zuhar"`
	AsrShafai int `json:"asarshafai"`
	AsrHanfi  int `json:"asrhanfi"`
	Maghrib   int `json:"maghrib"`
	Isha15    int `json:"isha15"`
	Isha18    int `json:"isha18"`
}

type SalahTimeYear struct {
	Year      int            `json:"year"`
	Salatimes [12][31][9]int `json:"salatimes"` //[month][day][salah]
}

func convertStrToMin(str string) int {
	parts := strings.Split(strings.Split(str, " ")[0], ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	return (h * 60) + m
}

func convertStrToSalaTime(str string) SalahTime {
	parts := strings.Split(strings.Split(str, " ")[0], ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	return SalahTime{Hour: h, Min: m}
}

type Multipart struct {
	Multipart string
}

type DailyIqamahTime struct {
	Month     int
	Day       int
	DayOfWeek string `json:",omitempty"`
	SalahName string
	Times     []int
	Relative  bool `json:",omitempty"`
}

const (
	fajar   = "fajar"
	zuhar   = "zuhar"
	asar    = "asar"
	maghrib = "maghrib"
	isha    = "isha"
	juma    = "juma"
	eid     = "eid"
	special = "special"
)

func tomin(h int, m int) int {
	return (h * 60) + m
}

func tomina(h int, m int) []int {
	return []int{(h * 60) + m}
}

func writeMasjidZakriyaTimes() {
	times := []DailyIqamahTime{
		{Month: 1, Day: 1, SalahName: fajar, Times: tomina(6, 30)},
		{Month: 1, Day: 1, SalahName: zuhar, Times: tomina(13, 0)},
		{Month: 1, Day: 1, SalahName: asar, Times: tomina(16, 0)},
		{Month: 1, Day: 1, SalahName: maghrib, Times: []int{2}, Relative: true},
		{Month: 1, Day: 1, SalahName: isha, Times: tomina(19, 30)},

		{Month: 2, Day: 1, SalahName: asar, Times: tomina(16, 15)},

		{Month: 2, Day: 11, SalahName: asar, Times: tomina(16, 30)},
		{Month: 2, Day: 11, SalahName: isha, Times: tomina(19, 45)},

		{Month: 2, Day: 21, SalahName: fajar, Times: tomina(6, 15)},
		{Month: 2, Day: 21, SalahName: asar, Times: tomina(16, 45)},
		{Month: 2, Day: 21, SalahName: isha, Times: tomina(20, 0)},

		{Month: 3, Day: 2, DayOfWeek: "sunday", SalahName: fajar, Times: tomina(6, 30)},
		{Month: 3, Day: 2, DayOfWeek: "sunday", SalahName: asar, Times: tomina(18, 30)},
		{Month: 3, Day: 2, DayOfWeek: "sunday", SalahName: isha, Times: tomina(21, 15)},

		{Month: 3, Day: 21, SalahName: isha, Times: tomina(21, 30)},

		{Month: 4, Day: 1, SalahName: fajar, Times: tomina(6, 15)},
		{Month: 4, Day: 1, SalahName: isha, Times: tomina(21, 45)},

		{Month: 4, Day: 11, SalahName: fajar, Times: tomina(6, 0)},
		{Month: 4, Day: 11, SalahName: isha, Times: tomina(22, 0)},

		{Month: 4, Day: 21, SalahName: fajar, Times: tomina(5, 45)},
		{Month: 4, Day: 21, SalahName: asar, Times: tomina(6, 45)},

		{Month: 5, Day: 1, SalahName: fajar, Times: tomina(5, 30)},
		{Month: 5, Day: 1, SalahName: asar, Times: tomina(19, 0)},
		{Month: 5, Day: 1, SalahName: isha, Times: tomina(22, 15)},

		{Month: 5, Day: 11, SalahName: fajar, Times: tomina(5, 15)},
		{Month: 5, Day: 11, SalahName: isha, Times: tomina(22, 30)},

		{Month: 5, Day: 21, SalahName: fajar, Times: tomina(5, 0)},
		{Month: 5, Day: 21, SalahName: isha, Times: tomina(22, 45)},

		{Month: 6, Day: 1, SalahName: isha, Times: tomina(23, 0)},

		{Month: 6, Day: 11, SalahName: fajar, Times: tomina(5, 15)},
		{Month: 6, Day: 11, SalahName: isha, Times: tomina(22, 50)},

		{Month: 7, Day: 21, SalahName: fajar, Times: tomina(5, 30)},
		{Month: 7, Day: 21, SalahName: isha, Times: tomina(22, 40)},

		{Month: 8, Day: 1, SalahName: isha, Times: tomina(22, 30)},

		{Month: 8, Day: 11, SalahName: fajar, Times: tomina(6, 0)},
		{Month: 8, Day: 11, SalahName: isha, Times: tomina(22, 0)},

		{Month: 8, Day: 21, SalahName: fajar, Times: tomina(5, 45)},
		{Month: 8, Day: 21, SalahName: isha, Times: tomina(22, 15)},

		{Month: 9, Day: 1, SalahName: fajar, Times: tomina(6, 15)},
		{Month: 9, Day: 1, SalahName: asar, Times: tomina(18, 30)},
		{Month: 9, Day: 1, SalahName: isha, Times: tomina(21, 45)},

		{Month: 9, Day: 11, SalahName: asar, Times: tomina(18, 15)},
		{Month: 9, Day: 11, SalahName: isha, Times: tomina(21, 30)},

		{Month: 9, Day: 21, SalahName: asar, Times: tomina(18, 0)},
		{Month: 9, Day: 21, SalahName: isha, Times: tomina(21, 0)},

		{Month: 10, Day: 1, SalahName: asar, Times: tomina(17, 45)},
		{Month: 10, Day: 1, SalahName: isha, Times: tomina(20, 45)},

		{Month: 10, Day: 11, SalahName: fajar, Times: tomina(6, 30)},
		{Month: 10, Day: 11, SalahName: asar, Times: tomina(17, 30)},
		{Month: 10, Day: 11, SalahName: isha, Times: tomina(20, 30)},

		{Month: 10, Day: 21, SalahName: fajar, Times: tomina(6, 45)},
		{Month: 10, Day: 21, SalahName: asar, Times: tomina(17, 15)},
		{Month: 10, Day: 21, SalahName: isha, Times: tomina(20, 15)},

		{Month: 11, Day: 1, DayOfWeek: "sunday", SalahName: fajar, Times: tomina(6, 15)},
		{Month: 11, Day: 1, DayOfWeek: "sunday", SalahName: asar, Times: tomina(16, 15)},
		{Month: 11, Day: 1, DayOfWeek: "sunday", SalahName: isha, Times: tomina(19, 30)},

		{Month: 12, Day: 11, SalahName: fajar, Times: tomina(6, 30)},
	}

	f, err := os.OpenFile("./1/iqamahtimes.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	t, _ := json.Marshal(Multipart{"begin"})
	if _, err := f.WriteString(fmt.Sprintf("%s\n", string(t))); err != nil {
		log.Println(err)
	}

	for _, v := range times {
		t, _ = json.Marshal(v)
		if _, err := f.WriteString(fmt.Sprintf("%s\n", string(t))); err != nil {
			log.Println(err)
		}
	}

	t, _ = json.Marshal(Multipart{"end"})
	if _, err := f.WriteString(fmt.Sprintf("%s\n", string(t))); err != nil {
		log.Println(err)
	}
}
func main() {

	writeMasjidZakriyaTimes()

	//year := 2020
	//for i:=1; i < 13; i++ {
	//	data := getMonthData2( year)
	//	file, _ := json.MarshalIndent(data, "", " ")
	//	_ = ioutil.WriteFile(fmt.Sprintf("%v.json", i), file, 0644)
	//}

	//getMonthData2( year)
	//file, _ := json.MarshalIndent(times, "", " ")
	//_ = ioutil.WriteFile(fmt.Sprintf("salahtimes.json"), file, 0644)

}

func writeMultipart(begin bool) {
	tag := Multipart{"begin"}
	fileMode := os.O_CREATE | os.O_WRONLY
	if !begin {
		tag.Multipart = "end"
		fileMode = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	}
	t, _ := json.Marshal(tag)
	f, err := os.OpenFile("./1/salahtimes/salahtimes.json", fileMode, 0644)
	if err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString(fmt.Sprintf("%s\n", string(t))); err != nil {
		log.Println(err)
	}
}

func getMonthData2(year int) {
	//var times SalahTimeYear
	writeMultipart(true)
	for month := 0; month < 12; month++ {
		resp1, err := http.Get("http://api.aladhan.com/v1/calendarByCity?city=Buffalo&country=USA&state=NY&school=0&method=99&methodSettings=15,null,15" + fmt.Sprintf("&year=%v&month=%v", year, month+1))
		resp2, err := http.Get("http://api.aladhan.com/v1/calendarByCity?city=Buffalo&country=USA&state=NY&school=1&method=99&methodSettings=18,null,18" + fmt.Sprintf("&year=%v&month=%v", year, month+1))

		if err != nil {
			log.Fatalln(err)
		}

		var d1 AutoGenerated
		json.NewDecoder(resp1.Body).Decode(&d1)

		var d2 AutoGenerated
		json.NewDecoder(resp2.Body).Decode(&d2)

		//fmt.Printf("%+v", decoded)

		if d1.Code == 200 && d2.Code == 200 {
			for i, _ := range d1.Data {
				//fmt.Printf("%v:%v\n", v.Date.Gregorian.Month.Number,  v.Date.Gregorian.Day)

				//times.Salatimes[month][i][0] = convertStrToMin(d1.Data[i].Timings.Fajr)
				//times.Salatimes[month][i][1] = convertStrToMin(d2.Data[i].Timings.Fajr)
				//
				//times.Salatimes[month][i][2] = convertStrToMin(d1.Data[i].Timings.Sunrise)
				//
				//times.Salatimes[month][i][3] = convertStrToMin(d1.Data[i].Timings.Dhuhr)
				//
				//times.Salatimes[month][i][4] = convertStrToMin(d1.Data[i].Timings.Asr)
				//times.Salatimes[month][i][5] = convertStrToMin(d2.Data[i].Timings.Asr)
				//
				//times.Salatimes[month][i][6] = convertStrToMin(d1.Data[i].Timings.Sunset)
				//
				//times.Salatimes[month][i][7] = convertStrToMin(d1.Data[i].Timings.Isha)
				//times.Salatimes[month][i][8] = convertStrToMin(d2.Data[i].Timings.Isha)

				var times Salatimes
				times.Month = month + 1
				times.Day, _ = strconv.Atoi(d1.Data[i].Date.Gregorian.Day)
				times.Fajar15 = convertStrToMin(d1.Data[i].Timings.Fajr)
				times.Fajar18 = convertStrToMin(d2.Data[i].Timings.Fajr)

				times.Sunrise = convertStrToMin(d1.Data[i].Timings.Sunrise)

				times.Zuhar = convertStrToMin(d1.Data[i].Timings.Dhuhr)

				times.AsrShafai = convertStrToMin(d1.Data[i].Timings.Asr)
				times.AsrHanfi = convertStrToMin(d2.Data[i].Timings.Asr)

				times.Maghrib = convertStrToMin(d1.Data[i].Timings.Sunset)

				times.Isha15 = convertStrToMin(d1.Data[i].Timings.Isha)
				times.Isha18 = convertStrToMin(d2.Data[i].Timings.Isha)

				fmt.Printf("%v %+v\n", d1.Data[i].Date.Gregorian.Day, times)
				//st.Salatimes = append(st.Salatimes, times)
				//fmt.Printf("index: %v, day %v/%v, Fajar15 %v\t Fajar 18 %v\n", i, d1.Data[i].Date.Gregorian.Day, d2.Data[i].Date.Gregorian.Day, convertStrToSalaTime(d1.Data[i].Timings.Fajr), convertStrToSalaTime(d2.Data[i].Timings.Fajr));

				t, _ := json.Marshal(times)
				f, err := os.OpenFile("./1/salahtimes/salahtimes.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Println(err)
				}
				if _, err := f.WriteString(fmt.Sprintf("%s\n", string(t))); err != nil {
					log.Println(err)
				}
			}
		}
	}

	writeMultipart(false)
	//file, _ := json.Marshal(times.Salatimes)
	//_ = ioutil.WriteFile(fmt.Sprintf("./1/salahtimes/salahtimes.json"), file, 0644)

}

func getMonthData(month int, year int) SalahTimeRecord {
	resp1, err := http.Get("http://api.aladhan.com/v1/calendarByCity?city=Buffalo&country=USA&state=NY&school=0&method=99&methodSettings=15,null,15" + fmt.Sprintf("&year=%v&month=%v", year, month))
	resp2, err := http.Get("http://api.aladhan.com/v1/calendarByCity?city=Buffalo&country=USA&state=NY&school=1&method=99&methodSettings=18,null,18" + fmt.Sprintf("&year=%v&month=%v", year, month))

	st := SalahTimeRecord{Month: month, Year: year}
	if err != nil {
		log.Fatalln(err)
	}

	var d1 AutoGenerated
	json.NewDecoder(resp1.Body).Decode(&d1)

	var d2 AutoGenerated
	json.NewDecoder(resp2.Body).Decode(&d2)

	//fmt.Printf("%+v", decoded)

	if d1.Code == 200 && d2.Code == 200 {
		for i, _ := range d1.Data {
			//fmt.Printf("%v:%v\n", v.Date.Gregorian.Month.Number,  v.Date.Gregorian.Day)
			var times Salatimes
			times.Fajar15 = convertStrToMin(d1.Data[i].Timings.Fajr)
			times.Fajar18 = convertStrToMin(d2.Data[i].Timings.Fajr)

			times.Sunrise = convertStrToMin(d1.Data[i].Timings.Sunrise)

			times.Zuhar = convertStrToMin(d1.Data[i].Timings.Dhuhr)

			times.AsrShafai = convertStrToMin(d1.Data[i].Timings.Asr)
			times.AsrHanfi = convertStrToMin(d2.Data[i].Timings.Asr)

			times.Maghrib = convertStrToMin(d1.Data[i].Timings.Sunset)

			times.Isha15 = convertStrToMin(d1.Data[i].Timings.Isha)
			times.Isha18 = convertStrToMin(d2.Data[i].Timings.Isha)

			fmt.Printf("%v %+v\n", d1.Data[i].Date.Gregorian.Day, times)
			st.Salatimes = append(st.Salatimes, times)
			//fmt.Printf("index: %v, day %v/%v, Fajar15 %v\t Fajar 18 %v\n", i, d1.Data[i].Date.Gregorian.Day, d2.Data[i].Date.Gregorian.Day, convertStrToSalaTime(d1.Data[i].Timings.Fajr), convertStrToSalaTime(d2.Data[i].Timings.Fajr));

		}
	}
	return st
}

type AutoGenerated struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []Data `json:"data"`
}
type Timings struct {
	Fajr     string `json:"Fajr"`
	Sunrise  string `json:"Sunrise"`
	Dhuhr    string `json:"Dhuhr"`
	Asr      string `json:"Asr"`
	Sunset   string `json:"Sunset"`
	Maghrib  string `json:"Maghrib"`
	Isha     string `json:"Isha"`
	Imsak    string `json:"Imsak"`
	Midnight string `json:"Midnight"`
}
type Weekday struct {
	En string `json:"en"`
}
type Month struct {
	Number int    `json:"number"`
	En     string `json:"en"`
}

//type Designation struct {
//	Abbreviated string `json:"abbreviated"`
//	Expanded    string `json:"expanded"`
//}
type Gregorian struct {
	Date    string  `json:"date"`
	Format  string  `json:"format"`
	Day     string  `json:"day"`
	Weekday Weekday `json:"weekday"`
	Month   Month   `json:"month"`
	Year    string  `json:"year"`
	//Designation Designation `json:"designation"`
}

//type Weekday struct {
//	En string `json:"en"`
//	Ar string `json:"ar"`
//}
//type Month struct {
//	Number int    `json:"number"`
//	En     string `json:"en"`
//	Ar     string `json:"ar"`
//}
//type Hijri struct {
//	Date        string      `json:"date"`
//	Format      string      `json:"format"`
//	Day         string      `json:"day"`
//	Weekday     Weekday     `json:"weekday"`
//	Month       Month       `json:"month"`
//	Year        string      `json:"year"`
//	Designation Designation `json:"designation"`
//	Holidays    []string    `json:"holidays"`
//}
type Date struct {
	Readable  string    `json:"readable"`
	Timestamp string    `json:"timestamp"`
	Gregorian Gregorian `json:"gregorian"`
	//Hijri     Hijri     `json:"hijri"`
}
type Params struct {
	Fajr int `json:"Fajr"`
	Isha int `json:"Isha"`
}
type Method struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Params Params `json:"params"`
}
type Offset struct {
	Imsak    int `json:"Imsak"`
	Fajr     int `json:"Fajr"`
	Sunrise  int `json:"Sunrise"`
	Dhuhr    int `json:"Dhuhr"`
	Asr      int `json:"Asr"`
	Maghrib  int `json:"Maghrib"`
	Sunset   int `json:"Sunset"`
	Isha     int `json:"Isha"`
	Midnight int `json:"Midnight"`
}
type Meta struct {
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	Timezone                 string  `json:"timezone"`
	Method                   Method  `json:"method"`
	LatitudeAdjustmentMethod string  `json:"latitudeAdjustmentMethod"`
	MidnightMode             string  `json:"midnightMode"`
	School                   string  `json:"school"`
	Offset                   Offset  `json:"offset"`
}
type Data struct {
	Timings Timings `json:"timings"`
	Date    Date    `json:"date"`
	Meta    Meta    `json:"meta"`
}
