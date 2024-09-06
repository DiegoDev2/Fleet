package main

// Libmnl - Minimalistic user-space library oriented to Netlink developers
// Homepage: https://www.netfilter.org/projects/libmnl

import (
	"fmt"
	
	"os/exec"
)

func installLibmnl() {
	// Método 1: Descargar y extraer .tar.gz
	libmnl_tar_url := "https://www.netfilter.org/projects/libmnl/files/libmnl-1.0.5.tar.bz2"
	libmnl_cmd_tar := exec.Command("curl", "-L", libmnl_tar_url, "-o", "package.tar.gz")
	err := libmnl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmnl_zip_url := "https://www.netfilter.org/projects/libmnl/files/libmnl-1.0.5.tar.bz2"
	libmnl_cmd_zip := exec.Command("curl", "-L", libmnl_zip_url, "-o", "package.zip")
	err = libmnl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmnl_bin_url := "https://www.netfilter.org/projects/libmnl/files/libmnl-1.0.5.tar.bz2"
	libmnl_cmd_bin := exec.Command("curl", "-L", libmnl_bin_url, "-o", "binary.bin")
	err = libmnl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmnl_src_url := "https://www.netfilter.org/projects/libmnl/files/libmnl-1.0.5.tar.bz2"
	libmnl_cmd_src := exec.Command("curl", "-L", libmnl_src_url, "-o", "source.tar.gz")
	err = libmnl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmnl_cmd_direct := exec.Command("./binary")
	err = libmnl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
