package main

// Asuka - Gemini Project client written in Rust with NCurses
// Homepage: https://sr.ht/~julienxx/Asuka/

import (
	"fmt"
	
	"os/exec"
)

func installAsuka() {
	// Método 1: Descargar y extraer .tar.gz
	asuka_tar_url := "https://git.sr.ht/~julienxx/asuka/archive/0.8.5.tar.gz"
	asuka_cmd_tar := exec.Command("curl", "-L", asuka_tar_url, "-o", "package.tar.gz")
	err := asuka_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asuka_zip_url := "https://git.sr.ht/~julienxx/asuka/archive/0.8.5.zip"
	asuka_cmd_zip := exec.Command("curl", "-L", asuka_zip_url, "-o", "package.zip")
	err = asuka_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asuka_bin_url := "https://git.sr.ht/~julienxx/asuka/archive/0.8.5.bin"
	asuka_cmd_bin := exec.Command("curl", "-L", asuka_bin_url, "-o", "binary.bin")
	err = asuka_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asuka_src_url := "https://git.sr.ht/~julienxx/asuka/archive/0.8.5.src.tar.gz"
	asuka_cmd_src := exec.Command("curl", "-L", asuka_src_url, "-o", "source.tar.gz")
	err = asuka_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asuka_cmd_direct := exec.Command("./binary")
	err = asuka_cmd_direct.Run()
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
