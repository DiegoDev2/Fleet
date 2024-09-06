package main

// Nave - Virtual environments for Node.js
// Homepage: https://github.com/isaacs/nave

import (
	"fmt"
	
	"os/exec"
)

func installNave() {
	// Método 1: Descargar y extraer .tar.gz
	nave_tar_url := "https://github.com/isaacs/nave/archive/refs/tags/v3.5.2.tar.gz"
	nave_cmd_tar := exec.Command("curl", "-L", nave_tar_url, "-o", "package.tar.gz")
	err := nave_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nave_zip_url := "https://github.com/isaacs/nave/archive/refs/tags/v3.5.2.zip"
	nave_cmd_zip := exec.Command("curl", "-L", nave_zip_url, "-o", "package.zip")
	err = nave_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nave_bin_url := "https://github.com/isaacs/nave/archive/refs/tags/v3.5.2.bin"
	nave_cmd_bin := exec.Command("curl", "-L", nave_bin_url, "-o", "binary.bin")
	err = nave_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nave_src_url := "https://github.com/isaacs/nave/archive/refs/tags/v3.5.2.src.tar.gz"
	nave_cmd_src := exec.Command("curl", "-L", nave_src_url, "-o", "source.tar.gz")
	err = nave_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nave_cmd_direct := exec.Command("./binary")
	err = nave_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
