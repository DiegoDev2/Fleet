package main

// Rswift - Get strong typed, autocompleted resources like images, fonts and segues
// Homepage: https://github.com/mac-cain13/R.swift

import (
	"fmt"
	
	"os/exec"
)

func installRswift() {
	// Método 1: Descargar y extraer .tar.gz
	rswift_tar_url := "https://github.com/mac-cain13/R.swift/releases/download/7.5.0/rswift-7.5.0-source.tar.gz"
	rswift_cmd_tar := exec.Command("curl", "-L", rswift_tar_url, "-o", "package.tar.gz")
	err := rswift_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rswift_zip_url := "https://github.com/mac-cain13/R.swift/releases/download/7.5.0/rswift-7.5.0-source.zip"
	rswift_cmd_zip := exec.Command("curl", "-L", rswift_zip_url, "-o", "package.zip")
	err = rswift_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rswift_bin_url := "https://github.com/mac-cain13/R.swift/releases/download/7.5.0/rswift-7.5.0-source.bin"
	rswift_cmd_bin := exec.Command("curl", "-L", rswift_bin_url, "-o", "binary.bin")
	err = rswift_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rswift_src_url := "https://github.com/mac-cain13/R.swift/releases/download/7.5.0/rswift-7.5.0-source.src.tar.gz"
	rswift_cmd_src := exec.Command("curl", "-L", rswift_src_url, "-o", "source.tar.gz")
	err = rswift_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rswift_cmd_direct := exec.Command("./binary")
	err = rswift_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
