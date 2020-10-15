### go-str

tiny package to work with golang strings as flows
and remove spaghetti code fro yjouse operations

Classic style

 str := "SaMe bAd stRing"
 str := strings.Title(strings.Replace(strings.ToLower(str), "bad", "good"))

Str Style:

` 
  str := str.New("SaMe bAd stRing").
          ToLower().
          Replace("bad", "good").
          Title().
          String() 
          
`
