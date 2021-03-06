package main

import (
	dolores_corecode "github.com/OpenChaos/dolores/corecode"
	dolores_gitlab "github.com/OpenChaos/dolores/drives/gitlab"
	dolores_slack "github.com/OpenChaos/dolores/loops/slack"
	dolores_memories "github.com/OpenChaos/dolores/memories"

	"github.com/jasonlvhit/gocron"
)

func prepareScheduler() {
	scheduler := gocron.NewScheduler()
	scheduler.Every(1).Hours().Do(dolores_memories.GcloudComputeInstances)
	scheduler.Every(1).Hours().Do(dolores_gitlab.MarkUsersInternal)
	<-scheduler.Start()
	//scheduler.Every(3).Minutes().Do(task)
	// more examples: https://github.com/jasonlvhit/gocron/blob/master/example/example.go#L19
}

func main() {
	config := dolores_corecode.ConfigFromFlags()

	go prepareScheduler() // spawn cron scheduler jobs
	dolores_slack.LoopRTMEvents(config)
}
