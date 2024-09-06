package main

// DependencyCheck - OWASP dependency-check
// Homepage: https://owasp.org/www-project-dependency-check/

import (
	"fmt"
	
	"os/exec"
)

func installDependencyCheck() {
	// Método 1: Descargar y extraer .tar.gz
	dependencycheck_tar_url := "https://github.com/jeremylong/DependencyCheck/releases/download/v10.0.2/dependency-check-10.0.2-release.zip"
	dependencycheck_cmd_tar := exec.Command("curl", "-L", dependencycheck_tar_url, "-o", "package.tar.gz")
	err := dependencycheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dependencycheck_zip_url := "https://github.com/jeremylong/DependencyCheck/releases/download/v10.0.2/dependency-check-10.0.2-release.zip"
	dependencycheck_cmd_zip := exec.Command("curl", "-L", dependencycheck_zip_url, "-o", "package.zip")
	err = dependencycheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dependencycheck_bin_url := "https://github.com/jeremylong/DependencyCheck/releases/download/v10.0.2/dependency-check-10.0.2-release.zip"
	dependencycheck_cmd_bin := exec.Command("curl", "-L", dependencycheck_bin_url, "-o", "binary.bin")
	err = dependencycheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dependencycheck_src_url := "https://github.com/jeremylong/DependencyCheck/releases/download/v10.0.2/dependency-check-10.0.2-release.zip"
	dependencycheck_cmd_src := exec.Command("curl", "-L", dependencycheck_src_url, "-o", "source.tar.gz")
	err = dependencycheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dependencycheck_cmd_direct := exec.Command("./binary")
	err = dependencycheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
