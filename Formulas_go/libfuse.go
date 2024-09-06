package main

// Libfuse - Reference implementation of the Linux FUSE interface
// Homepage: https://github.com/libfuse/libfuse

import (
	"fmt"
	
	"os/exec"
)

func installLibfuse() {
	// Método 1: Descargar y extraer .tar.gz
	libfuse_tar_url := "https://github.com/libfuse/libfuse/releases/download/fuse-3.16.2/fuse-3.16.2.tar.gz"
	libfuse_cmd_tar := exec.Command("curl", "-L", libfuse_tar_url, "-o", "package.tar.gz")
	err := libfuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfuse_zip_url := "https://github.com/libfuse/libfuse/releases/download/fuse-3.16.2/fuse-3.16.2.zip"
	libfuse_cmd_zip := exec.Command("curl", "-L", libfuse_zip_url, "-o", "package.zip")
	err = libfuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfuse_bin_url := "https://github.com/libfuse/libfuse/releases/download/fuse-3.16.2/fuse-3.16.2.bin"
	libfuse_cmd_bin := exec.Command("curl", "-L", libfuse_bin_url, "-o", "binary.bin")
	err = libfuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfuse_src_url := "https://github.com/libfuse/libfuse/releases/download/fuse-3.16.2/fuse-3.16.2.src.tar.gz"
	libfuse_cmd_src := exec.Command("curl", "-L", libfuse_src_url, "-o", "source.tar.gz")
	err = libfuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfuse_cmd_direct := exec.Command("./binary")
	err = libfuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
