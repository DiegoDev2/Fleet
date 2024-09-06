package main

// Maclaunch - Manage your macOS startup items
// Homepage: https://github.com/hazcod/maclaunch

import (
	"fmt"
	
	"os/exec"
)

func installMaclaunch() {
	// Método 1: Descargar y extraer .tar.gz
	maclaunch_tar_url := "https://github.com/hazcod/maclaunch/archive/refs/tags/2.5.tar.gz"
	maclaunch_cmd_tar := exec.Command("curl", "-L", maclaunch_tar_url, "-o", "package.tar.gz")
	err := maclaunch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	maclaunch_zip_url := "https://github.com/hazcod/maclaunch/archive/refs/tags/2.5.zip"
	maclaunch_cmd_zip := exec.Command("curl", "-L", maclaunch_zip_url, "-o", "package.zip")
	err = maclaunch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	maclaunch_bin_url := "https://github.com/hazcod/maclaunch/archive/refs/tags/2.5.bin"
	maclaunch_cmd_bin := exec.Command("curl", "-L", maclaunch_bin_url, "-o", "binary.bin")
	err = maclaunch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	maclaunch_src_url := "https://github.com/hazcod/maclaunch/archive/refs/tags/2.5.src.tar.gz"
	maclaunch_cmd_src := exec.Command("curl", "-L", maclaunch_src_url, "-o", "source.tar.gz")
	err = maclaunch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	maclaunch_cmd_direct := exec.Command("./binary")
	err = maclaunch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
