package main

// Uv - Extremely fast Python package installer and resolver, written in Rust
// Homepage: https://github.com/astral-sh/uv

import (
	"fmt"
	
	"os/exec"
)

func installUv() {
	// Método 1: Descargar y extraer .tar.gz
	uv_tar_url := "https://github.com/astral-sh/uv/archive/refs/tags/0.4.6.tar.gz"
	uv_cmd_tar := exec.Command("curl", "-L", uv_tar_url, "-o", "package.tar.gz")
	err := uv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uv_zip_url := "https://github.com/astral-sh/uv/archive/refs/tags/0.4.6.zip"
	uv_cmd_zip := exec.Command("curl", "-L", uv_zip_url, "-o", "package.zip")
	err = uv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uv_bin_url := "https://github.com/astral-sh/uv/archive/refs/tags/0.4.6.bin"
	uv_cmd_bin := exec.Command("curl", "-L", uv_bin_url, "-o", "binary.bin")
	err = uv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uv_src_url := "https://github.com/astral-sh/uv/archive/refs/tags/0.4.6.src.tar.gz"
	uv_cmd_src := exec.Command("curl", "-L", uv_src_url, "-o", "source.tar.gz")
	err = uv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uv_cmd_direct := exec.Command("./binary")
	err = uv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: bzip2")
exec.Command("latte", "install", "bzip2")
}
