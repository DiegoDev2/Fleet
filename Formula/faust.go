package main

// Faust - Functional programming language for real time signal processing
// Homepage: https://faust.grame.fr

import (
	"fmt"
	
	"os/exec"
)

func installFaust() {
	// Método 1: Descargar y extraer .tar.gz
	faust_tar_url := "https://github.com/grame-cncm/faust/releases/download/2.72.14/faust-2.72.14.tar.gz"
	faust_cmd_tar := exec.Command("curl", "-L", faust_tar_url, "-o", "package.tar.gz")
	err := faust_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	faust_zip_url := "https://github.com/grame-cncm/faust/releases/download/2.72.14/faust-2.72.14.zip"
	faust_cmd_zip := exec.Command("curl", "-L", faust_zip_url, "-o", "package.zip")
	err = faust_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	faust_bin_url := "https://github.com/grame-cncm/faust/releases/download/2.72.14/faust-2.72.14.bin"
	faust_cmd_bin := exec.Command("curl", "-L", faust_bin_url, "-o", "binary.bin")
	err = faust_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	faust_src_url := "https://github.com/grame-cncm/faust/releases/download/2.72.14/faust-2.72.14.src.tar.gz"
	faust_cmd_src := exec.Command("curl", "-L", faust_src_url, "-o", "source.tar.gz")
	err = faust_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	faust_cmd_direct := exec.Command("./binary")
	err = faust_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libmicrohttpd")
	exec.Command("latte", "install", "libmicrohttpd").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
