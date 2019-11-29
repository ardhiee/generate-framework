package config

import "os"

func WriteMiddleware(path string) {

	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	str := []string{
		"package middleware\n",
		"\n",
		"import (\n",
		"	\"github.com/ardhiee/utils\"\n",
		"	\"github.com/labstack/echo\"\n",
		"	\"github.com/labstack/echo/middleware\"\n",
		")\n",
		"\n",
		" //GoMiddleware represent the data-struct for middleware\n",
		"type GoMiddleware struct {\n",
		"	 //another stuff , may be needed by middleware\n",
		"}\n",
		"\n",
		"var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{\n",
		"	SigningKey: []byte(\"secret\"),\n",
		"})\n",
		"\n",
		" //CORS will handle the CORS middleware\n",
		"func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {\n",
		"	return func(c echo.Context) error {\n",
		"		c.Response().Header().Set(\"Access-Control-Allow-Origin\", \"*\")\n",
		"		return next(c)\n",
		"	}\n",
		"}\n",
		"\n",
		"func (m *GoMiddleware) LOG(next echo.HandlerFunc) echo.HandlerFunc {\n",
		"	return func(c echo.Context) error {\n",
		"		utils.MakeLogEntry(c, nil).Info(\"Incoming HTTP request\")\n",
		"		return next(c)\n",
		"	}\n",
		"}\n",
		"\n",
		" //InitMiddleware intialize the middleware\n",
		"func InitMiddleware() *GoMiddleware {\n",
		"	return &GoMiddleware{}\n",
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
