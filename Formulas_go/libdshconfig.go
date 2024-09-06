package main

// Libdshconfig - Distributed shell library
// Homepage: https://www.netfort.gr.jp/~dancer/software/dsh.html.en

import (
	"fmt"
	
	"os/exec"
)

func installLibdshconfig() {
	// Método 1: Descargar y extraer .tar.gz
	libdshconfig_tar_url := "https://www.netfort.gr.jp/~dancer/software/downloads/libdshconfig-0.20.13.tar.gz"
	libdshconfig_cmd_tar := exec.Command("curl", "-L", libdshconfig_tar_url, "-o", "package.tar.gz")
	err := libdshconfig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdshconfig_zip_url := "https://www.netfort.gr.jp/~dancer/software/downloads/libdshconfig-0.20.13.zip"
	libdshconfig_cmd_zip := exec.Command("curl", "-L", libdshconfig_zip_url, "-o", "package.zip")
	err = libdshconfig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdshconfig_bin_url := "https://www.netfort.gr.jp/~dancer/software/downloads/libdshconfig-0.20.13.bin"
	libdshconfig_cmd_bin := exec.Command("curl", "-L", libdshconfig_bin_url, "-o", "binary.bin")
	err = libdshconfig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdshconfig_src_url := "https://www.netfort.gr.jp/~dancer/software/downloads/libdshconfig-0.20.13.src.tar.gz"
	libdshconfig_cmd_src := exec.Command("curl", "-L", libdshconfig_src_url, "-o", "source.tar.gz")
	err = libdshconfig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdshconfig_cmd_direct := exec.Command("./binary")
	err = libdshconfig_cmd_direct.Run()
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
}
