html-query: A fluent and functional approach to querying HTML
=============================================================

html-query is a Go package that provides a fluent and functional interface for
querying HTML. It is based on code.google.com/p/go.net/html.

Examples
========
1. A simple example (under "examples" directory)
```
    r := get(`http://blog.golang.org/index`)
    defer r.Close()
    root, err := query.Parse(r)
    checkError(err)
    root.Div(Id("content")).Children(Class("blogtitle")).For(func(item *query.Node) {
        href := item.Ahref().Href()
        date := item.Span(Class("date")).Text()
        tags := item.Span(Class("tags")).Text()
        // ......
    })
```

2. Generator of html-query (under "gen" directory)

   A large part of html-query is automatically generated from HTML spec. The
spec is in HTML format. So the generator parses it using html-query itself.

Design
======
Here is a simple explanation of the design of html-query.
###Functional query expressions
All functional definitions are defined in html-query/expr package.

1. Checker and checker composition

   A checker is a function that accept and conditionally returns a *html.Node.
```
    type Checker func(*html.Node) *html.Node
```
   Here are some checker examples:
```
    Id("id1")
    Class("c1")
    Div
    Abbr
    H1
    H2
```
   Checkers can be combined as boolean expressions:
```
    And(Id("id1"), Class("c1"))
    Or(Class("c1"), Class("c2"))
    And(Class("c1"), Not(Class("c2")))
```
2. Checker builder

   A checker builder is a function that returns a checker. "Id", "Class", "And",
   "Or", "Not" shown above are all checker builders. There are also some checker
   builder builder (function that returns a checker builder) defined in
   html-query when needed.

###Fluent interface
Fluent interface (http://en.wikipedia.org/wiki/Fluent_interface) are defined in
html-query package.

1. Root node

   Function Parse returns the root node of an html document.

2. Node finder

   Method Node.Find implements a BFS search for a node, e.g.
```
    node.Find(Div, Class("id1"))
```
   But usually you can write the short form:
```
    node.Div(Class("id1"))
```
3. Attribute getter

   Method Node.Attr can be used to get the value (or a regular expression
   submatch of the value) of a node, e.g.
```
    node.Attr("Id")
    node.Attr("href", "\(.*)")
```
   But usually you can write the short form:
```
    node.Id()
    node.Href("\(.*)")
```
4. Node iterator

   Method Node.Children and Node.Descendants each returns a node iterator
   (NodeIter). Method NodeIter.For can be used to loop through these nodes.

Alternative
===========
If you prefer a jquery like DSL rather than functional way, you might want to
try goquery: https://github.com/PuerkitoBio/goquery.
