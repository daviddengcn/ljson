LJSON
========

Loose JSON
----------
Loose JSON(LJSON) is a superset of JSON which allows some imperfect formatting(i.e. the marshaller should be more robust).

Format differences to standard JSON:

(A standard JSON data:

```javascript
{
    "color": "red",
    "elems": [1,2,3,4],
    "weight": 10
}
```
)

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

This is useful especially when the data is written by hand, e.g. as a configure file.

Decoder
-------
This package implements decoder for LJSON. (For encoding, use build-in `encodeing/json` package.

[Godoc for LJSON](http://godoc.org/github.com/daviddengcn/ljson)
