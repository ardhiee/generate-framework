package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// WriteDomainRepo will create Repository Domain
func WriteDomainRepo(path, domain, projectName string) {

	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	content := viper.GetString(`models`)
	header := []string{
		"package repository\n\n",
		"import (\n",
		"	\"context\"\n",
		"	\"database/sql\"\n",
		"	\"github.com/ardhiee/" + projectName + "/" + domain + "\"\n",
		"	\"github.com/ardhiee/" + projectName + "/models\"\n",
		"	\"github.com/ardhiee/utils\"\n",
		")\n\n",
		"type " + domain + "MysqlRepository struct {\n",
		"	Conn *sql.DB\n",
		"}\n\n",
		"func New" + strings.Title(domain) + "MysqlRepository(Conn *sql.DB) " + domain + ".Repository {\n",
		"	return &" + domain + "MysqlRepository{Conn}\n",
		"}",
		"\n\n\n\n",
		"func (a *" + domain + "MysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models." + strings.Title(domain) + ", error) {\n",
		"	rows, err := a.Conn.QueryContext(ctx, query, args...)\n",
		"	if err != nil {\n",
		"		utils.MakeLogEntry(nil, nil).Error(err)\n",
		"		return nil, err\n",
		"	}\n",
		"\n",
		"	defer func() {\n",
		"		err := rows.Close()\n",
		"		if err != nil {\n",
		"			utils.MakeLogEntry(nil, nil).Error(err)\n",
		"		}\n",
		"	}()\n",
		"\n",
		"	result := make([]*models." + strings.Title(domain) + ", 0)\n",
		"	for rows.Next() {\n",
		"		t := new(models." + strings.Title(domain) + ")\n",
		"		err = rows.Scan(\n",
	}

	footer := []string{
		"	)\n",
		"\n",
		"		if err != nil {\n",
		"			utils.MakeLogEntry(nil, nil).Error(err)\n",
		"			return nil, err\n",
		"		}\n",
		"\n",
		"		result = append(result, t)\n",
		"}\n",
		"\n",
		"return result, nil\n",
		"}\n",
	}

	//
	headerOne := []string{
		"func (a *" + domain + "MysqlRepository) GetByID(ctx context.Context, id int, l *utils.Logdata) (res *models." + strings.Title(domain) + ", err error) {\n",
		"	utils.MakeLogEntry(nil, l).Info(\"Repository receive the request\")\n",
		"\n",
	}

	footerOne := []string{
		"\n",
		"	list, err := a.fetch(ctx, query, id)\n",
		"	if err != nil {\n",
		"		return nil, err\n",
		"	}\n",
		"\n",
		"	if len(list) > 0 {\n",
		"		res = list[0]\n",
		"	} else {\n",
		"		return nil, models.ErrNotFound\n",
		"	}\n",
		"\n",
		"	utils.MakeLogEntry(nil, l).Info(\"Repository respond the request\")\n",
		"	return\n",
		"}\n",
	}

	for _, str := range header {
		_, err = file.WriteString(str)
	}
	var columnTableQueries string
	c := strings.Split(content, ",")
	for _, d := range c {
		e := strings.Split(d, ":")
		columnTableQueries = columnTableQueries + e[0] + ","
		_, err = file.WriteString("		&t." + strings.ToUpper(e[0]) + ",\n")
	}
	for _, str := range footer {
		_, err = file.WriteString(str)
	}

	for _, str := range headerOne {
		_, err = file.WriteString(str)
	}

	s := columnTableQueries
	if last := len(s) - 1; last >= 0 && s[last] == ',' {
		s = s[:last]
	}

	query := "	query := `SELECT " + s + "\n		FROM " + domain + " WHERE ID = ?`\n"
	_, err = file.WriteString(query)

	for _, str := range footerOne {
		_, err = file.WriteString(str)
	}

	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

}

// WriteDomainRepoInterface will create repo interface
func WriteDomainRepoInterface(path, domain, projectName string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	//
	str := []string{
		"package " + domain,
		"\n",
		"import (\n",
		"	\"context\"\n",
		"\n",
		"	\"github.com/ardhiee/" + projectName + "/models\"\n",
		"	\"github.com/ardhiee/utils\"\n",
		")\n",
		"\n",
		"type Repository interface {\n",
		"	GetByID(ctx context.Context, id int, l *utils.Logdata) (*models." + strings.Title(domain) + ", error)\n",
		"}\n",
	}
	//
	for _, s := range str {
		_, err = file.WriteString(s)
		if isError(err) {
			return
		}
	}
	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
