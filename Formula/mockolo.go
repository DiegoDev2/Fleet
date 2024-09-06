package main

// Mockolo - Efficient Mock Generator for Swift
// Homepage: https://github.com/uber/mockolo

import (
	"fmt"
	
	"os/exec"
)

func installMockolo() {
	// Método 1: Descargar y extraer .tar.gz
	mockolo_tar_url := "https://github.com/uber/mockolo/archive/refs/tags/2.1.1.tar.gz"
	mockolo_cmd_tar := exec.Command("curl", "-L", mockolo_tar_url, "-o", "package.tar.gz")
	err := mockolo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mockolo_zip_url := "https://github.com/uber/mockolo/archive/refs/tags/2.1.1.zip"
	mockolo_cmd_zip := exec.Command("curl", "-L", mockolo_zip_url, "-o", "package.zip")
	err = mockolo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mockolo_bin_url := "https://github.com/uber/mockolo/archive/refs/tags/2.1.1.bin"
	mockolo_cmd_bin := exec.Command("curl", "-L", mockolo_bin_url, "-o", "binary.bin")
	err = mockolo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mockolo_src_url := "https://github.com/uber/mockolo/archive/refs/tags/2.1.1.src.tar.gz"
	mockolo_cmd_src := exec.Command("curl", "-L", mockolo_src_url, "-o", "source.tar.gz")
	err = mockolo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mockolo_cmd_direct := exec.Command("./binary")
	err = mockolo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
