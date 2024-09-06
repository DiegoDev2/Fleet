package main

// Xtitle - Set window title and icon for your X terminal
// Homepage: https://kinzler.com/me/xtitle/

import (
	"fmt"
	
	"os/exec"
)

func installXtitle() {
	// Método 1: Descargar y extraer .tar.gz
	xtitle_tar_url := "https://kinzler.com/me/xtitle/xtitle-1.0.4.tgz"
	xtitle_cmd_tar := exec.Command("curl", "-L", xtitle_tar_url, "-o", "package.tar.gz")
	err := xtitle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xtitle_zip_url := "https://kinzler.com/me/xtitle/xtitle-1.0.4.tgz"
	xtitle_cmd_zip := exec.Command("curl", "-L", xtitle_zip_url, "-o", "package.zip")
	err = xtitle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xtitle_bin_url := "https://kinzler.com/me/xtitle/xtitle-1.0.4.tgz"
	xtitle_cmd_bin := exec.Command("curl", "-L", xtitle_bin_url, "-o", "binary.bin")
	err = xtitle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xtitle_src_url := "https://kinzler.com/me/xtitle/xtitle-1.0.4.tgz"
	xtitle_cmd_src := exec.Command("curl", "-L", xtitle_src_url, "-o", "source.tar.gz")
	err = xtitle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xtitle_cmd_direct := exec.Command("./binary")
	err = xtitle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
