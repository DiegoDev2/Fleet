package main

// NoMoreSecrets - Recreates the SETEC ASTRONOMY effect from 'Sneakers'
// Homepage: https://github.com/bartobri/no-more-secrets

import (
	"fmt"
	
	"os/exec"
)

func installNoMoreSecrets() {
	// Método 1: Descargar y extraer .tar.gz
	nomoresecrets_tar_url := "https://github.com/bartobri/no-more-secrets/archive/refs/tags/v1.0.1.tar.gz"
	nomoresecrets_cmd_tar := exec.Command("curl", "-L", nomoresecrets_tar_url, "-o", "package.tar.gz")
	err := nomoresecrets_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nomoresecrets_zip_url := "https://github.com/bartobri/no-more-secrets/archive/refs/tags/v1.0.1.zip"
	nomoresecrets_cmd_zip := exec.Command("curl", "-L", nomoresecrets_zip_url, "-o", "package.zip")
	err = nomoresecrets_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nomoresecrets_bin_url := "https://github.com/bartobri/no-more-secrets/archive/refs/tags/v1.0.1.bin"
	nomoresecrets_cmd_bin := exec.Command("curl", "-L", nomoresecrets_bin_url, "-o", "binary.bin")
	err = nomoresecrets_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nomoresecrets_src_url := "https://github.com/bartobri/no-more-secrets/archive/refs/tags/v1.0.1.src.tar.gz"
	nomoresecrets_cmd_src := exec.Command("curl", "-L", nomoresecrets_src_url, "-o", "source.tar.gz")
	err = nomoresecrets_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nomoresecrets_cmd_direct := exec.Command("./binary")
	err = nomoresecrets_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
