package main

// Ffuf - Fast web fuzzer written in Go
// Homepage: https://github.com/ffuf/ffuf

import (
	"fmt"
	
	"os/exec"
)

func installFfuf() {
	// Método 1: Descargar y extraer .tar.gz
	ffuf_tar_url := "https://github.com/ffuf/ffuf/archive/refs/tags/v2.1.0.tar.gz"
	ffuf_cmd_tar := exec.Command("curl", "-L", ffuf_tar_url, "-o", "package.tar.gz")
	err := ffuf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ffuf_zip_url := "https://github.com/ffuf/ffuf/archive/refs/tags/v2.1.0.zip"
	ffuf_cmd_zip := exec.Command("curl", "-L", ffuf_zip_url, "-o", "package.zip")
	err = ffuf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ffuf_bin_url := "https://github.com/ffuf/ffuf/archive/refs/tags/v2.1.0.bin"
	ffuf_cmd_bin := exec.Command("curl", "-L", ffuf_bin_url, "-o", "binary.bin")
	err = ffuf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ffuf_src_url := "https://github.com/ffuf/ffuf/archive/refs/tags/v2.1.0.src.tar.gz"
	ffuf_cmd_src := exec.Command("curl", "-L", ffuf_src_url, "-o", "source.tar.gz")
	err = ffuf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ffuf_cmd_direct := exec.Command("./binary")
	err = ffuf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
