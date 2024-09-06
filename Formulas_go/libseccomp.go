package main

// Libseccomp - Interface to the Linux Kernel's syscall filtering mechanism
// Homepage: https://github.com/seccomp/libseccomp

import (
	"fmt"
	
	"os/exec"
)

func installLibseccomp() {
	// Método 1: Descargar y extraer .tar.gz
	libseccomp_tar_url := "https://github.com/seccomp/libseccomp/releases/download/v2.5.5/libseccomp-2.5.5.tar.gz"
	libseccomp_cmd_tar := exec.Command("curl", "-L", libseccomp_tar_url, "-o", "package.tar.gz")
	err := libseccomp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libseccomp_zip_url := "https://github.com/seccomp/libseccomp/releases/download/v2.5.5/libseccomp-2.5.5.zip"
	libseccomp_cmd_zip := exec.Command("curl", "-L", libseccomp_zip_url, "-o", "package.zip")
	err = libseccomp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libseccomp_bin_url := "https://github.com/seccomp/libseccomp/releases/download/v2.5.5/libseccomp-2.5.5.bin"
	libseccomp_cmd_bin := exec.Command("curl", "-L", libseccomp_bin_url, "-o", "binary.bin")
	err = libseccomp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libseccomp_src_url := "https://github.com/seccomp/libseccomp/releases/download/v2.5.5/libseccomp-2.5.5.src.tar.gz"
	libseccomp_cmd_src := exec.Command("curl", "-L", libseccomp_src_url, "-o", "source.tar.gz")
	err = libseccomp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libseccomp_cmd_direct := exec.Command("./binary")
	err = libseccomp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: gperf")
exec.Command("latte", "install", "gperf")
}
