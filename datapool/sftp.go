package datapool

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type DataPool struct {
	Client *sftp.Client
	Target TargetConfig
}

func (n *DataPool) UploadFile(name string, reader io.Reader) error {
	dstPath := filepath.Join(n.Target.Path, name)
	tmpPath := dstPath + ".temp"

	// 重命名文件
	defer n.Client.Rename(tmpPath, dstPath)

	dstFile, err := n.Client.Create(tmpPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, reader)

	return err
}

func New(target TargetConfig) *DataPool {
	var (
		ssh *ssh.Client
		ftp *sftp.Client
		err error
	)

	auth := []ssh.AuthMethod{ssh.Password(target.Password)}

	conf := &ssh.ClientConfig{
		User:            target.Username,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%d", target.Host, target.Port)

	if ssh, err = ssh.Dial("tcp", addr, conf); err != nil {
		log.Fatalln("Ssh dial error: ", err)
		return nil
	}

	if ftp, err = sftp.NewClient(ssh); err != nil {
		log.Fatalln("SFtp new error: ", err)
		return nil
	}

	net := &Network{
		Client: ftp,
		Target: target,
	}

	return net
}
