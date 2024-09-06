package main

// Guetzli - Perceptual JPEG encoder
// Homepage: https://github.com/google/guetzli

import (
	"fmt"
	
	"os/exec"
)

func installGuetzli() {
	// Método 1: Descargar y extraer .tar.gz
	guetzli_tar_url := "https://github.com/google/guetzli/archive/refs/tags/v1.0.1.tar.gz"
	guetzli_cmd_tar := exec.Command("curl", "-L", guetzli_tar_url, "-o", "package.tar.gz")
	err := guetzli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	guetzli_zip_url := "https://github.com/google/guetzli/archive/refs/tags/v1.0.1.zip"
	guetzli_cmd_zip := exec.Command("curl", "-L", guetzli_zip_url, "-o", "package.zip")
	err = guetzli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	guetzli_bin_url := "https://github.com/google/guetzli/archive/refs/tags/v1.0.1.bin"
	guetzli_cmd_bin := exec.Command("curl", "-L", guetzli_bin_url, "-o", "binary.bin")
	err = guetzli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	guetzli_src_url := "https://github.com/google/guetzli/archive/refs/tags/v1.0.1.src.tar.gz"
	guetzli_cmd_src := exec.Command("curl", "-L", guetzli_src_url, "-o", "source.tar.gz")
	err = guetzli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	guetzli_cmd_direct := exec.Command("./binary")
	err = guetzli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
