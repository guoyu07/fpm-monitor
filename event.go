package main

type Event struct {
}

type FPMStatus struct {
	Pool               string `json:"pool"`
	ProcessManager     string `json:"process manager"`
	StartTime          int    `json:"start time"`
	StartSince         int    `json:"start since"`
	AcceptedConn       int    `json:"accepted conn"`
	ListenQueue        int    `json:"listen queue"`
	MaxListenQueue     int    `json:"max listen queue"`
	ListenQueueLen     int    `json:"listen queue len"`
	IdleProcesses      int    `json:"idle processes"`
	ActiveProcesses    int    `json:"active processes"`
	TotalProcesses     int    `json:"total processes"`
	MaxActiveProcesses int    `json:"max active processes"`
	MaxChildrenReached int    `json:"max children reached"`
	SlowRequests       int    `json:"slow requests"`
	Processes          []struct {
		Pid               int     `json:"pid"`
		State             string  `json:"state"`
		StartTime         int     `json:"start time"`
		StartSince        int     `json:"start since"`
		Requests          int     `json:"requests"`
		RequestDuration   int     `json:"request duration"`
		RequestMethod     string  `json:"request method"`
		RequestURI        string  `json:"request uri"`
		ContentLength     int     `json:"content length"`
		User              string  `json:"user"`
		Script            string  `json:"script"`
		LastRequestCPU    float64 `json:"last request cpu"`
		LastRequestMemory int     `json:"last request memory"`
	} `json:"processes"`
}
