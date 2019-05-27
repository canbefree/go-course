package consul

type Consul struct {
	mode      string
	configDir string
}

// mode agent client

// -agent -client

func (consul *Consul) start() {
	command := ""
	if consul.mode == "client" {
		command += " -client"
	}
}
