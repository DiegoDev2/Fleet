package main

// Vcftools - Tools for working with VCF files
// Homepage: https://vcftools.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installVcftools() {
	// Método 1: Descargar y extraer .tar.gz
	vcftools_tar_url := "https://github.com/vcftools/vcftools/releases/download/v0.1.16/vcftools-0.1.16.tar.gz"
	vcftools_cmd_tar := exec.Command("curl", "-L", vcftools_tar_url, "-o", "package.tar.gz")
	err := vcftools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vcftools_zip_url := "https://github.com/vcftools/vcftools/releases/download/v0.1.16/vcftools-0.1.16.zip"
	vcftools_cmd_zip := exec.Command("curl", "-L", vcftools_zip_url, "-o", "package.zip")
	err = vcftools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vcftools_bin_url := "https://github.com/vcftools/vcftools/releases/download/v0.1.16/vcftools-0.1.16.bin"
	vcftools_cmd_bin := exec.Command("curl", "-L", vcftools_bin_url, "-o", "binary.bin")
	err = vcftools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vcftools_src_url := "https://github.com/vcftools/vcftools/releases/download/v0.1.16/vcftools-0.1.16.src.tar.gz"
	vcftools_cmd_src := exec.Command("curl", "-L", vcftools_src_url, "-o", "source.tar.gz")
	err = vcftools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vcftools_cmd_direct := exec.Command("./binary")
	err = vcftools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: htslib")
exec.Command("latte", "install", "htslib")
}
