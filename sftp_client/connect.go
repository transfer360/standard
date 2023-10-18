package sftp_client

import (
	"context"
	"golang.org/x/crypto/ssh"
	"time"
)

func ConnectFromCredentials(ctx context.Context, secretPath string) (sshc *ssh.Client, err error) {

	cnt, err := LoadSFTPSecrets(ctx, secretPath)
	if err != nil {
		return nil, err
	}

	return Connect(cnt.Host, cnt.Username, cnt.Password)

}

func Connect(host, username, password string) (sshc *ssh.Client, err error) {

	return ssh.Dial("tcp", host+":22", &ssh.ClientConfig{
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	})
	//if err != nil {
	//	return err, fn
	//}
	//
	//// Create a new SFTP client
	//sftpClient, err := sftp.NewClient(client)
	//if err != nil {
	//	return err, fn
	//}
	//defer sftpClient.Close()
	//
	//fl, err := sftpClient.ReadDir("/Outbound")
	//if err != nil {
	//	return err, fn
	//}
	//
	//remoteFile := ""
	//localFile := ""
	//for _, rFile := range fl {
	//
	//	if !rFile.IsDir() {
	//		if strings.Contains(rFile.Name(), fileDate) {
	//			remoteFile = fmt.Sprintf("/Outbound/%s", rFile.Name())
	//			localFile = fmt.Sprintf("/tmp/%s", rFile.Name())
	//			fn = rFile.Name()
	//			break
	//		}
	//	}
	//}
	//
	//if len(remoteFile) == 0 {
	//	log.Warn("no file found with:", fileDate)
	//	return nil, fn
	//}
	//
	//// Open the remote file
	//srcFile, err := sftpClient.Open(remoteFile)
	//if err != nil {
	//	return err, fn
	//}
	//defer srcFile.Close()
	//
	//// Create the local file
	//dstFile, err := os.Create(localFile)
	//if err != nil {
	//	return err, fn
	//}
	//defer dstFile.Close()
	//
	//// Copy the contents of the remote file to the local file
	//_, err = io.Copy(dstFile, srcFile)
	//if err != nil {
	//	return err, fn
	//}
	//
	//return nil, fn
}
