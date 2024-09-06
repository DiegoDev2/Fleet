package main

// Libimagequant - Palette quantization library extracted from pnquant2
// Homepage: https://pngquant.org/lib/

import (
	"fmt"
	
	"os/exec"
)

func installLibimagequant() {
	// Método 1: Descargar y extraer .tar.gz
	libimagequant_tar_url := "https://github.com/ImageOptim/libimagequant/archive/refs/tags/4.3.3.tar.gz"
	libimagequant_cmd_tar := exec.Command("curl", "-L", libimagequant_tar_url, "-o", "package.tar.gz")
	err := libimagequant_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libimagequant_zip_url := "https://github.com/ImageOptim/libimagequant/archive/refs/tags/4.3.3.zip"
	libimagequant_cmd_zip := exec.Command("curl", "-L", libimagequant_zip_url, "-o", "package.zip")
	err = libimagequant_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libimagequant_bin_url := "https://github.com/ImageOptim/libimagequant/archive/refs/tags/4.3.3.bin"
	libimagequant_cmd_bin := exec.Command("curl", "-L", libimagequant_bin_url, "-o", "binary.bin")
	err = libimagequant_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libimagequant_src_url := "https://github.com/ImageOptim/libimagequant/archive/refs/tags/4.3.3.src.tar.gz"
	libimagequant_cmd_src := exec.Command("curl", "-L", libimagequant_src_url, "-o", "source.tar.gz")
	err = libimagequant_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libimagequant_cmd_direct := exec.Command("./binary")
	err = libimagequant_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cargo-c")
exec.Command("latte", "install", "cargo-c")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
