package main

// Dockward - Port forwarding tool for Docker containers
// Homepage: https://github.com/abiosoft/dockward

import (
	"fmt"
	
	"os/exec"
)

func installDockward() {
	// Método 1: Descargar y extraer .tar.gz
	dockward_tar_url := "https://github.com/abiosoft/dockward/archive/refs/tags/0.0.4.tar.gz"
	dockward_cmd_tar := exec.Command("curl", "-L", dockward_tar_url, "-o", "package.tar.gz")
	err := dockward_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockward_zip_url := "https://github.com/abiosoft/dockward/archive/refs/tags/0.0.4.zip"
	dockward_cmd_zip := exec.Command("curl", "-L", dockward_zip_url, "-o", "package.zip")
	err = dockward_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockward_bin_url := "https://github.com/abiosoft/dockward/archive/refs/tags/0.0.4.bin"
	dockward_cmd_bin := exec.Command("curl", "-L", dockward_bin_url, "-o", "binary.bin")
	err = dockward_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockward_src_url := "https://github.com/abiosoft/dockward/archive/refs/tags/0.0.4.src.tar.gz"
	dockward_cmd_src := exec.Command("curl", "-L", dockward_src_url, "-o", "source.tar.gz")
	err = dockward_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockward_cmd_direct := exec.Command("./binary")
	err = dockward_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
