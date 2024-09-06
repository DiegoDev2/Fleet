package main

// Utf8proc - Clean C library for processing UTF-8 Unicode data
// Homepage: https://juliastrings.github.io/utf8proc/

import (
	"fmt"
	
	"os/exec"
)

func installUtf8proc() {
	// Método 1: Descargar y extraer .tar.gz
	utf8proc_tar_url := "https://github.com/JuliaStrings/utf8proc/archive/refs/tags/v2.9.0.tar.gz"
	utf8proc_cmd_tar := exec.Command("curl", "-L", utf8proc_tar_url, "-o", "package.tar.gz")
	err := utf8proc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	utf8proc_zip_url := "https://github.com/JuliaStrings/utf8proc/archive/refs/tags/v2.9.0.zip"
	utf8proc_cmd_zip := exec.Command("curl", "-L", utf8proc_zip_url, "-o", "package.zip")
	err = utf8proc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	utf8proc_bin_url := "https://github.com/JuliaStrings/utf8proc/archive/refs/tags/v2.9.0.bin"
	utf8proc_cmd_bin := exec.Command("curl", "-L", utf8proc_bin_url, "-o", "binary.bin")
	err = utf8proc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	utf8proc_src_url := "https://github.com/JuliaStrings/utf8proc/archive/refs/tags/v2.9.0.src.tar.gz"
	utf8proc_cmd_src := exec.Command("curl", "-L", utf8proc_src_url, "-o", "source.tar.gz")
	err = utf8proc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	utf8proc_cmd_direct := exec.Command("./binary")
	err = utf8proc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
