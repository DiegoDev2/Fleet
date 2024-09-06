package main

// Pciutils - PCI utilities
// Homepage: https://github.com/pciutils/pciutils

import (
	"fmt"
	
	"os/exec"
)

func installPciutils() {
	// Método 1: Descargar y extraer .tar.gz
	pciutils_tar_url := "https://github.com/pciutils/pciutils/archive/refs/tags/v3.13.0.tar.gz"
	pciutils_cmd_tar := exec.Command("curl", "-L", pciutils_tar_url, "-o", "package.tar.gz")
	err := pciutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pciutils_zip_url := "https://github.com/pciutils/pciutils/archive/refs/tags/v3.13.0.zip"
	pciutils_cmd_zip := exec.Command("curl", "-L", pciutils_zip_url, "-o", "package.zip")
	err = pciutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pciutils_bin_url := "https://github.com/pciutils/pciutils/archive/refs/tags/v3.13.0.bin"
	pciutils_cmd_bin := exec.Command("curl", "-L", pciutils_bin_url, "-o", "binary.bin")
	err = pciutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pciutils_src_url := "https://github.com/pciutils/pciutils/archive/refs/tags/v3.13.0.src.tar.gz"
	pciutils_cmd_src := exec.Command("curl", "-L", pciutils_src_url, "-o", "source.tar.gz")
	err = pciutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pciutils_cmd_direct := exec.Command("./binary")
	err = pciutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
}
