package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ardhiee/generate-framework/config"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
func prepareFolder() {
	projectName := viper.GetString(`projectName`)
	path := viper.GetString(`path`) + projectName
	domain := viper.GetString(`domain`)

	projectSubFolder := []string{domain + "/delivery/http", domain + "/repository", domain + "/usecase", "models", "middleware"}

	for _, folder := range projectSubFolder {
		CreateDirIfNotExist(path + "/" + folder)
		if strings.Contains(folder, "http") {
			http := path + "/" + folder + "/" + domain + "_handler.go"
			createFile(http)
			config.WriteDomainDeliveryHTTP(http, domain, projectName)
		} else if strings.Contains(folder, "repository") {
			repo := path + "/" + folder + "/" + domain + "_repository.go"
			createFile(repo)
			config.WriteDomainRepo(repo, domain, projectName)
		} else if strings.Contains(folder, "usecase") {
			uc := path + "/" + folder + "/" + domain + "_usecase.go"
			createFile(uc)
			config.WriteDomainUsecase(uc, domain, projectName)
		} else if strings.Contains(folder, "models") {
			models := path + "/" + folder + "/" + domain + ".go"
			createFile(models)
			config.WriteModel(models, domain)

			errs := path + "/" + folder + "/errors.go"
			createFile(errs)
			config.WriteModelErrors(errs)

			lg := path + "/" + folder + "/log.go"
			createFile(lg)
			config.WriteModelLog(lg)

		} else if strings.Contains(folder, "middleware") {
			mdl := path + "/" + folder + "/" + "middleware.go"
			createFile(mdl)
			config.WriteMiddleware(mdl)
		}
		// fmt.Println("DONE :", path+"/"+folder)
	}
	ri := path + "/" + domain + "/" + "repository_interface.go"
	createFile(ri)
	config.WriteDomainRepoInterface(ri, domain, projectName)

	ui := path + "/" + domain + "/" + "usecase_interface.go"
	createFile(ui)
	config.WriteDomainUsecaseInterface(ui, domain, projectName)

	mn := path + "/" + "main.go"
	createFile(mn)
	config.WriteMain(mn, domain, projectName)

	cfg := path + "/" + "config.json"
	createFile(cfg)
	config.WriteConfigJSON(cfg, domain)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile(path string) {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func main() {

	//path := viper.GetString(`path`) + viper.GetString(`projectName`)
	//domain := viper.GetString(`domain`)
	prepareFolder()
}
