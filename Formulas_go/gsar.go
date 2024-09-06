package main

// Gsar - General Search And Replace on files
// Homepage: https://tjaberg.com/

import (
	"fmt"
	
	"os/exec"
)

func installGsar() {
	// Método 1: Descargar y extraer .tar.gz
	gsar_tar_url := "https://tjaberg.com/gsar151.zip"
	gsar_cmd_tar := exec.Command("curl", "-L", gsar_tar_url, "-o", "package.tar.gz")
	err := gsar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gsar_zip_url := "https://tjaberg.com/gsar151.zip"
	gsar_cmd_zip := exec.Command("curl", "-L", gsar_zip_url, "-o", "package.zip")
	err = gsar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gsar_bin_url := "https://tjaberg.com/gsar151.zip"
	gsar_cmd_bin := exec.Command("curl", "-L", gsar_bin_url, "-o", "binary.bin")
	err = gsar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gsar_src_url := "https://tjaberg.com/gsar151.zip"
	gsar_cmd_src := exec.Command("curl", "-L", gsar_src_url, "-o", "source.tar.gz")
	err = gsar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gsar_cmd_direct := exec.Command("./binary")
	err = gsar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
