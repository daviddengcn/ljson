LJSON [![GoSearch](http://go-search.org/badge?id=github.com%2Fdaviddengcn%2Fljson)](http://go-search.org/view?id=github.com%2Fdaviddengcn%2Fljson)
========

Loose JSON
----------
Loose JSON(LJSON) is a superset of JSON which allows some imperfect formatting(i.e. the marshaller should be more robust).

Format differences to standard JSON:

A standard JSON data:

```javascript
{
    "color": "red",
    "elems": [1,2,3,4],
    "weight": 10
}
```


1) Redundant commas are allowed in object and array definition. e.g.

```javascript
{
    "color": "red",,,
    "elems": [1,,,,2,3,4,,,],
    "weight": 10,
}
```

2) Commas can be omitted in definition of object and array. e.g.

```javascript
{
    "color": "red"
    "elems": [1 2 3 4]
    "weight": 10
}
```

3) Support naked-key. If the key part of an element in an object, which is a string, 
contains naked-key-valid characters only, the embracing double quotes can be omitted(the key becomes a naked-key).
naked-key-valid characters are printable characters other than the following:
`"`, `'`, `:`, `[`, `]`, `{`, `}`, `\`, and `,`. (This is similar to grammar in javascript)

```javascript
{
    color : "red"
    elems : [1 2 3 4]
    weight: 10
}
```

This is useful especially when the data is written by hand, e.g. as a configure file.

Implemenation
-------------
This package implements decoder for LJSON. (For encoding, use build-in `encoding/json` package. Visit [Godoc for LJSON](http://godoc.org/github.com/daviddengcn/ljson) for more usage information.

The code is modified from `encoding/json` package in [go 1.0.3](https://code.google.com/p/go/source/browse/?name=go1.0.3#hg%2Fsrc%2Fpkg%2Fencoding%2Fjson).

LICENSE
-------
This library is distributed under BSD license.
