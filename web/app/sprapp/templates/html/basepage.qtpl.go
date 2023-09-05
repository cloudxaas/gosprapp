// Code generated by qtc from "basepage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.
//

//line templates/html/basepage.qtpl:3
package html

//line templates/html/basepage.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/html/basepage.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/html/basepage.qtpl:4
type Page interface {
//line templates/html/basepage.qtpl:4
	Title() string
//line templates/html/basepage.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
//line templates/html/basepage.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
//line templates/html/basepage.qtpl:4
	Body() string
//line templates/html/basepage.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
//line templates/html/basepage.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line templates/html/basepage.qtpl:4
}

// Page prints a page implementing Page interface.

//line templates/html/basepage.qtpl:12
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line templates/html/basepage.qtpl:12
	qw422016.N().S(`
<html>
	<head>
		<title>`)
//line templates/html/basepage.qtpl:15
	p.StreamTitle(qw422016)
//line templates/html/basepage.qtpl:15
	qw422016.N().S(`</title>
		<link rel="stylesheet" type="text/css" href="/m/a.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	</head>
	<body>
  <!-- Navigation -->
  <nav class="noprint">
    <ul>
      <li>
        <a href="/"><img src="/m/s.png" width=120></a>
      </li>
      <li class="float-right sticky"><a onclick="addFontSize(-1)">ᴀ-</a>|<a onclick="addFontSize(1)">A+</a></li>
      <li class="float-right sticky"><a onclick="toggleDarkMode(this)">☪</a></li>
      <li><a href="#demo">Demo</a></li>
      <li><a href="#basic">Reference ▾</a>
        <ul>
          <li><a href="#basic">Basic Style</a></li>
          <li><a href="#extra">Extra Style</a></li>
          <li><a href="#bootless">Classes</a></li>
          <li><a href="#sec-addons">Add-Ons</a></li>
        </ul>
      </li>
      <li class="float-right"><a href="https://github.com/emareg/classlesscss">@GitHub</a></li>
    </ul>
  </nav>
		`)
//line templates/html/basepage.qtpl:40
	p.StreamBody(qw422016)
//line templates/html/basepage.qtpl:40
	qw422016.N().S(`


<br>&nbsp;<br>
<hr>
<small>
&copy; 2023 SPRAPP. All Rights Reserved.
      <div class="float-right">Help :: Privacy :: Terms</div>
</small>



	</body>
</html>
`)
//line templates/html/basepage.qtpl:54
}

//line templates/html/basepage.qtpl:54
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line templates/html/basepage.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/html/basepage.qtpl:54
	StreamPageTemplate(qw422016, p)
//line templates/html/basepage.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line templates/html/basepage.qtpl:54
}

//line templates/html/basepage.qtpl:54
func PageTemplate(p Page) string {
//line templates/html/basepage.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/html/basepage.qtpl:54
	WritePageTemplate(qb422016, p)
//line templates/html/basepage.qtpl:54
	qs422016 := string(qb422016.B)
//line templates/html/basepage.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/html/basepage.qtpl:54
	return qs422016
//line templates/html/basepage.qtpl:54
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line templates/html/basepage.qtpl:59
type BasePage struct{}

//line templates/html/basepage.qtpl:60
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/html/basepage.qtpl:60
	qw422016.N().S(`This is a base title`)
//line templates/html/basepage.qtpl:60
}

//line templates/html/basepage.qtpl:60
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/html/basepage.qtpl:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/html/basepage.qtpl:60
	p.StreamTitle(qw422016)
//line templates/html/basepage.qtpl:60
	qt422016.ReleaseWriter(qw422016)
//line templates/html/basepage.qtpl:60
}

//line templates/html/basepage.qtpl:60
func (p *BasePage) Title() string {
//line templates/html/basepage.qtpl:60
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/html/basepage.qtpl:60
	p.WriteTitle(qb422016)
//line templates/html/basepage.qtpl:60
	qs422016 := string(qb422016.B)
//line templates/html/basepage.qtpl:60
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/html/basepage.qtpl:60
	return qs422016
//line templates/html/basepage.qtpl:60
}

//line templates/html/basepage.qtpl:61
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/html/basepage.qtpl:61
	qw422016.N().S(`This is a base body`)
//line templates/html/basepage.qtpl:61
}

//line templates/html/basepage.qtpl:61
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/html/basepage.qtpl:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/html/basepage.qtpl:61
	p.StreamBody(qw422016)
//line templates/html/basepage.qtpl:61
	qt422016.ReleaseWriter(qw422016)
//line templates/html/basepage.qtpl:61
}

//line templates/html/basepage.qtpl:61
func (p *BasePage) Body() string {
//line templates/html/basepage.qtpl:61
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/html/basepage.qtpl:61
	p.WriteBody(qb422016)
//line templates/html/basepage.qtpl:61
	qs422016 := string(qb422016.B)
//line templates/html/basepage.qtpl:61
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/html/basepage.qtpl:61
	return qs422016
//line templates/html/basepage.qtpl:61
}