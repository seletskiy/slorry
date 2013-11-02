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
import "image/png"
import "slorry/kpch/crack"


var CaptchaLinkRe *regexp.Regexp = regexp.MustCompile(
        `/CaptchaService/CaptchaImage.ashx\?Id=(\d+)`)

var TrackTableRe *regexp.Regexp = regexp.MustCompile(
    `(?s)tbl_track_id_info[^>]+>(.*)</table>`)

var TrackTableTr *regexp.Regexp = regexp.MustCompile(
    `(?s)<tr.*?>(.*?)</tr>`)

var TrackTableTd *regexp.Regexp = regexp.MustCompile(
    `(?s)<td.*?>(.*?)</td>`)

var LineSplitRe *regexp.Regexp = regexp.MustCompile(
    `(.{10})(?: |$)`)

var BaseUrl string = "http://www.russianpost.ru"

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

func GetTrack(number string) []*TrackStage {
    resp, err := http.Get(BaseUrl)
    if err != nil {
        log.Fatal(err)
    }

    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    captchaInfo := CaptchaLinkRe.FindSubmatch(bytes)
    captchaLink := string(captchaInfo[0])
    captchaId := string(captchaInfo[1])

    resp, err = http.Get(BaseUrl + captchaLink)
    if err != nil {
        log.Fatal(err)
    }

    img, err := png.Decode(resp.Body)

    captchaCode := ""
    for _, sym := range crack.Crack(img) {
        captchaCode += sym.Sym.Char
    }

    resp, err = http.PostForm(BaseUrl +
        "/resp_engine.aspx?Path=rp/servise/ru/home/postuslug/trackingpo",
        url.Values{
            "CaptchaId": {captchaId},
            "InputedCaptchaCode": {captchaCode},
            "BarCode": {strings.ToUpper(number)},
            "searchsign": {"1"},
        })
    if err != nil {
        log.Fatal(err)
    }

    bytes, err = ioutil.ReadAll(resp.Body)

    trackTableMatch := TrackTableRe.FindStringSubmatch(string(bytes))
    if len(trackTableMatch) < 1 {
        log.Fatal("отправление не найдено")
    }

    trackTable := trackTableMatch[1]

    tracking := make([]*TrackStage, 0)

    for _, tr := range TrackTableTr.FindAllStringSubmatch(trackTable, -1) {
        tds := TrackTableTd.FindAllStringSubmatch(tr[1], -1)
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

        date, _ := time.ParseInLocation("02.01.2006 15:04", tds[1][1],
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
