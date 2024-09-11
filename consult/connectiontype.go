package consult

type Status struct {
	LocalAddr string `json:"local_addr"`
	Status    bool   `json:"status"`
}

type Address struct {
	RemoteAddr         string `json:"remoteAddr"`
	Remoteprotocoltype string `json:"remoteprotocoltype"`
}
