package main

// Dsocks - SOCKS client wrapper for *BSD/macOS
// Homepage: https://monkey.org/~dugsong/dsocks/

import (
	"fmt"
	
	"os/exec"
)

func installDsocks() {
	// Método 1: Descargar y extraer .tar.gz
	dsocks_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/dsocks/dsocks-1.8.tar.gz"
	dsocks_cmd_tar := exec.Command("curl", "-L", dsocks_tar_url, "-o", "package.tar.gz")
	err := dsocks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dsocks_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/dsocks/dsocks-1.8.zip"
	dsocks_cmd_zip := exec.Command("curl", "-L", dsocks_zip_url, "-o", "package.zip")
	err = dsocks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dsocks_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/dsocks/dsocks-1.8.bin"
	dsocks_cmd_bin := exec.Command("curl", "-L", dsocks_bin_url, "-o", "binary.bin")
	err = dsocks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dsocks_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/dsocks/dsocks-1.8.src.tar.gz"
	dsocks_cmd_src := exec.Command("curl", "-L", dsocks_src_url, "-o", "source.tar.gz")
	err = dsocks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dsocks_cmd_direct := exec.Command("./binary")
	err = dsocks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
