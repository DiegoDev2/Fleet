package main

// GitRemoteHg - Transparent bidirectional bridge between Git and Mercurial
// Homepage: https://github.com/felipec/git-remote-hg

import (
	"fmt"
	
	"os/exec"
)

func installGitRemoteHg() {
	// Método 1: Descargar y extraer .tar.gz
	gitremotehg_tar_url := "https://github.com/felipec/git-remote-hg/archive/refs/tags/v0.6.tar.gz"
	gitremotehg_cmd_tar := exec.Command("curl", "-L", gitremotehg_tar_url, "-o", "package.tar.gz")
	err := gitremotehg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitremotehg_zip_url := "https://github.com/felipec/git-remote-hg/archive/refs/tags/v0.6.zip"
	gitremotehg_cmd_zip := exec.Command("curl", "-L", gitremotehg_zip_url, "-o", "package.zip")
	err = gitremotehg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitremotehg_bin_url := "https://github.com/felipec/git-remote-hg/archive/refs/tags/v0.6.bin"
	gitremotehg_cmd_bin := exec.Command("curl", "-L", gitremotehg_bin_url, "-o", "binary.bin")
	err = gitremotehg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitremotehg_src_url := "https://github.com/felipec/git-remote-hg/archive/refs/tags/v0.6.src.tar.gz"
	gitremotehg_cmd_src := exec.Command("curl", "-L", gitremotehg_src_url, "-o", "source.tar.gz")
	err = gitremotehg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitremotehg_cmd_direct := exec.Command("./binary")
	err = gitremotehg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
	exec.Command("latte", "install", "asciidoctor").Run()
	fmt.Println("Instalando dependencia: mercurial")
	exec.Command("latte", "install", "mercurial").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
