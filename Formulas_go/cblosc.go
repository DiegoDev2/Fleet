package main

// CBlosc - Blocking, shuffling and loss-less compression library
// Homepage: https://www.blosc.org/

import (
	"fmt"
	
	"os/exec"
)

func installCBlosc() {
	// Método 1: Descargar y extraer .tar.gz
	cblosc_tar_url := "https://github.com/Blosc/c-blosc/archive/refs/tags/v1.21.6.tar.gz"
	cblosc_cmd_tar := exec.Command("curl", "-L", cblosc_tar_url, "-o", "package.tar.gz")
	err := cblosc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cblosc_zip_url := "https://github.com/Blosc/c-blosc/archive/refs/tags/v1.21.6.zip"
	cblosc_cmd_zip := exec.Command("curl", "-L", cblosc_zip_url, "-o", "package.zip")
	err = cblosc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cblosc_bin_url := "https://github.com/Blosc/c-blosc/archive/refs/tags/v1.21.6.bin"
	cblosc_cmd_bin := exec.Command("curl", "-L", cblosc_bin_url, "-o", "binary.bin")
	err = cblosc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cblosc_src_url := "https://github.com/Blosc/c-blosc/archive/refs/tags/v1.21.6.src.tar.gz"
	cblosc_cmd_src := exec.Command("curl", "-L", cblosc_src_url, "-o", "source.tar.gz")
	err = cblosc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cblosc_cmd_direct := exec.Command("./binary")
	err = cblosc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
