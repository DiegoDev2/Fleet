package main

// Btfs - BitTorrent filesystem based on FUSE
// Homepage: https://github.com/johang/btfs

import (
	"fmt"
	
	"os/exec"
)

func installBtfs() {
	// Método 1: Descargar y extraer .tar.gz
	btfs_tar_url := "https://github.com/johang/btfs/archive/refs/tags/v2.24.tar.gz"
	btfs_cmd_tar := exec.Command("curl", "-L", btfs_tar_url, "-o", "package.tar.gz")
	err := btfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	btfs_zip_url := "https://github.com/johang/btfs/archive/refs/tags/v2.24.zip"
	btfs_cmd_zip := exec.Command("curl", "-L", btfs_zip_url, "-o", "package.zip")
	err = btfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	btfs_bin_url := "https://github.com/johang/btfs/archive/refs/tags/v2.24.bin"
	btfs_cmd_bin := exec.Command("curl", "-L", btfs_bin_url, "-o", "binary.bin")
	err = btfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	btfs_src_url := "https://github.com/johang/btfs/archive/refs/tags/v2.24.src.tar.gz"
	btfs_cmd_src := exec.Command("curl", "-L", btfs_src_url, "-o", "source.tar.gz")
	err = btfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	btfs_cmd_direct := exec.Command("./binary")
	err = btfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
	fmt.Println("Instalando dependencia: libtorrent-rasterbar")
	exec.Command("latte", "install", "libtorrent-rasterbar").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
