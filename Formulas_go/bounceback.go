package main

// Bounceback - Stealth redirector for red team operation security
// Homepage: https://github.com/D00Movenok/BounceBack

import (
	"fmt"
	
	"os/exec"
)

func installBounceback() {
	// Método 1: Descargar y extraer .tar.gz
	bounceback_tar_url := "https://github.com/D00Movenok/BounceBack/archive/refs/tags/v1.5.1.tar.gz"
	bounceback_cmd_tar := exec.Command("curl", "-L", bounceback_tar_url, "-o", "package.tar.gz")
	err := bounceback_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bounceback_zip_url := "https://github.com/D00Movenok/BounceBack/archive/refs/tags/v1.5.1.zip"
	bounceback_cmd_zip := exec.Command("curl", "-L", bounceback_zip_url, "-o", "package.zip")
	err = bounceback_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bounceback_bin_url := "https://github.com/D00Movenok/BounceBack/archive/refs/tags/v1.5.1.bin"
	bounceback_cmd_bin := exec.Command("curl", "-L", bounceback_bin_url, "-o", "binary.bin")
	err = bounceback_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bounceback_src_url := "https://github.com/D00Movenok/BounceBack/archive/refs/tags/v1.5.1.src.tar.gz"
	bounceback_cmd_src := exec.Command("curl", "-L", bounceback_src_url, "-o", "source.tar.gz")
	err = bounceback_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bounceback_cmd_direct := exec.Command("./binary")
	err = bounceback_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
