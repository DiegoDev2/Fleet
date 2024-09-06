package main

// Trunk - Build, bundle & ship your Rust WASM application to the web
// Homepage: https://trunkrs.dev/

import (
	"fmt"
	
	"os/exec"
)

func installTrunk() {
	// Método 1: Descargar y extraer .tar.gz
	trunk_tar_url := "https://github.com/trunk-rs/trunk/archive/refs/tags/v0.20.3.tar.gz"
	trunk_cmd_tar := exec.Command("curl", "-L", trunk_tar_url, "-o", "package.tar.gz")
	err := trunk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trunk_zip_url := "https://github.com/trunk-rs/trunk/archive/refs/tags/v0.20.3.zip"
	trunk_cmd_zip := exec.Command("curl", "-L", trunk_zip_url, "-o", "package.zip")
	err = trunk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trunk_bin_url := "https://github.com/trunk-rs/trunk/archive/refs/tags/v0.20.3.bin"
	trunk_cmd_bin := exec.Command("curl", "-L", trunk_bin_url, "-o", "binary.bin")
	err = trunk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trunk_src_url := "https://github.com/trunk-rs/trunk/archive/refs/tags/v0.20.3.src.tar.gz"
	trunk_cmd_src := exec.Command("curl", "-L", trunk_src_url, "-o", "source.tar.gz")
	err = trunk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trunk_cmd_direct := exec.Command("./binary")
	err = trunk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
