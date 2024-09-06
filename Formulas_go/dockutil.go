package main

// Dockutil - Tool for managing dock items
// Homepage: https://github.com/kcrawford/dockutil

import (
	"fmt"
	
	"os/exec"
)

func installDockutil() {
	// Método 1: Descargar y extraer .tar.gz
	dockutil_tar_url := "https://github.com/kcrawford/dockutil/archive/refs/tags/3.1.3.tar.gz"
	dockutil_cmd_tar := exec.Command("curl", "-L", dockutil_tar_url, "-o", "package.tar.gz")
	err := dockutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockutil_zip_url := "https://github.com/kcrawford/dockutil/archive/refs/tags/3.1.3.zip"
	dockutil_cmd_zip := exec.Command("curl", "-L", dockutil_zip_url, "-o", "package.zip")
	err = dockutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockutil_bin_url := "https://github.com/kcrawford/dockutil/archive/refs/tags/3.1.3.bin"
	dockutil_cmd_bin := exec.Command("curl", "-L", dockutil_bin_url, "-o", "binary.bin")
	err = dockutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockutil_src_url := "https://github.com/kcrawford/dockutil/archive/refs/tags/3.1.3.src.tar.gz"
	dockutil_cmd_src := exec.Command("curl", "-L", dockutil_src_url, "-o", "source.tar.gz")
	err = dockutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockutil_cmd_direct := exec.Command("./binary")
	err = dockutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
