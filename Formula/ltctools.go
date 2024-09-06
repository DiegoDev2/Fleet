package main

// LtcTools - Tools to deal with linear-timecode (LTC)
// Homepage: https://github.com/x42/ltc-tools

import (
	"fmt"
	
	"os/exec"
)

func installLtcTools() {
	// Método 1: Descargar y extraer .tar.gz
	ltctools_tar_url := "https://github.com/x42/ltc-tools/archive/refs/tags/v0.7.0.tar.gz"
	ltctools_cmd_tar := exec.Command("curl", "-L", ltctools_tar_url, "-o", "package.tar.gz")
	err := ltctools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ltctools_zip_url := "https://github.com/x42/ltc-tools/archive/refs/tags/v0.7.0.zip"
	ltctools_cmd_zip := exec.Command("curl", "-L", ltctools_zip_url, "-o", "package.zip")
	err = ltctools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ltctools_bin_url := "https://github.com/x42/ltc-tools/archive/refs/tags/v0.7.0.bin"
	ltctools_cmd_bin := exec.Command("curl", "-L", ltctools_bin_url, "-o", "binary.bin")
	err = ltctools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ltctools_src_url := "https://github.com/x42/ltc-tools/archive/refs/tags/v0.7.0.src.tar.gz"
	ltctools_cmd_src := exec.Command("curl", "-L", ltctools_src_url, "-o", "source.tar.gz")
	err = ltctools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ltctools_cmd_direct := exec.Command("./binary")
	err = ltctools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jack")
	exec.Command("latte", "install", "jack").Run()
	fmt.Println("Instalando dependencia: libltc")
	exec.Command("latte", "install", "libltc").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
}
