package config

import "os"

import "strings"

func WriteMain(path, domain, projectName string) {

	//buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	str := []string{
		"package main\n",
		"\n",
		"import (\n",
		"	\"database/sql\"\n",
		"	\"fmt\"\n",
		"	\"log\"\n",
		"	\"net/url\"\n",
		"	\"os\"\n",
		"	\"time\"\n",
		"\n",
		"	\"github.com/ardhiee/" + projectName + "/middleware\"\n",
		"	_ \"github.com/go-sql-driver/mysql\"\n",
		"	\"github.com/labstack/echo\"\n",
		"	\"github.com/spf13/viper\"\n",
		"\n",
		"	_" + domain + "HttpDeliver \"github.com/ardhiee/" + projectName + "/" + domain + "/delivery/http\"\n",
		"	_" + domain + "Repo \"github.com/ardhiee/" + projectName + "/" + domain + "/repository\"\n",
		"	_" + domain + "Usecase \"github.com/ardhiee/" + projectName + "/" + domain + "/usecase\"\n",
		")\n",
		"\n",
		"func init() {\n",
		"	viper.SetConfigFile(`config.json`)\n",
		"	err := viper.ReadInConfig()\n",
		"	if err != nil {\n",
		"		panic(err)\n",
		"	}\n",
		"\n",
		"	if viper.GetBool(`debug`) {\n",
		"		fmt.Println(\"Service RUN on DEBUG mode\")\n",
		"	}\n",
		"}\n",
		"\n",
		"func main() {\n",
		"\n",
		"	dbHost := viper.GetString(`database.host`)\n",
		"	dbPort := viper.GetString(`database.port`)\n",
		"	dbUser := viper.GetString(`database.user`)\n",
		"	dbPass := viper.GetString(`database.pass`)\n",
		"	dbName := viper.GetString(`database.name`)\n",
		"	connection := fmt.Sprintf(\"%s:%s@tcp(%s:%s)/%s\", dbUser, dbPass, dbHost, dbPort, dbName)\n",
		"	val := url.Values{}\n",
		"	val.Add(\"parseTime\", \"1\")\n",
		"	val.Add(\"loc\", \"Asia/Jakarta\")\n",
		"	dsn := fmt.Sprintf(\"%s?%s\", connection, val.Encode())\n",
		"	dbConn, err := sql.Open(`mysql`, dsn)\n",
		"	if err != nil && viper.GetBool(\"debug\") {\n",
		"		fmt.Println(err)\n",
		"	}\n",
		"	err = dbConn.Ping()\n",
		"	if err != nil {\n",
		"		log.Fatal(err)\n",
		"		os.Exit(1)\n",
		"	}\n",
		"\n",
		"	defer func() {\n",
		"		err := dbConn.Close()\n",
		"		if err != nil {\n",
		"			log.Fatal(err)\n",
		"		}\n",
		"	}()\n",
		"\n",
		"	e := echo.New()\n",
		"	middL := middleware.InitMiddleware()\n",
		"	e.Use(middL.CORS, middL.LOG, middleware.IsLoggedIn)\n",
		"\n",
		"	ar := _" + domain + "Repo.New" + strings.Title(domain) + "MysqlRepository(dbConn)\n",
		"\n",
		"	timeoutContext := time.Duration(viper.GetInt(\"context.timeout\")) * time.Second\n",
		"\n",
		"	au := _" + domain + "Usecase.New" + strings.Title(domain) + "Usecase(ar, timeoutContext)\n",
		"\n",
		"	_" + domain + "HttpDeliver.New" + strings.Title(domain) + "Handler(e, au)\n",
		"\n",
		"	log.Fatal(e.Start(viper.GetString(\"server.address\")))\n",
		"\n",
		"}\n",
	}

	for _, s := range str {
		_, err = file.WriteString(s)
		if isError(err) {
			return
		}
	}

	//simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

}

func WriteConfigJSON(path, domain string) {

	//buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	str := []string{
		"{\n",
		"	\"debug\": true,\n",
		"	\"server\": {\n",
		"	  \"address\": \":9091\"\n",
		"	},\n",
		"	\"context\":{\n",
		"	  \"timeout\":2\n",
		"	},\n",
		"	\"database\": {\n",
		"		\"host\": \"localhost\",\n",
		"		\"port\": \"3306\",\n",
		"		\"user\": \"root\",\n",
		"		\"pass\": \"password\",\n",
		"		\"name\": \"" + domain + "\"\n",
		"	}\n",
		"}\n",
	}
	for _, s := range str {
		_, err = file.WriteString(s)
		if isError(err) {
			return
		}
	}

	//simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

}
