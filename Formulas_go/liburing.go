package main

// Liburing - Helpers to setup and teardown io_uring instances
// Homepage: https://github.com/axboe/liburing

import (
	"fmt"
	
	"os/exec"
)

func installLiburing() {
	// Método 1: Descargar y extraer .tar.gz
	liburing_tar_url := "https://github.com/axboe/liburing/archive/refs/tags/liburing-2.6.tar.gz"
	liburing_cmd_tar := exec.Command("curl", "-L", liburing_tar_url, "-o", "package.tar.gz")
	err := liburing_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liburing_zip_url := "https://github.com/axboe/liburing/archive/refs/tags/liburing-2.6.zip"
	liburing_cmd_zip := exec.Command("curl", "-L", liburing_zip_url, "-o", "package.zip")
	err = liburing_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liburing_bin_url := "https://github.com/axboe/liburing/archive/refs/tags/liburing-2.6.bin"
	liburing_cmd_bin := exec.Command("curl", "-L", liburing_bin_url, "-o", "binary.bin")
	err = liburing_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liburing_src_url := "https://github.com/axboe/liburing/archive/refs/tags/liburing-2.6.src.tar.gz"
	liburing_cmd_src := exec.Command("curl", "-L", liburing_src_url, "-o", "source.tar.gz")
	err = liburing_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liburing_cmd_direct := exec.Command("./binary")
	err = liburing_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
