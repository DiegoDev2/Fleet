package main

// Fend - Arbitrary-precision unit-aware calculator
// Homepage: https://printfn.github.io/fend

import (
	"fmt"
	
	"os/exec"
)

func installFend() {
	// Método 1: Descargar y extraer .tar.gz
	fend_tar_url := "https://github.com/printfn/fend/archive/refs/tags/v1.5.1.tar.gz"
	fend_cmd_tar := exec.Command("curl", "-L", fend_tar_url, "-o", "package.tar.gz")
	err := fend_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fend_zip_url := "https://github.com/printfn/fend/archive/refs/tags/v1.5.1.zip"
	fend_cmd_zip := exec.Command("curl", "-L", fend_zip_url, "-o", "package.zip")
	err = fend_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fend_bin_url := "https://github.com/printfn/fend/archive/refs/tags/v1.5.1.bin"
	fend_cmd_bin := exec.Command("curl", "-L", fend_bin_url, "-o", "binary.bin")
	err = fend_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fend_src_url := "https://github.com/printfn/fend/archive/refs/tags/v1.5.1.src.tar.gz"
	fend_cmd_src := exec.Command("curl", "-L", fend_src_url, "-o", "source.tar.gz")
	err = fend_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fend_cmd_direct := exec.Command("./binary")
	err = fend_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
