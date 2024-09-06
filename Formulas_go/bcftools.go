package main

// Bcftools - Tools for BCF/VCF files and variant calling from samtools
// Homepage: https://www.htslib.org/

import (
	"fmt"
	
	"os/exec"
)

func installBcftools() {
	// Método 1: Descargar y extraer .tar.gz
	bcftools_tar_url := "https://github.com/samtools/bcftools/releases/download/1.20/bcftools-1.20.tar.bz2"
	bcftools_cmd_tar := exec.Command("curl", "-L", bcftools_tar_url, "-o", "package.tar.gz")
	err := bcftools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bcftools_zip_url := "https://github.com/samtools/bcftools/releases/download/1.20/bcftools-1.20.tar.bz2"
	bcftools_cmd_zip := exec.Command("curl", "-L", bcftools_zip_url, "-o", "package.zip")
	err = bcftools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bcftools_bin_url := "https://github.com/samtools/bcftools/releases/download/1.20/bcftools-1.20.tar.bz2"
	bcftools_cmd_bin := exec.Command("curl", "-L", bcftools_bin_url, "-o", "binary.bin")
	err = bcftools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bcftools_src_url := "https://github.com/samtools/bcftools/releases/download/1.20/bcftools-1.20.tar.bz2"
	bcftools_cmd_src := exec.Command("curl", "-L", bcftools_src_url, "-o", "source.tar.gz")
	err = bcftools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bcftools_cmd_direct := exec.Command("./binary")
	err = bcftools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: htslib")
exec.Command("latte", "install", "htslib")
}
