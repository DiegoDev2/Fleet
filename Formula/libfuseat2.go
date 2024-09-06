package main

// LibfuseAT2 - Reference implementation of the Linux FUSE interface
// Homepage: https://github.com/libfuse/libfuse

import (
	"fmt"
	
	"os/exec"
)

func installLibfuseAT2() {
	// Método 1: Descargar y extraer .tar.gz
	libfuseat2_tar_url := "https://github.com/libfuse/libfuse/releases/download/fuse-2.9.9/fuse-2.9.9.tar.gz"
	libfuseat2_cmd_tar := exec.Command("curl", "-L", libfuseat2_tar_url, "-o", "package.tar.gz")
	err := libfuseat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfuseat2_zip_url := "https://github.com/libfuse/libfuse/releases/download/fuse-2.9.9/fuse-2.9.9.zip"
	libfuseat2_cmd_zip := exec.Command("curl", "-L", libfuseat2_zip_url, "-o", "package.zip")
	err = libfuseat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfuseat2_bin_url := "https://github.com/libfuse/libfuse/releases/download/fuse-2.9.9/fuse-2.9.9.bin"
	libfuseat2_cmd_bin := exec.Command("curl", "-L", libfuseat2_bin_url, "-o", "binary.bin")
	err = libfuseat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfuseat2_src_url := "https://github.com/libfuse/libfuse/releases/download/fuse-2.9.9/fuse-2.9.9.src.tar.gz"
	libfuseat2_cmd_src := exec.Command("curl", "-L", libfuseat2_src_url, "-o", "source.tar.gz")
	err = libfuseat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfuseat2_cmd_direct := exec.Command("./binary")
	err = libfuseat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
