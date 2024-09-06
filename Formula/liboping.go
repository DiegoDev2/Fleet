package main

// Liboping - C library to generate ICMP echo requests
// Homepage: https://noping.cc/

import (
	"fmt"
	
	"os/exec"
)

func installLiboping() {
	// Método 1: Descargar y extraer .tar.gz
	liboping_tar_url := "https://noping.cc/files/liboping-1.10.0.tar.bz2"
	liboping_cmd_tar := exec.Command("curl", "-L", liboping_tar_url, "-o", "package.tar.gz")
	err := liboping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liboping_zip_url := "https://noping.cc/files/liboping-1.10.0.tar.bz2"
	liboping_cmd_zip := exec.Command("curl", "-L", liboping_zip_url, "-o", "package.zip")
	err = liboping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liboping_bin_url := "https://noping.cc/files/liboping-1.10.0.tar.bz2"
	liboping_cmd_bin := exec.Command("curl", "-L", liboping_bin_url, "-o", "binary.bin")
	err = liboping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liboping_src_url := "https://noping.cc/files/liboping-1.10.0.tar.bz2"
	liboping_cmd_src := exec.Command("curl", "-L", liboping_src_url, "-o", "source.tar.gz")
	err = liboping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liboping_cmd_direct := exec.Command("./binary")
	err = liboping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
