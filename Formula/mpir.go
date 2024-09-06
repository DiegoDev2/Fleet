package main

// Mpir - Multiple Precision Integers and Rationals (fork of GMP)
// Homepage: https://web.archive.org/web/20221207200514/https://mpir.org/

import (
	"fmt"
	
	"os/exec"
)

func installMpir() {
	// Método 1: Descargar y extraer .tar.gz
	mpir_tar_url := "https://web.archive.org/web/20220224004857/https://mpir.org/mpir-3.0.0.tar.bz2"
	mpir_cmd_tar := exec.Command("curl", "-L", mpir_tar_url, "-o", "package.tar.gz")
	err := mpir_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpir_zip_url := "https://web.archive.org/web/20220224004857/https://mpir.org/mpir-3.0.0.tar.bz2"
	mpir_cmd_zip := exec.Command("curl", "-L", mpir_zip_url, "-o", "package.zip")
	err = mpir_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpir_bin_url := "https://web.archive.org/web/20220224004857/https://mpir.org/mpir-3.0.0.tar.bz2"
	mpir_cmd_bin := exec.Command("curl", "-L", mpir_bin_url, "-o", "binary.bin")
	err = mpir_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpir_src_url := "https://web.archive.org/web/20220224004857/https://mpir.org/mpir-3.0.0.tar.bz2"
	mpir_cmd_src := exec.Command("curl", "-L", mpir_src_url, "-o", "source.tar.gz")
	err = mpir_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpir_cmd_direct := exec.Command("./binary")
	err = mpir_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: yasm")
	exec.Command("latte", "install", "yasm").Run()
}
