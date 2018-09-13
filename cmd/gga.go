// Command-line tools, currentlly for download file form remote, after may be add more feature to here.
//
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

type newProject struct{}

// download url by different os type(Linux/Windows).
const (
	WINURL = "https://github.com/nightlegend/apigateway/files/2365358/apigateway-template.zip"
	LUXURL = "https://github.com/nightlegend/apigateway/files/2365349/apigateway-template.tar.gz"
)

var requestURL string

// newCmd is new commmand tools. And define all sub-command.
func (n *newProject) newCmd() *cobra.Command {
	// root command
	cmd := &cobra.Command{
		Use:   "gga [command]",
		Short: "A very helpful command line tools about apigateway",
		Long:  `A cool tools about init a project, It is can help you do some base thing`,
	}

	// gga new [path+filename]
	cmdNewProject := &cobra.Command{
		Use:   "clone [path+filename]",
		Short: "clone sample code to your path.",
		Long:  `Create a new content file. It will guess which kind of file to create based on the path provided.`,
		RunE:  n.downloadFile,
	}
	// add sub-command to root command
	cmd.AddCommand(cmdNewProject)

	return cmd
}

//download file from remote(github)
func (n *newProject) downloadFile(cmd *cobra.Command, args []string) error {
	createpath, _ := filepath.Abs(filepath.Clean(args[0]))
	fmt.Println(createpath)
	if runtime.GOOS == "windows" {
		requestURL = WINURL
	} else {
		requestURL = LUXURL
	}
	out, err := os.Create(createpath + ".zip")
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(requestURL)
	if err != nil {
		fmt.Println("Same happen error, please try again:)")
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	jww.FEEDBACK.Printf("Congratulations! Your download is created in %s.\n\n", createpath+".zip")
	jww.FEEDBACK.Println(nextStepsText())
	return nil
}

func nextStepsText() string {
	var nextStepsText bytes.Buffer
	nextStepsText.WriteString(`Just a few more steps and you're ready to go:
1. Unzip or tar download file.
2. Rename folder. `)
	nextStepsText.WriteString(`
3. Start add your code to here.`)
	return nextStepsText.String()
}

func main() {
	n := &newProject{}
	n.newCmd().Execute()
}
