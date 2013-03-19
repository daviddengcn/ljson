ljson
========

Loose JSON is a superset of JSON which allows some imperfect format.

Format differences to standard JSON:

1) Redundant commas are allow in object and array definition. e.g.

```javascript
{
    "color": "red",,,
    "elems": [1,,,,2,3,4,,,],
    "weight": 10,
}
```

2) Commas can be omitted in definition of object. e.g.

```javascript
{
    "color": "red"
    "elems": [1,2,3,4]
    "weight": 10
}
```
