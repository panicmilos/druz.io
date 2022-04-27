package setup

import (
	"UserService/cron_tasks"

	"gopkg.in/robfig/cron.v2"
)

func SetupCronTasks() {

	c := cron.New()
	c.AddFunc("@every 1h", cron_tasks.DeletePasswordRecoveryRequests)

	c.Start()
}
