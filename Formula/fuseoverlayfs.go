package main

// FuseOverlayfs - FUSE implementation for overlayfs
// Homepage: https://github.com/containers/fuse-overlayfs

import (
	"fmt"
	
	"os/exec"
)

func installFuseOverlayfs() {
	// Método 1: Descargar y extraer .tar.gz
	fuseoverlayfs_tar_url := "https://github.com/containers/fuse-overlayfs/archive/refs/tags/v1.14.tar.gz"
	fuseoverlayfs_cmd_tar := exec.Command("curl", "-L", fuseoverlayfs_tar_url, "-o", "package.tar.gz")
	err := fuseoverlayfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fuseoverlayfs_zip_url := "https://github.com/containers/fuse-overlayfs/archive/refs/tags/v1.14.zip"
	fuseoverlayfs_cmd_zip := exec.Command("curl", "-L", fuseoverlayfs_zip_url, "-o", "package.zip")
	err = fuseoverlayfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fuseoverlayfs_bin_url := "https://github.com/containers/fuse-overlayfs/archive/refs/tags/v1.14.bin"
	fuseoverlayfs_cmd_bin := exec.Command("curl", "-L", fuseoverlayfs_bin_url, "-o", "binary.bin")
	err = fuseoverlayfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fuseoverlayfs_src_url := "https://github.com/containers/fuse-overlayfs/archive/refs/tags/v1.14.src.tar.gz"
	fuseoverlayfs_cmd_src := exec.Command("curl", "-L", fuseoverlayfs_src_url, "-o", "source.tar.gz")
	err = fuseoverlayfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fuseoverlayfs_cmd_direct := exec.Command("./binary")
	err = fuseoverlayfs_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libfuse")
	exec.Command("latte", "install", "libfuse").Run()
}
