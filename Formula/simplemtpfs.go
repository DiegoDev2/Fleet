package main

// SimpleMtpfs - Simple MTP fuse filesystem driver
// Homepage: https://github.com/phatina/simple-mtpfs

import (
	"fmt"
	
	"os/exec"
)

func installSimpleMtpfs() {
	// Método 1: Descargar y extraer .tar.gz
	simplemtpfs_tar_url := "https://github.com/phatina/simple-mtpfs/archive/refs/tags/v0.4.0.tar.gz"
	simplemtpfs_cmd_tar := exec.Command("curl", "-L", simplemtpfs_tar_url, "-o", "package.tar.gz")
	err := simplemtpfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simplemtpfs_zip_url := "https://github.com/phatina/simple-mtpfs/archive/refs/tags/v0.4.0.zip"
	simplemtpfs_cmd_zip := exec.Command("curl", "-L", simplemtpfs_zip_url, "-o", "package.zip")
	err = simplemtpfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simplemtpfs_bin_url := "https://github.com/phatina/simple-mtpfs/archive/refs/tags/v0.4.0.bin"
	simplemtpfs_cmd_bin := exec.Command("curl", "-L", simplemtpfs_bin_url, "-o", "binary.bin")
	err = simplemtpfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simplemtpfs_src_url := "https://github.com/phatina/simple-mtpfs/archive/refs/tags/v0.4.0.src.tar.gz"
	simplemtpfs_cmd_src := exec.Command("curl", "-L", simplemtpfs_src_url, "-o", "source.tar.gz")
	err = simplemtpfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simplemtpfs_cmd_direct := exec.Command("./binary")
	err = simplemtpfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: autoconf-archive")
	exec.Command("latte", "install", "autoconf-archive").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
	fmt.Println("Instalando dependencia: libmtp")
	exec.Command("latte", "install", "libmtp").Run()
}
