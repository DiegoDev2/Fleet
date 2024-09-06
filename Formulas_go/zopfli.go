package main

// Zopfli - New zlib (gzip, deflate) compatible compressor
// Homepage: https://github.com/google/zopfli

import (
	"fmt"
	
	"os/exec"
)

func installZopfli() {
	// Método 1: Descargar y extraer .tar.gz
	zopfli_tar_url := "https://github.com/google/zopfli/archive/refs/tags/zopfli-1.0.3.tar.gz"
	zopfli_cmd_tar := exec.Command("curl", "-L", zopfli_tar_url, "-o", "package.tar.gz")
	err := zopfli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zopfli_zip_url := "https://github.com/google/zopfli/archive/refs/tags/zopfli-1.0.3.zip"
	zopfli_cmd_zip := exec.Command("curl", "-L", zopfli_zip_url, "-o", "package.zip")
	err = zopfli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zopfli_bin_url := "https://github.com/google/zopfli/archive/refs/tags/zopfli-1.0.3.bin"
	zopfli_cmd_bin := exec.Command("curl", "-L", zopfli_bin_url, "-o", "binary.bin")
	err = zopfli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zopfli_src_url := "https://github.com/google/zopfli/archive/refs/tags/zopfli-1.0.3.src.tar.gz"
	zopfli_cmd_src := exec.Command("curl", "-L", zopfli_src_url, "-o", "source.tar.gz")
	err = zopfli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zopfli_cmd_direct := exec.Command("./binary")
	err = zopfli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
