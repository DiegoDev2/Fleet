package main

// Imagejs - Tool to hide JavaScript inside valid image files
// Homepage: https://github.com/jklmnn/imagejs

import (
	"fmt"
	
	"os/exec"
)

func installImagejs() {
	// Método 1: Descargar y extraer .tar.gz
	imagejs_tar_url := "https://github.com/jklmnn/imagejs/archive/refs/tags/0.7.2.tar.gz"
	imagejs_cmd_tar := exec.Command("curl", "-L", imagejs_tar_url, "-o", "package.tar.gz")
	err := imagejs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imagejs_zip_url := "https://github.com/jklmnn/imagejs/archive/refs/tags/0.7.2.zip"
	imagejs_cmd_zip := exec.Command("curl", "-L", imagejs_zip_url, "-o", "package.zip")
	err = imagejs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imagejs_bin_url := "https://github.com/jklmnn/imagejs/archive/refs/tags/0.7.2.bin"
	imagejs_cmd_bin := exec.Command("curl", "-L", imagejs_bin_url, "-o", "binary.bin")
	err = imagejs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imagejs_src_url := "https://github.com/jklmnn/imagejs/archive/refs/tags/0.7.2.src.tar.gz"
	imagejs_cmd_src := exec.Command("curl", "-L", imagejs_src_url, "-o", "source.tar.gz")
	err = imagejs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imagejs_cmd_direct := exec.Command("./binary")
	err = imagejs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
