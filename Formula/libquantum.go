package main

// Libquantum - C library for the simulation of quantum mechanics
// Homepage: http://www.libquantum.de/

import (
	"fmt"
	
	"os/exec"
)

func installLibquantum() {
	// Método 1: Descargar y extraer .tar.gz
	libquantum_tar_url := "http://www.libquantum.de/files/libquantum-1.0.0.tar.gz"
	libquantum_cmd_tar := exec.Command("curl", "-L", libquantum_tar_url, "-o", "package.tar.gz")
	err := libquantum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libquantum_zip_url := "http://www.libquantum.de/files/libquantum-1.0.0.zip"
	libquantum_cmd_zip := exec.Command("curl", "-L", libquantum_zip_url, "-o", "package.zip")
	err = libquantum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libquantum_bin_url := "http://www.libquantum.de/files/libquantum-1.0.0.bin"
	libquantum_cmd_bin := exec.Command("curl", "-L", libquantum_bin_url, "-o", "binary.bin")
	err = libquantum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libquantum_src_url := "http://www.libquantum.de/files/libquantum-1.0.0.src.tar.gz"
	libquantum_cmd_src := exec.Command("curl", "-L", libquantum_src_url, "-o", "source.tar.gz")
	err = libquantum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libquantum_cmd_direct := exec.Command("./binary")
	err = libquantum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
