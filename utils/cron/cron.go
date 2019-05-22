/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-22
 * Time: 00:41
 */
package cron

import (
	"github.com/izghua/zgh"
	"github.com/robfig/cron"
)


func ZgCron(spec string,f func()) {
	c := cron.New()
	_ = c.AddFunc(spec, func() {
		f()
		zgh.ZLog().Info("ZgCron","ZgCron","Function",f)
	})

	c.Start()

	//go func() {
	//	for {
	//		f()
	//		now := time.Now()
	//		next := now.Add(duration)
	//		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
	//		t := time.NewTimer(next.Sub(now))
	//		<-t.C
	//	}
	//}()
}



