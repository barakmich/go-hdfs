package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/user"
	"time"

	"github.com/colinmarc/hdfs"
	"github.com/colinmarc/hdfs/hadoopconf"
)

var cachedClients = make(map[string]*hdfs.Client)

func getClient(namenode string) (*hdfs.Client, error) {
	if cachedClients[namenode] != nil {
		return cachedClients[namenode], nil
	}

	if namenode == "" {
		namenode = os.Getenv("HADOOP_NAMENODE")
	}

	conf, err := hadoopconf.LoadFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("Problem loading configuration: %s", err)
	}

	options := hdfs.ClientOptionsFromConf(conf)
	if namenode != "" {
		options.Addresses = []string{namenode}
	}

	if options.Addresses == nil {
		return nil, errors.New("couldn't find a namenode to connect to. You should specify hdfs://<namenode>:<port> in your paths. Alternatively, set HADOOP_NAMENODE or HADOOP_CONF_DIR in your environment")
	}

	if options.KerberosClient != nil {
		options.KerberosClient, err = getKerberosClient()
		if err != nil {
			return nil, fmt.Errorf("Problem with kerberos authentication: %s", err)
		}
	} else {
		options.User = os.Getenv("HADOOP_USER_NAME")
		if options.User == "" {
			u, err := user.Current()
			if err != nil {
				return nil, fmt.Errorf("Couldn't determine user: %s", err)
			}

			options.User = u.Username
		}
	}

	// Set some basic defaults.
	dialFunc := (&net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 5 * time.Second,
		DualStack: true,
	}).DialContext

	options.NamenodeDialFunc = dialFunc
	options.DatanodeDialFunc = dialFunc

	c, err := hdfs.NewClient(options)
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to namenode: %s", err)
	}

	cachedClients[namenode] = c
	return c, nil
}

func formatBytes(i uint64) string {
	switch {
	case i > (1024 * 1024 * 1024 * 1024):
		return fmt.Sprintf("%#.1fT", float64(i)/1024/1024/1024/1024)
	case i > (1024 * 1024 * 1024):
		return fmt.Sprintf("%#.1fG", float64(i)/1024/1024/1024)
	case i > (1024 * 1024):
		return fmt.Sprintf("%#.1fM", float64(i)/1024/1024)
	case i > 1024:
		return fmt.Sprintf("%#.1fK", float64(i)/1024)
	default:
		return fmt.Sprintf("%dB", i)
	}
}
