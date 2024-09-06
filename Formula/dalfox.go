package main

// Dalfox - XSS scanner and utility focused on automation
// Homepage: https://dalfox.hahwul.com

import (
	"fmt"
	
	"os/exec"
)

func installDalfox() {
	// Método 1: Descargar y extraer .tar.gz
	dalfox_tar_url := "https://github.com/hahwul/dalfox/archive/refs/tags/v2.9.3.tar.gz"
	dalfox_cmd_tar := exec.Command("curl", "-L", dalfox_tar_url, "-o", "package.tar.gz")
	err := dalfox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dalfox_zip_url := "https://github.com/hahwul/dalfox/archive/refs/tags/v2.9.3.zip"
	dalfox_cmd_zip := exec.Command("curl", "-L", dalfox_zip_url, "-o", "package.zip")
	err = dalfox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dalfox_bin_url := "https://github.com/hahwul/dalfox/archive/refs/tags/v2.9.3.bin"
	dalfox_cmd_bin := exec.Command("curl", "-L", dalfox_bin_url, "-o", "binary.bin")
	err = dalfox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dalfox_src_url := "https://github.com/hahwul/dalfox/archive/refs/tags/v2.9.3.src.tar.gz"
	dalfox_cmd_src := exec.Command("curl", "-L", dalfox_src_url, "-o", "source.tar.gz")
	err = dalfox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dalfox_cmd_direct := exec.Command("./binary")
	err = dalfox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
