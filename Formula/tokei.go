package main

// Tokei - Program that allows you to count code, quickly
// Homepage: https://github.com/XAMPPRocky/tokei

import (
	"fmt"
	
	"os/exec"
)

func installTokei() {
	// Método 1: Descargar y extraer .tar.gz
	tokei_tar_url := "https://github.com/XAMPPRocky/tokei/archive/refs/tags/v12.1.2.tar.gz"
	tokei_cmd_tar := exec.Command("curl", "-L", tokei_tar_url, "-o", "package.tar.gz")
	err := tokei_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tokei_zip_url := "https://github.com/XAMPPRocky/tokei/archive/refs/tags/v12.1.2.zip"
	tokei_cmd_zip := exec.Command("curl", "-L", tokei_zip_url, "-o", "package.zip")
	err = tokei_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tokei_bin_url := "https://github.com/XAMPPRocky/tokei/archive/refs/tags/v12.1.2.bin"
	tokei_cmd_bin := exec.Command("curl", "-L", tokei_bin_url, "-o", "binary.bin")
	err = tokei_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tokei_src_url := "https://github.com/XAMPPRocky/tokei/archive/refs/tags/v12.1.2.src.tar.gz"
	tokei_cmd_src := exec.Command("curl", "-L", tokei_src_url, "-o", "source.tar.gz")
	err = tokei_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tokei_cmd_direct := exec.Command("./binary")
	err = tokei_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
