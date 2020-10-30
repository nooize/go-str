### go-str

tiny package to work with golang strings as flows
and remove spaghetti code from your string operations

classic style with _strings_ package

```golang
str := "SaMe bAd stRing"
str := strings.Title(strings.Replace(strings.ToLower(str), "bad", "good"))
```

__str__ style:

```golang
  str := str.New("SaMe bAd stRing").
          ToLower().
          Replace("bad", "good").
          Title().
          String() 
```

also __str__ has same additional functions:


