package main

// Restic - Fast, efficient and secure backup program
// Homepage: https://restic.net/

import (
	"fmt"
	
	"os/exec"
)

func installRestic() {
	// Método 1: Descargar y extraer .tar.gz
	restic_tar_url := "https://github.com/restic/restic/archive/refs/tags/v0.17.0.tar.gz"
	restic_cmd_tar := exec.Command("curl", "-L", restic_tar_url, "-o", "package.tar.gz")
	err := restic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	restic_zip_url := "https://github.com/restic/restic/archive/refs/tags/v0.17.0.zip"
	restic_cmd_zip := exec.Command("curl", "-L", restic_zip_url, "-o", "package.zip")
	err = restic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	restic_bin_url := "https://github.com/restic/restic/archive/refs/tags/v0.17.0.bin"
	restic_cmd_bin := exec.Command("curl", "-L", restic_bin_url, "-o", "binary.bin")
	err = restic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	restic_src_url := "https://github.com/restic/restic/archive/refs/tags/v0.17.0.src.tar.gz"
	restic_cmd_src := exec.Command("curl", "-L", restic_src_url, "-o", "source.tar.gz")
	err = restic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	restic_cmd_direct := exec.Command("./binary")
	err = restic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
