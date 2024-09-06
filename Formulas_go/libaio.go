package main

// Libaio - Linux-native asynchronous I/O access library
// Homepage: https://pagure.io/libaio

import (
	"fmt"
	
	"os/exec"
)

func installLibaio() {
	// Método 1: Descargar y extraer .tar.gz
	libaio_tar_url := "https://pagure.io/libaio/archive/libaio-0.3.113/libaio-libaio-0.3.113.tar.gz"
	libaio_cmd_tar := exec.Command("curl", "-L", libaio_tar_url, "-o", "package.tar.gz")
	err := libaio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libaio_zip_url := "https://pagure.io/libaio/archive/libaio-0.3.113/libaio-libaio-0.3.113.zip"
	libaio_cmd_zip := exec.Command("curl", "-L", libaio_zip_url, "-o", "package.zip")
	err = libaio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libaio_bin_url := "https://pagure.io/libaio/archive/libaio-0.3.113/libaio-libaio-0.3.113.bin"
	libaio_cmd_bin := exec.Command("curl", "-L", libaio_bin_url, "-o", "binary.bin")
	err = libaio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libaio_src_url := "https://pagure.io/libaio/archive/libaio-0.3.113/libaio-libaio-0.3.113.src.tar.gz"
	libaio_cmd_src := exec.Command("curl", "-L", libaio_src_url, "-o", "source.tar.gz")
	err = libaio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libaio_cmd_direct := exec.Command("./binary")
	err = libaio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
