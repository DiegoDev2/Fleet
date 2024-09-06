package main

// Fasttext - Library for fast text representation and classification
// Homepage: https://fasttext.cc

import (
	"fmt"
	
	"os/exec"
)

func installFasttext() {
	// Método 1: Descargar y extraer .tar.gz
	fasttext_tar_url := "https://github.com/facebookresearch/fastText/archive/refs/tags/v0.9.2.tar.gz"
	fasttext_cmd_tar := exec.Command("curl", "-L", fasttext_tar_url, "-o", "package.tar.gz")
	err := fasttext_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fasttext_zip_url := "https://github.com/facebookresearch/fastText/archive/refs/tags/v0.9.2.zip"
	fasttext_cmd_zip := exec.Command("curl", "-L", fasttext_zip_url, "-o", "package.zip")
	err = fasttext_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fasttext_bin_url := "https://github.com/facebookresearch/fastText/archive/refs/tags/v0.9.2.bin"
	fasttext_cmd_bin := exec.Command("curl", "-L", fasttext_bin_url, "-o", "binary.bin")
	err = fasttext_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fasttext_src_url := "https://github.com/facebookresearch/fastText/archive/refs/tags/v0.9.2.src.tar.gz"
	fasttext_cmd_src := exec.Command("curl", "-L", fasttext_src_url, "-o", "source.tar.gz")
	err = fasttext_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fasttext_cmd_direct := exec.Command("./binary")
	err = fasttext_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
