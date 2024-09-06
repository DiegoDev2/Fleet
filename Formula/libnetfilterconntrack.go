package main

// LibnetfilterConntrack - Library providing an API to the in-kernel connection tracking state table
// Homepage: https://www.netfilter.org/projects/libnetfilter_conntrack/

import (
	"fmt"
	
	"os/exec"
)

func installLibnetfilterConntrack() {
	// Método 1: Descargar y extraer .tar.gz
	libnetfilterconntrack_tar_url := "https://www.netfilter.org/pub/libnetfilter_conntrack/libnetfilter_conntrack-1.0.9.tar.bz2"
	libnetfilterconntrack_cmd_tar := exec.Command("curl", "-L", libnetfilterconntrack_tar_url, "-o", "package.tar.gz")
	err := libnetfilterconntrack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnetfilterconntrack_zip_url := "https://www.netfilter.org/pub/libnetfilter_conntrack/libnetfilter_conntrack-1.0.9.tar.bz2"
	libnetfilterconntrack_cmd_zip := exec.Command("curl", "-L", libnetfilterconntrack_zip_url, "-o", "package.zip")
	err = libnetfilterconntrack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnetfilterconntrack_bin_url := "https://www.netfilter.org/pub/libnetfilter_conntrack/libnetfilter_conntrack-1.0.9.tar.bz2"
	libnetfilterconntrack_cmd_bin := exec.Command("curl", "-L", libnetfilterconntrack_bin_url, "-o", "binary.bin")
	err = libnetfilterconntrack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnetfilterconntrack_src_url := "https://www.netfilter.org/pub/libnetfilter_conntrack/libnetfilter_conntrack-1.0.9.tar.bz2"
	libnetfilterconntrack_cmd_src := exec.Command("curl", "-L", libnetfilterconntrack_src_url, "-o", "source.tar.gz")
	err = libnetfilterconntrack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnetfilterconntrack_cmd_direct := exec.Command("./binary")
	err = libnetfilterconntrack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libmnl")
	exec.Command("latte", "install", "libmnl").Run()
	fmt.Println("Instalando dependencia: libnfnetlink")
	exec.Command("latte", "install", "libnfnetlink").Run()
}
