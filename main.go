package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"text/template"
	"time"
)

type Members map[string]*Member

type Member struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Blog        string  `json:"blog"`
	Feed        string  `json:"feed"`
	Twitter     string  `json:"twitter"`
	Date        RssTime `json:"date_joined"`
	GravatarURL string  `json:"-"`
}

func (m *Member) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.Name, start)
}

func parseMembers(filename string) (Members, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var m Members
	dec := json.NewDecoder(f)
	err = dec.Decode(&m)
	if err != nil {
		return nil, err
	}

	for _, member := range m {
		member.GravatarURL = gravatarURL(member.Email)
	}
	return m, nil
}

func gravatarURL(email string) string {
	h := md5.New()
	h.Write([]byte(email))
	return fmt.Sprintf("https://www.gravatar.com/avatar/%v.jpg", hex.EncodeToString(h.Sum(nil)))
}

type RssFeed struct {
	XMLName xml.Name    `xml:"rss"`
	Version string      `xml:"version,attr"`
	Channel *RssChannel `xml:"channel"`
}

type RssChannel struct {
	Title         string     `xml:"title"`
	Link          string     `xml:"link"`
	Description   string     `xml:"description"`
	Language      string     `xml:"language,omitempty"`
	LastBuildDate RssTime    `xml:"lastBuildDate,omitempty"`
	Items         []*RssItem `xml:"item"`
}

type RssItem struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Author      *Member  `xml:"author,omitempty"`
	Category    []string `xml:"category,omitempty"`
	Guid        string   `xml:"guid,omitempty"`
	PubDate     RssTime  `xml:"pubDate,omitempty"`
	Source      string   `xml:"source,omitempty"`
}

type RssTime struct {
	time.Time
}

func (t *RssTime) String() string {
	return t.Format("02 Jan 2006")
}

func (t *RssTime) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	parse, err := time.Parse("2006-01-02", v)
	if err != nil {
		return err
	}

	*t = RssTime{parse}
	return nil
}

func (t *RssTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(time.RFC1123Z, v)
	if err == nil {
		*t = RssTime{parse}
		return nil
	}

	parse, err = time.Parse(time.RFC1123, v)
	if err == nil {
		*t = RssTime{parse}
		return nil
	}

	return fmt.Errorf("unknown time format: %s", v)
}

type SortedRssItems []*RssItem

func (ri SortedRssItems) Len() int           { return len(ri) }
func (ri SortedRssItems) Swap(i, j int)      { ri[i], ri[j] = ri[j], ri[i] }
func (ri SortedRssItems) Less(i, j int) bool { return ri[i].PubDate.Unix() > ri[j].PubDate.Unix() }

var (
	// flag options
	membersFile = flag.String("m", "members.json", "Arquivo com informações dos membros")
	staticDir   = flag.String("d", "./static", "Static directory")
	port        = flag.Int("p", 9000, "HTTP port")

	// global vars
	thcFeed = &RssFeed{
		Version: "2.0",
		Channel: &RssChannel{
			Title:         "TeresinaHC Planet",
			Link:          "http://planet.teresinahc.org/",
			Description:   "TeresinaHC Planet",
			Language:      "pt-BR",
			LastBuildDate: RssTime{time.Now()},
		},
	}

	htmlPage = template.Must(template.New("html").ParseFiles(path.Join(*staticDir, "index.html")))
)

func feed(w http.ResponseWriter, req *http.Request) {
	log.Printf("Serving feed to %s\n", req.RemoteAddr)

	w.Header().Set("Content-Type", "text/xml; charset=UTF-8")
	enc := xml.NewEncoder(w)
	err := enc.Encode(thcFeed)
	if err != nil {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func homePage(w http.ResponseWriter, req *http.Request) {
	log.Printf("Visit from %s\n", req.RemoteAddr)

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	err := htmlPage.ExecuteTemplate(w, "index.html", map[string]interface{}{
		"Items": thcFeed.Channel.Items,
	})
	if err != nil {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse()
	members, err := parseMembers(*membersFile)
	if err != nil {
		log.Fatalf("ERROR Invalid member file '%s': %v", *membersFile, err)
	}

	var items []*RssItem
	for m, mm := range members {
		f, _ := os.Open(fmt.Sprintf("feeds/%s.xml", m))
		d := xml.NewDecoder(f)

		var rss2 RssFeed
		err := d.Decode(&rss2)
		if err != nil {
			log.Fatalf("%s>> %v\n", m, err)
		}
		for _, i := range rss2.Channel.Items {
			i.Author = mm
		}
		items = append(items, rss2.Channel.Items...)
	}
	sort.Sort(SortedRssItems(items))
	thcFeed.Channel.Items = items

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(*staticDir))))
	http.HandleFunc("/feed", feed)
	http.HandleFunc("/", homePage)

	log.Printf("Starting HTTP server at :%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
