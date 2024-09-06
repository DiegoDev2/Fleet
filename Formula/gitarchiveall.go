package main

// GitArchiveAll - Archive a project and its submodules
// Homepage: https://github.com/Kentzo/git-archive-all

import (
	"fmt"
	
	"os/exec"
)

func installGitArchiveAll() {
	// Método 1: Descargar y extraer .tar.gz
	gitarchiveall_tar_url := "https://github.com/Kentzo/git-archive-all/archive/refs/tags/1.23.1.tar.gz"
	gitarchiveall_cmd_tar := exec.Command("curl", "-L", gitarchiveall_tar_url, "-o", "package.tar.gz")
	err := gitarchiveall_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitarchiveall_zip_url := "https://github.com/Kentzo/git-archive-all/archive/refs/tags/1.23.1.zip"
	gitarchiveall_cmd_zip := exec.Command("curl", "-L", gitarchiveall_zip_url, "-o", "package.zip")
	err = gitarchiveall_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitarchiveall_bin_url := "https://github.com/Kentzo/git-archive-all/archive/refs/tags/1.23.1.bin"
	gitarchiveall_cmd_bin := exec.Command("curl", "-L", gitarchiveall_bin_url, "-o", "binary.bin")
	err = gitarchiveall_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitarchiveall_src_url := "https://github.com/Kentzo/git-archive-all/archive/refs/tags/1.23.1.src.tar.gz"
	gitarchiveall_cmd_src := exec.Command("curl", "-L", gitarchiveall_src_url, "-o", "source.tar.gz")
	err = gitarchiveall_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitarchiveall_cmd_direct := exec.Command("./binary")
	err = gitarchiveall_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
