package main

// Bindfs - FUSE file system for mounting to another location
// Homepage: https://bindfs.org/

import (
	"fmt"
	
	"os/exec"
)

func installBindfs() {
	// Método 1: Descargar y extraer .tar.gz
	bindfs_tar_url := "https://bindfs.org/downloads/bindfs-1.17.7.tar.gz"
	bindfs_cmd_tar := exec.Command("curl", "-L", bindfs_tar_url, "-o", "package.tar.gz")
	err := bindfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bindfs_zip_url := "https://bindfs.org/downloads/bindfs-1.17.7.zip"
	bindfs_cmd_zip := exec.Command("curl", "-L", bindfs_zip_url, "-o", "package.zip")
	err = bindfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bindfs_bin_url := "https://bindfs.org/downloads/bindfs-1.17.7.bin"
	bindfs_cmd_bin := exec.Command("curl", "-L", bindfs_bin_url, "-o", "binary.bin")
	err = bindfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bindfs_src_url := "https://bindfs.org/downloads/bindfs-1.17.7.src.tar.gz"
	bindfs_cmd_src := exec.Command("curl", "-L", bindfs_src_url, "-o", "source.tar.gz")
	err = bindfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bindfs_cmd_direct := exec.Command("./binary")
	err = bindfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libfuse")
	exec.Command("latte", "install", "libfuse").Run()
}
