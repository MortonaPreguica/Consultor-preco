package scrapper

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

type Produto struct {
	nome  string
	preco float64
}

func Scrapper() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org", "www.amazon.com.br"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)
	// On every a element which has href attribute call callback
	c.OnHTML("#centerCol", func(e *colly.HTMLElement) {
		nomeProduto := e.ChildText("#productTitle")

		valorRaw := e.ChildText("#sns-base-price")
		re := regexp.MustCompile("[0-9]+")
		valores := re.FindAllString(valorRaw, -1)
		valor := fmt.Sprintf("%s.%s", valores[0], valores[1])
		num, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			fmt.Println("Erro na conversão:", err)
			return
		}
		produto := &Produto{nome: nomeProduto, preco: num}
		// Print link
		fmt.Println(produto)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		fmt.Println("UserAgent", r.Headers.Get("User-Agent"))
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.amazon.com.br/dp/B07LBM4STJ/ref=sspa_dk_detail_2?psc=1&pd_rd_i=B07LBM4STJ&pd_rd_w=EyG0c&content-id=amzn1.sym.dd14805d-2347-4e14-a8a3-4a68eee2b6c8&pf_rd_p=dd14805d-2347-4e14-a8a3-4a68eee2b6c8&pf_rd_r=QZ1EASF9F3YFQ1YKJ4SZ&pd_rd_wg=A5kAg&pd_rd_r=42f1f4ca-fc55-4757-9269-d98f5459e500&s=hpc&sp_csd=d2lkZ2V0TmFtZT1zcF9kZXRhaWw")
}
