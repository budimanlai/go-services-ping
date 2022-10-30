package ping

type PingInterface interface {
	// send ping when services is start running
	Start()

	// send periodic ping when service is running
	Update()

	// send ping when services is stop or terminated
	Stop()
}
