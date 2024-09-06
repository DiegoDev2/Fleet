package main

// Libcerf - Numeric library for complex error functions
// Homepage: https://jugit.fz-juelich.de/mlz/libcerf

import (
	"fmt"
	
	"os/exec"
)

func installLibcerf() {
	// Método 1: Descargar y extraer .tar.gz
	libcerf_tar_url := "https://jugit.fz-juelich.de/mlz/libcerf/-/archive/v2.4/libcerf-v2.4.tar.gz"
	libcerf_cmd_tar := exec.Command("curl", "-L", libcerf_tar_url, "-o", "package.tar.gz")
	err := libcerf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcerf_zip_url := "https://jugit.fz-juelich.de/mlz/libcerf/-/archive/v2.4/libcerf-v2.4.zip"
	libcerf_cmd_zip := exec.Command("curl", "-L", libcerf_zip_url, "-o", "package.zip")
	err = libcerf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcerf_bin_url := "https://jugit.fz-juelich.de/mlz/libcerf/-/archive/v2.4/libcerf-v2.4.bin"
	libcerf_cmd_bin := exec.Command("curl", "-L", libcerf_bin_url, "-o", "binary.bin")
	err = libcerf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcerf_src_url := "https://jugit.fz-juelich.de/mlz/libcerf/-/archive/v2.4/libcerf-v2.4.src.tar.gz"
	libcerf_cmd_src := exec.Command("curl", "-L", libcerf_src_url, "-o", "source.tar.gz")
	err = libcerf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcerf_cmd_direct := exec.Command("./binary")
	err = libcerf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
