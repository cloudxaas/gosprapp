package sprapp

import (
	//	"bytes"

	"arena"
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/netip"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	cxstrconv "github.com/cloudxaas/gostrconv"
	//qtc "github.com/valyala/quicktemplate/qtc"
	templatesHTML "github.com/cloudxaas/gosprapp/web/app/sprapp/templates/html"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"github.com/zeebo/blake3"
)

var (
	m *minify.M

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

func init() {
	m = minify.New()
	m.AddFunc("text/html", html.Minify)
}

type H struct {
	Data []byte
	Err  error
}
type U struct {
	User user.User
	Err  error
}

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

	switch string(reqURI) {

	case "/m/s.png":
		log.Printf("gfdsugigd")

		ctx.Response.Header.SetStatusCode(200)
		ctx.Response.Header.SetContentType("image/png")
		fsHandler(context.Background(), ctx)
		return

	case "/m/u.js":

		ctx.Response.Header.SetStatusCode(200)
		ctx.Response.Header.SetContentType("text/javascript")
		fsHandler(context.Background(), ctx)

		mem := arena.NewArena()
		o := arena.New[H](mem)

		//o.Num = 123 // incorrect: use after free
		///*
		o.Data, o.Err = m.Bytes("text/javascript", ctx.Response.Body())
		if o.Err != nil {
			log.Printf("minify1 %s", o.Err)
		} else {
			//log.Printf("o.Data = %s", o.Data)
			ctx.Response.SetBody(o.Data)
		}

		mem.Free()
		return
	case "/m/a.css":

		ctx.Response.Header.SetStatusCode(200)
		ctx.Response.Header.SetContentType("text/css")
		//ctx := &fasthttp.RequestCtx{}
		//resp.Reset()
		//ctx.Init(ctx, nil, nil)

		//log.Printf("ngfuds9guheugfhe")

		fsHandler(context.Background(), ctx)

		mem := arena.NewArena()
		o := arena.New[H](mem)

		//o.Num = 123 // incorrect: use after free
		///*
		o.Data, o.Err = m.Bytes("text/css", ctx.Response.Body())
		if o.Err != nil {
			log.Printf("minify1 %s", o.Err)
		} else {
			//log.Printf("o.Data = %s", o.Data)
			ctx.Response.SetBody(o.Data)
		}

		//*/

		//ctx.Response.SetBody(out)
		mem.Free()
		return

	}
	host := ctx.Request.Header.Host()

	// Split the Host header by '.' to separate subdomain and domain
	parts := bytes.Split(host, []byte("."))

	// Check if there are at least two parts (subdomain and domain)
	if len(parts) >= 2 {
		log.Printf("ngusaigsafd")
		// The subdomain is the first part of the split
		//subdomain :=
		if bytes.EqualFold(parts[0], []byte("user")) {
			//fmt.Fprintf(ctx, "Subdomain: %s", parts[0])
			switch string(reqURI) {
			case "/1/s":

				//		sprappconf.Display()

				log.Printf("hudsoiafuhdsafsa")
				mem := arena.NewArena()
				u := arena.New[U](mem)

				defer mem.Free()

				//userData := &user.User{}

				/*
					// Assuming you're receiving form data
					c.PostArgs().VisitAll(func(key, value []byte) {
						if string(key) == "name" {
							fmt.Printf("This is %s!", string(value))
						}
					})
				*/

				u.User.Email = string(ctx.FormValue("e"))
				u.User.Password = ctx.FormValue("p")
				u.User.FirstName = string(ctx.FormValue("f"))
				u.User.LastName = string(ctx.FormValue("l"))
				u.User.DOB = uint32(cxstrconv.Atoi(string(ctx.FormValue("l"))))
				gender, _ := strconv.Atoi(string(ctx.FormValue("g")))
				u.User.Gender = uint8(gender)
				pronouns, _ := strconv.Atoi(string(ctx.FormValue("n")))
				u.User.Pronouns = uint8(pronouns)
				//log.Printf("k fgosdigjfdsoigjfdsogd = %s = %s = %s", email, fname, lastname)
				/*
					ctx.PostArgs().VisitAll(func(key, value []byte) {

						log.Printf("k = %s, v = %s", key, value)
						switch string(key) {

						case "e":
							u.User.Email = string(value)
						case "p":
							u.User.Password = string(value)
						case "f":
							u.User.FirstName = string(value)
						case "l":
							u.User.LastName = string(value)
						case "b":
							u.User.DOB = uint32(cxstrconv.Atoi(string(value))) // assuming date format as "yyyy-mm-dd"
						case "g":
							gender, _ := strconv.Atoi(string(value))
							u.User.Gender = uint8(gender)
						case "n":
							pronouns, _ := strconv.Atoi(string(value))
							u.User.Pronouns = uint8(pronouns)
						}
					})
				*/
				//snowflakex.GenerateID([]byte(u.User.Email))
				//u.User.ID = cxstrconv.Int64toa(int64(snowflakex.GenerateID([]byte(u.User.Email))))
				/*
					if err != nil {
						log.Printf("Error parsing form data: %s", err)
						return
					}
				*/

				log.Printf("eeeeeeeeeeeeee = %s, p = %s, f = %s, l = %s, b = %s, g = %v, n = %v", u.User.Email, u.User.Password, u.User.FirstName, u.User.LastName, u.User.DOB, u.User.Gender, u.User.Pronouns)

				// Now call CreateUser function with parsed data
				retval := user.RegisterUser(&u.User)
				if retval > 0 {
					ctx.Response.Header.SetStatusCode(422)
					//ctx.Response.Header.SetStatusCode(200)
					ctx.Response.Header.SetContentType("text/json")

					if retval == 2 {
						ctx.Response.SetBody([]byte(`{"errors": { "email": "Email is already taken" }}`))
					} else if retval == 1 {
						ctx.Response.SetBody([]byte(`{"errors": { "general": "Database connection error. Please try again later." }}`))
					}

					log.Printf("Error creating user: %v", retval)
					return
				}
				ctx.Response.Header.SetStatusCode(200)

				//user.CreateUserHandler(ctx)
			case "/1/r":
				ctx.Response.Header.SetStatusCode(200)
				ctx.Response.Header.SetContentType("text/html")
				p := &templatesHTML.ErrorPage{
					Path: ctx.Path(),
				}
				templatesHTML.WritePageTemplate(ctx, p)
				//ctx.SetStatusCode(fasthttp.StatusBadRequest)
				//user.ResetPasswordHandler(ctx)
			case "/1/l":
				ctx.Response.Header.SetStatusCode(200)
				ctx.Response.Header.SetContentType("text/html")
				//UserLogin(ctx, rw)
				rowsCount := 10
				if rowsCount == 0 {
					rowsCount = 10
				}
				p := &templatesHTML.TablePage{
					Rows: generateRows(rowsCount),
				}
				templatesHTML.WritePageTemplate(ctx, p)
				///*
			default:
				ctx.Response.Header.SetStatusCode(200)
				ctx.Response.Header.SetContentType("text/html")
				/*
					names := []string{"Kate", "Go", "John", "Brad"}

					// qtc creates Write* function for each template function.
					// Such functions accept io.Writer as first parameter:
					var buf bytes.Buffer
					templates.WritePageTemplate(&buf, names)
					ctx.Response.SetBody(buf.Bytes())
				*/

				///*
				p := &templatesHTML.MainPage{
					CTX: ctx,
				}
				templatesHTML.WritePageTemplate(ctx, p)

			}
		}
	} else {
		/*
			resp.SetBody([]byte(`User-agent: *
			asda111sdsadasdasd ` + string(reqURI) + `
			Sitemap: https://bcz.com/sitemap.xml`))
		*/
		// caution: don't pass in c.GetResponse() as it return a copy of response
		//rw := adaptor.GetCompatResponseWriter(&ctx.Response)
		/*
				categories := []Category{
				{
					Name: "Books",
					URL:  "/books",
					Subcategories: []Subcategory{
						{Name: "Fiction", URL: "/books/fiction"},
						{Name: "Non-Fiction", URL: "/books/non-fiction"},
					},
				},
				// Additional categories...
			}
			//*/

		switch string(reqURI) {
		default:

			log.Printf("bngfduiou32hr2")
			ctx.Response.Header.SetStatusCode(200)
			ctx.Response.Header.SetContentType("text/html")
			/*
				names := []string{"Kate", "Go", "John", "Brad"}

				// qtc creates Write* function for each template function.
				// Such functions accept io.Writer as first parameter:
				var buf bytes.Buffer
				templates.WritePageTemplate(&buf, names)
				ctx.Response.SetBody(buf.Bytes())
			*/

			///*
			p := &templatesHTML.MainPage{
				CTX: ctx,
			}
			templatesHTML.WritePageTemplate(ctx, p)

			/*
				buffer := bytebufferpool.Get()

				_, err := m.Writer("text/html", buffer).Write(ctx.Response.Body())
				if err != nil {
					log.Printf("%s", err)
				} else {
					ctx.Response.SetBodyStream(buffer, buffer.Len())
				}

				bytebufferpool.Put(buffer)
			//*/

			//MinifyFilter("text/html", rw)
			/*
				if _, err := io.WriteString(rw, "<p class="message"> This HTTP response will be minified. </p>"); err != nil {
					panic(err)
				}
				if err := rw.Close(); err != nil {
					panic(err)
				}
			*/
			/*
				mem := arena.NewArena()
				o := arena.New[H](mem)

				//o.Num = 123 // incorrect: use after free
				///*
				o.Data, o.Err = m.Bytes("text/html", ctx.Response.Body())
				if o.Err != nil {
					log.Printf("minify1 %s", o.Err)
				} else {
					//log.Printf("o.Data = %s", o.Data)
					ctx.Response.SetBody(o.Data)
				}

				//*/
			/*
				//ctx.Response.SetBody(out)
				mem.Free()

				//rw = MinifyFilter("text/html", rw)
				//rw = MinifyFilter("text/html", rw)
				//fmt.Printf("buf=\n%s", buf.Bytes())

				/*
					var userList = []string{
						"Alice",
						"Bob",/Greet
						"Tom",
					}

					// Had better use buffer pool. Hero exports `GetBuffer` and `PutBuffer` for this.
					//
					// For convenience, hero also supports `io.Writer`. For example, you can also define
					// the function to `func UserList(userList []string, w io.Writer) (int, error)`,
					// and then:
					//
					//   template.UserList(userList, w)
					//
					buffer := new(bytes.Buffer)
					template.UserList(userList, buffer)
					rw.Write(buffer.Bytes())
			*/
			//qtc.WriteIndex(rw, "My Page Title", categories, currentYear)
			//templates.Greetings("my page title", categories, time.Now().Year())

			//		Index(cx.B2s(reqURI), nil, time.Now().Year(), rw)
			//JadeHeader(cx.B2s(reqURI), nil, rw)
			//JadeHeader(cx.B2s(reqURI), nil, rw)
		}

		//Index(string(reqURI), categories, rw)
		//Index("Hertz", rw)

		// No subdomain found
		fmt.Fprintf(ctx, "No subdomain found")
	}

}

func generateRows(rowsCount int) []string {
	var rows []string
	for i := 0; i < rowsCount; i++ {
		r := fmt.Sprintf("row %d", i)
		if rand.Intn(20) == 0 {
			r = "bingo"
		}
		rows = append(rows, r)
	}
	return rows
}

func createSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	h := blake3.New()
	_, _ = h.Write(append([]byte(password), salt...))
	hashedPassword := h.Sum(nil)
	return hex.EncodeToString(hashedPassword)
}

func checkPassword(password string, hashedPassword string, salt []byte) bool {
	return hashPassword(password, salt) == hashedPassword
}

