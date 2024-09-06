package main

// Pdfrip - Multi-threaded PDF password cracking utility
// Homepage: https://github.com/mufeedvh/pdfrip

import (
	"fmt"
	
	"os/exec"
)

func installPdfrip() {
	// Método 1: Descargar y extraer .tar.gz
	pdfrip_tar_url := "https://github.com/mufeedvh/pdfrip/archive/refs/tags/v2.0.1.tar.gz"
	pdfrip_cmd_tar := exec.Command("curl", "-L", pdfrip_tar_url, "-o", "package.tar.gz")
	err := pdfrip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdfrip_zip_url := "https://github.com/mufeedvh/pdfrip/archive/refs/tags/v2.0.1.zip"
	pdfrip_cmd_zip := exec.Command("curl", "-L", pdfrip_zip_url, "-o", "package.zip")
	err = pdfrip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdfrip_bin_url := "https://github.com/mufeedvh/pdfrip/archive/refs/tags/v2.0.1.bin"
	pdfrip_cmd_bin := exec.Command("curl", "-L", pdfrip_bin_url, "-o", "binary.bin")
	err = pdfrip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdfrip_src_url := "https://github.com/mufeedvh/pdfrip/archive/refs/tags/v2.0.1.src.tar.gz"
	pdfrip_cmd_src := exec.Command("curl", "-L", pdfrip_src_url, "-o", "source.tar.gz")
	err = pdfrip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdfrip_cmd_direct := exec.Command("./binary")
	err = pdfrip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
