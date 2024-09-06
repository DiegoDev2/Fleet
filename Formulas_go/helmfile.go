package main

// Helmfile - Deploy Kubernetes Helm Charts
// Homepage: https://github.com/helmfile/helmfile

import (
	"fmt"
	
	"os/exec"
)

func installHelmfile() {
	// Método 1: Descargar y extraer .tar.gz
	helmfile_tar_url := "https://github.com/helmfile/helmfile/archive/refs/tags/v0.167.1.tar.gz"
	helmfile_cmd_tar := exec.Command("curl", "-L", helmfile_tar_url, "-o", "package.tar.gz")
	err := helmfile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	helmfile_zip_url := "https://github.com/helmfile/helmfile/archive/refs/tags/v0.167.1.zip"
	helmfile_cmd_zip := exec.Command("curl", "-L", helmfile_zip_url, "-o", "package.zip")
	err = helmfile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	helmfile_bin_url := "https://github.com/helmfile/helmfile/archive/refs/tags/v0.167.1.bin"
	helmfile_cmd_bin := exec.Command("curl", "-L", helmfile_bin_url, "-o", "binary.bin")
	err = helmfile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	helmfile_src_url := "https://github.com/helmfile/helmfile/archive/refs/tags/v0.167.1.src.tar.gz"
	helmfile_cmd_src := exec.Command("curl", "-L", helmfile_src_url, "-o", "source.tar.gz")
	err = helmfile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	helmfile_cmd_direct := exec.Command("./binary")
	err = helmfile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: helm")
exec.Command("latte", "install", "helm")
}
