package main

// Um - Command-line utility for creating and maintaining personal man pages
// Homepage: https://github.com/sinclairtarget/um

import (
	"fmt"
	
	"os/exec"
)

func installUm() {
	// Método 1: Descargar y extraer .tar.gz
	um_tar_url := "https://github.com/sinclairtarget/um/archive/refs/tags/4.2.0.tar.gz"
	um_cmd_tar := exec.Command("curl", "-L", um_tar_url, "-o", "package.tar.gz")
	err := um_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	um_zip_url := "https://github.com/sinclairtarget/um/archive/refs/tags/4.2.0.zip"
	um_cmd_zip := exec.Command("curl", "-L", um_zip_url, "-o", "package.zip")
	err = um_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	um_bin_url := "https://github.com/sinclairtarget/um/archive/refs/tags/4.2.0.bin"
	um_cmd_bin := exec.Command("curl", "-L", um_bin_url, "-o", "binary.bin")
	err = um_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	um_src_url := "https://github.com/sinclairtarget/um/archive/refs/tags/4.2.0.src.tar.gz"
	um_cmd_src := exec.Command("curl", "-L", um_src_url, "-o", "source.tar.gz")
	err = um_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	um_cmd_direct := exec.Command("./binary")
	err = um_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
