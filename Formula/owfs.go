package main

// Owfs - Monitor and control physical environment using Dallas/Maxim 1-wire system
// Homepage: https://owfs.org/

import (
	"fmt"
	
	"os/exec"
)

func installOwfs() {
	// Método 1: Descargar y extraer .tar.gz
	owfs_tar_url := "https://github.com/owfs/owfs/releases/download/v3.2p4/owfs-3.2p4.tar.gz"
	owfs_cmd_tar := exec.Command("curl", "-L", owfs_tar_url, "-o", "package.tar.gz")
	err := owfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	owfs_zip_url := "https://github.com/owfs/owfs/releases/download/v3.2p4/owfs-3.2p4.zip"
	owfs_cmd_zip := exec.Command("curl", "-L", owfs_zip_url, "-o", "package.zip")
	err = owfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	owfs_bin_url := "https://github.com/owfs/owfs/releases/download/v3.2p4/owfs-3.2p4.bin"
	owfs_cmd_bin := exec.Command("curl", "-L", owfs_bin_url, "-o", "binary.bin")
	err = owfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	owfs_src_url := "https://github.com/owfs/owfs/releases/download/v3.2p4/owfs-3.2p4.src.tar.gz"
	owfs_cmd_src := exec.Command("curl", "-L", owfs_src_url, "-o", "source.tar.gz")
	err = owfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	owfs_cmd_direct := exec.Command("./binary")
	err = owfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libftdi")
	exec.Command("latte", "install", "libftdi").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
