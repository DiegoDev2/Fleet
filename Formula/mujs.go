package main

// Mujs - Embeddable Javascript interpreter
// Homepage: https://www.mujs.com/

import (
	"fmt"
	
	"os/exec"
)

func installMujs() {
	// Método 1: Descargar y extraer .tar.gz
	mujs_tar_url := "https://mujs.com/downloads/mujs-1.3.5.tar.gz"
	mujs_cmd_tar := exec.Command("curl", "-L", mujs_tar_url, "-o", "package.tar.gz")
	err := mujs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mujs_zip_url := "https://mujs.com/downloads/mujs-1.3.5.zip"
	mujs_cmd_zip := exec.Command("curl", "-L", mujs_zip_url, "-o", "package.zip")
	err = mujs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mujs_bin_url := "https://mujs.com/downloads/mujs-1.3.5.bin"
	mujs_cmd_bin := exec.Command("curl", "-L", mujs_bin_url, "-o", "binary.bin")
	err = mujs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mujs_src_url := "https://mujs.com/downloads/mujs-1.3.5.src.tar.gz"
	mujs_cmd_src := exec.Command("curl", "-L", mujs_src_url, "-o", "source.tar.gz")
	err = mujs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mujs_cmd_direct := exec.Command("./binary")
	err = mujs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
