package main

// Drill - HTTP load testing application written in Rust
// Homepage: https://github.com/fcsonline/drill

import (
	"fmt"
	
	"os/exec"
)

func installDrill() {
	// Método 1: Descargar y extraer .tar.gz
	drill_tar_url := "https://github.com/fcsonline/drill/archive/refs/tags/0.8.3.tar.gz"
	drill_cmd_tar := exec.Command("curl", "-L", drill_tar_url, "-o", "package.tar.gz")
	err := drill_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	drill_zip_url := "https://github.com/fcsonline/drill/archive/refs/tags/0.8.3.zip"
	drill_cmd_zip := exec.Command("curl", "-L", drill_zip_url, "-o", "package.zip")
	err = drill_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	drill_bin_url := "https://github.com/fcsonline/drill/archive/refs/tags/0.8.3.bin"
	drill_cmd_bin := exec.Command("curl", "-L", drill_bin_url, "-o", "binary.bin")
	err = drill_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	drill_src_url := "https://github.com/fcsonline/drill/archive/refs/tags/0.8.3.src.tar.gz"
	drill_cmd_src := exec.Command("curl", "-L", drill_src_url, "-o", "source.tar.gz")
	err = drill_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	drill_cmd_direct := exec.Command("./binary")
	err = drill_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
