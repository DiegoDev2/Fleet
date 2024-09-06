package main

// Phoneinfoga - Information gathering framework for phone numbers
// Homepage: https://sundowndev.github.io/phoneinfoga/

import (
	"fmt"
	
	"os/exec"
)

func installPhoneinfoga() {
	// Método 1: Descargar y extraer .tar.gz
	phoneinfoga_tar_url := "https://github.com/sundowndev/phoneinfoga/archive/refs/tags/v2.11.0.tar.gz"
	phoneinfoga_cmd_tar := exec.Command("curl", "-L", phoneinfoga_tar_url, "-o", "package.tar.gz")
	err := phoneinfoga_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phoneinfoga_zip_url := "https://github.com/sundowndev/phoneinfoga/archive/refs/tags/v2.11.0.zip"
	phoneinfoga_cmd_zip := exec.Command("curl", "-L", phoneinfoga_zip_url, "-o", "package.zip")
	err = phoneinfoga_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phoneinfoga_bin_url := "https://github.com/sundowndev/phoneinfoga/archive/refs/tags/v2.11.0.bin"
	phoneinfoga_cmd_bin := exec.Command("curl", "-L", phoneinfoga_bin_url, "-o", "binary.bin")
	err = phoneinfoga_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phoneinfoga_src_url := "https://github.com/sundowndev/phoneinfoga/archive/refs/tags/v2.11.0.src.tar.gz"
	phoneinfoga_cmd_src := exec.Command("curl", "-L", phoneinfoga_src_url, "-o", "source.tar.gz")
	err = phoneinfoga_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phoneinfoga_cmd_direct := exec.Command("./binary")
	err = phoneinfoga_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: yarn")
exec.Command("latte", "install", "yarn")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
