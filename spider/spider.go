package spider

import (
	"bytes"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/go-playground/pool.v1"

	"github.com/zwh8800/getall/conf"
	"github.com/zwh8800/getall/event"
	"github.com/zwh8800/getall/util"
)

type spider struct {
	pool           *pool.Pool
	baseUrl        *url.URL
	rewriteBaseUrl *url.URL
	workDir        string
	client         *http.Client
}

func NewSpider() *spider {
	baseUrl, err := url.Parse(conf.Conf.BaseUrl)
	if err != nil {
		panic(err)
	}

	rewriteBaseUrl, err := url.Parse(conf.Conf.RewriteBaseUrl)
	if err != nil {
		panic(err)
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	jar.SetCookies(baseUrl, util.ParseCookies(conf.Conf.Cookies))

	return &spider{
		pool:           pool.NewPool(4, 1000),
		baseUrl:        baseUrl,
		rewriteBaseUrl: rewriteBaseUrl,
		workDir:        conf.Conf.WorkDir,
		client: &http.Client{
			Jar: jar,
		},
	}
}

func spiderFn(job *pool.Job) {
	s := job.Params()[0].(*spider)
	u := job.Params()[1].(*url.URL)

	resp, err := s.client.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reader := tee(s, resp)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)
	}

	doc.Html()
}

func tee(s *spider, resp *http.Response) io.Reader {
	id := util.GenerateId()
	event.Server.Publish(event.ProgressStart, id, "")
	defer event.Server.Publish(event.ProgressFinish, id)

	reader := &bytes.Buffer{}
	totalLength := resp.ContentLength
	var readLen int64 = 0
	for {
		n, err := io.CopyN(reader, resp.Body, 4096)
		if err != nil {
			break
		}
		readLen += n

		if totalLength != -1 {
			event.Server.Publish(event.ProgressUpdate, id, int(float64(readLen)/float64(totalLength)*100))
		}
	}

	return reader
}

func (s *spider) Queue(u *url.URL) {
	s.pool.Queue(spiderFn, s, u)
}

func (s *spider) Go() {
	s.Queue(s.baseUrl)
}

func (s *spider) Stop() {

}

var defaultSpider = NewSpider()

func Go() {
	defaultSpider.Go()
}

func Stop() {
	defaultSpider.Stop()
}
