package main

// Libprotoident - Performs application layer protocol identification for flows
// Homepage: https://github.com/LibtraceTeam/libprotoident

import (
	"fmt"
	
	"os/exec"
)

func installLibprotoident() {
	// Método 1: Descargar y extraer .tar.gz
	libprotoident_tar_url := "https://github.com/LibtraceTeam/libprotoident/archive/refs/tags/2.0.15-2.tar.gz"
	libprotoident_cmd_tar := exec.Command("curl", "-L", libprotoident_tar_url, "-o", "package.tar.gz")
	err := libprotoident_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libprotoident_zip_url := "https://github.com/LibtraceTeam/libprotoident/archive/refs/tags/2.0.15-2.zip"
	libprotoident_cmd_zip := exec.Command("curl", "-L", libprotoident_zip_url, "-o", "package.zip")
	err = libprotoident_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libprotoident_bin_url := "https://github.com/LibtraceTeam/libprotoident/archive/refs/tags/2.0.15-2.bin"
	libprotoident_cmd_bin := exec.Command("curl", "-L", libprotoident_bin_url, "-o", "binary.bin")
	err = libprotoident_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libprotoident_src_url := "https://github.com/LibtraceTeam/libprotoident/archive/refs/tags/2.0.15-2.src.tar.gz"
	libprotoident_cmd_src := exec.Command("curl", "-L", libprotoident_src_url, "-o", "source.tar.gz")
	err = libprotoident_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libprotoident_cmd_direct := exec.Command("./binary")
	err = libprotoident_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libflowmanager")
	exec.Command("latte", "install", "libflowmanager").Run()
	fmt.Println("Instalando dependencia: libtrace")
	exec.Command("latte", "install", "libtrace").Run()
}
