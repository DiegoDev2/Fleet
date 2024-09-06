package main

// Fpart - Sorts file trees and packs them into bags
// Homepage: https://github.com/martymac/fpart/

import (
	"fmt"
	
	"os/exec"
)

func installFpart() {
	// Método 1: Descargar y extraer .tar.gz
	fpart_tar_url := "https://github.com/martymac/fpart/archive/refs/tags/fpart-1.6.0.tar.gz"
	fpart_cmd_tar := exec.Command("curl", "-L", fpart_tar_url, "-o", "package.tar.gz")
	err := fpart_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fpart_zip_url := "https://github.com/martymac/fpart/archive/refs/tags/fpart-1.6.0.zip"
	fpart_cmd_zip := exec.Command("curl", "-L", fpart_zip_url, "-o", "package.zip")
	err = fpart_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fpart_bin_url := "https://github.com/martymac/fpart/archive/refs/tags/fpart-1.6.0.bin"
	fpart_cmd_bin := exec.Command("curl", "-L", fpart_bin_url, "-o", "binary.bin")
	err = fpart_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fpart_src_url := "https://github.com/martymac/fpart/archive/refs/tags/fpart-1.6.0.src.tar.gz"
	fpart_cmd_src := exec.Command("curl", "-L", fpart_src_url, "-o", "source.tar.gz")
	err = fpart_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fpart_cmd_direct := exec.Command("./binary")
	err = fpart_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
