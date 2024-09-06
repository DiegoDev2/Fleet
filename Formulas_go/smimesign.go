package main

// Smimesign - S/MIME signing utility for use with Git
// Homepage: https://github.com/github/smimesign

import (
	"fmt"
	
	"os/exec"
)

func installSmimesign() {
	// Método 1: Descargar y extraer .tar.gz
	smimesign_tar_url := "https://github.com/github/smimesign/archive/refs/tags/v0.2.0.tar.gz"
	smimesign_cmd_tar := exec.Command("curl", "-L", smimesign_tar_url, "-o", "package.tar.gz")
	err := smimesign_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smimesign_zip_url := "https://github.com/github/smimesign/archive/refs/tags/v0.2.0.zip"
	smimesign_cmd_zip := exec.Command("curl", "-L", smimesign_zip_url, "-o", "package.zip")
	err = smimesign_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smimesign_bin_url := "https://github.com/github/smimesign/archive/refs/tags/v0.2.0.bin"
	smimesign_cmd_bin := exec.Command("curl", "-L", smimesign_bin_url, "-o", "binary.bin")
	err = smimesign_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smimesign_src_url := "https://github.com/github/smimesign/archive/refs/tags/v0.2.0.src.tar.gz"
	smimesign_cmd_src := exec.Command("curl", "-L", smimesign_src_url, "-o", "source.tar.gz")
	err = smimesign_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smimesign_cmd_direct := exec.Command("./binary")
	err = smimesign_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
