package conf

import (
	"flag"
	"os"
)

type config struct {
	WorkDir        string
	RewriteBaseUrl string
	BaseUrl        string
	Cookies        string
}

var Conf config

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	flag.StringVar(&Conf.WorkDir, "d", pwd, "specify a directory that mirror data will store in")
	flag.StringVar(&Conf.RewriteBaseUrl, "r", "", "specify a url to overwrite the links in crawled html")
	flag.StringVar(&Conf.Cookies, "c", "", "specify cookies")
	flag.Parse()
	if Conf.RewriteBaseUrl == "" {
		Conf.RewriteBaseUrl = Conf.WorkDir
	}

	Conf.BaseUrl = flag.Arg(flag.NArg() - 1)
}
