package config

import (
	"os"
	"strings"
)

func WriteDomainUsecase(path, domain, projectName string) {

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	str := []string{
		"package usecase\n",
		"\n",
		"import (\n",
		"	\"context\"\n",
		"	\"time\"\n",
		"\n",
		"	\"github.com/ardhiee/" + projectName + "/" + domain + "\"\n",
		"	\"github.com/ardhiee/" + projectName + "/models\"\n",
		"	\"github.com/ardhiee/utils\"\n",
		")\n",

		"type " + domain + "Usecase struct {\n",
		"	" + domain + "Repo    " + domain + ".Repository\n",
		"	contextTimeout time.Duration\n",
		"}\n",

		"func New" + strings.Title(domain) + "Usecase(a " + domain + ".Repository, timeout time.Duration) " + domain + ".Usecase {\n",
		"	return &" + domain + "Usecase{\n",
		"		" + domain + "Repo:    a,\n",
		"		contextTimeout: timeout,\n",
		"	}\n",
		"}\n",
		"\n",
		"func (a *" + domain + "Usecase) GetByID(c context.Context, id int, l *utils.Logdata) (*models." + strings.Title(domain) + ", error) {\n",
		"\n",
		"	utils.MakeLogEntry(nil, l).Info(\"Usecase receive the request\")\n",
		"	ctx, cancel := context.WithTimeout(c, a.contextTimeout)\n",
		"	defer cancel()\n",
		"\n",
		"	res, err := a." + domain + "Repo.GetByID(ctx, id, l)\n",
		"	if err != nil {\n",
		"		return nil, err\n",
		"	}\n",
		"	utils.MakeLogEntry(nil, l).Info(\"Usecase respond the request\")\n",
		"	return res, nil\n",
		"}\n",
	}

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

// WriteDomainRepoInterface will create repo interface
func WriteDomainUsecaseInterface(path, domain, projectName string) {
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
		"type Usecase interface {\n",
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
