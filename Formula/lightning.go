package main

// Lightning - Generates assembly language code at run-time
// Homepage: https://www.gnu.org/software/lightning/

import (
	"fmt"
	
	"os/exec"
)

func installLightning() {
	// Método 1: Descargar y extraer .tar.gz
	lightning_tar_url := "https://ftp.gnu.org/gnu/lightning/lightning-2.2.3.tar.gz"
	lightning_cmd_tar := exec.Command("curl", "-L", lightning_tar_url, "-o", "package.tar.gz")
	err := lightning_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lightning_zip_url := "https://ftp.gnu.org/gnu/lightning/lightning-2.2.3.zip"
	lightning_cmd_zip := exec.Command("curl", "-L", lightning_zip_url, "-o", "package.zip")
	err = lightning_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lightning_bin_url := "https://ftp.gnu.org/gnu/lightning/lightning-2.2.3.bin"
	lightning_cmd_bin := exec.Command("curl", "-L", lightning_bin_url, "-o", "binary.bin")
	err = lightning_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lightning_src_url := "https://ftp.gnu.org/gnu/lightning/lightning-2.2.3.src.tar.gz"
	lightning_cmd_src := exec.Command("curl", "-L", lightning_src_url, "-o", "source.tar.gz")
	err = lightning_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lightning_cmd_direct := exec.Command("./binary")
	err = lightning_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: binutils")
	exec.Command("latte", "install", "binutils").Run()
}
