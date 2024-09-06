package main

// AllRepos - Clone all your repositories and apply sweeping changes
// Homepage: https://github.com/asottile/all-repos

import (
	"fmt"
	
	"os/exec"
)

func installAllRepos() {
	// Método 1: Descargar y extraer .tar.gz
	allrepos_tar_url := "https://files.pythonhosted.org/packages/a6/56/29006be2546b897a5c62a3d4a7e613abf5a3533554d948b0e0af27546f1b/all_repos-1.27.0.tar.gz"
	allrepos_cmd_tar := exec.Command("curl", "-L", allrepos_tar_url, "-o", "package.tar.gz")
	err := allrepos_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	allrepos_zip_url := "https://files.pythonhosted.org/packages/a6/56/29006be2546b897a5c62a3d4a7e613abf5a3533554d948b0e0af27546f1b/all_repos-1.27.0.zip"
	allrepos_cmd_zip := exec.Command("curl", "-L", allrepos_zip_url, "-o", "package.zip")
	err = allrepos_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	allrepos_bin_url := "https://files.pythonhosted.org/packages/a6/56/29006be2546b897a5c62a3d4a7e613abf5a3533554d948b0e0af27546f1b/all_repos-1.27.0.bin"
	allrepos_cmd_bin := exec.Command("curl", "-L", allrepos_bin_url, "-o", "binary.bin")
	err = allrepos_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	allrepos_src_url := "https://files.pythonhosted.org/packages/a6/56/29006be2546b897a5c62a3d4a7e613abf5a3533554d948b0e0af27546f1b/all_repos-1.27.0.src.tar.gz"
	allrepos_cmd_src := exec.Command("curl", "-L", allrepos_src_url, "-o", "source.tar.gz")
	err = allrepos_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	allrepos_cmd_direct := exec.Command("./binary")
	err = allrepos_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
