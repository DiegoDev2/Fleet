package main

// Convox - Command-line interface for the Convox PaaS
// Homepage: https://convox.com/

import (
	"fmt"
	
	"os/exec"
)

func installConvox() {
	// Método 1: Descargar y extraer .tar.gz
	convox_tar_url := "https://github.com/convox/convox/archive/refs/tags/3.18.10.tar.gz"
	convox_cmd_tar := exec.Command("curl", "-L", convox_tar_url, "-o", "package.tar.gz")
	err := convox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	convox_zip_url := "https://github.com/convox/convox/archive/refs/tags/3.18.10.zip"
	convox_cmd_zip := exec.Command("curl", "-L", convox_zip_url, "-o", "package.zip")
	err = convox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	convox_bin_url := "https://github.com/convox/convox/archive/refs/tags/3.18.10.bin"
	convox_cmd_bin := exec.Command("curl", "-L", convox_bin_url, "-o", "binary.bin")
	err = convox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	convox_src_url := "https://github.com/convox/convox/archive/refs/tags/3.18.10.src.tar.gz"
	convox_cmd_src := exec.Command("curl", "-L", convox_src_url, "-o", "source.tar.gz")
	err = convox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	convox_cmd_direct := exec.Command("./binary")
	err = convox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
