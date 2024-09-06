package main

// Libnfs - C client library for NFS
// Homepage: https://github.com/sahlberg/libnfs

import (
	"fmt"
	
	"os/exec"
)

func installLibnfs() {
	// Método 1: Descargar y extraer .tar.gz
	libnfs_tar_url := "https://github.com/sahlberg/libnfs/archive/refs/tags/libnfs-5.0.3.tar.gz"
	libnfs_cmd_tar := exec.Command("curl", "-L", libnfs_tar_url, "-o", "package.tar.gz")
	err := libnfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnfs_zip_url := "https://github.com/sahlberg/libnfs/archive/refs/tags/libnfs-5.0.3.zip"
	libnfs_cmd_zip := exec.Command("curl", "-L", libnfs_zip_url, "-o", "package.zip")
	err = libnfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnfs_bin_url := "https://github.com/sahlberg/libnfs/archive/refs/tags/libnfs-5.0.3.bin"
	libnfs_cmd_bin := exec.Command("curl", "-L", libnfs_bin_url, "-o", "binary.bin")
	err = libnfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnfs_src_url := "https://github.com/sahlberg/libnfs/archive/refs/tags/libnfs-5.0.3.src.tar.gz"
	libnfs_cmd_src := exec.Command("curl", "-L", libnfs_src_url, "-o", "source.tar.gz")
	err = libnfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnfs_cmd_direct := exec.Command("./binary")
	err = libnfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
