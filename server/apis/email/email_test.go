package email

import "testing"

func TestEmailReceive(t *testing.T) {
	name := "yujin"
	email := "winter.kwon@gbike.io"
	number := "010-9965-9409"
	title := "문의 제목"
	content := "내용"
	receiver := "contact.youjinkwon@gmail.com"
	pwd := "tkkrzchbpqldhobo"
	if err := ReceiveMail3(name, email, number, title, content, receiver, pwd); err != nil {
		t.Error(err)
	}
}
