package main

// Dylibbundler - Utility to bundle libraries into executables for macOS
// Homepage: https://github.com/auriamg/macdylibbundler

import (
	"fmt"
	
	"os/exec"
)

func installDylibbundler() {
	// Método 1: Descargar y extraer .tar.gz
	dylibbundler_tar_url := "https://github.com/auriamg/macdylibbundler/archive/refs/tags/1.0.5.tar.gz"
	dylibbundler_cmd_tar := exec.Command("curl", "-L", dylibbundler_tar_url, "-o", "package.tar.gz")
	err := dylibbundler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dylibbundler_zip_url := "https://github.com/auriamg/macdylibbundler/archive/refs/tags/1.0.5.zip"
	dylibbundler_cmd_zip := exec.Command("curl", "-L", dylibbundler_zip_url, "-o", "package.zip")
	err = dylibbundler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dylibbundler_bin_url := "https://github.com/auriamg/macdylibbundler/archive/refs/tags/1.0.5.bin"
	dylibbundler_cmd_bin := exec.Command("curl", "-L", dylibbundler_bin_url, "-o", "binary.bin")
	err = dylibbundler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dylibbundler_src_url := "https://github.com/auriamg/macdylibbundler/archive/refs/tags/1.0.5.src.tar.gz"
	dylibbundler_cmd_src := exec.Command("curl", "-L", dylibbundler_src_url, "-o", "source.tar.gz")
	err = dylibbundler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dylibbundler_cmd_direct := exec.Command("./binary")
	err = dylibbundler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
