package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Startup Executes necessary functions before starting server.
func Startup() *gin.Engine {

	styledStartup()

	app := getGinEngine()

	return app
}

/* styledStartup Shows an ascii stylish startup message ;)
https://patorjk.com/software/taag/#p=display*/
func styledStartup() {
	var message = `
	 ██████╗  ██████╗      █████╗ ██████╗ ██╗
	██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██║
	██║  ███╗██║   ██║    ███████║██████╔╝██║   
	██║   ██║██║   ██║    ██╔══██║██╔═══╝ ██║   
	╚██████╔╝╚██████╔╝    ██║  ██║██║     ██║    
	 ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝     ╚═╝`
	fmt.Println(message)
}

/* getGinEngine Returns a new gin engine with default format.
 */
func getGinEngine() *gin.Engine {
	return gin.New()
}
