# GOAuth
[under development] Golang OAuth system like for desktop applicattions.

### Configuration
Create a file `config.go` like this:

```golang
package main

func init() {
	CONFIG = Config{
		OAuthResponse: "success",
		OAuthURL: "http://yoursupremesite/oauth",
	}
}
```

### How it Works
1) Application will create a HTTP server listening on `http://localhost:8080/`;
2) Your site must make a request to `http://localhost:8080/`;
3) Application will get your request and respond!

### Example in Your Site
```javascript
async function login() {
    const user = await (await fetch("http://mysite/login")).json();
    const auth = await fetch(`http://localhost:8080/?token=${user.jwt}`);
    if (auth == "success") alert(`Your device it's connected to your account!`);
}
```