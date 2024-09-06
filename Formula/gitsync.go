package main

// GitSync - Clones a git repository and keeps it synchronized with the upstream
// Homepage: https://github.com/kubernetes/git-sync

import (
	"fmt"
	
	"os/exec"
)

func installGitSync() {
	// Método 1: Descargar y extraer .tar.gz
	gitsync_tar_url := "https://github.com/kubernetes/git-sync/archive/refs/tags/v4.2.4.tar.gz"
	gitsync_cmd_tar := exec.Command("curl", "-L", gitsync_tar_url, "-o", "package.tar.gz")
	err := gitsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitsync_zip_url := "https://github.com/kubernetes/git-sync/archive/refs/tags/v4.2.4.zip"
	gitsync_cmd_zip := exec.Command("curl", "-L", gitsync_zip_url, "-o", "package.zip")
	err = gitsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitsync_bin_url := "https://github.com/kubernetes/git-sync/archive/refs/tags/v4.2.4.bin"
	gitsync_cmd_bin := exec.Command("curl", "-L", gitsync_bin_url, "-o", "binary.bin")
	err = gitsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitsync_src_url := "https://github.com/kubernetes/git-sync/archive/refs/tags/v4.2.4.src.tar.gz"
	gitsync_cmd_src := exec.Command("curl", "-L", gitsync_src_url, "-o", "source.tar.gz")
	err = gitsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitsync_cmd_direct := exec.Command("./binary")
	err = gitsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
}
