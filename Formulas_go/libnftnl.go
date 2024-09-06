package main

// Libnftnl - Netfilter library providing interface to the nf_tables subsystem
// Homepage: https://netfilter.org/projects/libnftnl/

import (
	"fmt"
	
	"os/exec"
)

func installLibnftnl() {
	// Método 1: Descargar y extraer .tar.gz
	libnftnl_tar_url := "https://www.netfilter.org/pub/libnftnl/libnftnl-1.2.7.tar.xz"
	libnftnl_cmd_tar := exec.Command("curl", "-L", libnftnl_tar_url, "-o", "package.tar.gz")
	err := libnftnl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnftnl_zip_url := "https://www.netfilter.org/pub/libnftnl/libnftnl-1.2.7.tar.xz"
	libnftnl_cmd_zip := exec.Command("curl", "-L", libnftnl_zip_url, "-o", "package.zip")
	err = libnftnl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnftnl_bin_url := "https://www.netfilter.org/pub/libnftnl/libnftnl-1.2.7.tar.xz"
	libnftnl_cmd_bin := exec.Command("curl", "-L", libnftnl_bin_url, "-o", "binary.bin")
	err = libnftnl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnftnl_src_url := "https://www.netfilter.org/pub/libnftnl/libnftnl-1.2.7.tar.xz"
	libnftnl_cmd_src := exec.Command("curl", "-L", libnftnl_src_url, "-o", "source.tar.gz")
	err = libnftnl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnftnl_cmd_direct := exec.Command("./binary")
	err = libnftnl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libmnl")
exec.Command("latte", "install", "libmnl")
}
