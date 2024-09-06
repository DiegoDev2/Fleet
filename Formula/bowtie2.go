package main

// Bowtie2 - Fast and sensitive gapped read aligner
// Homepage: https://bowtie-bio.sourceforge.net/bowtie2/index.shtml

import (
	"fmt"
	
	"os/exec"
)

func installBowtie2() {
	// Método 1: Descargar y extraer .tar.gz
	bowtie2_tar_url := "https://github.com/BenLangmead/bowtie2/archive/refs/tags/v2.5.4.tar.gz"
	bowtie2_cmd_tar := exec.Command("curl", "-L", bowtie2_tar_url, "-o", "package.tar.gz")
	err := bowtie2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bowtie2_zip_url := "https://github.com/BenLangmead/bowtie2/archive/refs/tags/v2.5.4.zip"
	bowtie2_cmd_zip := exec.Command("curl", "-L", bowtie2_zip_url, "-o", "package.zip")
	err = bowtie2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bowtie2_bin_url := "https://github.com/BenLangmead/bowtie2/archive/refs/tags/v2.5.4.bin"
	bowtie2_cmd_bin := exec.Command("curl", "-L", bowtie2_bin_url, "-o", "binary.bin")
	err = bowtie2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bowtie2_src_url := "https://github.com/BenLangmead/bowtie2/archive/refs/tags/v2.5.4.src.tar.gz"
	bowtie2_cmd_src := exec.Command("curl", "-L", bowtie2_src_url, "-o", "source.tar.gz")
	err = bowtie2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bowtie2_cmd_direct := exec.Command("./binary")
	err = bowtie2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: simde")
	exec.Command("latte", "install", "simde").Run()
}
