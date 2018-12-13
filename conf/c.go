package conf

const (
	DBHOST = "127.0.0.1"
	DBPORT = "3306"
	DBPASSWORD = "Passw0rd"
	DBUSERNAME = "root"
	DBDATABASE = "izghua"

	ALARMCRITICAL = "critical"
	ALARMWARNING  = "warning"
	ALARMALERT    = "alert"

	MAIlTYPE = "html"

	HASHIDSALT = "salt"
	HASHIDMINLENGTH = 8


	REDISADDR = ""
	REDISPWD = ""
	REDISDB = 0
)

// Log
const  (
	LOGFILEPATH = "./log"
	LOGFILENAME = "zog"
	LOGFILESUFFIX = "log"
	LOGFILEMAXSIZE = 0
	LOGFILEMAXNSIZE = 1
	LOGTIMEZONE = "Asia/Chongqing"
)
