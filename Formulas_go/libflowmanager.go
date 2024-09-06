package main

// Libflowmanager - Flow-based measurement tasks with packet-based inputs
// Homepage: https://github.com/LibtraceTeam/libflowmanager

import (
	"fmt"
	
	"os/exec"
)

func installLibflowmanager() {
	// Método 1: Descargar y extraer .tar.gz
	libflowmanager_tar_url := "https://github.com/LibtraceTeam/libflowmanager/archive/refs/tags/v3.0.0.tar.gz"
	libflowmanager_cmd_tar := exec.Command("curl", "-L", libflowmanager_tar_url, "-o", "package.tar.gz")
	err := libflowmanager_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libflowmanager_zip_url := "https://github.com/LibtraceTeam/libflowmanager/archive/refs/tags/v3.0.0.zip"
	libflowmanager_cmd_zip := exec.Command("curl", "-L", libflowmanager_zip_url, "-o", "package.zip")
	err = libflowmanager_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libflowmanager_bin_url := "https://github.com/LibtraceTeam/libflowmanager/archive/refs/tags/v3.0.0.bin"
	libflowmanager_cmd_bin := exec.Command("curl", "-L", libflowmanager_bin_url, "-o", "binary.bin")
	err = libflowmanager_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libflowmanager_src_url := "https://github.com/LibtraceTeam/libflowmanager/archive/refs/tags/v3.0.0.src.tar.gz"
	libflowmanager_cmd_src := exec.Command("curl", "-L", libflowmanager_src_url, "-o", "source.tar.gz")
	err = libflowmanager_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libflowmanager_cmd_direct := exec.Command("./binary")
	err = libflowmanager_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libtrace")
exec.Command("latte", "install", "libtrace")
}
