package cheat

import (
	"github.com/spf13/cobra"

	"github.com/devopshub-org/devopshub-cli/internal/output"
)

// linuxEntry represents a single Linux cheatsheet entry.
type linuxEntry struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// linuxSection groups related Linux commands.
type linuxSection struct {
	Title   string       `json:"title"`
	Entries []linuxEntry `json:"entries"`
}

// linuxCheatsheet is the canonical Linux cheatsheet embedded in the binary.
var linuxCheatsheet = []linuxSection{
	{
		Title: "File & Directory Operations",
		Entries: []linuxEntry{
			{"ls -lah", "List directory contents (long, human-readable, hidden)"},
			{"ls -lhS", "List files sorted by size (largest first)"},
			{"find /path -name '*.log' -mtime +7", "Find .log files older than 7 days"},
			{"find / -type f -size +100M", "Find files larger than 100MB"},
			{"cp -av src/ dst/", "Copy directory verbosely, preserving attributes"},
			{"mv file.txt /new/path/", "Move or rename a file"},
			{"rm -rf /path/to/dir", "Recursively remove directory (use carefully)"},
			{"mkdir -p /a/b/c", "Create nested directories"},
			{"tree -L 2", "Show directory tree 2 levels deep"},
			{"du -sh *", "Show disk usage of items in current directory"},
			{"df -h", "Show filesystem disk space usage"},
			{"stat file.txt", "Show file metadata (permissions, timestamps)"},
		},
	},
	{
		Title: "Text Processing",
		Entries: []linuxEntry{
			{"grep -rn 'pattern' /path", "Recursive case-sensitive search with line numbers"},
			{"grep -ri 'pattern' /path", "Case-insensitive recursive search"},
			{"grep -v 'exclude'", "Invert match (exclude lines)"},
			{"awk '{print $1, $3}'", "Print columns 1 and 3"},
			{"awk -F: '{print $1}' /etc/passwd", "Use custom field separator"},
			{"sed 's/old/new/g' file.txt", "Replace all occurrences of 'old' with 'new'"},
			{"sed -i 's/old/new/g' file.txt", "In-place replacement"},
			{"cut -d',' -f2 file.csv", "Extract column 2 from CSV"},
			{"sort -k2 -n file.txt", "Sort numerically by column 2"},
			{"sort | uniq -c | sort -rn", "Count unique lines, sorted by frequency"},
			{"wc -l file.txt", "Count lines in a file"},
			{"head -n 20 file.txt", "Print first 20 lines"},
			{"tail -n 20 file.txt", "Print last 20 lines"},
			{"tail -f /var/log/syslog", "Follow log file in real-time"},
			{"diff file1.txt file2.txt", "Show differences between two files"},
			{"comm -23 <(sort f1) <(sort f2)", "Lines in f1 but not f2"},
		},
	},
	{
		Title: "Process Management",
		Entries: []linuxEntry{
			{"ps aux", "List all running processes"},
			{"ps aux | grep <name>", "Find process by name"},
			{"pgrep -fl nginx", "Find process IDs by name pattern"},
			{"kill -15 <PID>", "Gracefully terminate process (SIGTERM)"},
			{"kill -9 <PID>", "Force-kill process (SIGKILL)"},
			{"pkill -f 'process-name'", "Kill processes matching a pattern"},
			{"top", "Interactive process viewer"},
			{"htop", "Enhanced interactive process viewer"},
			{"nice -n 10 <command>", "Run command with lower CPU priority"},
			{"renice +5 -p <PID>", "Adjust priority of running process"},
			{"nohup <command> &", "Run command immune to hangups"},
			{"<command> &", "Run command in background"},
			{"jobs", "List background jobs"},
			{"fg %1", "Bring job 1 to foreground"},
			{"wait <PID>", "Wait for process to finish"},
		},
	},
	{
		Title: "Networking",
		Entries: []linuxEntry{
			{"ip addr show", "Show network interfaces and IP addresses"},
			{"ip route show", "Show routing table"},
			{"ss -tulnp", "List listening sockets with process info"},
			{"netstat -tulnp", "Legacy: list listening sockets"},
			{"curl -I https://example.com", "Fetch HTTP response headers only"},
			{"curl -v -o /dev/null https://example.com", "Debug HTTP connection"},
			{"wget -q -O - https://example.com/file", "Download silently to stdout"},
			{"ping -c 4 8.8.8.8", "Send 4 ICMP packets to Google DNS"},
			{"traceroute 8.8.8.8", "Trace network hops"},
			{"nmap -p 22,80,443 <host>", "Scan specific ports"},
			{"dig example.com A", "Query DNS A record"},
			{"dig +short example.com", "Short DNS query output"},
			{"host example.com", "Simple DNS lookup"},
			{"tcpdump -i eth0 port 80", "Capture HTTP traffic on eth0"},
			{"iptables -L -n -v", "List iptables rules verbosely"},
		},
	},
	{
		Title: "System Information",
		Entries: []linuxEntry{
			{"uname -a", "Print all kernel information"},
			{"lsb_release -a", "Show Linux distribution info"},
			{"cat /etc/os-release", "Show OS details"},
			{"uptime", "Show system uptime and load averages"},
			{"free -h", "Show memory usage (human-readable)"},
			{"vmstat 1 5", "Virtual memory stats (5 samples, 1s interval)"},
			{"lscpu", "Show CPU architecture info"},
			{"lsblk", "List block devices"},
			{"lsof -i :8080", "List processes using port 8080"},
			{"lsof -p <PID>", "List files opened by a process"},
			{"dmesg | tail -50", "Show recent kernel messages"},
			{"journalctl -u nginx --since '1 hour ago'", "systemd service logs"},
			{"systemctl status nginx", "Check service status"},
			{"systemctl restart nginx", "Restart a systemd service"},
		},
	},
	{
		Title: "User & Permissions",
		Entries: []linuxEntry{
			{"whoami", "Print current user"},
			{"id", "Show user and group IDs"},
			{"sudo -i", "Switch to root shell"},
			{"su - <user>", "Switch to another user"},
			{"chmod 755 file", "Set rwxr-xr-x permissions"},
			{"chmod -R 644 /path/", "Recursively set permissions"},
			{"chown user:group file", "Change file owner and group"},
			{"chown -R user:group /path/", "Recursively change ownership"},
			{"passwd <user>", "Change user password"},
			{"useradd -m -s /bin/bash <user>", "Create user with home directory"},
			{"usermod -aG sudo <user>", "Add user to sudo group"},
			{"groups <user>", "List user's groups"},
			{"visudo", "Safely edit sudoers file"},
		},
	},
	{
		Title: "Archiving & Compression",
		Entries: []linuxEntry{
			{"tar -czf archive.tar.gz /path/", "Create gzip-compressed archive"},
			{"tar -cjf archive.tar.bz2 /path/", "Create bzip2-compressed archive"},
			{"tar -xzf archive.tar.gz -C /dest/", "Extract gzip archive to /dest/"},
			{"tar -tf archive.tar.gz", "List contents without extracting"},
			{"zip -r archive.zip /path/", "Create zip archive"},
			{"unzip archive.zip -d /dest/", "Extract zip to /dest/"},
			{"gzip file.txt", "Compress file in-place (creates file.txt.gz)"},
			{"gunzip file.txt.gz", "Decompress a gzip file"},
		},
	},
	{
		Title: "SSH & Remote",
		Entries: []linuxEntry{
			{"ssh user@host -p 22", "SSH into remote host"},
			{"ssh -i ~/.ssh/id_rsa user@host", "SSH with specific key"},
			{"ssh -L 8080:localhost:80 user@host", "Local port forwarding tunnel"},
			{"ssh -R 8080:localhost:80 user@host", "Remote port forwarding"},
			{"scp -r /local/path user@host:/remote/", "Copy files to remote over SSH"},
			{"rsync -avz /src/ user@host:/dst/", "Sync files efficiently with compression"},
			{"rsync -avz --delete /src/ /dst/", "Sync and remove deleted files"},
			{"ssh-keygen -t ed25519 -C 'email@example.com'", "Generate Ed25519 SSH key pair"},
			{"ssh-copy-id user@host", "Copy public key to remote host"},
		},
	},
}

// linuxSubCmd implements `devopshub cheat linux`.
type linuxSubCmd struct {
	cmd     *cobra.Command
	printer *output.Printer
}

func newLinuxCmd(printer *output.Printer) *linuxSubCmd {
	l := &linuxSubCmd{printer: printer}

	l.cmd = &cobra.Command{
		Use:     "linux",
		Short:   "Essential Linux commands cheatsheet",
		Aliases: []string{"unix", "shell"},
		Long: `A curated, categorized cheatsheet of essential Linux commands for DevOps engineers.

Covers file operations, text processing, process management, networking,
system info, user management, archiving, and SSH.

Examples:
  devopshub cheat linux
  devopshub cheat shell
  devopshub cheat linux --json`,
		Args: cobra.NoArgs,
		RunE: l.run,
	}

	return l
}

func (l *linuxSubCmd) run(_ *cobra.Command, _ []string) error {
	p := l.printer

	if p.JSONMode() {
		return p.JSON(linuxCheatsheet)
	}

	p.Header("Linux Cheatsheet for DevOps Engineers")
	p.Println("")

	for _, section := range linuxCheatsheet {
		p.Section(section.Title)
		for _, entry := range section.Entries {
			p.Command(entry.Command, entry.Description)
		}
	}

	p.Println("")
	p.Info("Tip: Use --json to get machine-readable output")
	return nil
}
