package main

// Trufflehog - Find and verify credentials
// Homepage: https://trufflesecurity.com/

import (
	"fmt"
	
	"os/exec"
)

func installTrufflehog() {
	// Método 1: Descargar y extraer .tar.gz
	trufflehog_tar_url := "https://github.com/trufflesecurity/trufflehog/archive/refs/tags/v3.81.10.tar.gz"
	trufflehog_cmd_tar := exec.Command("curl", "-L", trufflehog_tar_url, "-o", "package.tar.gz")
	err := trufflehog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trufflehog_zip_url := "https://github.com/trufflesecurity/trufflehog/archive/refs/tags/v3.81.10.zip"
	trufflehog_cmd_zip := exec.Command("curl", "-L", trufflehog_zip_url, "-o", "package.zip")
	err = trufflehog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trufflehog_bin_url := "https://github.com/trufflesecurity/trufflehog/archive/refs/tags/v3.81.10.bin"
	trufflehog_cmd_bin := exec.Command("curl", "-L", trufflehog_bin_url, "-o", "binary.bin")
	err = trufflehog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trufflehog_src_url := "https://github.com/trufflesecurity/trufflehog/archive/refs/tags/v3.81.10.src.tar.gz"
	trufflehog_cmd_src := exec.Command("curl", "-L", trufflehog_src_url, "-o", "source.tar.gz")
	err = trufflehog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trufflehog_cmd_direct := exec.Command("./binary")
	err = trufflehog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
