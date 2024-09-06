package main

// Ext2fuse - Compact implementation of ext2 file system using FUSE
// Homepage: https://sourceforge.net/projects/ext2fuse/

import (
	"fmt"
	
	"os/exec"
)

func installExt2fuse() {
	// Método 1: Descargar y extraer .tar.gz
	ext2fuse_tar_url := "https://downloads.sourceforge.net/project/ext2fuse/ext2fuse/0.8.1/ext2fuse-src-0.8.1.tar.gz"
	ext2fuse_cmd_tar := exec.Command("curl", "-L", ext2fuse_tar_url, "-o", "package.tar.gz")
	err := ext2fuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ext2fuse_zip_url := "https://downloads.sourceforge.net/project/ext2fuse/ext2fuse/0.8.1/ext2fuse-src-0.8.1.zip"
	ext2fuse_cmd_zip := exec.Command("curl", "-L", ext2fuse_zip_url, "-o", "package.zip")
	err = ext2fuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ext2fuse_bin_url := "https://downloads.sourceforge.net/project/ext2fuse/ext2fuse/0.8.1/ext2fuse-src-0.8.1.bin"
	ext2fuse_cmd_bin := exec.Command("curl", "-L", ext2fuse_bin_url, "-o", "binary.bin")
	err = ext2fuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ext2fuse_src_url := "https://downloads.sourceforge.net/project/ext2fuse/ext2fuse/0.8.1/ext2fuse-src-0.8.1.src.tar.gz"
	ext2fuse_cmd_src := exec.Command("curl", "-L", ext2fuse_src_url, "-o", "source.tar.gz")
	err = ext2fuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ext2fuse_cmd_direct := exec.Command("./binary")
	err = ext2fuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: e2fsprogs")
exec.Command("latte", "install", "e2fsprogs")
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
}
