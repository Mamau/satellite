package collector

import (
	"github.com/gookit/color"
	"os"
	"os/exec"
	"regexp"
	"satellite/internal/config"
	"satellite/internal/entity"
	"satellite/internal/validator"
	"satellite/pkg"
	"strings"
)

type Service struct {
	config *config.Config
}

func NewService(c *config.Config) *Service {
	return &Service{
		config: c,
	}
}

func (s *Service) GetMacrosList() []config.Macros {
	return s.config.GetMacros()
}

func (s *Service) GetMacros(name string) *config.Macros {
	for _, v := range s.config.Macros {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (s *Service) ServicesList() map[string]entity.Runner {
	data := make(map[string]entity.Runner)

	for _, v := range s.config.GetDockerCompose().GetCommands() {
		s.checkCommandConflict(data, v)
		data[v.GetName()] = v
	}

	for _, v := range s.config.GetDocker().GetCommands() {
		s.checkCommandConflict(data, v)
		data[v.GetName()] = v
	}

	return data
}

func (s *Service) checkCommandConflict(data map[string]entity.Runner, e entity.Runner) {
	if _, ok := data[e.GetName()]; ok {
		color.Red.Printf("Conflict command name %q\n", e.GetName())
		os.Exit(1)
	}
}

// FindCommand - FindService был
func (s *Service) FindCommand(name string) entity.Runner {
	var service entity.Runner

	for _, v := range s.ServicesList() {
		if v.GetName() == name {
			service = v
			break
		}
	}
	if service == nil {
		return service
	}

	valid := validator.NewValidator()
	errs, isValid := valid.Validate(service)
	if isValid {
		return service
	}
	for _, v := range errs {
		color.Red.Printf("Service %s error: %s\n", service.GetName(), v)
	}
	os.Exit(1)
	return nil
}

func (s *Service) GetMacrosCommands(macrosList []string) [][]string {
	var commandList [][]string
	for _, v := range macrosList {
		cml := strings.Split(v, " ")
		if serviceName := s.FindCommand(cml[0]); serviceName != nil {
			commandList = append(commandList, cml)
			continue
		}
		color.Danger.Printf("Service %q not found\n", cml[0])
		os.Exit(1)
	}

	return commandList
}

func (s *Service) ExecuteCommand(strategy entity.Runner, args []string) *exec.Cmd {
	replacedEnv := pkg.ReplaceEnvVariables(strategy.ToCommand(args))
	replacedPwd := pkg.ReplaceInternalVariables("\\$(\\(pwd\\))", pkg.GetPwd(), replacedEnv)
	replaceGateWay := s.getReplaceGateWay(replacedPwd)

	dcCommand := exec.Command(strategy.GetExecCommand(), replaceGateWay...)
	color.Info.Printf("Running command: %v\n", dcCommand.String())
	return dcCommand
}

func (s *Service) getReplaceGateWay(data []string) []string {
	from := "\\$(\\(gatewayHost\\))"
	r := regexp.MustCompile(from)
	if found := r.Find([]byte(strings.Join(data, " "))); found == nil {
		return data
	}

	inspectData := pkg.DockerExec([]string{"network", "inspect", "bridge"})
	host := pkg.RetrieveGatewayHost(inspectData)
	return pkg.ReplaceInternalVariables(from, host, data)
}
