package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	ServiceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       ServiceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: ServiceAddress + "/services",
		HeartbeatURL:     ServiceAddress + "/heartbeat",
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatalln(err)
	}

	//接收context传递的信号
	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
