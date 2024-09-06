package main

// Squashfs - Compressed read-only file system for Linux
// Homepage: https://github.com/plougher/squashfs-tools

import (
	"fmt"
	
	"os/exec"
)

func installSquashfs() {
	// Método 1: Descargar y extraer .tar.gz
	squashfs_tar_url := "https://github.com/plougher/squashfs-tools/archive/refs/tags/4.6.1.tar.gz"
	squashfs_cmd_tar := exec.Command("curl", "-L", squashfs_tar_url, "-o", "package.tar.gz")
	err := squashfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	squashfs_zip_url := "https://github.com/plougher/squashfs-tools/archive/refs/tags/4.6.1.zip"
	squashfs_cmd_zip := exec.Command("curl", "-L", squashfs_zip_url, "-o", "package.zip")
	err = squashfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	squashfs_bin_url := "https://github.com/plougher/squashfs-tools/archive/refs/tags/4.6.1.bin"
	squashfs_cmd_bin := exec.Command("curl", "-L", squashfs_bin_url, "-o", "binary.bin")
	err = squashfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	squashfs_src_url := "https://github.com/plougher/squashfs-tools/archive/refs/tags/4.6.1.src.tar.gz"
	squashfs_cmd_src := exec.Command("curl", "-L", squashfs_src_url, "-o", "source.tar.gz")
	err = squashfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	squashfs_cmd_direct := exec.Command("./binary")
	err = squashfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
