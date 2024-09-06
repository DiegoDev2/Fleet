package main

// DerAscii - Reversible DER and BER pretty-printer
// Homepage: https://github.com/google/der-ascii

import (
	"fmt"
	
	"os/exec"
)

func installDerAscii() {
	// Método 1: Descargar y extraer .tar.gz
	derascii_tar_url := "https://github.com/google/der-ascii/archive/refs/tags/v0.3.0.tar.gz"
	derascii_cmd_tar := exec.Command("curl", "-L", derascii_tar_url, "-o", "package.tar.gz")
	err := derascii_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	derascii_zip_url := "https://github.com/google/der-ascii/archive/refs/tags/v0.3.0.zip"
	derascii_cmd_zip := exec.Command("curl", "-L", derascii_zip_url, "-o", "package.zip")
	err = derascii_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	derascii_bin_url := "https://github.com/google/der-ascii/archive/refs/tags/v0.3.0.bin"
	derascii_cmd_bin := exec.Command("curl", "-L", derascii_bin_url, "-o", "binary.bin")
	err = derascii_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	derascii_src_url := "https://github.com/google/der-ascii/archive/refs/tags/v0.3.0.src.tar.gz"
	derascii_cmd_src := exec.Command("curl", "-L", derascii_src_url, "-o", "source.tar.gz")
	err = derascii_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	derascii_cmd_direct := exec.Command("./binary")
	err = derascii_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
