package main

// CBlosc2 - Fast, compressed, persistent binary data store library for C
// Homepage: https://www.blosc.org

import (
	"fmt"
	
	"os/exec"
)

func installCBlosc2() {
	// Método 1: Descargar y extraer .tar.gz
	cblosc2_tar_url := "https://github.com/Blosc/c-blosc2/archive/refs/tags/v2.15.1.tar.gz"
	cblosc2_cmd_tar := exec.Command("curl", "-L", cblosc2_tar_url, "-o", "package.tar.gz")
	err := cblosc2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cblosc2_zip_url := "https://github.com/Blosc/c-blosc2/archive/refs/tags/v2.15.1.zip"
	cblosc2_cmd_zip := exec.Command("curl", "-L", cblosc2_zip_url, "-o", "package.zip")
	err = cblosc2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cblosc2_bin_url := "https://github.com/Blosc/c-blosc2/archive/refs/tags/v2.15.1.bin"
	cblosc2_cmd_bin := exec.Command("curl", "-L", cblosc2_bin_url, "-o", "binary.bin")
	err = cblosc2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cblosc2_src_url := "https://github.com/Blosc/c-blosc2/archive/refs/tags/v2.15.1.src.tar.gz"
	cblosc2_cmd_src := exec.Command("curl", "-L", cblosc2_src_url, "-o", "source.tar.gz")
	err = cblosc2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cblosc2_cmd_direct := exec.Command("./binary")
	err = cblosc2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
