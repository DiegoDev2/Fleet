package main

// Ripsecrets - Prevent committing secret keys into your source code
// Homepage: https://github.com/sirwart/ripsecrets

import (
	"fmt"
	
	"os/exec"
)

func installRipsecrets() {
	// Método 1: Descargar y extraer .tar.gz
	ripsecrets_tar_url := "https://github.com/sirwart/ripsecrets/archive/refs/tags/v0.1.8.tar.gz"
	ripsecrets_cmd_tar := exec.Command("curl", "-L", ripsecrets_tar_url, "-o", "package.tar.gz")
	err := ripsecrets_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ripsecrets_zip_url := "https://github.com/sirwart/ripsecrets/archive/refs/tags/v0.1.8.zip"
	ripsecrets_cmd_zip := exec.Command("curl", "-L", ripsecrets_zip_url, "-o", "package.zip")
	err = ripsecrets_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ripsecrets_bin_url := "https://github.com/sirwart/ripsecrets/archive/refs/tags/v0.1.8.bin"
	ripsecrets_cmd_bin := exec.Command("curl", "-L", ripsecrets_bin_url, "-o", "binary.bin")
	err = ripsecrets_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ripsecrets_src_url := "https://github.com/sirwart/ripsecrets/archive/refs/tags/v0.1.8.src.tar.gz"
	ripsecrets_cmd_src := exec.Command("curl", "-L", ripsecrets_src_url, "-o", "source.tar.gz")
	err = ripsecrets_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ripsecrets_cmd_direct := exec.Command("./binary")
	err = ripsecrets_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
