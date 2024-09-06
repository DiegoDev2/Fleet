package main

// CreateDmg - Shell script to build fancy DMGs
// Homepage: https://github.com/create-dmg/create-dmg

import (
	"fmt"
	
	"os/exec"
)

func installCreateDmg() {
	// Método 1: Descargar y extraer .tar.gz
	createdmg_tar_url := "https://github.com/create-dmg/create-dmg/archive/refs/tags/v1.2.2.tar.gz"
	createdmg_cmd_tar := exec.Command("curl", "-L", createdmg_tar_url, "-o", "package.tar.gz")
	err := createdmg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	createdmg_zip_url := "https://github.com/create-dmg/create-dmg/archive/refs/tags/v1.2.2.zip"
	createdmg_cmd_zip := exec.Command("curl", "-L", createdmg_zip_url, "-o", "package.zip")
	err = createdmg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	createdmg_bin_url := "https://github.com/create-dmg/create-dmg/archive/refs/tags/v1.2.2.bin"
	createdmg_cmd_bin := exec.Command("curl", "-L", createdmg_bin_url, "-o", "binary.bin")
	err = createdmg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	createdmg_src_url := "https://github.com/create-dmg/create-dmg/archive/refs/tags/v1.2.2.src.tar.gz"
	createdmg_cmd_src := exec.Command("curl", "-L", createdmg_src_url, "-o", "source.tar.gz")
	err = createdmg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	createdmg_cmd_direct := exec.Command("./binary")
	err = createdmg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
