# will_it_redirect [![CircleCI](https://circleci.com/gh/asicsdigital/will_it_redirect.svg?style=svg)](https://circleci.com/gh/asicsdigital/will_it_redirect)

To install:
1. Install and Configure Go [here](https://golang.org/doc/install)
2. Set $GOPATH and add $GOPATH/bin to your $PATH in your ~/.bash_profile
3. Reload bash or **run**: source ~/.bash_profile
4. **Run**: go get github.com/asicsdigital/will_it_redirect
5. You will now be able to **run**: will_it_redirect <url>


```sh
$ ./will_it_redirect https://asics.com/technology
2018-05-29 16:28:52 INFO  site: https://asics.com/technology
2018-05-29 16:28:52 INFO  redirect: https://georedirect.asics.com/2017-technology
2018-05-29 16:28:52 INFO  redirect: https://www.asics.com/us/en-us/technology
200	https://www.asics.com/us/en-us/technology		2
```
