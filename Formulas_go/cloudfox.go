package main

// Cloudfox - Automating situational awareness for cloud penetration tests
// Homepage: https://github.com/BishopFox/cloudfox

import (
	"fmt"
	
	"os/exec"
)

func installCloudfox() {
	// Método 1: Descargar y extraer .tar.gz
	cloudfox_tar_url := "https://github.com/BishopFox/cloudfox/archive/refs/tags/v1.14.2.tar.gz"
	cloudfox_cmd_tar := exec.Command("curl", "-L", cloudfox_tar_url, "-o", "package.tar.gz")
	err := cloudfox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudfox_zip_url := "https://github.com/BishopFox/cloudfox/archive/refs/tags/v1.14.2.zip"
	cloudfox_cmd_zip := exec.Command("curl", "-L", cloudfox_zip_url, "-o", "package.zip")
	err = cloudfox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudfox_bin_url := "https://github.com/BishopFox/cloudfox/archive/refs/tags/v1.14.2.bin"
	cloudfox_cmd_bin := exec.Command("curl", "-L", cloudfox_bin_url, "-o", "binary.bin")
	err = cloudfox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudfox_src_url := "https://github.com/BishopFox/cloudfox/archive/refs/tags/v1.14.2.src.tar.gz"
	cloudfox_cmd_src := exec.Command("curl", "-L", cloudfox_src_url, "-o", "source.tar.gz")
	err = cloudfox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudfox_cmd_direct := exec.Command("./binary")
	err = cloudfox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
