package main

// Pyoxidizer - Modern Python application packaging and distribution tool
// Homepage: https://github.com/indygreg/PyOxidizer

import (
	"fmt"
	
	"os/exec"
)

func installPyoxidizer() {
	// Método 1: Descargar y extraer .tar.gz
	pyoxidizer_tar_url := "https://github.com/indygreg/PyOxidizer/archive/refs/tags/pyoxidizer/0.24.0.tar.gz"
	pyoxidizer_cmd_tar := exec.Command("curl", "-L", pyoxidizer_tar_url, "-o", "package.tar.gz")
	err := pyoxidizer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyoxidizer_zip_url := "https://github.com/indygreg/PyOxidizer/archive/refs/tags/pyoxidizer/0.24.0.zip"
	pyoxidizer_cmd_zip := exec.Command("curl", "-L", pyoxidizer_zip_url, "-o", "package.zip")
	err = pyoxidizer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyoxidizer_bin_url := "https://github.com/indygreg/PyOxidizer/archive/refs/tags/pyoxidizer/0.24.0.bin"
	pyoxidizer_cmd_bin := exec.Command("curl", "-L", pyoxidizer_bin_url, "-o", "binary.bin")
	err = pyoxidizer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyoxidizer_src_url := "https://github.com/indygreg/PyOxidizer/archive/refs/tags/pyoxidizer/0.24.0.src.tar.gz"
	pyoxidizer_cmd_src := exec.Command("curl", "-L", pyoxidizer_src_url, "-o", "source.tar.gz")
	err = pyoxidizer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyoxidizer_cmd_direct := exec.Command("./binary")
	err = pyoxidizer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
