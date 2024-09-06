package main

// Dockly - Immersive terminal interface for managing docker containers and services
// Homepage: https://lirantal.github.io/dockly/

import (
	"fmt"
	
	"os/exec"
)

func installDockly() {
	// Método 1: Descargar y extraer .tar.gz
	dockly_tar_url := "https://registry.npmjs.org/dockly/-/dockly-3.24.3.tgz"
	dockly_cmd_tar := exec.Command("curl", "-L", dockly_tar_url, "-o", "package.tar.gz")
	err := dockly_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockly_zip_url := "https://registry.npmjs.org/dockly/-/dockly-3.24.3.tgz"
	dockly_cmd_zip := exec.Command("curl", "-L", dockly_zip_url, "-o", "package.zip")
	err = dockly_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockly_bin_url := "https://registry.npmjs.org/dockly/-/dockly-3.24.3.tgz"
	dockly_cmd_bin := exec.Command("curl", "-L", dockly_bin_url, "-o", "binary.bin")
	err = dockly_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockly_src_url := "https://registry.npmjs.org/dockly/-/dockly-3.24.3.tgz"
	dockly_cmd_src := exec.Command("curl", "-L", dockly_src_url, "-o", "source.tar.gz")
	err = dockly_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockly_cmd_direct := exec.Command("./binary")
	err = dockly_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: xsel")
	exec.Command("latte", "install", "xsel").Run()
}
