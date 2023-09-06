package sprapp

import (
	//	"bytes"

	//"arena"
	//"bytes"
	"context"
	//"encoding/hex"
	//"fmt"
	//"log"
	//"math/rand"
	"net/netip"
	//"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	//cxstrconv "github.com/cloudxaas/gostrconv"
	//qtc "github.com/valyala/quicktemplate/qtc"
	//templatesHTML "github.com/cloudxaas/gosprapp/web/app/sprapp/templates/html"
	//"github.com/tdewolff/minify/v2"
	//"github.com/tdewolff/minify/v2/html"
	//"github.com/zeebo/blake3"
)

var (
	//m *minify.M

	fs = &app.FS{
		Root: "/var/web/site/sprapp/public_html",
		///IndexNames:         []string{"index.html"},
		//GenerateIndexPages: true,
		//PathRewrite: app.NewPathSlashesStripper(1),
		PathNotFound: func(_ context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusNotFound, "The requested resource does n2ot exist")
		},
		CacheDuration: time.Second * 5,
		//IndexNames:           indexNames,
		//Compress:             true,
		//CompressedFileSuffix: "hertz",
		//AcceptByteRange:      true,
	}
	fsHandler = fs.NewRequestHandler()

	topEmailProviders = []string{
		"gmail.com",
		"yahoo.com",
		"hotmail.com",
		"aol.com",
		"hotmail.co.uk",
		"hotmail.fr",
		"msn.com",
		"yahoo.fr",
		"wanadoo.fr",
		"orange.fr",
		"comcast.net",
		"yahoo.co.uk",
		"yahoo.com.br",
		"yahoo.co.in",
		"live.com",
		"rediffmail.com",
		"free.fr",
		"gmx.de",
		"web.de",
		"yandex.ru",
		"ymail.com",
		"libero.it",
		"outlook.com",
		"uol.com.br",
		"bol.com.br",
		"mail.ru",
		"cox.net",
		"hotmail.it",
		"sbcglobal.net",
		"sfr.fr",
		"live.fr",
		"verizon.net",
		"live.co.uk",
		"googlemail.com",
		"yahoo.es",
		"inbox.com",
		"icloud.com",
		"me.com",
		"mail.com",
		"planet.nl",
		"tin.it",
		"live.nl",
		"yahoo.it",
		"hotmail.de",
		"att.net",
		"laposte.net",
		"yahoo.de",
		"alice.it",
		"rocketmail.com",
		"yahoo.com.mx",
		"voila.fr",
		"gmx.com",
		"mail.dk",
		"email.com",
		"btinternet.com",
		"charter.net",
		"sky.com",
		"earthlink.net",
		"optonline.net",
		"freenet.de",
		"t-online.de",
		"aliceadsl.fr",
		"virgilio.it",
		"home.nl",
		"qq.com",
		"telenet.be",
		"tiscali.it",
		"zoho.com",
		"cfl.rr.com",
		"tele2.nl",
		"163.com",
		"bellsouth.net",
		"shaw.ca",
		"aim.com",
		"netscape.net",
		"ziggo.nl",
		"nate.com",
		"yahoo.com.ph",
		"yahoo.com.au",
		"xtra.co.nz",
		"mindspring.com",
		"ntlworld.com",
		"blueyonder.co.uk",
		"virginmedia.com",
		"live.se",
		"telus.net",
		"centrum.cz",
		"yahoo.com.hk",
		"rambler.ru",
		"abv.bg",
		"hotmail.be",
		"onet.pl",
		"bigpond.com",
		"protonmail.com",
		"yahoo.com.sg",
		"comcast.com",
		"yahoo.co.id",
		"frontier.com",
		"rogers.com",
		"optusnet.com.au",
		"sympatico.ca",
		"cox.com",
		"bt.com",
		"hotmail.se",
		"interia.pl",
		"hotmail.no",
		"o2.pl",
		"yahoo.com.my",
		"telefonica.net",
		"yahoo.com.vn",
		"yahoo.com.ve",
		"wp.pl",
		"bigpond.net.au",
		"seznam.cz",
		"iinet.net.au",
		"yahoo.com.co",
		"yahoo.gr",
		"mac.com",
		"bluewin.ch",
		"sky.it",
		"telstra.com",
		"gmx.fr",
		"naver.com",
		"verizon.com",
		"sasktel.net",
		"chello.nl",
		"tiscali.co.uk",
		"fastmail.fm",
		"pacbell.net",
		"yahoo.com.pe",
		"hushmail.com",
		"laposte.fr",
		"excite.com",
		"skynet.be",
		"bell.net",
		"yahoo.com.pk",
		"hetnet.nl",
		"live.jp",
		"yahoo.co.th",
		"wanadoo.es",
		"live.it",
		"walla.com",
		"orange.net",
		"gmx.net",
		"yahoo.co.nz",
		"hotmail.es",
		"lycos.com",
		"hotmail.ca",
		"yahoo.co.kr",
		"yandex.ua",
		"o2.com",
		"poczta.onet.pl",
		"yahoo.com.tw",
		"eircom.net",
		"libertyglobal.com",
		"kpnmail.nl",
		"xs4all.nl",
		"gmx.at",
		"tiscali.fr",
		"pt.lu",
		"live.be",
		"juno.com",
		"talktalk.net",
		"yahoo.ca",
		"neuf.fr",
		"live.ca",
		"hotmail.gr",
		"aol.co.uk",
		"terra.com.br",
		"ig.com.br",
		"itelefonica.com.br",
		"r7.com",
		"oi.com.br",
		"globo.com",
		"globomail.com",
		"pop.com.br",
		"sina.com",
		"gmw.cn",
		"126.com",
		"yeah.net",
		"sohu.com",
		"aliyun.com",
		"daum.net",
		"hanmail.net",
		"hotmail.com.br",
		"outlook.com.br",
		"zipmail.com.br",
		"hotmail.com.mx",
		"prodigy.net.mx",
		"hotmail.com.au",
		"hotmail.co.nz",
		"sina.cn",
		"21cn.com",
		"139.com",
		"hotmail.ru",
		"gmx.fr",
		"hotmail.com.hk",
		"hotmail.com.tr",
		"hotmail.cl",
		"hotmail.co.il",
		"walla.co.il",
		"hotmail.co.za",
		"yahoo.co.za",
		"ya.ru",
		"list.ru",
		"live.no",
		"hotmail.dk",
		"live.dk",
		"finn.no",
		"start.no",
		"online.no",
		"luukku.com",
		"kolumbus.fi",
		"netti.fi",
		"welho.com",
		"pp.inet.fi",
		"hotmail.pt",
		"live.com.pt",
		"mail.pt",
		"sapo.pt",
		"iol.pt",
		"yahooxtra.co.nz",
		"att.net",
	}
)
/*
func init() {
	m = minify.New()
	m.AddFunc("text/html", html.Minify)
}

type H struct {
	Data []byte
	Err  error
}
*/
/*
type U struct {
	User user.User
	Err  error
}
*/
/*
type MinifyResponseWriter struct {
	http.ResponseWriter
	io.WriteCloser
}

func (m MinifyResponseWriter) Write(b []byte) (int, error) {
	return m.WriteCloser.Write(b)
}

// MinifyResponseWriter must be closed explicitly by calling site.
func MinifyFilter(mediatype string, res http.ResponseWriter) MinifyResponseWriter {
	m := minify.New()
	// add minfiers

	mw := m.Writer(mediatype, res)
	return MinifyResponseWriter{res, mw}
}

//*/

func MainHTTP(ctx *app.RequestContext, reqMethod uint8, ipAddr netip.Addr, reqHost, reqURI []byte) {


	//resp.Reset()
	//resp.Header.SetStatusCode(200)


	//switch string(reqURI) {
		//default:

		//	log.Printf("bngfduiou32hr2")
			ctx.Response.Header.SetStatusCode(200)
			ctx.Response.Header.SetContentType("text/plain")
			ctx.Response.SetBodyRaw([]byte("Hello World!"))
			/*
			p := &templatesHTML.MainPage{
				CTX: ctx,
			}
			templatesHTML.WritePageTemplate(ctx, p)
			*/
		//fmt.Fprintf(ctx, "No subdomain found")
	//}

}
