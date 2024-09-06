package main

// Pius - PGP individual UID signer
// Homepage: https://www.phildev.net/pius/

import (
	"fmt"
	
	"os/exec"
)

func installPius() {
	// Método 1: Descargar y extraer .tar.gz
	pius_tar_url := "https://github.com/jaymzh/pius/archive/refs/tags/v3.0.0.tar.gz"
	pius_cmd_tar := exec.Command("curl", "-L", pius_tar_url, "-o", "package.tar.gz")
	err := pius_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pius_zip_url := "https://github.com/jaymzh/pius/archive/refs/tags/v3.0.0.zip"
	pius_cmd_zip := exec.Command("curl", "-L", pius_zip_url, "-o", "package.zip")
	err = pius_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pius_bin_url := "https://github.com/jaymzh/pius/archive/refs/tags/v3.0.0.bin"
	pius_cmd_bin := exec.Command("curl", "-L", pius_bin_url, "-o", "binary.bin")
	err = pius_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pius_src_url := "https://github.com/jaymzh/pius/archive/refs/tags/v3.0.0.src.tar.gz"
	pius_cmd_src := exec.Command("curl", "-L", pius_src_url, "-o", "source.tar.gz")
	err = pius_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pius_cmd_direct := exec.Command("./binary")
	err = pius_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
