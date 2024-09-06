package main

// Glab - Open-source GitLab command-line tool
// Homepage: https://gitlab.com/gitlab-org/cli

import (
	"fmt"
	
	"os/exec"
)

func installGlab() {
	// Método 1: Descargar y extraer .tar.gz
	glab_tar_url := "https://gitlab.com/gitlab-org/cli/-/archive/v1.46.0/cli-v1.46.0.tar.gz"
	glab_cmd_tar := exec.Command("curl", "-L", glab_tar_url, "-o", "package.tar.gz")
	err := glab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glab_zip_url := "https://gitlab.com/gitlab-org/cli/-/archive/v1.46.0/cli-v1.46.0.zip"
	glab_cmd_zip := exec.Command("curl", "-L", glab_zip_url, "-o", "package.zip")
	err = glab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glab_bin_url := "https://gitlab.com/gitlab-org/cli/-/archive/v1.46.0/cli-v1.46.0.bin"
	glab_cmd_bin := exec.Command("curl", "-L", glab_bin_url, "-o", "binary.bin")
	err = glab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glab_src_url := "https://gitlab.com/gitlab-org/cli/-/archive/v1.46.0/cli-v1.46.0.src.tar.gz"
	glab_cmd_src := exec.Command("curl", "-L", glab_src_url, "-o", "source.tar.gz")
	err = glab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glab_cmd_direct := exec.Command("./binary")
	err = glab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
