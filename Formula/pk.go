package main

// Pk - Field extractor command-line utility
// Homepage: https://github.com/johnmorrow/pk

import (
	"fmt"
	
	"os/exec"
)

func installPk() {
	// Método 1: Descargar y extraer .tar.gz
	pk_tar_url := "https://github.com/johnmorrow/pk/releases/download/v1.0.2/pk-1.0.2.tar.gz"
	pk_cmd_tar := exec.Command("curl", "-L", pk_tar_url, "-o", "package.tar.gz")
	err := pk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pk_zip_url := "https://github.com/johnmorrow/pk/releases/download/v1.0.2/pk-1.0.2.zip"
	pk_cmd_zip := exec.Command("curl", "-L", pk_zip_url, "-o", "package.zip")
	err = pk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pk_bin_url := "https://github.com/johnmorrow/pk/releases/download/v1.0.2/pk-1.0.2.bin"
	pk_cmd_bin := exec.Command("curl", "-L", pk_bin_url, "-o", "binary.bin")
	err = pk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pk_src_url := "https://github.com/johnmorrow/pk/releases/download/v1.0.2/pk-1.0.2.src.tar.gz"
	pk_cmd_src := exec.Command("curl", "-L", pk_src_url, "-o", "source.tar.gz")
	err = pk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pk_cmd_direct := exec.Command("./binary")
	err = pk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: argp-standalone")
	exec.Command("latte", "install", "argp-standalone").Run()
}
