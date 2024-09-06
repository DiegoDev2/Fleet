package main

// Viddy - Modern watch command
// Homepage: https://github.com/sachaos/viddy

import (
	"fmt"
	
	"os/exec"
)

func installViddy() {
	// Método 1: Descargar y extraer .tar.gz
	viddy_tar_url := "https://github.com/sachaos/viddy/archive/refs/tags/v1.1.2.tar.gz"
	viddy_cmd_tar := exec.Command("curl", "-L", viddy_tar_url, "-o", "package.tar.gz")
	err := viddy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	viddy_zip_url := "https://github.com/sachaos/viddy/archive/refs/tags/v1.1.2.zip"
	viddy_cmd_zip := exec.Command("curl", "-L", viddy_zip_url, "-o", "package.zip")
	err = viddy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	viddy_bin_url := "https://github.com/sachaos/viddy/archive/refs/tags/v1.1.2.bin"
	viddy_cmd_bin := exec.Command("curl", "-L", viddy_bin_url, "-o", "binary.bin")
	err = viddy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	viddy_src_url := "https://github.com/sachaos/viddy/archive/refs/tags/v1.1.2.src.tar.gz"
	viddy_cmd_src := exec.Command("curl", "-L", viddy_src_url, "-o", "source.tar.gz")
	err = viddy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	viddy_cmd_direct := exec.Command("./binary")
	err = viddy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
