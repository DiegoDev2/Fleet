package main

// Libansilove - Library for converting ANSI, ASCII, and other formats to PNG
// Homepage: https://www.ansilove.org

import (
	"fmt"
	
	"os/exec"
)

func installLibansilove() {
	// Método 1: Descargar y extraer .tar.gz
	libansilove_tar_url := "https://github.com/ansilove/libansilove/releases/download/1.4.1/libansilove-1.4.1.tar.gz"
	libansilove_cmd_tar := exec.Command("curl", "-L", libansilove_tar_url, "-o", "package.tar.gz")
	err := libansilove_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libansilove_zip_url := "https://github.com/ansilove/libansilove/releases/download/1.4.1/libansilove-1.4.1.zip"
	libansilove_cmd_zip := exec.Command("curl", "-L", libansilove_zip_url, "-o", "package.zip")
	err = libansilove_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libansilove_bin_url := "https://github.com/ansilove/libansilove/releases/download/1.4.1/libansilove-1.4.1.bin"
	libansilove_cmd_bin := exec.Command("curl", "-L", libansilove_bin_url, "-o", "binary.bin")
	err = libansilove_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libansilove_src_url := "https://github.com/ansilove/libansilove/releases/download/1.4.1/libansilove-1.4.1.src.tar.gz"
	libansilove_cmd_src := exec.Command("curl", "-L", libansilove_src_url, "-o", "source.tar.gz")
	err = libansilove_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libansilove_cmd_direct := exec.Command("./binary")
	err = libansilove_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
}
