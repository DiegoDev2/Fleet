package main

// Samtools - Tools for manipulating next-generation sequencing data
// Homepage: https://www.htslib.org/

import (
	"fmt"
	
	"os/exec"
)

func installSamtools() {
	// Método 1: Descargar y extraer .tar.gz
	samtools_tar_url := "https://github.com/samtools/samtools/releases/download/1.20/samtools-1.20.tar.bz2"
	samtools_cmd_tar := exec.Command("curl", "-L", samtools_tar_url, "-o", "package.tar.gz")
	err := samtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	samtools_zip_url := "https://github.com/samtools/samtools/releases/download/1.20/samtools-1.20.tar.bz2"
	samtools_cmd_zip := exec.Command("curl", "-L", samtools_zip_url, "-o", "package.zip")
	err = samtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	samtools_bin_url := "https://github.com/samtools/samtools/releases/download/1.20/samtools-1.20.tar.bz2"
	samtools_cmd_bin := exec.Command("curl", "-L", samtools_bin_url, "-o", "binary.bin")
	err = samtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	samtools_src_url := "https://github.com/samtools/samtools/releases/download/1.20/samtools-1.20.tar.bz2"
	samtools_cmd_src := exec.Command("curl", "-L", samtools_src_url, "-o", "source.tar.gz")
	err = samtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	samtools_cmd_direct := exec.Command("./binary")
	err = samtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: htslib")
	exec.Command("latte", "install", "htslib").Run()
}
