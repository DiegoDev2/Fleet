package main

// Fmt - Open-source formatting library for C++
// Homepage: https://fmt.dev/

import (
	"fmt"
	
	"os/exec"
)

func installFmt() {
	// Método 1: Descargar y extraer .tar.gz
	fmt_tar_url := "https://github.com/fmtlib/fmt/releases/download/11.0.2/fmt-11.0.2.zip"
	fmt_cmd_tar := exec.Command("curl", "-L", fmt_tar_url, "-o", "package.tar.gz")
	err := fmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fmt_zip_url := "https://github.com/fmtlib/fmt/releases/download/11.0.2/fmt-11.0.2.zip"
	fmt_cmd_zip := exec.Command("curl", "-L", fmt_zip_url, "-o", "package.zip")
	err = fmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fmt_bin_url := "https://github.com/fmtlib/fmt/releases/download/11.0.2/fmt-11.0.2.zip"
	fmt_cmd_bin := exec.Command("curl", "-L", fmt_bin_url, "-o", "binary.bin")
	err = fmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fmt_src_url := "https://github.com/fmtlib/fmt/releases/download/11.0.2/fmt-11.0.2.zip"
	fmt_cmd_src := exec.Command("curl", "-L", fmt_src_url, "-o", "source.tar.gz")
	err = fmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fmt_cmd_direct := exec.Command("./binary")
	err = fmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
