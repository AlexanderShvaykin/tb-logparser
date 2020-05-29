package mail_log

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"log"
	"regexp"
	"strings"
)

var conn *pg.DB

type Log struct {
	Name      string
	AccountId string
	Email     string
	DateTime  string
}

func (l Log) String() string {
	return fmt.Sprintf("Log<%v Email: %v AccountId: %v>", l.Name, l.Email, l.AccountId)
}

func Configure(user string, password string, database string) {
	conn = pg.Connect(&pg.Options{
		User: user, Password: password, Database: database,
	})
}

func Close() {
	defer conn.Close()
}

func Parse(line string) {
	argsPattern := `"arguments":\["(?P<mailer>[a-zA-Z]+)","(?P<type>[a-zA-Z_]+)`
	accountPattern := `"account_id":"(?P<accountId>[0-9]+)"`
	emailPattern := `"email":"(?P<email>[\S]+)",`
	datetimePattern := `completed_at\":\"(?P<datetime>[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2})`

	re := regexp.MustCompile(argsPattern)
	template := "$mailer#$type"
	name := string(re.ExpandString([]byte{}, template, line, re.FindStringSubmatchIndex(line)))

	re = regexp.MustCompile(accountPattern)
	template = "$accountId"
	accountId := string(re.ExpandString([]byte{}, template, line, re.FindStringSubmatchIndex(line)))

	re = regexp.MustCompile(datetimePattern)
	template = "$datetime"
	datetime := string(re.ExpandString([]byte{}, template, line, re.FindStringSubmatchIndex(line)))

	re = regexp.MustCompile(emailPattern)
	template = "$email"
	email := string(re.ExpandString([]byte{}, template, line, re.FindStringSubmatchIndex(line)))
	email = strings.Split(email, `"`)[0]

	mailLog := &Log{
		Name: name, Email: email, AccountId: accountId, DateTime: datetime,
	}

	fmt.Println(mailLog)
	err := conn.Insert(mailLog)
	if err != nil {
		fmt.Println(line)
		fmt.Println(mailLog)
		log.Fatal(err)
	}
}
