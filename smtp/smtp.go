package smtp

import (
	"SampleApp/config"
	"errors"
	"fmt"
	"github.com/emersion/go-smtp"
	"io"
	"io/ioutil"
	"log"
	"time"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// A Session is returned after EHLO.
type Session struct{}

var smtpConfig = &config.SMTPConfig{
	Host: "localhost",
	Port: 1024,
}

func (bkd *Backend) NewSession(_ smtp.ConnectionState, _ string) (smtp.Session, error) {
	return &Session{}, nil
}

func (bkd *Backend) Login(state *smtp.ConnectionState, username string, password string) (smtp.Session, error) {
	//Checking internal registered users
	for _, user := range smtpConfig.Registered {
		if user.Username == username {
			if password == user.Password {
				return &Session{}, nil
			}
		}
	}

	//todo check for users in smtp db

	return nil, errors.New("invalid username or password")
}

func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, errors.New("unauthorized")
}

func (s *Session) AuthPlain(username, password string) error {
	for _, user := range smtpConfig.Registered {
		if user.Username == username {
			if password == user.Password {
				return nil
			}
		}
	}
	return errors.New("invalid username or password")
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	//todo store mails from others
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		log.Println("Data:", string(b))
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func RunServer(config *config.SMTPConfig) {

	smtpConfig = config

	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = fmt.Sprintf(":%d", config.Port)
	s.Domain = config.Host
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024 * 10
	s.MaxRecipients = 50
	s.AllowInsecureAuth = false

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
