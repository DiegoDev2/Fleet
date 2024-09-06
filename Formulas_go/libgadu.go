package main

// Libgadu - Library for ICQ instant messenger protocol
// Homepage: https://libgadu.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibgadu() {
	// Método 1: Descargar y extraer .tar.gz
	libgadu_tar_url := "https://github.com/wojtekka/libgadu/releases/download/1.12.2/libgadu-1.12.2.tar.gz"
	libgadu_cmd_tar := exec.Command("curl", "-L", libgadu_tar_url, "-o", "package.tar.gz")
	err := libgadu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgadu_zip_url := "https://github.com/wojtekka/libgadu/releases/download/1.12.2/libgadu-1.12.2.zip"
	libgadu_cmd_zip := exec.Command("curl", "-L", libgadu_zip_url, "-o", "package.zip")
	err = libgadu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgadu_bin_url := "https://github.com/wojtekka/libgadu/releases/download/1.12.2/libgadu-1.12.2.bin"
	libgadu_cmd_bin := exec.Command("curl", "-L", libgadu_bin_url, "-o", "binary.bin")
	err = libgadu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgadu_src_url := "https://github.com/wojtekka/libgadu/releases/download/1.12.2/libgadu-1.12.2.src.tar.gz"
	libgadu_cmd_src := exec.Command("curl", "-L", libgadu_src_url, "-o", "source.tar.gz")
	err = libgadu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgadu_cmd_direct := exec.Command("./binary")
	err = libgadu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
