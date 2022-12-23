# Diffie Hellman Library

### "github.com/aead/ecdh" simple wrapper

> this library is just a very simple wrapper around ecdh so u can use it in your projects

```go
import (
    "fmt"
    "github.com/husseinamine/ncrypt"
)

var d1 ncrypt.DiffieHellman
var d2 ncrypt.DiffieHellman

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
