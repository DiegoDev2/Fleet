package main

// Ext4fuse - Read-only implementation of ext4 for FUSE
// Homepage: https://github.com/gerard/ext4fuse

import (
	"fmt"
	
	"os/exec"
)

func installExt4fuse() {
	// Método 1: Descargar y extraer .tar.gz
	ext4fuse_tar_url := "https://github.com/gerard/ext4fuse/archive/refs/tags/v0.1.3.tar.gz"
	ext4fuse_cmd_tar := exec.Command("curl", "-L", ext4fuse_tar_url, "-o", "package.tar.gz")
	err := ext4fuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ext4fuse_zip_url := "https://github.com/gerard/ext4fuse/archive/refs/tags/v0.1.3.zip"
	ext4fuse_cmd_zip := exec.Command("curl", "-L", ext4fuse_zip_url, "-o", "package.zip")
	err = ext4fuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ext4fuse_bin_url := "https://github.com/gerard/ext4fuse/archive/refs/tags/v0.1.3.bin"
	ext4fuse_cmd_bin := exec.Command("curl", "-L", ext4fuse_bin_url, "-o", "binary.bin")
	err = ext4fuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ext4fuse_src_url := "https://github.com/gerard/ext4fuse/archive/refs/tags/v0.1.3.src.tar.gz"
	ext4fuse_cmd_src := exec.Command("curl", "-L", ext4fuse_src_url, "-o", "source.tar.gz")
	err = ext4fuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ext4fuse_cmd_direct := exec.Command("./binary")
	err = ext4fuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
}
