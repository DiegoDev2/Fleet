package main

// Libmobi - C library for handling Kindle (MOBI) formats of ebook documents
// Homepage: https://github.com/bfabiszewski/libmobi

import (
	"fmt"
	
	"os/exec"
)

func installLibmobi() {
	// Método 1: Descargar y extraer .tar.gz
	libmobi_tar_url := "https://github.com/bfabiszewski/libmobi/releases/download/v0.12/libmobi-0.12.tar.gz"
	libmobi_cmd_tar := exec.Command("curl", "-L", libmobi_tar_url, "-o", "package.tar.gz")
	err := libmobi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmobi_zip_url := "https://github.com/bfabiszewski/libmobi/releases/download/v0.12/libmobi-0.12.zip"
	libmobi_cmd_zip := exec.Command("curl", "-L", libmobi_zip_url, "-o", "package.zip")
	err = libmobi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmobi_bin_url := "https://github.com/bfabiszewski/libmobi/releases/download/v0.12/libmobi-0.12.bin"
	libmobi_cmd_bin := exec.Command("curl", "-L", libmobi_bin_url, "-o", "binary.bin")
	err = libmobi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmobi_src_url := "https://github.com/bfabiszewski/libmobi/releases/download/v0.12/libmobi-0.12.src.tar.gz"
	libmobi_cmd_src := exec.Command("curl", "-L", libmobi_src_url, "-o", "source.tar.gz")
	err = libmobi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmobi_cmd_direct := exec.Command("./binary")
	err = libmobi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
