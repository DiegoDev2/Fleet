package main

// Xlsxio - C library for reading values from and writing values to .xlsx files
// Homepage: https://github.com/brechtsanders/xlsxio

import (
	"fmt"
	
	"os/exec"
)

func installXlsxio() {
	// Método 1: Descargar y extraer .tar.gz
	xlsxio_tar_url := "https://github.com/brechtsanders/xlsxio/archive/refs/tags/0.2.35.tar.gz"
	xlsxio_cmd_tar := exec.Command("curl", "-L", xlsxio_tar_url, "-o", "package.tar.gz")
	err := xlsxio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xlsxio_zip_url := "https://github.com/brechtsanders/xlsxio/archive/refs/tags/0.2.35.zip"
	xlsxio_cmd_zip := exec.Command("curl", "-L", xlsxio_zip_url, "-o", "package.zip")
	err = xlsxio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xlsxio_bin_url := "https://github.com/brechtsanders/xlsxio/archive/refs/tags/0.2.35.bin"
	xlsxio_cmd_bin := exec.Command("curl", "-L", xlsxio_bin_url, "-o", "binary.bin")
	err = xlsxio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xlsxio_src_url := "https://github.com/brechtsanders/xlsxio/archive/refs/tags/0.2.35.src.tar.gz"
	xlsxio_cmd_src := exec.Command("curl", "-L", xlsxio_src_url, "-o", "source.tar.gz")
	err = xlsxio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xlsxio_cmd_direct := exec.Command("./binary")
	err = xlsxio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
}
