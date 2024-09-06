package main

// Aftman - Toolchain manager for Roblox, the prodigal sequel to Foreman
// Homepage: https://github.com/LPGhatguy/aftman

import (
	"fmt"
	
	"os/exec"
)

func installAftman() {
	// Método 1: Descargar y extraer .tar.gz
	aftman_tar_url := "https://github.com/LPGhatguy/aftman/archive/refs/tags/v0.3.0.tar.gz"
	aftman_cmd_tar := exec.Command("curl", "-L", aftman_tar_url, "-o", "package.tar.gz")
	err := aftman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aftman_zip_url := "https://github.com/LPGhatguy/aftman/archive/refs/tags/v0.3.0.zip"
	aftman_cmd_zip := exec.Command("curl", "-L", aftman_zip_url, "-o", "package.zip")
	err = aftman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aftman_bin_url := "https://github.com/LPGhatguy/aftman/archive/refs/tags/v0.3.0.bin"
	aftman_cmd_bin := exec.Command("curl", "-L", aftman_bin_url, "-o", "binary.bin")
	err = aftman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aftman_src_url := "https://github.com/LPGhatguy/aftman/archive/refs/tags/v0.3.0.src.tar.gz"
	aftman_cmd_src := exec.Command("curl", "-L", aftman_src_url, "-o", "source.tar.gz")
	err = aftman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aftman_cmd_direct := exec.Command("./binary")
	err = aftman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
