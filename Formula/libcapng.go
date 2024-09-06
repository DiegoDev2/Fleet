package main

// LibcapNg - Library for Linux that makes using posix capabilities easy
// Homepage: https://people.redhat.com/sgrubb/libcap-ng

import (
	"fmt"
	
	"os/exec"
)

func installLibcapNg() {
	// Método 1: Descargar y extraer .tar.gz
	libcapng_tar_url := "https://people.redhat.com/sgrubb/libcap-ng/libcap-ng-0.8.5.tar.gz"
	libcapng_cmd_tar := exec.Command("curl", "-L", libcapng_tar_url, "-o", "package.tar.gz")
	err := libcapng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcapng_zip_url := "https://people.redhat.com/sgrubb/libcap-ng/libcap-ng-0.8.5.zip"
	libcapng_cmd_zip := exec.Command("curl", "-L", libcapng_zip_url, "-o", "package.zip")
	err = libcapng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcapng_bin_url := "https://people.redhat.com/sgrubb/libcap-ng/libcap-ng-0.8.5.bin"
	libcapng_cmd_bin := exec.Command("curl", "-L", libcapng_bin_url, "-o", "binary.bin")
	err = libcapng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcapng_src_url := "https://people.redhat.com/sgrubb/libcap-ng/libcap-ng-0.8.5.src.tar.gz"
	libcapng_cmd_src := exec.Command("curl", "-L", libcapng_src_url, "-o", "source.tar.gz")
	err = libcapng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcapng_cmd_direct := exec.Command("./binary")
	err = libcapng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: m4")
	exec.Command("latte", "install", "m4").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
}
