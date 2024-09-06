package main

// Ansilove - ANSI/ASCII art to PNG converter
// Homepage: https://www.ansilove.org

import (
	"fmt"
	
	"os/exec"
)

func installAnsilove() {
	// Método 1: Descargar y extraer .tar.gz
	ansilove_tar_url := "https://github.com/ansilove/ansilove/releases/download/4.2.0/ansilove-4.2.0.tar.gz"
	ansilove_cmd_tar := exec.Command("curl", "-L", ansilove_tar_url, "-o", "package.tar.gz")
	err := ansilove_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansilove_zip_url := "https://github.com/ansilove/ansilove/releases/download/4.2.0/ansilove-4.2.0.zip"
	ansilove_cmd_zip := exec.Command("curl", "-L", ansilove_zip_url, "-o", "package.zip")
	err = ansilove_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansilove_bin_url := "https://github.com/ansilove/ansilove/releases/download/4.2.0/ansilove-4.2.0.bin"
	ansilove_cmd_bin := exec.Command("curl", "-L", ansilove_bin_url, "-o", "binary.bin")
	err = ansilove_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansilove_src_url := "https://github.com/ansilove/ansilove/releases/download/4.2.0/ansilove-4.2.0.src.tar.gz"
	ansilove_cmd_src := exec.Command("curl", "-L", ansilove_src_url, "-o", "source.tar.gz")
	err = ansilove_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansilove_cmd_direct := exec.Command("./binary")
	err = ansilove_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libansilove")
exec.Command("latte", "install", "libansilove")
}
