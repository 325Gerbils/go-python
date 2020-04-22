package python

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var environment string

// Init sets a unique filename for the python environment's state
func Init() {
	rand.Seed(time.Now().UTC().UnixNano())
	environment = fmt.Sprintf("%x", rand.Int()) + ".pkl "
	cmd := exec.Command("python", "-c", `import dill
filename = '`+environment+`'
dill.dump_session(filename)`)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf(err)
	}
}

// Destroy deletes the files storing the python environment
func Destroy() error {
	err := os.RemoveAll(environment)
	return err
}

// Run executes the given python code in a persistent state interpreter
func Run(pycode string) (string, error) {

	execCode := `import dill

filename = '` + environment + `'
dill.load_session(filename)

` + pycode + `

dill.dump_session(filename)`

	cmd := exec.Command("python", "-c", execCode)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
