package cron

import (
	"strings"

	"github.com/Hamster601/fastweb/internal/repository/admin/cron_task"

	"github.com/spf13/cast"
)

func (s *server) AddTask(task *cron_task.CronTask) {
	spec := "0 " + strings.TrimSpace(task.Spec)
	name := cast.ToString(task.Id)

	s.cron.AddFunc(spec, s.AddJob(task), name)
}
