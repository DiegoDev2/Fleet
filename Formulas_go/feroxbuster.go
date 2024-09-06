package main

// Feroxbuster - Fast, simple, recursive content discovery tool written in Rust
// Homepage: https://epi052.github.io/feroxbuster

import (
	"fmt"
	
	"os/exec"
)

func installFeroxbuster() {
	// Método 1: Descargar y extraer .tar.gz
	feroxbuster_tar_url := "https://github.com/epi052/feroxbuster/archive/refs/tags/v2.10.4.tar.gz"
	feroxbuster_cmd_tar := exec.Command("curl", "-L", feroxbuster_tar_url, "-o", "package.tar.gz")
	err := feroxbuster_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	feroxbuster_zip_url := "https://github.com/epi052/feroxbuster/archive/refs/tags/v2.10.4.zip"
	feroxbuster_cmd_zip := exec.Command("curl", "-L", feroxbuster_zip_url, "-o", "package.zip")
	err = feroxbuster_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	feroxbuster_bin_url := "https://github.com/epi052/feroxbuster/archive/refs/tags/v2.10.4.bin"
	feroxbuster_cmd_bin := exec.Command("curl", "-L", feroxbuster_bin_url, "-o", "binary.bin")
	err = feroxbuster_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	feroxbuster_src_url := "https://github.com/epi052/feroxbuster/archive/refs/tags/v2.10.4.src.tar.gz"
	feroxbuster_cmd_src := exec.Command("curl", "-L", feroxbuster_src_url, "-o", "source.tar.gz")
	err = feroxbuster_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	feroxbuster_cmd_direct := exec.Command("./binary")
	err = feroxbuster_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: miniserve")
exec.Command("latte", "install", "miniserve")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
