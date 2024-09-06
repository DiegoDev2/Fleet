package main

// Curseofwar - Fast-paced action strategy game
// Homepage: https://a-nikolaev.github.io/curseofwar/

import (
	"fmt"
	
	"os/exec"
)

func installCurseofwar() {
	// Método 1: Descargar y extraer .tar.gz
	curseofwar_tar_url := "https://github.com/a-nikolaev/curseofwar/archive/refs/tags/v1.3.0.tar.gz"
	curseofwar_cmd_tar := exec.Command("curl", "-L", curseofwar_tar_url, "-o", "package.tar.gz")
	err := curseofwar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	curseofwar_zip_url := "https://github.com/a-nikolaev/curseofwar/archive/refs/tags/v1.3.0.zip"
	curseofwar_cmd_zip := exec.Command("curl", "-L", curseofwar_zip_url, "-o", "package.zip")
	err = curseofwar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	curseofwar_bin_url := "https://github.com/a-nikolaev/curseofwar/archive/refs/tags/v1.3.0.bin"
	curseofwar_cmd_bin := exec.Command("curl", "-L", curseofwar_bin_url, "-o", "binary.bin")
	err = curseofwar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	curseofwar_src_url := "https://github.com/a-nikolaev/curseofwar/archive/refs/tags/v1.3.0.src.tar.gz"
	curseofwar_cmd_src := exec.Command("curl", "-L", curseofwar_src_url, "-o", "source.tar.gz")
	err = curseofwar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	curseofwar_cmd_direct := exec.Command("./binary")
	err = curseofwar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
