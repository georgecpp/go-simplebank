package mail

import (
	"testing"

	"github.com/georgecpp/go-simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello World</h1>
	<p>This is a test message from <a href="https://github.com/georgecpp/go-simplebank">George Andronache</a></p>
	`

	to := []string{"georgeandronache.cpp@gmail.com"}
	attachFiles := []string{"../Makefile"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}