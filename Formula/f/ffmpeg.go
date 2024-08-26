package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installFFmpeg() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installFFmpegForMac()
	case "linux":
		installFFmpegForLinux()
	case "windows":
		installFFmpegForWindows()
	}
}

func installFFmpegForMac() {
	url := "https://evermeet.cx/ffmpeg/ffmpeg-4.3.2.zip"
	downloadAndInstall(url, "ffmpeg.zip", "unzip", "ffmpeg.zip", "&&", "sudo", "mv", "ffmpeg", "/usr/local/bin/")
}

func installFFmpegForLinux() {
	url := "https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz"
	downloadAndInstall(url, "ffmpeg.tar.xz", "tar", "-xvJf", "ffmpeg.tar.xz", "&&", "sudo", "mv", "ffmpeg-*/ffmpeg", "/usr/local/bin/")
}

func installFFmpegForWindows() {
	url := "https://ffmpeg.zeranoe.com/builds/win64/static/ffmpeg-20200831-4a11a6f-win64-static.zip"
	downloadAndInstall(url, "ffmpeg.zip", "unzip", "ffmpeg.zip", "-d", "C:\\ffmpeg", "&&", "set", "PATH=%PATH%;C:\\ffmpeg\\bin")
}

func downloadAndInstall(url, filename string, installCommand ...string) {
	downloadCmd := exec.Command("curl", "-o", filename, url)
	runCommand(downloadCmd)
	runCommand(exec.Command(installCommand[0], installCommand[1:]...))
}

func runCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	installFFmpeg()
}
