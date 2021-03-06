// Copyright 2017, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expr

import (
	"golang.org/x/net/html/atom"
)

var (
	A          = ElementChecker(atom.A)
	Abbr       = ElementChecker(atom.Abbr)
	Address    = ElementChecker(atom.Address)
	Area       = ElementChecker(atom.Area)
	Article    = ElementChecker(atom.Article)
	Aside      = ElementChecker(atom.Aside)
	Audio      = ElementChecker(atom.Audio)
	B          = ElementChecker(atom.B)
	Base       = ElementChecker(atom.Base)
	Bdi        = ElementChecker(atom.Bdi)
	Bdo        = ElementChecker(atom.Bdo)
	Blockquote = ElementChecker(atom.Blockquote)
	Body       = ElementChecker(atom.Body)
	Br         = ElementChecker(atom.Br)
	Button     = ElementChecker(atom.Button)
	Canvas     = ElementChecker(atom.Canvas)
	Caption    = ElementChecker(atom.Caption)
	Cite       = ElementChecker(atom.Cite)
	Code       = ElementChecker(atom.Code)
	Col        = ElementChecker(atom.Col)
	Colgroup   = ElementChecker(atom.Colgroup)
	Data       = ElementChecker(atom.Data)
	Datalist   = ElementChecker(atom.Datalist)
	Dd         = ElementChecker(atom.Dd)
	Del        = ElementChecker(atom.Del)
	Details    = ElementChecker(atom.Details)
	Dfn        = ElementChecker(atom.Dfn)
	Dialog     = ElementChecker(atom.Dialog)
	Div        = ElementChecker(atom.Div)
	Dl         = ElementChecker(atom.Dl)
	Dt         = ElementChecker(atom.Dt)
	Em         = ElementChecker(atom.Em)
	Embed      = ElementChecker(atom.Embed)
	Fieldset   = ElementChecker(atom.Fieldset)
	Figcaption = ElementChecker(atom.Figcaption)
	Figure     = ElementChecker(atom.Figure)
	Footer     = ElementChecker(atom.Footer)
	Form       = ElementChecker(atom.Form)
	H1         = ElementChecker(atom.H1)
	H2         = ElementChecker(atom.H2)
	H3         = ElementChecker(atom.H3)
	H4         = ElementChecker(atom.H4)
	H5         = ElementChecker(atom.H5)
	H6         = ElementChecker(atom.H6)
	Head       = ElementChecker(atom.Head)
	Header     = ElementChecker(atom.Header)
	Hgroup     = ElementChecker(atom.Hgroup)
	Hr         = ElementChecker(atom.Hr)
	Html       = ElementChecker(atom.Html)
	I          = ElementChecker(atom.I)
	Iframe     = ElementChecker(atom.Iframe)
	Img        = ElementChecker(atom.Img)
	Input      = ElementChecker(atom.Input)
	Ins        = ElementChecker(atom.Ins)
	Kbd        = ElementChecker(atom.Kbd)
	Label      = ElementChecker(atom.Label)
	Legend     = ElementChecker(atom.Legend)
	Li         = ElementChecker(atom.Li)
	Link       = ElementChecker(atom.Link)
	Map        = ElementChecker(atom.Map)
	Mark       = ElementChecker(atom.Mark)
	Menu       = ElementChecker(atom.Menu)
	Meta       = ElementChecker(atom.Meta)
	Meter      = ElementChecker(atom.Meter)
	Nav        = ElementChecker(atom.Nav)
	Noscript   = ElementChecker(atom.Noscript)
	Object     = ElementChecker(atom.Object)
	Ol         = ElementChecker(atom.Ol)
	Optgroup   = ElementChecker(atom.Optgroup)
	Option     = ElementChecker(atom.Option)
	Output     = ElementChecker(atom.Output)
	P          = ElementChecker(atom.P)
	Param      = ElementChecker(atom.Param)
	Pre        = ElementChecker(atom.Pre)
	Progress   = ElementChecker(atom.Progress)
	Q          = ElementChecker(atom.Q)
	Rp         = ElementChecker(atom.Rp)
	Rt         = ElementChecker(atom.Rt)
	Ruby       = ElementChecker(atom.Ruby)
	S          = ElementChecker(atom.S)
	Samp       = ElementChecker(atom.Samp)
	Script     = ElementChecker(atom.Script)
	Section    = ElementChecker(atom.Section)
	Select     = ElementChecker(atom.Select)
	Small      = ElementChecker(atom.Small)
	Source     = ElementChecker(atom.Source)
	Span       = ElementChecker(atom.Span)
	Strong     = ElementChecker(atom.Strong)
	Style      = ElementChecker(atom.Style)
	Sub        = ElementChecker(atom.Sub)
	Summary    = ElementChecker(atom.Summary)
	Sup        = ElementChecker(atom.Sup)
	Table      = ElementChecker(atom.Table)
	Tbody      = ElementChecker(atom.Tbody)
	Td         = ElementChecker(atom.Td)
	Textarea   = ElementChecker(atom.Textarea)
	Tfoot      = ElementChecker(atom.Tfoot)
	Th         = ElementChecker(atom.Th)
	Thead      = ElementChecker(atom.Thead)
	Time       = ElementChecker(atom.Time)
	Title      = ElementChecker(atom.Title)
	Tr         = ElementChecker(atom.Tr)
	Track      = ElementChecker(atom.Track)
	U          = ElementChecker(atom.U)
	Ul         = ElementChecker(atom.Ul)
	Var        = ElementChecker(atom.Var)
	Video      = ElementChecker(atom.Video)
	Wbr        = ElementChecker(atom.Wbr)
)

var (
	Abbr_               = AttrChecker("abbr")
	Accept              = AttrChecker("accept")
	AcceptCharset       = SeperatedAttrChecker("accept-charset", ' ')
	Accesskey           = SeperatedAttrChecker("accesskey", ' ')
	Action              = AttrChecker("action")
	Allowfullscreen     = AttrChecker("allowfullscreen")
	Allowpaymentrequest = AttrChecker("allowpaymentrequest")
	Allowusermedia      = AttrChecker("allowusermedia")
	Alt                 = AttrChecker("alt")
	As                  = AttrChecker("as")
	Async               = AttrChecker("async")
	Autocomplete        = AttrChecker("autocomplete")
	Autofocus           = AttrChecker("autofocus")
	Autoplay            = AttrChecker("autoplay")
	Charset             = AttrChecker("charset")
	Checked             = AttrChecker("checked")
	Cite_               = AttrChecker("cite")
	Class               = SeperatedAttrChecker("class", ' ')
	Color               = AttrChecker("color")
	Cols                = AttrChecker("cols")
	Colspan             = AttrChecker("colspan")
	Content             = AttrChecker("content")
	Contenteditable     = AttrChecker("contenteditable")
	Controls            = AttrChecker("controls")
	Coords              = AttrChecker("coords")
	Crossorigin         = AttrChecker("crossorigin")
	Data_               = AttrChecker("data")
	Datetime            = AttrChecker("datetime")
	Default             = AttrChecker("default")
	Defer               = AttrChecker("defer")
	Dir                 = AttrChecker("dir")
	Dirname             = AttrChecker("dirname")
	Disabled            = AttrChecker("disabled")
	Download            = AttrChecker("download")
	Draggable           = AttrChecker("draggable")
	Enctype             = AttrChecker("enctype")
	For                 = AttrChecker("for")
	Form_               = AttrChecker("form")
	Formaction          = AttrChecker("formaction")
	Formenctype         = AttrChecker("formenctype")
	Formmethod          = AttrChecker("formmethod")
	Formnovalidate      = AttrChecker("formnovalidate")
	Formtarget          = AttrChecker("formtarget")
	Headers             = SeperatedAttrChecker("headers", ' ')
	Height              = AttrChecker("height")
	Hidden              = AttrChecker("hidden")
	High                = AttrChecker("high")
	Href                = AttrChecker("href")
	Hreflang            = AttrChecker("hreflang")
	HttpEquiv           = AttrChecker("http-equiv")
	Id                  = AttrChecker("id")
	Inputmode           = AttrChecker("inputmode")
	Integrity           = AttrChecker("integrity")
	Is                  = AttrChecker("is")
	Ismap               = AttrChecker("ismap")
	Itemid              = AttrChecker("itemid")
	Itemprop            = SeperatedAttrChecker("itemprop", ' ')
	Itemref             = SeperatedAttrChecker("itemref", ' ')
	Itemscope           = AttrChecker("itemscope")
	Itemtype            = SeperatedAttrChecker("itemtype", ' ')
	Kind                = AttrChecker("kind")
	Label_              = AttrChecker("label")
	Lang                = AttrChecker("lang")
	List                = AttrChecker("list")
	Loop                = AttrChecker("loop")
	Low                 = AttrChecker("low")
	Manifest            = AttrChecker("manifest")
	Max                 = AttrChecker("max")
	Maxlength           = AttrChecker("maxlength")
	Media               = AttrChecker("media")
	Method              = AttrChecker("method")
	Min                 = AttrChecker("min")
	Minlength           = AttrChecker("minlength")
	Multiple            = AttrChecker("multiple")
	Muted               = AttrChecker("muted")
	Name                = AttrChecker("name")
	Nomodule            = AttrChecker("nomodule")
	Nonce               = AttrChecker("nonce")
	Novalidate          = AttrChecker("novalidate")
	Open                = AttrChecker("open")
	Optimum             = AttrChecker("optimum")
	Pattern             = AttrChecker("pattern")
	Ping                = SeperatedAttrChecker("ping", ' ')
	Placeholder         = AttrChecker("placeholder")
	Playsinline         = AttrChecker("playsinline")
	Poster              = AttrChecker("poster")
	Preload             = AttrChecker("preload")
	Readonly            = AttrChecker("readonly")
	Referrerpolicy      = AttrChecker("referrerpolicy")
	Rel                 = SeperatedAttrChecker("rel", ' ')
	Required            = AttrChecker("required")
	Reversed            = AttrChecker("reversed")
	Rows                = AttrChecker("rows")
	Rowspan             = AttrChecker("rowspan")
	Sandbox             = SeperatedAttrChecker("sandbox", ' ')
	Scope               = AttrChecker("scope")
	Selected            = AttrChecker("selected")
	Shape               = AttrChecker("shape")
	Size                = AttrChecker("size")
	Sizes               = SeperatedAttrChecker("sizes", ' ')
	Slot_               = AttrChecker("slot")
	Span_               = AttrChecker("span")
	Spellcheck          = AttrChecker("spellcheck")
	Src                 = AttrChecker("src")
	Srcdoc              = AttrChecker("srcdoc")
	Srclang             = AttrChecker("srclang")
	Srcset              = AttrChecker("srcset")
	Start               = AttrChecker("start")
	Step                = AttrChecker("step")
	Style_              = AttrChecker("style")
	Tabindex            = AttrChecker("tabindex")
	Target              = AttrChecker("target")
	Title_              = AttrChecker("title")
	Translate           = AttrChecker("translate")
	Type                = AttrChecker("type")
	Typemustmatch       = AttrChecker("typemustmatch")
	Updateviacache      = AttrChecker("updateviacache")
	Usemap              = AttrChecker("usemap")
	Value               = AttrChecker("value")
	Width               = AttrChecker("width")
	Workertype          = AttrChecker("workertype")
	Wrap                = AttrChecker("wrap")
)
