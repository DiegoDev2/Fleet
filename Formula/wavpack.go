package main

// Wavpack - Hybrid lossless audio compression
// Homepage: https://www.wavpack.com/

import (
	"fmt"
	
	"os/exec"
)

func installWavpack() {
	// Método 1: Descargar y extraer .tar.gz
	wavpack_tar_url := "https://www.wavpack.com/wavpack-5.7.0.tar.bz2"
	wavpack_cmd_tar := exec.Command("curl", "-L", wavpack_tar_url, "-o", "package.tar.gz")
	err := wavpack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wavpack_zip_url := "https://www.wavpack.com/wavpack-5.7.0.tar.bz2"
	wavpack_cmd_zip := exec.Command("curl", "-L", wavpack_zip_url, "-o", "package.zip")
	err = wavpack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wavpack_bin_url := "https://www.wavpack.com/wavpack-5.7.0.tar.bz2"
	wavpack_cmd_bin := exec.Command("curl", "-L", wavpack_bin_url, "-o", "binary.bin")
	err = wavpack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wavpack_src_url := "https://www.wavpack.com/wavpack-5.7.0.tar.bz2"
	wavpack_cmd_src := exec.Command("curl", "-L", wavpack_src_url, "-o", "source.tar.gz")
	err = wavpack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wavpack_cmd_direct := exec.Command("./binary")
	err = wavpack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
