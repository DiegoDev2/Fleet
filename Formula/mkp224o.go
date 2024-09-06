package main

// Mkp224o - Vanity address generator for tor onion v3 (ed25519) hidden services
// Homepage: https://github.com/cathugger/mkp224o

import (
	"fmt"
	
	"os/exec"
)

func installMkp224o() {
	// Método 1: Descargar y extraer .tar.gz
	mkp224o_tar_url := "https://github.com/cathugger/mkp224o/releases/download/v1.7.0/mkp224o-1.7.0-src.tar.gz"
	mkp224o_cmd_tar := exec.Command("curl", "-L", mkp224o_tar_url, "-o", "package.tar.gz")
	err := mkp224o_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkp224o_zip_url := "https://github.com/cathugger/mkp224o/releases/download/v1.7.0/mkp224o-1.7.0-src.zip"
	mkp224o_cmd_zip := exec.Command("curl", "-L", mkp224o_zip_url, "-o", "package.zip")
	err = mkp224o_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkp224o_bin_url := "https://github.com/cathugger/mkp224o/releases/download/v1.7.0/mkp224o-1.7.0-src.bin"
	mkp224o_cmd_bin := exec.Command("curl", "-L", mkp224o_bin_url, "-o", "binary.bin")
	err = mkp224o_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkp224o_src_url := "https://github.com/cathugger/mkp224o/releases/download/v1.7.0/mkp224o-1.7.0-src.src.tar.gz"
	mkp224o_cmd_src := exec.Command("curl", "-L", mkp224o_src_url, "-o", "source.tar.gz")
	err = mkp224o_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkp224o_cmd_direct := exec.Command("./binary")
	err = mkp224o_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
}
