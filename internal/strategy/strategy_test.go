package strategy

import (
	"context"
	"strings"
	"testing"

	"github.com/mamau/satellite/pkg"

	config2 "github.com/mamau/satellite/internal/config"
	docker2 "github.com/mamau/satellite/internal/config/docker"
)

func TestToCommand(t *testing.T) {
	pullStrategyToCommand(t)
	daemonStrategyToCommand(t)
	runStrategyToCommand(t)
}

func TestClientCommand(t *testing.T) {
	c := setConfig().GetService("composer")
	ctx := context.WithValue(context.Background(), "type", RunType)
	strategy := createRunStrategy(ctx, c, []string{"install --ignore-platform-reqs"})
	e := "composer install --ignore-platform-reqs"
	result := strategy.clientCommand()
	if e != strings.Join(result, " ") {
		t.Errorf("error on service type %q, expected: %q, got %q", strategy.GetContext().Value("type"), e, strategy.getArgs())
	}

	c = setConfig().GetService("composer-2")
	strategy = createRunStrategy(ctx, c, []string{"install --ignore-platform-reqs"})
	e = "/bin/bash -c git config --global http.sslVerify false; composer config -g http-basic.gitlab.com {GITLAB_USERNAME} {GITLAB_TOKEN}; composer install --ignore-platform-reqs; chown -R 501:501 /home/www-data"
	result = strategy.clientCommand()
	if e != strings.Join(result, " ") {
		t.Errorf("error on service type %q, expected: %q, got %q", strategy.GetContext().Value("type"), e, strategy.getArgs())
	}

	c = setConfig().GetService("composer-2")
	strategy = createRunStrategy(ctx, c, []string{"install --ignore-platform-reqs"})
	c.PreCommands = []string{}
	e = "/bin/bash -c composer install --ignore-platform-reqs; chown -R 501:501 /home/www-data"
	result = strategy.clientCommand()
	if e != strings.Join(result, " ") {
		t.Errorf("error on service type %q, expected: %q, got %q", strategy.GetContext().Value("type"), e, strategy.getArgs())
	}

	c = setConfig().GetService("composer-2")
	strategy = createRunStrategy(ctx, c, []string{"install --ignore-platform-reqs"})
	c.PreCommands = []string{"any command", "command 2"}
	c.PostCommands = []string{}
	e = "/bin/bash -c any command; command 2; composer install --ignore-platform-reqs"
	result = strategy.clientCommand()
	if e != strings.Join(result, " ") {
		t.Errorf("error on service type %q, expected: %q, got %q", strategy.GetContext().Value("type"), e, strategy.getArgs())
	}
}

func TestGetArgs(t *testing.T) {
	c := setConfig().GetService("composer")
	ctx := context.WithValue(context.Background(), "type", RunType)
	strategy := createRunStrategy(ctx, c, []string{"install --ignore-platform-reqs"})
	e := "install --ignore-platform-reqs"
	if e != strings.Join(strategy.getArgs(), " ") {
		t.Errorf("error get args on service type %q, expected: %q, got %q", strategy.GetContext().Value("type"), e, strategy.getArgs())
	}

	c = setConfig().GetService("composer-2")
	strategy = createRunStrategy(ctx, c, []string{"install --ignore-platform-reqs"})
	e = "composer install --ignore-platform-reqs"
	if e != strings.Join(strategy.getArgs(), " ") {
		t.Errorf("error get args on service type %q, expected: %q, got %q", strategy.GetContext().Value("type"), e, strategy.getArgs())
	}
}

func runStrategyToCommand(t *testing.T) {
	s := createStrategy("composer", RunType, []string{"--version"})
	result := s.ToCommand()
	e := "run -ti -u 501 --workdir=/home/www-data -v $(pwd):/home/www-data -v $(pwd)/cache:/tmp composer:2 composer --version"
	if e != strings.Join(result, " ") {
		name := s.GetContext().Value("type")
		t.Errorf("error to command %q service, expected: %q, got %q", name, e, strings.Join(result, " "))
	}

	s = createStrategy("composer-2", RunType, []string{"install", "--ignore-platform-reqs"})
	result = s.ToCommand()
	e = "run -ti -u 501 --workdir=/home/www-data -v $(pwd):/home/www-data -v $(pwd)/cache:/tmp composer-2:1.10 /bin/bash -c git config --global http.sslVerify false; composer config -g http-basic.gitlab.com {GITLAB_USERNAME} {GITLAB_TOKEN}; composer install --ignore-platform-reqs; chown -R 501:501 /home/www-data"
	if e != strings.Join(result, " ") {
		name := s.GetContext().Value("type")
		t.Errorf("error to command %q service, expected:\n %q,\n got:\n %q", name, e, strings.Join(result, " "))
	}
}

func daemonStrategyToCommand(t *testing.T) {
	s := createStrategy("my-image", DaemonType, []string{})
	e := "run -d --rm -e PHP_IDE_CONFIG=serverName=192.168.0.1 -p 127.0.0.1:443:443 -p 80:80 --dns=8.8.8.8 -v $(pwd):/home/www gitlab.com/my/image"
	result := s.ToCommand()
	if e != strings.Join(result, " ") {
		name := s.GetContext().Value("type")
		t.Errorf("error to command %q service, expected: %q, got %q", name, e, strings.Join(result, " "))
	}

	s = createStrategy("my-image-2", DaemonType, []string{})
	e = "run -d -v $(pwd):/home/www any-image:2"
	result = s.ToCommand()
	if e != strings.Join(result, " ") {
		name := s.GetContext().Value("type")
		t.Errorf("error to command %q service, expected: %q, got %q", name, e, strings.Join(result, " "))
	}
}

func pullStrategyToCommand(t *testing.T) {
	s := createStrategy("fresh-img", PullType, []string{})
	e := "pull node:10"
	result := s.ToCommand()

	if e != strings.Join(result, " ") {
		name := s.GetContext().Value("type")
		t.Errorf("error to command %q service, expected: %q, got %q", name, e, strings.Join(result, " "))
	}

	s = createStrategy("fresh-img-2", PullType, []string{})
	e = "pull node"
	result = s.ToCommand()

	if e != strings.Join(result, " ") {
		name := s.GetContext().Value("type")
		t.Errorf("error to command %q service, expected: %q, got %q", name, e, strings.Join(result, " "))
	}
}

func createStrategy(name, t string, args []string) Strategy {
	c := setConfig().GetService(name)
	ctx := context.WithValue(context.Background(), "type", t)

	switch t {
	case PullType:
		return NewPullStrategy(ctx, c)
	case DaemonType:
		return NewDaemonStrategy(ctx, c)
	case RunType:
		fallthrough
	default:
		return createRunStrategy(ctx, c, args)
	}
}

func createRunStrategy(ctx context.Context, c *docker2.Docker, args []string) *RunStrategy {
	return NewRunStrategy(ctx, c, args)
}

func setConfig() *config2.Config {
	config2.NewConfig(pkg.GetPwd() + "/testdata/satellite")
	c := config2.GetConfig()
	return c
}
