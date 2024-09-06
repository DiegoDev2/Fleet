package main

// ErofsUtils - Utilities for Enhanced Read-Only File System
// Homepage: https://git.kernel.org/pub/scm/linux/kernel/git/xiang/erofs-utils.git

import (
	"fmt"
	
	"os/exec"
)

func installErofsUtils() {
	// Método 1: Descargar y extraer .tar.gz
	erofsutils_tar_url := "https://git.kernel.org/pub/scm/linux/kernel/git/xiang/erofs-utils.git/snapshot/erofs-utils-1.8.1.tar.gz"
	erofsutils_cmd_tar := exec.Command("curl", "-L", erofsutils_tar_url, "-o", "package.tar.gz")
	err := erofsutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	erofsutils_zip_url := "https://git.kernel.org/pub/scm/linux/kernel/git/xiang/erofs-utils.git/snapshot/erofs-utils-1.8.1.zip"
	erofsutils_cmd_zip := exec.Command("curl", "-L", erofsutils_zip_url, "-o", "package.zip")
	err = erofsutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	erofsutils_bin_url := "https://git.kernel.org/pub/scm/linux/kernel/git/xiang/erofs-utils.git/snapshot/erofs-utils-1.8.1.bin"
	erofsutils_cmd_bin := exec.Command("curl", "-L", erofsutils_bin_url, "-o", "binary.bin")
	err = erofsutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	erofsutils_src_url := "https://git.kernel.org/pub/scm/linux/kernel/git/xiang/erofs-utils.git/snapshot/erofs-utils-1.8.1.src.tar.gz"
	erofsutils_cmd_src := exec.Command("curl", "-L", erofsutils_src_url, "-o", "source.tar.gz")
	err = erofsutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	erofsutils_cmd_direct := exec.Command("./binary")
	err = erofsutils_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
}
