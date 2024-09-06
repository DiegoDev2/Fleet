package main

// Loc - Count lines of code quickly
// Homepage: https://github.com/cgag/loc

import (
	"fmt"
	
	"os/exec"
)

func installLoc() {
	// Método 1: Descargar y extraer .tar.gz
	loc_tar_url := "https://github.com/cgag/loc/archive/refs/tags/v0.4.1.tar.gz"
	loc_cmd_tar := exec.Command("curl", "-L", loc_tar_url, "-o", "package.tar.gz")
	err := loc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	loc_zip_url := "https://github.com/cgag/loc/archive/refs/tags/v0.4.1.zip"
	loc_cmd_zip := exec.Command("curl", "-L", loc_zip_url, "-o", "package.zip")
	err = loc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	loc_bin_url := "https://github.com/cgag/loc/archive/refs/tags/v0.4.1.bin"
	loc_cmd_bin := exec.Command("curl", "-L", loc_bin_url, "-o", "binary.bin")
	err = loc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	loc_src_url := "https://github.com/cgag/loc/archive/refs/tags/v0.4.1.src.tar.gz"
	loc_cmd_src := exec.Command("curl", "-L", loc_src_url, "-o", "source.tar.gz")
	err = loc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	loc_cmd_direct := exec.Command("./binary")
	err = loc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
