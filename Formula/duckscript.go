package main

// Duckscript - Simple, extendable and embeddable scripting language
// Homepage: https://sagiegurari.github.io/duckscript

import (
	"fmt"
	
	"os/exec"
)

func installDuckscript() {
	// Método 1: Descargar y extraer .tar.gz
	duckscript_tar_url := "https://github.com/sagiegurari/duckscript/archive/refs/tags/0.9.3.tar.gz"
	duckscript_cmd_tar := exec.Command("curl", "-L", duckscript_tar_url, "-o", "package.tar.gz")
	err := duckscript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	duckscript_zip_url := "https://github.com/sagiegurari/duckscript/archive/refs/tags/0.9.3.zip"
	duckscript_cmd_zip := exec.Command("curl", "-L", duckscript_zip_url, "-o", "package.zip")
	err = duckscript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	duckscript_bin_url := "https://github.com/sagiegurari/duckscript/archive/refs/tags/0.9.3.bin"
	duckscript_cmd_bin := exec.Command("curl", "-L", duckscript_bin_url, "-o", "binary.bin")
	err = duckscript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	duckscript_src_url := "https://github.com/sagiegurari/duckscript/archive/refs/tags/0.9.3.src.tar.gz"
	duckscript_cmd_src := exec.Command("curl", "-L", duckscript_src_url, "-o", "source.tar.gz")
	err = duckscript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	duckscript_cmd_direct := exec.Command("./binary")
	err = duckscript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
