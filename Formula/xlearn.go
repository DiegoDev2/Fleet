package main

// Xlearn - High performance, easy-to-use, and scalable machine learning package
// Homepage: https://xlearn-doc.readthedocs.io/en/latest/index.html

import (
	"fmt"
	
	"os/exec"
)

func installXlearn() {
	// Método 1: Descargar y extraer .tar.gz
	xlearn_tar_url := "https://github.com/aksnzhy/xlearn/archive/refs/tags/v0.4.4.tar.gz"
	xlearn_cmd_tar := exec.Command("curl", "-L", xlearn_tar_url, "-o", "package.tar.gz")
	err := xlearn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xlearn_zip_url := "https://github.com/aksnzhy/xlearn/archive/refs/tags/v0.4.4.zip"
	xlearn_cmd_zip := exec.Command("curl", "-L", xlearn_zip_url, "-o", "package.zip")
	err = xlearn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xlearn_bin_url := "https://github.com/aksnzhy/xlearn/archive/refs/tags/v0.4.4.bin"
	xlearn_cmd_bin := exec.Command("curl", "-L", xlearn_bin_url, "-o", "binary.bin")
	err = xlearn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xlearn_src_url := "https://github.com/aksnzhy/xlearn/archive/refs/tags/v0.4.4.src.tar.gz"
	xlearn_cmd_src := exec.Command("curl", "-L", xlearn_src_url, "-o", "source.tar.gz")
	err = xlearn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xlearn_cmd_direct := exec.Command("./binary")
	err = xlearn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
