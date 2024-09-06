package main

// FastqTools - Small utilities for working with fastq sequence files
// Homepage: https://github.com/dcjones/fastq-tools

import (
	"fmt"
	
	"os/exec"
)

func installFastqTools() {
	// Método 1: Descargar y extraer .tar.gz
	fastqtools_tar_url := "https://github.com/dcjones/fastq-tools/archive/refs/tags/v0.8.3.tar.gz"
	fastqtools_cmd_tar := exec.Command("curl", "-L", fastqtools_tar_url, "-o", "package.tar.gz")
	err := fastqtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastqtools_zip_url := "https://github.com/dcjones/fastq-tools/archive/refs/tags/v0.8.3.zip"
	fastqtools_cmd_zip := exec.Command("curl", "-L", fastqtools_zip_url, "-o", "package.zip")
	err = fastqtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastqtools_bin_url := "https://github.com/dcjones/fastq-tools/archive/refs/tags/v0.8.3.bin"
	fastqtools_cmd_bin := exec.Command("curl", "-L", fastqtools_bin_url, "-o", "binary.bin")
	err = fastqtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastqtools_src_url := "https://github.com/dcjones/fastq-tools/archive/refs/tags/v0.8.3.src.tar.gz"
	fastqtools_cmd_src := exec.Command("curl", "-L", fastqtools_src_url, "-o", "source.tar.gz")
	err = fastqtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastqtools_cmd_direct := exec.Command("./binary")
	err = fastqtools_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
