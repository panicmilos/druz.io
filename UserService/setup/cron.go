package setup

import (
	"github.com/panicmilos/druz.io/UserService/cron_tasks"

	"gopkg.in/robfig/cron.v2"
)

func SetupCronTasks() {

	c := cron.New()
	c.AddFunc("@every 1h", cron_tasks.DeletePasswordRecoveryRequests)
	c.AddFunc("@every 1h", cron_tasks.DeleteUserReactivationsRequests)

	c.Start()
}
