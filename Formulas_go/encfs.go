package main

// Encfs - Encrypted pass-through FUSE file system
// Homepage: https://vgough.github.io/encfs/

import (
	"fmt"
	
	"os/exec"
)

func installEncfs() {
	// Método 1: Descargar y extraer .tar.gz
	encfs_tar_url := "https://github.com/vgough/encfs/archive/refs/tags/v1.9.5.tar.gz"
	encfs_cmd_tar := exec.Command("curl", "-L", encfs_tar_url, "-o", "package.tar.gz")
	err := encfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	encfs_zip_url := "https://github.com/vgough/encfs/archive/refs/tags/v1.9.5.zip"
	encfs_cmd_zip := exec.Command("curl", "-L", encfs_zip_url, "-o", "package.zip")
	err = encfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	encfs_bin_url := "https://github.com/vgough/encfs/archive/refs/tags/v1.9.5.bin"
	encfs_cmd_bin := exec.Command("curl", "-L", encfs_bin_url, "-o", "binary.bin")
	err = encfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	encfs_src_url := "https://github.com/vgough/encfs/archive/refs/tags/v1.9.5.src.tar.gz"
	encfs_cmd_src := exec.Command("curl", "-L", encfs_src_url, "-o", "source.tar.gz")
	err = encfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	encfs_cmd_direct := exec.Command("./binary")
	err = encfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: tinyxml2")
exec.Command("latte", "install", "tinyxml2")
}
