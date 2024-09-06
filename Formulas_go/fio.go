package main

// Fio - I/O benchmark and stress test
// Homepage: https://github.com/axboe/fio

import (
	"fmt"
	
	"os/exec"
)

func installFio() {
	// Método 1: Descargar y extraer .tar.gz
	fio_tar_url := "https://github.com/axboe/fio/archive/refs/tags/fio-3.37.tar.gz"
	fio_cmd_tar := exec.Command("curl", "-L", fio_tar_url, "-o", "package.tar.gz")
	err := fio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fio_zip_url := "https://github.com/axboe/fio/archive/refs/tags/fio-3.37.zip"
	fio_cmd_zip := exec.Command("curl", "-L", fio_zip_url, "-o", "package.zip")
	err = fio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fio_bin_url := "https://github.com/axboe/fio/archive/refs/tags/fio-3.37.bin"
	fio_cmd_bin := exec.Command("curl", "-L", fio_bin_url, "-o", "binary.bin")
	err = fio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fio_src_url := "https://github.com/axboe/fio/archive/refs/tags/fio-3.37.src.tar.gz"
	fio_cmd_src := exec.Command("curl", "-L", fio_src_url, "-o", "source.tar.gz")
	err = fio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fio_cmd_direct := exec.Command("./binary")
	err = fio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
