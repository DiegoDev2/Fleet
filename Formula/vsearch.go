package main

// Vsearch - Versatile open-source tool for microbiome analysis
// Homepage: https://github.com/torognes/vsearch

import (
	"fmt"
	
	"os/exec"
)

func installVsearch() {
	// Método 1: Descargar y extraer .tar.gz
	vsearch_tar_url := "https://github.com/torognes/vsearch/archive/refs/tags/v2.28.1.tar.gz"
	vsearch_cmd_tar := exec.Command("curl", "-L", vsearch_tar_url, "-o", "package.tar.gz")
	err := vsearch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vsearch_zip_url := "https://github.com/torognes/vsearch/archive/refs/tags/v2.28.1.zip"
	vsearch_cmd_zip := exec.Command("curl", "-L", vsearch_zip_url, "-o", "package.zip")
	err = vsearch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vsearch_bin_url := "https://github.com/torognes/vsearch/archive/refs/tags/v2.28.1.bin"
	vsearch_cmd_bin := exec.Command("curl", "-L", vsearch_bin_url, "-o", "binary.bin")
	err = vsearch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vsearch_src_url := "https://github.com/torognes/vsearch/archive/refs/tags/v2.28.1.src.tar.gz"
	vsearch_cmd_src := exec.Command("curl", "-L", vsearch_src_url, "-o", "source.tar.gz")
	err = vsearch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vsearch_cmd_direct := exec.Command("./binary")
	err = vsearch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
