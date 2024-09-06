package main

// Tinysearch - Tiny, full-text search engine for static websites built with Rust and Wasm
// Homepage: https://github.com/tinysearch/tinysearch

import (
	"fmt"
	
	"os/exec"
)

func installTinysearch() {
	// Método 1: Descargar y extraer .tar.gz
	tinysearch_tar_url := "https://github.com/tinysearch/tinysearch/archive/refs/tags/v0.8.2.tar.gz"
	tinysearch_cmd_tar := exec.Command("curl", "-L", tinysearch_tar_url, "-o", "package.tar.gz")
	err := tinysearch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tinysearch_zip_url := "https://github.com/tinysearch/tinysearch/archive/refs/tags/v0.8.2.zip"
	tinysearch_cmd_zip := exec.Command("curl", "-L", tinysearch_zip_url, "-o", "package.zip")
	err = tinysearch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tinysearch_bin_url := "https://github.com/tinysearch/tinysearch/archive/refs/tags/v0.8.2.bin"
	tinysearch_cmd_bin := exec.Command("curl", "-L", tinysearch_bin_url, "-o", "binary.bin")
	err = tinysearch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tinysearch_src_url := "https://github.com/tinysearch/tinysearch/archive/refs/tags/v0.8.2.src.tar.gz"
	tinysearch_cmd_src := exec.Command("curl", "-L", tinysearch_src_url, "-o", "source.tar.gz")
	err = tinysearch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tinysearch_cmd_direct := exec.Command("./binary")
	err = tinysearch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
	fmt.Println("Instalando dependencia: wasm-pack")
	exec.Command("latte", "install", "wasm-pack").Run()
}
