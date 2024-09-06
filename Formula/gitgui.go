package main

// GitGui - Tcl/Tk UI for the git revision control system
// Homepage: https://git-scm.com

import (
	"fmt"
	
	"os/exec"
)

func installGitGui() {
	// Método 1: Descargar y extraer .tar.gz
	gitgui_tar_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	gitgui_cmd_tar := exec.Command("curl", "-L", gitgui_tar_url, "-o", "package.tar.gz")
	err := gitgui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitgui_zip_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	gitgui_cmd_zip := exec.Command("curl", "-L", gitgui_zip_url, "-o", "package.zip")
	err = gitgui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitgui_bin_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	gitgui_cmd_bin := exec.Command("curl", "-L", gitgui_bin_url, "-o", "binary.bin")
	err = gitgui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitgui_src_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	gitgui_cmd_src := exec.Command("curl", "-L", gitgui_src_url, "-o", "source.tar.gz")
	err = gitgui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitgui_cmd_direct := exec.Command("./binary")
	err = gitgui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tcl-tk")
	exec.Command("latte", "install", "tcl-tk").Run()
}
