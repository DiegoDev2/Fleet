package main

// PySpy - Sampling profiler for Python programs
// Homepage: https://github.com/benfred/py-spy

import (
	"fmt"
	
	"os/exec"
)

func installPySpy() {
	// Método 1: Descargar y extraer .tar.gz
	pyspy_tar_url := "https://github.com/benfred/py-spy/archive/refs/tags/v0.3.14.tar.gz"
	pyspy_cmd_tar := exec.Command("curl", "-L", pyspy_tar_url, "-o", "package.tar.gz")
	err := pyspy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyspy_zip_url := "https://github.com/benfred/py-spy/archive/refs/tags/v0.3.14.zip"
	pyspy_cmd_zip := exec.Command("curl", "-L", pyspy_zip_url, "-o", "package.zip")
	err = pyspy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyspy_bin_url := "https://github.com/benfred/py-spy/archive/refs/tags/v0.3.14.bin"
	pyspy_cmd_bin := exec.Command("curl", "-L", pyspy_bin_url, "-o", "binary.bin")
	err = pyspy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyspy_src_url := "https://github.com/benfred/py-spy/archive/refs/tags/v0.3.14.src.tar.gz"
	pyspy_cmd_src := exec.Command("curl", "-L", pyspy_src_url, "-o", "source.tar.gz")
	err = pyspy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyspy_cmd_direct := exec.Command("./binary")
	err = pyspy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm@15")
	exec.Command("latte", "install", "llvm@15").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: libunwind")
	exec.Command("latte", "install", "libunwind").Run()
}
