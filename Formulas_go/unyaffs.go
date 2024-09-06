package main

// Unyaffs - Extract files from a YAFFS2 filesystem image
// Homepage: https://packages.debian.org/sid/unyaffs

import (
	"fmt"
	
	"os/exec"
)

func installUnyaffs() {
	// Método 1: Descargar y extraer .tar.gz
	unyaffs_tar_url := "https://deb.debian.org/debian/pool/main/u/unyaffs/unyaffs_0.9.7.orig.tar.gz"
	unyaffs_cmd_tar := exec.Command("curl", "-L", unyaffs_tar_url, "-o", "package.tar.gz")
	err := unyaffs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unyaffs_zip_url := "https://deb.debian.org/debian/pool/main/u/unyaffs/unyaffs_0.9.7.orig.zip"
	unyaffs_cmd_zip := exec.Command("curl", "-L", unyaffs_zip_url, "-o", "package.zip")
	err = unyaffs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unyaffs_bin_url := "https://deb.debian.org/debian/pool/main/u/unyaffs/unyaffs_0.9.7.orig.bin"
	unyaffs_cmd_bin := exec.Command("curl", "-L", unyaffs_bin_url, "-o", "binary.bin")
	err = unyaffs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unyaffs_src_url := "https://deb.debian.org/debian/pool/main/u/unyaffs/unyaffs_0.9.7.orig.src.tar.gz"
	unyaffs_cmd_src := exec.Command("curl", "-L", unyaffs_src_url, "-o", "source.tar.gz")
	err = unyaffs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unyaffs_cmd_direct := exec.Command("./binary")
	err = unyaffs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
