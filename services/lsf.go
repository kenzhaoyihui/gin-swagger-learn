package services

import (
	"fmt"
	"ginlearn/db"
	"ginlearn/models"
	"os"
	"os/exec"
	"strings"
)

func SubmitJob(jobName, cwd, cmds string) (string, error) {
	var err error
	args := []string{""}
	if jobName != "" {
		args = append(args, fmt.Sprintf("JOB_NAME=%s", jobName))
	}
	if cwd != "" {
		args = append(args, fmt.Sprintf("CURRENT_WORK_DIR=%s", cwd))
	}
	if cmds != "" {
		args = append(args, fmt.Sprintf("COMMAND=%s", cmds))
	}

	cmd := exec.Command("/bin/bash", "-c", "/Users/hyzsherry/go/src/ginlearn/test.sh")
	fmt.Println("before: ", cmd.Env)

	changeYourCmdEnvironment(cmd, args)
	fmt.Println("after: ", cmd.Env)

	output, err := cmd.Output()
	job := &models.Job{
		JobID: string(output),
		JobName: jobName,
		WorkDir: cwd,
		Command: cmds,
	}

	db.InsertJob(job)
	return string(output), err
}

func changeYourCmdEnvironment(cmd *exec.Cmd, cmdEnv []string) error {
	env := os.Environ()

	for _, e := range env {
		i := strings.Index(e, "=")
		if i > 0 && (e[:i] == "ENV_NAME") {

		} else {
			cmdEnv = append(cmdEnv, e)
		}
	}

	cmd.Env = cmdEnv

	return nil
}
