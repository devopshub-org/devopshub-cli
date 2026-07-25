package cheat

import (
	"github.com/spf13/cobra"

	"github.com/devopshub-org/devopshub-cli/internal/output"
)

// dockerEntry represents a single Docker cheatsheet entry.
type dockerEntry struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// dockerSection groups related Docker commands.
type dockerSection struct {
	Title   string        `json:"title"`
	Entries []dockerEntry `json:"entries"`
}

// dockerCheatsheet is the canonical Docker cheatsheet data embedded in the binary.
var dockerCheatsheet = []dockerSection{
	{
		Title: "Container Lifecycle",
		Entries: []dockerEntry{
			{"docker run -d --name <name> <image>", "Run container in detached mode"},
			{"docker run -it <image> /bin/bash", "Run container interactively"},
			{"docker run --rm <image>", "Run and auto-remove on exit"},
			{"docker start <container>", "Start a stopped container"},
			{"docker stop <container>", "Gracefully stop a running container"},
			{"docker kill <container>", "Force-kill a running container"},
			{"docker restart <container>", "Restart a container"},
			{"docker rm <container>", "Remove a stopped container"},
			{"docker rm -f <container>", "Force-remove a running container"},
		},
	},
	{
		Title: "Container Inspection",
		Entries: []dockerEntry{
			{"docker ps", "List running containers"},
			{"docker ps -a", "List all containers (including stopped)"},
			{"docker logs -f <container>", "Follow container logs"},
			{"docker logs --tail 100 <container>", "Show last 100 log lines"},
			{"docker inspect <container>", "Show detailed container metadata (JSON)"},
			{"docker stats", "Live resource usage (CPU, memory, net I/O)"},
			{"docker top <container>", "Show processes running inside container"},
			{"docker diff <container>", "Show filesystem changes in container"},
		},
	},
	{
		Title: "Images",
		Entries: []dockerEntry{
			{"docker images", "List local images"},
			{"docker pull <image>:<tag>", "Pull image from registry"},
			{"docker push <image>:<tag>", "Push image to registry"},
			{"docker build -t <name>:<tag> .", "Build image from Dockerfile"},
			{"docker build --no-cache -t <name> .", "Build without layer cache"},
			{"docker rmi <image>", "Remove a local image"},
			{"docker image prune -a", "Remove all unused images"},
			{"docker history <image>", "Show image layer history"},
			{"docker save -o image.tar <image>", "Export image to tar archive"},
			{"docker load -i image.tar", "Import image from tar archive"},
		},
	},
	{
		Title: "Networking",
		Entries: []dockerEntry{
			{"docker network ls", "List networks"},
			{"docker network create <name>", "Create a user-defined bridge network"},
			{"docker network inspect <network>", "Inspect network details"},
			{"docker network connect <network> <container>", "Connect container to network"},
			{"docker network disconnect <network> <container>", "Disconnect container from network"},
			{"docker run -p 8080:80 <image>", "Map host port 8080 → container port 80"},
			{"docker run --network host <image>", "Use host networking"},
		},
	},
	{
		Title: "Volumes",
		Entries: []dockerEntry{
			{"docker volume ls", "List volumes"},
			{"docker volume create <name>", "Create a named volume"},
			{"docker volume inspect <name>", "Inspect volume details"},
			{"docker volume rm <name>", "Remove a volume"},
			{"docker volume prune", "Remove all unused volumes"},
			{"docker run -v myvolume:/data <image>", "Mount named volume at /data"},
			{"docker run -v $(pwd):/app <image>", "Bind-mount current directory"},
		},
	},
	{
		Title: "Docker Compose",
		Entries: []dockerEntry{
			{"docker compose up -d", "Start all services in detached mode"},
			{"docker compose down", "Stop and remove containers/networks"},
			{"docker compose down -v", "Also remove volumes"},
			{"docker compose logs -f", "Follow logs from all services"},
			{"docker compose ps", "List compose service containers"},
			{"docker compose exec <svc> /bin/sh", "Open shell in running service"},
			{"docker compose build", "Build (or rebuild) services"},
			{"docker compose pull", "Pull latest images for services"},
			{"docker compose restart <svc>", "Restart a specific service"},
		},
	},
	{
		Title: "System & Cleanup",
		Entries: []dockerEntry{
			{"docker system df", "Show disk usage by Docker resources"},
			{"docker system prune", "Remove stopped containers, dangling images, unused networks"},
			{"docker system prune -a --volumes", "Full cleanup (all unused resources)"},
			{"docker events", "Stream real-time events from the Docker daemon"},
			{"docker info", "Display system-wide Docker information"},
			{"docker version", "Show Docker client and server version"},
		},
	},
}

// dockerSubCmd implements `devopshub cheat docker`.
type dockerSubCmd struct {
	cmd     *cobra.Command
	printer *output.Printer
}

func newDockerCmd(printer *output.Printer) *dockerSubCmd {
	d := &dockerSubCmd{printer: printer}

	d.cmd = &cobra.Command{
		Use:   "docker",
		Short: "Docker commands cheatsheet",
		Long: `A curated, categorized cheatsheet of essential Docker commands.

Covers container lifecycle, inspection, images, networking, volumes,
Docker Compose, and system maintenance.

Examples:
  devopshub cheat docker
  devopshub cheat docker --json`,
		Args: cobra.NoArgs,
		RunE: d.run,
	}

	return d
}

func (d *dockerSubCmd) run(_ *cobra.Command, _ []string) error {
	p := d.printer

	// JSON mode: emit the full structured dataset.
	if p.JSONMode() {
		return p.JSON(dockerCheatsheet)
	}

	p.Header("Docker Cheatsheet")
	p.Println("")

	for _, section := range dockerCheatsheet {
		p.Section(section.Title)
		for _, entry := range section.Entries {
			p.Command(entry.Command, entry.Description)
		}
	}

	p.Println("")
	p.Info("Tip: Use --json to get machine-readable output")
	return nil
}
