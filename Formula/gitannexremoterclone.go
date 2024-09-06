package main

// GitAnnexRemoteRclone - Use rclone supported cloud storage with git-annex
// Homepage: https://github.com/git-annex-remote-rclone/git-annex-remote-rclone

import (
	"fmt"
	
	"os/exec"
)

func installGitAnnexRemoteRclone() {
	// Método 1: Descargar y extraer .tar.gz
	gitannexremoterclone_tar_url := "https://github.com/git-annex-remote-rclone/git-annex-remote-rclone/archive/refs/tags/v0.8.tar.gz"
	gitannexremoterclone_cmd_tar := exec.Command("curl", "-L", gitannexremoterclone_tar_url, "-o", "package.tar.gz")
	err := gitannexremoterclone_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitannexremoterclone_zip_url := "https://github.com/git-annex-remote-rclone/git-annex-remote-rclone/archive/refs/tags/v0.8.zip"
	gitannexremoterclone_cmd_zip := exec.Command("curl", "-L", gitannexremoterclone_zip_url, "-o", "package.zip")
	err = gitannexremoterclone_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitannexremoterclone_bin_url := "https://github.com/git-annex-remote-rclone/git-annex-remote-rclone/archive/refs/tags/v0.8.bin"
	gitannexremoterclone_cmd_bin := exec.Command("curl", "-L", gitannexremoterclone_bin_url, "-o", "binary.bin")
	err = gitannexremoterclone_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitannexremoterclone_src_url := "https://github.com/git-annex-remote-rclone/git-annex-remote-rclone/archive/refs/tags/v0.8.src.tar.gz"
	gitannexremoterclone_cmd_src := exec.Command("curl", "-L", gitannexremoterclone_src_url, "-o", "source.tar.gz")
	err = gitannexremoterclone_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitannexremoterclone_cmd_direct := exec.Command("./binary")
	err = gitannexremoterclone_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: git-annex")
	exec.Command("latte", "install", "git-annex").Run()
	fmt.Println("Instalando dependencia: rclone")
	exec.Command("latte", "install", "rclone").Run()
}
