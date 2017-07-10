// Copyright 2017, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package query

import (
	. "h12.me/html-query/expr"
)

func (n *Node) A(cs ...Checker) *Node {
	return n.find(A, cs)
}

func (n *Node) Abbr(cs ...Checker) *Node {
	return n.find(Abbr, cs)
}

func (n *Node) Address(cs ...Checker) *Node {
	return n.find(Address, cs)
}

func (n *Node) Area(cs ...Checker) *Node {
	return n.find(Area, cs)
}

func (n *Node) Article(cs ...Checker) *Node {
	return n.find(Article, cs)
}

func (n *Node) Aside(cs ...Checker) *Node {
	return n.find(Aside, cs)
}

func (n *Node) Audio(cs ...Checker) *Node {
	return n.find(Audio, cs)
}

func (n *Node) B(cs ...Checker) *Node {
	return n.find(B, cs)
}

func (n *Node) Base(cs ...Checker) *Node {
	return n.find(Base, cs)
}

func (n *Node) Bdi(cs ...Checker) *Node {
	return n.find(Bdi, cs)
}

func (n *Node) Bdo(cs ...Checker) *Node {
	return n.find(Bdo, cs)
}

func (n *Node) Blockquote(cs ...Checker) *Node {
	return n.find(Blockquote, cs)
}

func (n *Node) Body(cs ...Checker) *Node {
	return n.find(Body, cs)
}

func (n *Node) Br(cs ...Checker) *Node {
	return n.find(Br, cs)
}

func (n *Node) Button(cs ...Checker) *Node {
	return n.find(Button, cs)
}

func (n *Node) Canvas(cs ...Checker) *Node {
	return n.find(Canvas, cs)
}

func (n *Node) Caption(cs ...Checker) *Node {
	return n.find(Caption, cs)
}

func (n *Node) Cite(cs ...Checker) *Node {
	return n.find(Cite, cs)
}

func (n *Node) Code(cs ...Checker) *Node {
	return n.find(Code, cs)
}

func (n *Node) Col(cs ...Checker) *Node {
	return n.find(Col, cs)
}

func (n *Node) Colgroup(cs ...Checker) *Node {
	return n.find(Colgroup, cs)
}

func (n *Node) Data(cs ...Checker) *Node {
	return n.find(Data, cs)
}

func (n *Node) Datalist(cs ...Checker) *Node {
	return n.find(Datalist, cs)
}

func (n *Node) Dd(cs ...Checker) *Node {
	return n.find(Dd, cs)
}

func (n *Node) Del(cs ...Checker) *Node {
	return n.find(Del, cs)
}

func (n *Node) Details(cs ...Checker) *Node {
	return n.find(Details, cs)
}

func (n *Node) Dfn(cs ...Checker) *Node {
	return n.find(Dfn, cs)
}

func (n *Node) Dialog(cs ...Checker) *Node {
	return n.find(Dialog, cs)
}

func (n *Node) Div(cs ...Checker) *Node {
	return n.find(Div, cs)
}

func (n *Node) Dl(cs ...Checker) *Node {
	return n.find(Dl, cs)
}

func (n *Node) Dt(cs ...Checker) *Node {
	return n.find(Dt, cs)
}

func (n *Node) Em(cs ...Checker) *Node {
	return n.find(Em, cs)
}

func (n *Node) Embed(cs ...Checker) *Node {
	return n.find(Embed, cs)
}

func (n *Node) Fieldset(cs ...Checker) *Node {
	return n.find(Fieldset, cs)
}

func (n *Node) Figcaption(cs ...Checker) *Node {
	return n.find(Figcaption, cs)
}

func (n *Node) Figure(cs ...Checker) *Node {
	return n.find(Figure, cs)
}

func (n *Node) Footer(cs ...Checker) *Node {
	return n.find(Footer, cs)
}

func (n *Node) Form(cs ...Checker) *Node {
	return n.find(Form, cs)
}

func (n *Node) H1(cs ...Checker) *Node {
	return n.find(H1, cs)
}

func (n *Node) H2(cs ...Checker) *Node {
	return n.find(H2, cs)
}

func (n *Node) H3(cs ...Checker) *Node {
	return n.find(H3, cs)
}

func (n *Node) H4(cs ...Checker) *Node {
	return n.find(H4, cs)
}

func (n *Node) H5(cs ...Checker) *Node {
	return n.find(H5, cs)
}

func (n *Node) H6(cs ...Checker) *Node {
	return n.find(H6, cs)
}

func (n *Node) Head(cs ...Checker) *Node {
	return n.find(Head, cs)
}

func (n *Node) Header(cs ...Checker) *Node {
	return n.find(Header, cs)
}

func (n *Node) Hgroup(cs ...Checker) *Node {
	return n.find(Hgroup, cs)
}

func (n *Node) Hr(cs ...Checker) *Node {
	return n.find(Hr, cs)
}

func (n *Node) Html(cs ...Checker) *Node {
	return n.find(Html, cs)
}

func (n *Node) I(cs ...Checker) *Node {
	return n.find(I, cs)
}

func (n *Node) Iframe(cs ...Checker) *Node {
	return n.find(Iframe, cs)
}

func (n *Node) Img(cs ...Checker) *Node {
	return n.find(Img, cs)
}

func (n *Node) Input(cs ...Checker) *Node {
	return n.find(Input, cs)
}

func (n *Node) Ins(cs ...Checker) *Node {
	return n.find(Ins, cs)
}

func (n *Node) Kbd(cs ...Checker) *Node {
	return n.find(Kbd, cs)
}

func (n *Node) Label(cs ...Checker) *Node {
	return n.find(Label, cs)
}

func (n *Node) Legend(cs ...Checker) *Node {
	return n.find(Legend, cs)
}

func (n *Node) Li(cs ...Checker) *Node {
	return n.find(Li, cs)
}

func (n *Node) Link(cs ...Checker) *Node {
	return n.find(Link, cs)
}

func (n *Node) Map(cs ...Checker) *Node {
	return n.find(Map, cs)
}

func (n *Node) Mark(cs ...Checker) *Node {
	return n.find(Mark, cs)
}

func (n *Node) Menu(cs ...Checker) *Node {
	return n.find(Menu, cs)
}

func (n *Node) Meta(cs ...Checker) *Node {
	return n.find(Meta, cs)
}

func (n *Node) Meter(cs ...Checker) *Node {
	return n.find(Meter, cs)
}

func (n *Node) Nav(cs ...Checker) *Node {
	return n.find(Nav, cs)
}

func (n *Node) Noscript(cs ...Checker) *Node {
	return n.find(Noscript, cs)
}

func (n *Node) Object(cs ...Checker) *Node {
	return n.find(Object, cs)
}

func (n *Node) Ol(cs ...Checker) *Node {
	return n.find(Ol, cs)
}

func (n *Node) Optgroup(cs ...Checker) *Node {
	return n.find(Optgroup, cs)
}

func (n *Node) Option(cs ...Checker) *Node {
	return n.find(Option, cs)
}

func (n *Node) Output(cs ...Checker) *Node {
	return n.find(Output, cs)
}

func (n *Node) P(cs ...Checker) *Node {
	return n.find(P, cs)
}

func (n *Node) Param(cs ...Checker) *Node {
	return n.find(Param, cs)
}

func (n *Node) Pre(cs ...Checker) *Node {
	return n.find(Pre, cs)
}

func (n *Node) Progress(cs ...Checker) *Node {
	return n.find(Progress, cs)
}

func (n *Node) Q(cs ...Checker) *Node {
	return n.find(Q, cs)
}

func (n *Node) Rp(cs ...Checker) *Node {
	return n.find(Rp, cs)
}

func (n *Node) Rt(cs ...Checker) *Node {
	return n.find(Rt, cs)
}

func (n *Node) Ruby(cs ...Checker) *Node {
	return n.find(Ruby, cs)
}

func (n *Node) S(cs ...Checker) *Node {
	return n.find(S, cs)
}

func (n *Node) Samp(cs ...Checker) *Node {
	return n.find(Samp, cs)
}

func (n *Node) Script(cs ...Checker) *Node {
	return n.find(Script, cs)
}

func (n *Node) Section(cs ...Checker) *Node {
	return n.find(Section, cs)
}

func (n *Node) Select(cs ...Checker) *Node {
	return n.find(Select, cs)
}

func (n *Node) Small(cs ...Checker) *Node {
	return n.find(Small, cs)
}

func (n *Node) Source(cs ...Checker) *Node {
	return n.find(Source, cs)
}

func (n *Node) Span(cs ...Checker) *Node {
	return n.find(Span, cs)
}

func (n *Node) Strong(cs ...Checker) *Node {
	return n.find(Strong, cs)
}

func (n *Node) Style(cs ...Checker) *Node {
	return n.find(Style, cs)
}

func (n *Node) Sub(cs ...Checker) *Node {
	return n.find(Sub, cs)
}

func (n *Node) Summary(cs ...Checker) *Node {
	return n.find(Summary, cs)
}

func (n *Node) Sup(cs ...Checker) *Node {
	return n.find(Sup, cs)
}

func (n *Node) Table(cs ...Checker) *Node {
	return n.find(Table, cs)
}

func (n *Node) Tbody(cs ...Checker) *Node {
	return n.find(Tbody, cs)
}

func (n *Node) Td(cs ...Checker) *Node {
	return n.find(Td, cs)
}

func (n *Node) Textarea(cs ...Checker) *Node {
	return n.find(Textarea, cs)
}

func (n *Node) Tfoot(cs ...Checker) *Node {
	return n.find(Tfoot, cs)
}

func (n *Node) Th(cs ...Checker) *Node {
	return n.find(Th, cs)
}

func (n *Node) Thead(cs ...Checker) *Node {
	return n.find(Thead, cs)
}

func (n *Node) Time(cs ...Checker) *Node {
	return n.find(Time, cs)
}

func (n *Node) Title(cs ...Checker) *Node {
	return n.find(Title, cs)
}

func (n *Node) Tr(cs ...Checker) *Node {
	return n.find(Tr, cs)
}

func (n *Node) Track(cs ...Checker) *Node {
	return n.find(Track, cs)
}

func (n *Node) U(cs ...Checker) *Node {
	return n.find(U, cs)
}

func (n *Node) Ul(cs ...Checker) *Node {
	return n.find(Ul, cs)
}

func (n *Node) Var(cs ...Checker) *Node {
	return n.find(Var, cs)
}

func (n *Node) Video(cs ...Checker) *Node {
	return n.find(Video, cs)
}

func (n *Node) Wbr(cs ...Checker) *Node {
	return n.find(Wbr, cs)
}

func (n *Node) Abbr_(pat ...string) *string {
	return n.Attr("abbr", pat...)
}

func (n *Node) Accept(pat ...string) *string {
	return n.Attr("accept", pat...)
}

func (n *Node) AcceptCharset(pat ...string) *string {
	return n.Attr("accept-charset", pat...)
}

func (n *Node) Accesskey(pat ...string) *string {
	return n.Attr("accesskey", pat...)
}

func (n *Node) Action(pat ...string) *string {
	return n.Attr("action", pat...)
}

func (n *Node) Allowfullscreen(pat ...string) *string {
	return n.Attr("allowfullscreen", pat...)
}

func (n *Node) Allowpaymentrequest(pat ...string) *string {
	return n.Attr("allowpaymentrequest", pat...)
}

func (n *Node) Allowusermedia(pat ...string) *string {
	return n.Attr("allowusermedia", pat...)
}

func (n *Node) Alt(pat ...string) *string {
	return n.Attr("alt", pat...)
}

func (n *Node) As(pat ...string) *string {
	return n.Attr("as", pat...)
}

func (n *Node) Async(pat ...string) *string {
	return n.Attr("async", pat...)
}

func (n *Node) Autocomplete(pat ...string) *string {
	return n.Attr("autocomplete", pat...)
}

func (n *Node) Autofocus(pat ...string) *string {
	return n.Attr("autofocus", pat...)
}

func (n *Node) Autoplay(pat ...string) *string {
	return n.Attr("autoplay", pat...)
}

func (n *Node) Charset(pat ...string) *string {
	return n.Attr("charset", pat...)
}

func (n *Node) Checked(pat ...string) *string {
	return n.Attr("checked", pat...)
}

func (n *Node) Cite_(pat ...string) *string {
	return n.Attr("cite", pat...)
}

func (n *Node) Class(pat ...string) *string {
	return n.Attr("class", pat...)
}

func (n *Node) Color(pat ...string) *string {
	return n.Attr("color", pat...)
}

func (n *Node) Cols(pat ...string) *string {
	return n.Attr("cols", pat...)
}

func (n *Node) Colspan(pat ...string) *string {
	return n.Attr("colspan", pat...)
}

func (n *Node) Content(pat ...string) *string {
	return n.Attr("content", pat...)
}

func (n *Node) Contenteditable(pat ...string) *string {
	return n.Attr("contenteditable", pat...)
}

func (n *Node) Controls(pat ...string) *string {
	return n.Attr("controls", pat...)
}

func (n *Node) Coords(pat ...string) *string {
	return n.Attr("coords", pat...)
}

func (n *Node) Crossorigin(pat ...string) *string {
	return n.Attr("crossorigin", pat...)
}

func (n *Node) Data_(pat ...string) *string {
	return n.Attr("data", pat...)
}

func (n *Node) Datetime(pat ...string) *string {
	return n.Attr("datetime", pat...)
}

func (n *Node) Default(pat ...string) *string {
	return n.Attr("default", pat...)
}

func (n *Node) Defer(pat ...string) *string {
	return n.Attr("defer", pat...)
}

func (n *Node) Dir(pat ...string) *string {
	return n.Attr("dir", pat...)
}

func (n *Node) Dirname(pat ...string) *string {
	return n.Attr("dirname", pat...)
}

func (n *Node) Disabled(pat ...string) *string {
	return n.Attr("disabled", pat...)
}

func (n *Node) Download(pat ...string) *string {
	return n.Attr("download", pat...)
}

func (n *Node) Draggable(pat ...string) *string {
	return n.Attr("draggable", pat...)
}

func (n *Node) Enctype(pat ...string) *string {
	return n.Attr("enctype", pat...)
}

func (n *Node) For(pat ...string) *string {
	return n.Attr("for", pat...)
}

func (n *Node) Form_(pat ...string) *string {
	return n.Attr("form", pat...)
}

func (n *Node) Formaction(pat ...string) *string {
	return n.Attr("formaction", pat...)
}

func (n *Node) Formenctype(pat ...string) *string {
	return n.Attr("formenctype", pat...)
}

func (n *Node) Formmethod(pat ...string) *string {
	return n.Attr("formmethod", pat...)
}

func (n *Node) Formnovalidate(pat ...string) *string {
	return n.Attr("formnovalidate", pat...)
}

func (n *Node) Formtarget(pat ...string) *string {
	return n.Attr("formtarget", pat...)
}

func (n *Node) Headers(pat ...string) *string {
	return n.Attr("headers", pat...)
}

func (n *Node) Height(pat ...string) *string {
	return n.Attr("height", pat...)
}

func (n *Node) Hidden(pat ...string) *string {
	return n.Attr("hidden", pat...)
}

func (n *Node) High(pat ...string) *string {
	return n.Attr("high", pat...)
}

func (n *Node) Href(pat ...string) *string {
	return n.Attr("href", pat...)
}

func (n *Node) Hreflang(pat ...string) *string {
	return n.Attr("hreflang", pat...)
}

func (n *Node) HttpEquiv(pat ...string) *string {
	return n.Attr("http-equiv", pat...)
}

func (n *Node) Id(pat ...string) *string {
	return n.Attr("id", pat...)
}

func (n *Node) Inputmode(pat ...string) *string {
	return n.Attr("inputmode", pat...)
}

func (n *Node) Integrity(pat ...string) *string {
	return n.Attr("integrity", pat...)
}

func (n *Node) Is(pat ...string) *string {
	return n.Attr("is", pat...)
}

func (n *Node) Ismap(pat ...string) *string {
	return n.Attr("ismap", pat...)
}

func (n *Node) Itemid(pat ...string) *string {
	return n.Attr("itemid", pat...)
}

func (n *Node) Itemprop(pat ...string) *string {
	return n.Attr("itemprop", pat...)
}

func (n *Node) Itemref(pat ...string) *string {
	return n.Attr("itemref", pat...)
}

func (n *Node) Itemscope(pat ...string) *string {
	return n.Attr("itemscope", pat...)
}

func (n *Node) Itemtype(pat ...string) *string {
	return n.Attr("itemtype", pat...)
}

func (n *Node) Kind(pat ...string) *string {
	return n.Attr("kind", pat...)
}

func (n *Node) Label_(pat ...string) *string {
	return n.Attr("label", pat...)
}

func (n *Node) Lang(pat ...string) *string {
	return n.Attr("lang", pat...)
}

func (n *Node) List(pat ...string) *string {
	return n.Attr("list", pat...)
}

func (n *Node) Loop(pat ...string) *string {
	return n.Attr("loop", pat...)
}

func (n *Node) Low(pat ...string) *string {
	return n.Attr("low", pat...)
}

func (n *Node) Manifest(pat ...string) *string {
	return n.Attr("manifest", pat...)
}

func (n *Node) Max(pat ...string) *string {
	return n.Attr("max", pat...)
}

func (n *Node) Maxlength(pat ...string) *string {
	return n.Attr("maxlength", pat...)
}

func (n *Node) Media(pat ...string) *string {
	return n.Attr("media", pat...)
}

func (n *Node) Method(pat ...string) *string {
	return n.Attr("method", pat...)
}

func (n *Node) Min(pat ...string) *string {
	return n.Attr("min", pat...)
}

func (n *Node) Minlength(pat ...string) *string {
	return n.Attr("minlength", pat...)
}

func (n *Node) Multiple(pat ...string) *string {
	return n.Attr("multiple", pat...)
}

func (n *Node) Muted(pat ...string) *string {
	return n.Attr("muted", pat...)
}

func (n *Node) Name(pat ...string) *string {
	return n.Attr("name", pat...)
}

func (n *Node) Nomodule(pat ...string) *string {
	return n.Attr("nomodule", pat...)
}

func (n *Node) Nonce(pat ...string) *string {
	return n.Attr("nonce", pat...)
}

func (n *Node) Novalidate(pat ...string) *string {
	return n.Attr("novalidate", pat...)
}

func (n *Node) Open(pat ...string) *string {
	return n.Attr("open", pat...)
}

func (n *Node) Optimum(pat ...string) *string {
	return n.Attr("optimum", pat...)
}

func (n *Node) Pattern(pat ...string) *string {
	return n.Attr("pattern", pat...)
}

func (n *Node) Ping(pat ...string) *string {
	return n.Attr("ping", pat...)
}

func (n *Node) Placeholder(pat ...string) *string {
	return n.Attr("placeholder", pat...)
}

func (n *Node) Playsinline(pat ...string) *string {
	return n.Attr("playsinline", pat...)
}

func (n *Node) Poster(pat ...string) *string {
	return n.Attr("poster", pat...)
}

func (n *Node) Preload(pat ...string) *string {
	return n.Attr("preload", pat...)
}

func (n *Node) Readonly(pat ...string) *string {
	return n.Attr("readonly", pat...)
}

func (n *Node) Referrerpolicy(pat ...string) *string {
	return n.Attr("referrerpolicy", pat...)
}

func (n *Node) Rel(pat ...string) *string {
	return n.Attr("rel", pat...)
}

func (n *Node) Required(pat ...string) *string {
	return n.Attr("required", pat...)
}

func (n *Node) Reversed(pat ...string) *string {
	return n.Attr("reversed", pat...)
}

func (n *Node) Rows(pat ...string) *string {
	return n.Attr("rows", pat...)
}

func (n *Node) Rowspan(pat ...string) *string {
	return n.Attr("rowspan", pat...)
}

func (n *Node) Sandbox(pat ...string) *string {
	return n.Attr("sandbox", pat...)
}

func (n *Node) Scope(pat ...string) *string {
	return n.Attr("scope", pat...)
}

func (n *Node) Selected(pat ...string) *string {
	return n.Attr("selected", pat...)
}

func (n *Node) Shape(pat ...string) *string {
	return n.Attr("shape", pat...)
}

func (n *Node) Size(pat ...string) *string {
	return n.Attr("size", pat...)
}

func (n *Node) Sizes(pat ...string) *string {
	return n.Attr("sizes", pat...)
}

func (n *Node) Slot_(pat ...string) *string {
	return n.Attr("slot", pat...)
}

func (n *Node) Span_(pat ...string) *string {
	return n.Attr("span", pat...)
}

func (n *Node) Spellcheck(pat ...string) *string {
	return n.Attr("spellcheck", pat...)
}

func (n *Node) Src(pat ...string) *string {
	return n.Attr("src", pat...)
}

func (n *Node) Srcdoc(pat ...string) *string {
	return n.Attr("srcdoc", pat...)
}

func (n *Node) Srclang(pat ...string) *string {
	return n.Attr("srclang", pat...)
}

func (n *Node) Srcset(pat ...string) *string {
	return n.Attr("srcset", pat...)
}

func (n *Node) Start(pat ...string) *string {
	return n.Attr("start", pat...)
}

func (n *Node) Step(pat ...string) *string {
	return n.Attr("step", pat...)
}

func (n *Node) Style_(pat ...string) *string {
	return n.Attr("style", pat...)
}

func (n *Node) Tabindex(pat ...string) *string {
	return n.Attr("tabindex", pat...)
}

func (n *Node) Target(pat ...string) *string {
	return n.Attr("target", pat...)
}

func (n *Node) Title_(pat ...string) *string {
	return n.Attr("title", pat...)
}

func (n *Node) Translate(pat ...string) *string {
	return n.Attr("translate", pat...)
}

func (n *Node) Type(pat ...string) *string {
	return n.Attr("type", pat...)
}

func (n *Node) Typemustmatch(pat ...string) *string {
	return n.Attr("typemustmatch", pat...)
}

func (n *Node) Updateviacache(pat ...string) *string {
	return n.Attr("updateviacache", pat...)
}

func (n *Node) Usemap(pat ...string) *string {
	return n.Attr("usemap", pat...)
}

func (n *Node) Value(pat ...string) *string {
	return n.Attr("value", pat...)
}

func (n *Node) Width(pat ...string) *string {
	return n.Attr("width", pat...)
}

func (n *Node) Workertype(pat ...string) *string {
	return n.Attr("workertype", pat...)
}

func (n *Node) Wrap(pat ...string) *string {
	return n.Attr("wrap", pat...)
}
