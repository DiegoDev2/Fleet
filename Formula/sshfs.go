package main

// Sshfs - File system client based on SSH File Transfer Protocol
// Homepage: https://github.com/libfuse/sshfs

import (
	"fmt"
	
	"os/exec"
)

func installSshfs() {
	// Método 1: Descargar y extraer .tar.gz
	sshfs_tar_url := "https://github.com/libfuse/sshfs/archive/refs/tags/sshfs-3.7.3.tar.gz"
	sshfs_cmd_tar := exec.Command("curl", "-L", sshfs_tar_url, "-o", "package.tar.gz")
	err := sshfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshfs_zip_url := "https://github.com/libfuse/sshfs/archive/refs/tags/sshfs-3.7.3.zip"
	sshfs_cmd_zip := exec.Command("curl", "-L", sshfs_zip_url, "-o", "package.zip")
	err = sshfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshfs_bin_url := "https://github.com/libfuse/sshfs/archive/refs/tags/sshfs-3.7.3.bin"
	sshfs_cmd_bin := exec.Command("curl", "-L", sshfs_bin_url, "-o", "binary.bin")
	err = sshfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshfs_src_url := "https://github.com/libfuse/sshfs/archive/refs/tags/sshfs-3.7.3.src.tar.gz"
	sshfs_cmd_src := exec.Command("curl", "-L", sshfs_src_url, "-o", "source.tar.gz")
	err = sshfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshfs_cmd_direct := exec.Command("./binary")
	err = sshfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libfuse")
	exec.Command("latte", "install", "libfuse").Run()
}
