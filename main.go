package main

import (
  // "log"
  // "os"
  "fmt"
  "flag"
  // "encoding/json"
  // "io/ioutil"

  "github.com/minimalchat/chatctl/commands"
  "github.com/minimalchat/mnml-daemon/operator"
)

type configuration struct {
  ConfigFile string
  Host string
}

type chatrc struct {
  operator operator.Operator
}

var config configuration
// var commands map[int]string


func help() {
  fmt.Println("chatctl controls the Chat daemon\n")
  fmt.Println("Find more information at https://github.com/letschat/chatctl\n")

  fmt.Println("Usage:")
  fmt.Println("\tchatctl [flags] COMMAND")
  fmt.Println("\tchatctl COMMAND\n")

  fmt.Println("Commands:")
  fmt.Println("  get\tDisplay one or many resources")
  fmt.Println("  config\tModifies config files")
  fmt.Println("")

  fmt.Println("Flags:")
  flag.PrintDefaults()
}

func init() {
  // Configuration
  flag.StringVar(&config.Host, "api-server", "http://localhost:8000", "The `address` of the API server")
  flag.StringVar(&config.ConfigFile, "config", "~/.chatrc", "Path to chat config `file`, specifying operator and how to authenticate to the API server")
}

func main() {

  flag.Parse()
  argv := flag.Args()

  // // Try to load config file
  // if config.ConfigFile != "" {
  //   // Open file
  //   // log.Println(fmt.Sprintf("Opening %s ...", config.ConfigFile))
  //   confData, _ = ioutil.ReadFile(config.ConfigFile)
  //   // if err != nil {
  //   //     panic(err)
  //   // }
  // } else {
  //   // Open file
  //   // log.Println("Opening ~/.chatrc ...")
  //   confData, _ = ioutil.ReadFile("~/.chatrc")
  //   // if err != nil {
  //   //     panic(err)
  //   // }
  // }
  // // fmt.Println(string(confData))
  // err := json.Unmarshal(confData, &confObject)
  // if err != nil {
  //   log.Fatal(err)
  // }


  if len(argv) > 0 {
    fmt.Println(argv[0])

    switch {
      case argv[0] == "get": cmds.Get(argv)
      // case argv[0] == "config": cmds.Config(argv)
      default: help()
    }
  } else {
    help()
  }

}
