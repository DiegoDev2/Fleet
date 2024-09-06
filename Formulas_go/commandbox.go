package main

// Commandbox - CFML embedded server, package manager, and app scaffolding tools
// Homepage: https://www.ortussolutions.com/products/commandbox

import (
	"fmt"
	
	"os/exec"
)

func installCommandbox() {
	// Método 1: Descargar y extraer .tar.gz
	commandbox_tar_url := "https://downloads.ortussolutions.com/ortussolutions/commandbox/6.0.0/commandbox-bin-6.0.0.zip"
	commandbox_cmd_tar := exec.Command("curl", "-L", commandbox_tar_url, "-o", "package.tar.gz")
	err := commandbox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	commandbox_zip_url := "https://downloads.ortussolutions.com/ortussolutions/commandbox/6.0.0/commandbox-bin-6.0.0.zip"
	commandbox_cmd_zip := exec.Command("curl", "-L", commandbox_zip_url, "-o", "package.zip")
	err = commandbox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	commandbox_bin_url := "https://downloads.ortussolutions.com/ortussolutions/commandbox/6.0.0/commandbox-bin-6.0.0.zip"
	commandbox_cmd_bin := exec.Command("curl", "-L", commandbox_bin_url, "-o", "binary.bin")
	err = commandbox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	commandbox_src_url := "https://downloads.ortussolutions.com/ortussolutions/commandbox/6.0.0/commandbox-bin-6.0.0.zip"
	commandbox_cmd_src := exec.Command("curl", "-L", commandbox_src_url, "-o", "source.tar.gz")
	err = commandbox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	commandbox_cmd_direct := exec.Command("./binary")
	err = commandbox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
