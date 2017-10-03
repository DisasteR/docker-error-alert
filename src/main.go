package main

import (
    "context"
    "fmt"
    "strconv"
    "time"
    "os"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
    "gopkg.in/telegram-bot-api.v4"
)

func main() {
  waiting := false
  id, err := strconv.ParseInt(os.Getenv("TELEGRAM_ID"), 10, 64)
  if err != nil { panic(err) }
  bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
  if err != nil { panic(err) }

  for {
    cli, err := client.NewEnvClient()
    containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
    if err != nil {
        panic(err)
    }
    for _, container := range containers {
      if (container.State != "running") {
        if err != nil { panic(err) }
        waiting = true
        text := fmt.Sprintf("** %s **\n Images : %s\n State : %s\n Status : `%s` \n\n", container.Names, container.Image, container.State, container.Status)
        msg := tgbotapi.NewMessage(id, text)
        msg.ParseMode = "markdown"
        bot.Send(msg)
      }
    }
    if (waiting) {
      fmt.Println("Mute for 60min.. Waiting...")
      time.Sleep(60 * time.Minute)
      waiting = false
    }
    time.Sleep(5 * time.Second)
    fmt.Println("OK.")

  }
}
