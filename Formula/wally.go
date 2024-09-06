package main

// Wally - Modern package manager for Roblox projects inspired by Cargo
// Homepage: https://github.com/UpliftGames/wally

import (
	"fmt"
	
	"os/exec"
)

func installWally() {
	// Método 1: Descargar y extraer .tar.gz
	wally_tar_url := "https://github.com/UpliftGames/wally/archive/refs/tags/v0.3.2.tar.gz"
	wally_cmd_tar := exec.Command("curl", "-L", wally_tar_url, "-o", "package.tar.gz")
	err := wally_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wally_zip_url := "https://github.com/UpliftGames/wally/archive/refs/tags/v0.3.2.zip"
	wally_cmd_zip := exec.Command("curl", "-L", wally_zip_url, "-o", "package.zip")
	err = wally_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wally_bin_url := "https://github.com/UpliftGames/wally/archive/refs/tags/v0.3.2.bin"
	wally_cmd_bin := exec.Command("curl", "-L", wally_bin_url, "-o", "binary.bin")
	err = wally_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wally_src_url := "https://github.com/UpliftGames/wally/archive/refs/tags/v0.3.2.src.tar.gz"
	wally_cmd_src := exec.Command("curl", "-L", wally_src_url, "-o", "source.tar.gz")
	err = wally_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wally_cmd_direct := exec.Command("./binary")
	err = wally_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
