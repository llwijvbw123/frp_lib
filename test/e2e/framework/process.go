package framework

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	flog "frp_lib/pkg/util/log"
	"frp_lib/test/e2e/pkg/process"
)

// RunProcesses run multiple processes from templates.
// The first template should always be frps.
func (f *Framework) RunProcesses(serverTemplates []string, clientTemplates []string) ([]*process.Process, []*process.Process) {
	templates := make([]string, 0, len(serverTemplates)+len(clientTemplates))
	templates = append(templates, serverTemplates...)
	templates = append(templates, clientTemplates...)
	outs, ports, err := f.RenderTemplates(templates)
	ExpectNoError(err)
	ExpectTrue(len(templates) > 0)

	for name, port := range ports {
		f.usedPorts[name] = port
	}

	currentServerProcesses := make([]*process.Process, 0, len(serverTemplates))
	for i := range serverTemplates {
		path := filepath.Join(f.TempDirectory, fmt.Sprintf("frp-e2e-server-%d", i))
		err = os.WriteFile(path, []byte(outs[i]), 0o666)
		ExpectNoError(err)
		flog.Trace("[%s] %s", path, outs[i])

		p := process.NewWithEnvs(TestContext.FRPServerPath, []string{"-c", path}, f.osEnvs)
		f.serverConfPaths = append(f.serverConfPaths, path)
		f.serverProcesses = append(f.serverProcesses, p)
		currentServerProcesses = append(currentServerProcesses, p)
		err = p.Start()
		ExpectNoError(err)
	}
	time.Sleep(1 * time.Second)

	currentClientProcesses := make([]*process.Process, 0, len(clientTemplates))
	for i := range clientTemplates {
		index := i + len(serverTemplates)
		path := filepath.Join(f.TempDirectory, fmt.Sprintf("frp-e2e-client-%d", i))
		err = os.WriteFile(path, []byte(outs[index]), 0o666)
		ExpectNoError(err)
		flog.Trace("[%s] %s", path, outs[index])

		p := process.NewWithEnvs(TestContext.FRPClientPath, []string{"-c", path}, f.osEnvs)
		f.clientConfPaths = append(f.clientConfPaths, path)
		f.clientProcesses = append(f.clientProcesses, p)
		currentClientProcesses = append(currentClientProcesses, p)
		err = p.Start()
		ExpectNoError(err)
		time.Sleep(500 * time.Millisecond)
	}
	time.Sleep(3 * time.Second)

	return currentServerProcesses, currentClientProcesses
}

func (f *Framework) RunFrps(args ...string) (*process.Process, string, error) {
	p := process.NewWithEnvs(TestContext.FRPServerPath, args, f.osEnvs)
	f.serverProcesses = append(f.serverProcesses, p)
	err := p.Start()
	if err != nil {
		return p, p.StdOutput(), err
	}
	// sleep for a while to get std output
	time.Sleep(time.Second)
	return p, p.StdOutput(), nil
}

func (f *Framework) RunFrpc(args ...string) (*process.Process, string, error) {
	p := process.NewWithEnvs(TestContext.FRPClientPath, args, f.osEnvs)
	f.clientProcesses = append(f.clientProcesses, p)
	err := p.Start()
	if err != nil {
		return p, p.StdOutput(), err
	}
	time.Sleep(time.Second)
	return p, p.StdOutput(), nil
}

func (f *Framework) GenerateConfigFile(content string) string {
	f.configFileIndex++
	path := filepath.Join(f.TempDirectory, fmt.Sprintf("frp-e2e-config-%d", f.configFileIndex))
	err := os.WriteFile(path, []byte(content), 0o666)
	ExpectNoError(err)
	return path
}
