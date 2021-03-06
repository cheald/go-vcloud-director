# go-vcloud-director [![Build Status](https://travis-ci.org/vmware/go-vcloud-director.svg?branch=master)](https://travis-ci.org/vmware/go-vcloud-director) [![Coverage Status](https://coveralls.io/repos/vmware/go-vcloud-director/badge.svg?branch=master&service=github)](https://coveralls.io/github/vmware/go-vcloud-director?branch=master) [![GoDoc](https://godoc.org/github.com/vmware/go-vcloud-director?status.svg)](http://godoc.org/github.com/vmware/go-vcloud-director) [![Chat](https://img.shields.io/badge/chat-on%20slack-brightgreen.svg)](https://vmwarecode.slack.com/messages/CBBBXVB16)

This repo contains the `go-vcloud-director` package which implements
an SDK for vCloud Director. The project serves the needs of Golang
developers who need to integrate with vCloud Director. It is also the
basis of the [vCD Terraform
Provider](https://github.com/terraform-providers/terraform-provider-vcd).

## Contributions ##

Contributions to `go-vcloud-director` are gladly welcome and range
from participating in community discussions to submitting pull
requests.  Please see the [contributing guide](CONTRIBUTING.md) for
details on joining the community.

### Install and Build ###

This project started using Go [modules](https://github.com/golang/go/wiki/Modules)
starting with version 2.1 and `vendor` directory is no longer bundled.

As per the [modules documentation](https://github.com/golang/go/wiki/Modules)
you no longer need to use `GOPATH`. You can clone the branch in any directory
you like and go will fetch dependencies specified in the `go.mod` file:
```
cd ~/Documents/mycode
git clone https://github.com/vmware/go-vcloud-director.git
cd go-vcloud-director/govcd
go build
```

**Note** Do not forget to set `GO111MODULE=on` if you are in your GOPATH.
[Read more about this.](https://github.com/golang/go/wiki/Modules#how-to-install-and-activate-module-support)

#### Example code ####

To show the SDK in action run the example:
```
mkdir ~/govcd_example
go mod init govcd_example
go get github.com/vmware/go-vcloud-director/v2@master
go build -o example
./example user_name "password" org_name vcd_IP vdc_name 
```

Here's the code:
```go
package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/vmware/go-vcloud-director/v2/govcd"
)

type Config struct {
	User     string
	Password string
	Org      string
	Href     string
	VDC      string
	Insecure bool
}

func (c *Config) Client() (*govcd.VCDClient, error) {
	u, err := url.ParseRequestURI(c.Href)
	if err != nil {
		return nil, fmt.Errorf("unable to pass url: %s", err)
	}

	vcdclient := govcd.NewVCDClient(*u, c.Insecure)
	err = vcdclient.Authenticate(c.User, c.Password, c.Org)
	if err != nil {
		return nil, fmt.Errorf("unable to authenticate: %s", err)
	}
	return vcdclient, nil
}

func main() {
	if len(os.Args) < 6 {
		fmt.Println("Syntax: example user password org VCD_IP VDC ")
		os.Exit(1)
	}
	config := Config{
		User:     os.Args[1],
		Password: os.Args[2],
		Org:      os.Args[3],
		Href:     fmt.Sprintf("https://%s/api", os.Args[4]),
		VDC:      os.Args[5],
		Insecure: true,
	}

	client, err := config.Client() // We now have a client
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	org, err := govcd.GetOrgByName(client, config.Org)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vdc, err := org.GetVdcByName(config.VDC)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Org URL: %s\n", org.Org.HREF)
	fmt.Printf("VDC URL: %s\n", vdc.Vdc.HREF)
}

```
