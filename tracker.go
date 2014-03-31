package main

import "os"
import "fmt"
import "log"
import "strings"
import "regexp"
import "time"
import "strconv"
import "io/ioutil"
import "net/url"
import "net/http"
import "net/http/cookiejar"
import "image/png"
import "slorry/kpch/crack"

var TrackTableRe *regexp.Regexp = regexp.MustCompile(
    `(?s)tbl_track_id_info[^>]+>(.*)</table>`)

var TrackTableTrRe *regexp.Regexp = regexp.MustCompile(
    `(?s)<tr.*?>(.*?)</tr>`)

var TrackTableTdRe *regexp.Regexp = regexp.MustCompile(
    `(?s)<td.*?>(.*?)</td>`)

var LineSplitRe *regexp.Regexp = regexp.MustCompile(
    `(.{10})(?: |$)`)

var CaptchaErrorRe *regexp.Regexp = regexp.MustCompile(
	`CaptchaErrorCodeContainer">([^<]+)`)

var BaseUrl string = "http://www.russianpost.ru"
var TrackUrl string = BaseUrl + "/tracking20"
var CaptchaUrl string = TrackUrl + "/Code/Code.png.ashx"

type TrackStage struct {
    Name string
    Date time.Time
    IndexWhere string
    NameWhere string
    Attr string
    Weight float64
    Value float64
    Price float64
    IndexTarget string
    NameTarget string
}

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Usage: tracker <tracking number>\n")
        return
    }

    PrintTextTrack(GetTrack(os.Args[1]))
}

func TryToBreakCaptcha(number string) string {
    u, err := url.Parse(TrackUrl)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Get(TrackUrl)
    if err != nil {
        log.Fatal(err)
    }

    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    resp, err = http.Get(CaptchaUrl)
    if err != nil {
        log.Fatal(err)
    }

    img, err := png.Decode(resp.Body)

    captchaCode := ""
    for _, sym := range crack.Crack(img) {
        captchaCode += sym.Sym.Char
    }

    jarOpts := cookiejar.Options{
        PublicSuffixList: nil,
    }
    jar, err := cookiejar.New(&jarOpts)
    if err != nil {
        log.Fatal(err)
    }
    jar.SetCookies(u, resp.Cookies())

    c := &http.Client{Jar: jar}
    resp, err = c.PostForm(TrackUrl,
        url.Values{
            "CaptchaCode": {captchaCode},
            "BarCode": {strings.ToUpper(number)},
            "searchsign": {"1"},
        })
    if err != nil {
        log.Fatal(err)
    }

    bytes, err = ioutil.ReadAll(resp.Body)

	body := string(bytes)

	capchaErrorMatch := CaptchaErrorRe.FindStringSubmatch(body)
	if len(capchaErrorMatch) > 0 {
		log.Println("неправильно прочитал капчу")
		log.Printf("ссылка на капчу: %s", CaptchaUrl)
		log.Printf("прочитанный код: %s", captchaCode)

		return ""
	}

	return body
}

func GetTrack(number string) []*TrackStage {
	var body string

	for i := 0; i < 5; i++ {
		body = TryToBreakCaptcha(number)
		if body != "" {
			break
		}
	}

	if body == "" {
		log.Fatal("не могу взломать капчу")
	} else {
		log.Println("ok, капча разгадана")
	}

    trackTableMatch := TrackTableRe.FindStringSubmatch(body)
    if len(trackTableMatch) < 1 {
        log.Fatal("отправление не найдено")
    }

    trackTable := trackTableMatch[1]

    tracking := make([]*TrackStage, 0)

    for _, tr := range TrackTableTrRe.FindAllStringSubmatch(trackTable, -1) {
        tds := TrackTableTdRe.FindAllStringSubmatch(tr[1], -1)
        if len(tds) != 10 {
            continue
        }

        weight := 0.0
        if tds[5][1] != "-" {
            weight, _ = strconv.ParseFloat(
                strings.Replace(tds[5][1], ",", ".", -1), 64)
        }

        value := 0.0
        if tds[6][1] != "-" {
            value, _ = strconv.ParseFloat(
                strings.Replace(tds[6][1], ",", ".", -1), 64)
        }

        price := 0.0
        if tds[7][1] != "-" {
            price, _ = strconv.ParseFloat(
                strings.Replace(tds[7][1], ",", ".", -1), 64)
        }

		date, _ := time.ParseInLocation("02.01.2006 15:04:05", tds[1][1],
            time.Local)

        tracking = append(tracking, &TrackStage{
            Name: strings.TrimSpace(tds[0][1]),
            Date: date,
            IndexWhere: tds[2][1],
            NameWhere: tds[3][1],
            Attr: tds[4][1],
            Weight: weight,
            Value: value,
            Price: price,
            IndexTarget: tds[8][1],
            NameTarget: tds[9][1],
        })
    }

    return tracking
}

func PrintTextTrack(track []*TrackStage)() {
    for _, track := range track {
        days := int(time.Now().Sub(track.Date).Hours() / 24)

        if days == 0 {
            fmt.Println("Сегондя:")
        } else {
            fmt.Printf("%d %s назад:\n", days,
                Decline(days, "день", "дня", "дней"))
        }

        if track.Attr != "" {
            fmt.Printf("%s (%s)\n", track.Name, track.Attr)
        } else {
            fmt.Printf("%s\n", track.Name)
        }
        fmt.Println(track.NameWhere)

        if track.Weight > 0 {
            fmt.Printf("%3.2f кг.\n", track.Weight)
        }

        fmt.Println()
    }
}

func Decline(count int, one string, two string, five string) string {
    if count > 10 && count < 20 {
        return five
    }

    if count % 10 == 1 {
        return one
    }

    if count % 10 > 1 && count % 10 < 5 {
        return two
    }

    return five
}
