# Diffie Hellman Library

### "github.com/aead/ecdh" simple wrapper

> this library is just a very simple wrapper around ecdh so u can use it in your projects

```go
import (
    "fmt"
    "github.com/husseinraed/diffie"
)

var d1 diffie.DiffieHellman
var d2 diffie.DiffieHellman

d1_publickey, _ := d1.GenPublicKey()
d2_publickey, _ := d2.GenPublicKey()

d1_sharedkey, _ := d1.GenSharedKey(d2_publickey)
d2_sharedkey, _ := d2.GenSharedKey(d1_publickey)

fmt.Println(d1_sharedkey == d2_sharedkey)
```
```
output:

true
```
