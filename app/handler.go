package app

import (
	"hbcase/config"
	"hbcase/pkg"
	"hbcase/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var cmdService pkg.CommandService

func HandleCustomData(c *gin.Context) {
	file, err := c.FormFile("custom_file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "file not uploaded"})
		return
	}

	if !strings.HasSuffix(file.Filename, ".xlsx") {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "file extension must be xlsx."})
		return
	}

	path := "./" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "file could not be saved"})
		return
	}
	defer os.Remove(path)

	list := utils.ReadFromFile(path)

	commands := pkg.GenerateCommands(list)

	output := pkg.ExecuteCommands(commands, cmdService)

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Custom data implemented successfully",
		"output":  output,
	})
}

func HandleSampleData(c *gin.Context) {
	list := utils.ReadFromFile(config.GetFromEnv("SAMPLE_FILE"))

	commands := pkg.GenerateCommands(list)

	output := pkg.ExecuteCommands(commands, cmdService)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Sample data implemented successfully",
		"output":  output,
	})
}
