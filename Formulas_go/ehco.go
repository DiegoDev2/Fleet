package main

// Ehco - Network relay tool and a typo :)
// Homepage: https://github.com/Ehco1996/ehco

import (
	"fmt"
	
	"os/exec"
)

func installEhco() {
	// Método 1: Descargar y extraer .tar.gz
	ehco_tar_url := "https://github.com/Ehco1996/ehco/archive/refs/tags/v1.1.4.tar.gz"
	ehco_cmd_tar := exec.Command("curl", "-L", ehco_tar_url, "-o", "package.tar.gz")
	err := ehco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ehco_zip_url := "https://github.com/Ehco1996/ehco/archive/refs/tags/v1.1.4.zip"
	ehco_cmd_zip := exec.Command("curl", "-L", ehco_zip_url, "-o", "package.zip")
	err = ehco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ehco_bin_url := "https://github.com/Ehco1996/ehco/archive/refs/tags/v1.1.4.bin"
	ehco_cmd_bin := exec.Command("curl", "-L", ehco_bin_url, "-o", "binary.bin")
	err = ehco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ehco_src_url := "https://github.com/Ehco1996/ehco/archive/refs/tags/v1.1.4.src.tar.gz"
	ehco_cmd_src := exec.Command("curl", "-L", ehco_src_url, "-o", "source.tar.gz")
	err = ehco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ehco_cmd_direct := exec.Command("./binary")
	err = ehco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
