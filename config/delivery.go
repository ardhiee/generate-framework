package config

import (
	"os"
	"strings"
)

func WriteDomainDeliveryHTTP(path, domain, projectName string) {

	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	str := []string{
		" package http\n",
		"\n",
		" import (\n",
		" 	\"context\"\n",
		" 	\"net/http\"\n",
		" 	\"strconv\"\n",
		"\n",
		" 	\"github.com/ardhiee/" + projectName + "/" + domain + "\"\n",
		" 	\"github.com/ardhiee/" + projectName + "/models\"\n",
		" 	\"github.com/ardhiee/utils\"\n",
		" 	\"github.com/labstack/echo\"\n",
		" )\n",
		"\n",
		" type ResponseError struct {\n",
		" 	Message string `json:\"message\"`\n",
		" }\n",
		"\n",
		" type " + strings.Title(domain) + "Handler struct {\n",
		" 	" + strings.Title(domain) + "Usecase " + domain + ".Usecase\n",
		" }\n",
		"\n",
		" func New" + strings.Title(domain) + "Handler(e *echo.Echo, auc " + domain + ".Usecase) {\n",
		" 	handler := &" + strings.Title(domain) + "Handler{\n",
		" 		" + strings.Title(domain) + "Usecase: auc,\n",
		" 	}\n",
		"\n",
		" 	e.GET(\"/" + domain + "/:id\", handler.GetByID)\n",
		" }\n",
		"\n",
		" func (a *" + strings.Title(domain) + "Handler) GetByID(c echo.Context) error {\n",
		"\n",
		" 	idP, err := strconv.Atoi(c.Param(\"id\"))\n",
		" 	if err != nil {\n",
		" 		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())\n",
		" 	}\n",
		"\n",
		" 	id := idP\n",
		"\n",
		" 	ctx := c.Request().Context()\n",
		" 	if ctx == nil {\n",
		" 		ctx = context.Background()\n",
		" 	}\n",
		"\n",
		" 	logdata := utils.Logdata{\n",
		" 		RequestID: c.Request().Header.Get(\"x-request-id\"),\n",
		" 	}\n",
		"\n",
		" 	utils.MakeLogEntry(nil, &logdata).Info(\"HTTP Handler receive the request\")\n",
		"\n",
		" 	// call the usecase\n",
		" 	art, err := a." + strings.Title(domain) + "Usecase.GetByID(ctx, id, &logdata)\n",
		" 	if err != nil {\n",
		" 		return c.JSON(getStatusCode(err, logdata), ResponseError{Message: err.Error()})\n",
		" 	}\n",
		" 	utils.MakeLogEntry(nil, &logdata).Info(\"HTTP Handler respond the request\")\n",
		" 	return c.JSON(http.StatusOK, art)\n",
		"\n",
		" }\n",
		"\n",
		" func getStatusCode(err error, l utils.Logdata) int {\n",
		" 	if err == nil {\n",
		" 		return http.StatusOK\n",
		" 	}\n",
		" 	utils.MakeLogEntry(nil, &l).Error(err)\n",
		" 	switch err {\n",
		" 	case models.ErrInternalServerError:\n",
		" 		return http.StatusInternalServerError\n",
		" 	case models.ErrNotFound:\n",
		" 		return http.StatusNotFound\n",
		" 	case models.ErrConflict:\n",
		" 		return http.StatusConflict\n",
		" 	default:\n",
		" 		return http.StatusInternalServerError\n",
		" 	}\n",
		" }\n",
	}

	for _, s := range str {
		_, err = file.WriteString(s)
		if isError(err) {
			return
		}
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
