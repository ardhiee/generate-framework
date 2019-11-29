package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// WriteModel to create model struct
func WriteModel(path string, domain string) {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	content := viper.GetString(`models`)
	header := "package models\n\ntype " + strings.Title(domain) + " struct {\n"
	_, err = file.WriteString(header)

	c := strings.Split(content, ",")

	for _, d := range c {
		e := strings.Split(d, ":")
		_, err = file.WriteString(" 			" + strings.ToUpper(e[0]) + " " + e[1] + "\n")
	}

	_, err = file.WriteString("}")

	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	// fmt.Println("==> file berhasil di isi")
}

func WriteModelLog(path string) {

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	str := []string{

		"package models\n",
		"\n",
		"type Logdata struct {\n",
		"	RequestID string\n",
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

func WriteModelErrors(path string) {

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	str := []string{
		"package models\n",
		"\n",
		"import \"errors\"\n",
		"\n",
		"var (\n",
		"	// ErrInternalServerError will throw if any the Internal Server Error happen\n",
		"	ErrInternalServerError = errors.New(\"Internal Server Error\")\n",
		"	// ErrNotFound will throw if the requested item is not exists\n",
		"	ErrNotFound = errors.New(\"Your requested Item is not found\")\n",
		"	// ErrConflict will throw if the current action already exists\n",
		"	ErrConflict = errors.New(\"Your Item already exist\")\n",
		"	// ErrBadParamInput will throw if the given request-body or params is not valid\n",
		"	ErrBadParamInput = errors.New(\"Given Param is not valid\")\n",
		")\n",
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
