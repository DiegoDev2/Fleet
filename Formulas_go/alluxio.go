package main

// Alluxio - Open Source Memory Speed Virtual Distributed Storage
// Homepage: https://www.alluxio.io/

import (
	"fmt"
	
	"os/exec"
)

func installAlluxio() {
	// Método 1: Descargar y extraer .tar.gz
	alluxio_tar_url := "https://downloads.alluxio.io/downloads/files/2.9.5/alluxio-2.9.5-bin.tar.gz"
	alluxio_cmd_tar := exec.Command("curl", "-L", alluxio_tar_url, "-o", "package.tar.gz")
	err := alluxio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alluxio_zip_url := "https://downloads.alluxio.io/downloads/files/2.9.5/alluxio-2.9.5-bin.zip"
	alluxio_cmd_zip := exec.Command("curl", "-L", alluxio_zip_url, "-o", "package.zip")
	err = alluxio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alluxio_bin_url := "https://downloads.alluxio.io/downloads/files/2.9.5/alluxio-2.9.5-bin.bin"
	alluxio_cmd_bin := exec.Command("curl", "-L", alluxio_bin_url, "-o", "binary.bin")
	err = alluxio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alluxio_src_url := "https://downloads.alluxio.io/downloads/files/2.9.5/alluxio-2.9.5-bin.src.tar.gz"
	alluxio_cmd_src := exec.Command("curl", "-L", alluxio_src_url, "-o", "source.tar.gz")
	err = alluxio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alluxio_cmd_direct := exec.Command("./binary")
	err = alluxio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
