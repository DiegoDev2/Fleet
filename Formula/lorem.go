package main

// Lorem - Python generator for the console
// Homepage: https://github.com/per9000/lorem

import (
	"fmt"
	
	"os/exec"
)

func installLorem() {
	// Método 1: Descargar y extraer .tar.gz
	lorem_tar_url := "https://github.com/per9000/lorem/archive/refs/tags/v0.8.0.tar.gz"
	lorem_cmd_tar := exec.Command("curl", "-L", lorem_tar_url, "-o", "package.tar.gz")
	err := lorem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lorem_zip_url := "https://github.com/per9000/lorem/archive/refs/tags/v0.8.0.zip"
	lorem_cmd_zip := exec.Command("curl", "-L", lorem_zip_url, "-o", "package.zip")
	err = lorem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lorem_bin_url := "https://github.com/per9000/lorem/archive/refs/tags/v0.8.0.bin"
	lorem_cmd_bin := exec.Command("curl", "-L", lorem_bin_url, "-o", "binary.bin")
	err = lorem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lorem_src_url := "https://github.com/per9000/lorem/archive/refs/tags/v0.8.0.src.tar.gz"
	lorem_cmd_src := exec.Command("curl", "-L", lorem_src_url, "-o", "source.tar.gz")
	err = lorem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lorem_cmd_direct := exec.Command("./binary")
	err = lorem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
