package main

import (
  // "errors"
  "fmt"
  "net/http"
  "net/url"
  "os"

  "github.com/jcelliott/lumber"
  "gopkg.in/urfave/cli.v1" // imports as package "cli"
)

type parsedResponse struct {
  code int
  request string
  response string
  content string
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
  lumber.Info("%s\n", arg)
  // try to parse the arg as a url
  parsedUrl, err := url.Parse(arg)

  // if it parses, get it
  parsed, err := getUrl(parsedUrl)

  return parsed, err
}

func getUrl(u *url.URL) (parsedResponse, error) {
  client := &http.Client{
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
    },
  }

  req, err := http.NewRequest("GET", u.String(), nil)
  resp, err := client.Do(req)

  code := resp.StatusCode
  request := resp.Request.URL.String()
  response := resp.Header.Get("location")
  content := ""

  defer resp.Body.Close()

  return parsedResponse{code: code, request: request, response:response, content: content}, err
}

func printResp(resp parsedResponse) error {
  fmt.Printf("%d\t%s\t%s\n", resp.code, resp.request, resp.response)

  return nil
}
