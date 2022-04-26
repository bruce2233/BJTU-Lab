package myshell

import (
	"fmt"
	"mysyscall/gocall"
	"syscall"
	"time"

	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

// func main() {
// 	println("Hello, I'm myapp from MYPATH")
// }
func InitCmd() *grumble.App {
	cwdstr := gocall.Getcwd()
	var app = grumble.New(&grumble.Config{
		Name:        "shell",
		Description: "A simple shell by 19281030-张云鹏",
		Prompt:      "19281030:" + cwdstr + "# ",
		PromptColor: color.New(color.FgGreen, color.Bold),
		Flags: func(f *grumble.Flags) {
			f.String("d", "directory", "DEFAULT", "set an alternative directory path")
			f.Bool("v", "verbose", false, "enable verbose mode")
		},
	})
	app.AddCommand(&grumble.Command{
		Name: "cd",
		Help: "cd a path",
		Args: func(a *grumble.Args) {
			a.String("cdpath", "absoulute or relative path")
		},

		Run: func(c *grumble.Context) error {
			// Args.
			err := syscall.Chdir(c.Args.String("cdpath"))
			if err != nil {
				print(err, "wrong path")
			}
			// c.App.Println("cdpath:", c.Args.String("cdpath"))
			c.App.SetPrompt("19281030:" + gocall.Getcwd() + "# ")
			return nil
		},
	})
	app.AddCommand(&grumble.Command{
		Name: "ls",
		Help: "ls a path",

		Flags: func(f *grumble.Flags) {
			f.Duration("t", "timeout", time.Second, "timeout duration")
		},

		Args: func(a *grumble.Args) {
			a.String("lspath", "absoulute or relative path", grumble.Default(""))
		},

		Run: func(c *grumble.Context) error {
			// Args.
			myPath := c.Args.String("lspath")
			// c.App.Println("myPath:", myPath)
			names := ls(myPath)
			fmt.Println(names)
			return nil
		},
	})
	app.AddCommand(&grumble.Command{
		Name: "run",
		Help: "run a external file",

		Flags: func(f *grumble.Flags) {
			f.Duration("t", "timeout", time.Second, "timeout duration")
		},

		Args: func(a *grumble.Args) {
			a.String("file", "file name", grumble.Default(""))
		},

		Run: func(c *grumble.Context) error {
			// Args.
			file := c.Args.String("file")
			// c.App.Println("myPath:", myPath)
			_, err := run(file)
			if err != nil {
				println("run cmd fail")
			}
			print(" ")
			// println("pid: ", pid)
			return nil
		},
	})
	app.AddCommand(&grumble.Command{
		Name: "mkdir",
		Help: "make dir absolute or relative",

		Args: func(a *grumble.Args) {
			a.String("dirpath", "dirpath")
		},

		Run: func(c *grumble.Context) error {
			// Args.
			dirpath := c.Args.String("dirpath")
			// c.App.Println("myPath:", myPath)
			err := syscall.Mkdir(dirpath, syscall.O_WRONLY)
			if err != nil {
				println("mkdir fail")
			}
			return nil
		},
	})
	app.AddCommand(&grumble.Command{
		Name: "creat",
		Help: "create a file",

		Args: func(a *grumble.Args) {
			a.String("filepath", "filepath")
		},

		Run: func(c *grumble.Context) error {
			// Args.
			filepath := c.Args.String("filepath")
			// c.App.Println("myPath:", myPath)
			err := creat(filepath)
			return err
		},
	})
	app.AddCommand(&grumble.Command{
		Name: "cp",
		Help: "copy a filet to dest",

		Args: func(a *grumble.Args) {
			a.String("sourcePath", "sourcepath")
			a.String("destPath", "destpath")
		},

		Run: func(c *grumble.Context) error {
			// Args.
			sourcePath := c.Args.String("sourcePath")
			destPath := c.Args.String("destPath")
			sourcefd, err := syscall.Open(sourcePath, syscall.O_RDONLY, 0)
			if sourcefd == -1 {
				println("source path not found!")
				return err
			}
			destfd, err := syscall.Open(destPath, syscall.O_RDONLY, 0)
			if destfd != -1 {
				println("dest file exists!")
				return err
			}
			destfd, _ = syscall.Creat(destPath, syscall.O_WRONLY)
			for {
				p := make([]byte, 2048)
				n, _ := syscall.Read(sourcefd, p)
				syscall.Write(destfd, p[0:n])
				if n < 2048 {
					break
				}
			}
			return nil
		},
	})
	return app
}
