# Diffie Hellman Library

### "github.com/aead/ecdh" simple wrapper

> this library is just a very simple wrapper around ecdh so u can use it in your projects

```go
import (
    "fmt"
    "diffie"
)

var d1 diffie.DiffieHellman
var d2 diffie.DiffieHellman

d1_publickey := d1.GenPublicKey()
d2_publickey := d2.GenPublicKey()

d1_sharedkey := d1.GenSharedKey(d2_publickey)
d2_sharedkey := d2.GenSharedKey(d1_publickey)

fmt.Println(d1_sharedkey == d2_sharedkey)
```
```
output:

true
```