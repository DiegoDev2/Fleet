package main

// Libtrace - Library for trace processing supporting multiple inputs
// Homepage: https://github.com/LibtraceTeam/libtrace

import (
	"fmt"
	
	"os/exec"
)

func installLibtrace() {
	// Método 1: Descargar y extraer .tar.gz
	libtrace_tar_url := "https://github.com/LibtraceTeam/libtrace/archive/refs/tags/4.0.26-1.tar.gz"
	libtrace_cmd_tar := exec.Command("curl", "-L", libtrace_tar_url, "-o", "package.tar.gz")
	err := libtrace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtrace_zip_url := "https://github.com/LibtraceTeam/libtrace/archive/refs/tags/4.0.26-1.zip"
	libtrace_cmd_zip := exec.Command("curl", "-L", libtrace_zip_url, "-o", "package.zip")
	err = libtrace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtrace_bin_url := "https://github.com/LibtraceTeam/libtrace/archive/refs/tags/4.0.26-1.bin"
	libtrace_cmd_bin := exec.Command("curl", "-L", libtrace_bin_url, "-o", "binary.bin")
	err = libtrace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtrace_src_url := "https://github.com/LibtraceTeam/libtrace/archive/refs/tags/4.0.26-1.src.tar.gz"
	libtrace_cmd_src := exec.Command("curl", "-L", libtrace_src_url, "-o", "source.tar.gz")
	err = libtrace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtrace_cmd_direct := exec.Command("./binary")
	err = libtrace_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: wandio")
	exec.Command("latte", "install", "wandio").Run()
}
