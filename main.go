package main

import (
  // "errors"
  "fmt"
  "net/http"
  "net/url"
  "os"

  "github.com/golang-collections/collections/stack"
  "github.com/jcelliott/lumber"
  "gopkg.in/urfave/cli.v1" // imports as package "cli"
)

var redirectsStack *stack.Stack = stack.New()

type parsedResponse struct {
  code int
  request string
  response string
  content string
  redirects *stack.Stack
}

func main() {
  lumber.Info("initialized")


  app := cli.NewApp()
  app.EnableBashCompletion = true

  app.Action = func(c *cli.Context) error {
    arg := c.Args().Get(0)
    resp, err := doAction(arg)

    if err != nil {
      lumber.Fatal(err.Error())
      os.Exit(1)
    } else {
      err = printResp(resp)
    }

    return err
  }

  err := app.Run(os.Args)
  if err != nil {
    lumber.Fatal(err.Error())
    os.Exit(1)
  }
}

func doAction(arg string) (parsedResponse, error) {
  lumber.Info("arg: %s\n", arg)
  // try to parse the arg as a url
  parsedUrl, err := url.Parse(arg)

  // if it parses, get it
  parsed, err := getUrl(parsedUrl)

  return parsed, err
}

func getUrl(u *url.URL) (parsedResponse, error) {
  client := &http.Client{
    CheckRedirect: checkRedirect(),
  }

  req, err := http.NewRequest("GET", u.String(), nil)
  resp, err := client.Do(req)

  code := resp.StatusCode
  request := resp.Request.URL.String()
  response := resp.Header.Get("location")
  content := ""
  redirects := redirectsStack

  defer resp.Body.Close()

  return parsedResponse{code: code, request: request, response:response, content: content, redirects: redirects}, err
}

type checkRedirectFunction func(*http.Request, []*http.Request) error

func checkRedirect() checkRedirectFunction {
  return func(req *http.Request, via []*http.Request) error {
    lumber.Info("redirected to: %s\n", req.URL.String())
    redirectsStack.Push(req.URL)
    return nil
  }
}

func printResp(resp parsedResponse) error {
  fmt.Printf("%d\t%s\t%s\t%d\n", resp.code, resp.request, resp.response, resp.redirects.Len())

  return nil
}
