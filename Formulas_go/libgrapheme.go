package main

// Libgrapheme - Unicode string library
// Homepage: https://libs.suckless.org/libgrapheme/

import (
	"fmt"
	
	"os/exec"
)

func installLibgrapheme() {
	// Método 1: Descargar y extraer .tar.gz
	libgrapheme_tar_url := "https://dl.suckless.org/libgrapheme/libgrapheme-2.0.2.tar.gz"
	libgrapheme_cmd_tar := exec.Command("curl", "-L", libgrapheme_tar_url, "-o", "package.tar.gz")
	err := libgrapheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgrapheme_zip_url := "https://dl.suckless.org/libgrapheme/libgrapheme-2.0.2.zip"
	libgrapheme_cmd_zip := exec.Command("curl", "-L", libgrapheme_zip_url, "-o", "package.zip")
	err = libgrapheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgrapheme_bin_url := "https://dl.suckless.org/libgrapheme/libgrapheme-2.0.2.bin"
	libgrapheme_cmd_bin := exec.Command("curl", "-L", libgrapheme_bin_url, "-o", "binary.bin")
	err = libgrapheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgrapheme_src_url := "https://dl.suckless.org/libgrapheme/libgrapheme-2.0.2.src.tar.gz"
	libgrapheme_cmd_src := exec.Command("curl", "-L", libgrapheme_src_url, "-o", "source.tar.gz")
	err = libgrapheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgrapheme_cmd_direct := exec.Command("./binary")
	err = libgrapheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
