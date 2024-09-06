package main

// Sz81 - ZX80/81 emulator
// Homepage: https://sz81.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSz81() {
	// Método 1: Descargar y extraer .tar.gz
	sz81_tar_url := "https://downloads.sourceforge.net/project/sz81/sz81/2.1.7/sz81-2.1.7-source.tar.gz"
	sz81_cmd_tar := exec.Command("curl", "-L", sz81_tar_url, "-o", "package.tar.gz")
	err := sz81_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sz81_zip_url := "https://downloads.sourceforge.net/project/sz81/sz81/2.1.7/sz81-2.1.7-source.zip"
	sz81_cmd_zip := exec.Command("curl", "-L", sz81_zip_url, "-o", "package.zip")
	err = sz81_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sz81_bin_url := "https://downloads.sourceforge.net/project/sz81/sz81/2.1.7/sz81-2.1.7-source.bin"
	sz81_cmd_bin := exec.Command("curl", "-L", sz81_bin_url, "-o", "binary.bin")
	err = sz81_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sz81_src_url := "https://downloads.sourceforge.net/project/sz81/sz81/2.1.7/sz81-2.1.7-source.src.tar.gz"
	sz81_cmd_src := exec.Command("curl", "-L", sz81_src_url, "-o", "source.tar.gz")
	err = sz81_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sz81_cmd_direct := exec.Command("./binary")
	err = sz81_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
}
