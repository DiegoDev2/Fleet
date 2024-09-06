package main

// Libestr - C library for string handling (and a bit more)
// Homepage: https://libestr.adiscon.com/

import (
	"fmt"
	
	"os/exec"
)

func installLibestr() {
	// Método 1: Descargar y extraer .tar.gz
	libestr_tar_url := "https://libestr.adiscon.com/files/download/libestr-0.1.11.tar.gz"
	libestr_cmd_tar := exec.Command("curl", "-L", libestr_tar_url, "-o", "package.tar.gz")
	err := libestr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libestr_zip_url := "https://libestr.adiscon.com/files/download/libestr-0.1.11.zip"
	libestr_cmd_zip := exec.Command("curl", "-L", libestr_zip_url, "-o", "package.zip")
	err = libestr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libestr_bin_url := "https://libestr.adiscon.com/files/download/libestr-0.1.11.bin"
	libestr_cmd_bin := exec.Command("curl", "-L", libestr_bin_url, "-o", "binary.bin")
	err = libestr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libestr_src_url := "https://libestr.adiscon.com/files/download/libestr-0.1.11.src.tar.gz"
	libestr_cmd_src := exec.Command("curl", "-L", libestr_src_url, "-o", "source.tar.gz")
	err = libestr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libestr_cmd_direct := exec.Command("./binary")
	err = libestr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
