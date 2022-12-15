package query

type Config struct {
	QueryInterval int    `json:"query_interval"` // query interval, seconds
	QueryCount    int    `json:"query_count"`    // query count
	NeedEmail     bool   `json:"need_email"`     //if true, need to fill the below fields
	FromEmail     string `json:"from_email"`
	ToEmail       string `json:"to_email"`
	EmailSubject  string `json:"email_subject"`
	Smtp          string `json:"smtp"`
	Port          int    `json:"port"`
	Password      string `json:"password"` // 密码/授权码，腾讯邮箱可以在账户设置里面开启POP3/SMTP服务器并得到一个授权码
}
