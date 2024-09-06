package main

// Libfaketime - Report faked system time to programs
// Homepage: https://github.com/wolfcw/libfaketime

import (
	"fmt"
	
	"os/exec"
)

func installLibfaketime() {
	// Método 1: Descargar y extraer .tar.gz
	libfaketime_tar_url := "https://github.com/wolfcw/libfaketime/archive/refs/tags/v0.9.10.tar.gz"
	libfaketime_cmd_tar := exec.Command("curl", "-L", libfaketime_tar_url, "-o", "package.tar.gz")
	err := libfaketime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfaketime_zip_url := "https://github.com/wolfcw/libfaketime/archive/refs/tags/v0.9.10.zip"
	libfaketime_cmd_zip := exec.Command("curl", "-L", libfaketime_zip_url, "-o", "package.zip")
	err = libfaketime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfaketime_bin_url := "https://github.com/wolfcw/libfaketime/archive/refs/tags/v0.9.10.bin"
	libfaketime_cmd_bin := exec.Command("curl", "-L", libfaketime_bin_url, "-o", "binary.bin")
	err = libfaketime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfaketime_src_url := "https://github.com/wolfcw/libfaketime/archive/refs/tags/v0.9.10.src.tar.gz"
	libfaketime_cmd_src := exec.Command("curl", "-L", libfaketime_src_url, "-o", "source.tar.gz")
	err = libfaketime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfaketime_cmd_direct := exec.Command("./binary")
	err = libfaketime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
}
