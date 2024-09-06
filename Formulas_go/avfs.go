package main

// Avfs - Virtual file system that facilitates looking inside archives
// Homepage: https://avf.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installAvfs() {
	// Método 1: Descargar y extraer .tar.gz
	avfs_tar_url := "https://downloads.sourceforge.net/project/avf/avfs/1.1.5/avfs-1.1.5.tar.bz2"
	avfs_cmd_tar := exec.Command("curl", "-L", avfs_tar_url, "-o", "package.tar.gz")
	err := avfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	avfs_zip_url := "https://downloads.sourceforge.net/project/avf/avfs/1.1.5/avfs-1.1.5.tar.bz2"
	avfs_cmd_zip := exec.Command("curl", "-L", avfs_zip_url, "-o", "package.zip")
	err = avfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	avfs_bin_url := "https://downloads.sourceforge.net/project/avf/avfs/1.1.5/avfs-1.1.5.tar.bz2"
	avfs_cmd_bin := exec.Command("curl", "-L", avfs_bin_url, "-o", "binary.bin")
	err = avfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	avfs_src_url := "https://downloads.sourceforge.net/project/avf/avfs/1.1.5/avfs-1.1.5.tar.bz2"
	avfs_cmd_src := exec.Command("curl", "-L", avfs_src_url, "-o", "source.tar.gz")
	err = avfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	avfs_cmd_direct := exec.Command("./binary")
	err = avfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: bzip2")
exec.Command("latte", "install", "bzip2")
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
}
