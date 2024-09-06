package main

// GambitScheme - Implementation of the Scheme Language
// Homepage: https://github.com/gambit/gambit

import (
	"fmt"
	
	"os/exec"
)

func installGambitScheme() {
	// Método 1: Descargar y extraer .tar.gz
	gambitscheme_tar_url := "https://github.com/gambit/gambit/archive/refs/tags/v4.9.5.tar.gz"
	gambitscheme_cmd_tar := exec.Command("curl", "-L", gambitscheme_tar_url, "-o", "package.tar.gz")
	err := gambitscheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gambitscheme_zip_url := "https://github.com/gambit/gambit/archive/refs/tags/v4.9.5.zip"
	gambitscheme_cmd_zip := exec.Command("curl", "-L", gambitscheme_zip_url, "-o", "package.zip")
	err = gambitscheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gambitscheme_bin_url := "https://github.com/gambit/gambit/archive/refs/tags/v4.9.5.bin"
	gambitscheme_cmd_bin := exec.Command("curl", "-L", gambitscheme_bin_url, "-o", "binary.bin")
	err = gambitscheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gambitscheme_src_url := "https://github.com/gambit/gambit/archive/refs/tags/v4.9.5.src.tar.gz"
	gambitscheme_cmd_src := exec.Command("curl", "-L", gambitscheme_src_url, "-o", "source.tar.gz")
	err = gambitscheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gambitscheme_cmd_direct := exec.Command("./binary")
	err = gambitscheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
