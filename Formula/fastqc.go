package main

// Fastqc - Quality control tool for high throughput sequence data
// Homepage: https://www.bioinformatics.babraham.ac.uk/projects/fastqc/

import (
	"fmt"
	
	"os/exec"
)

func installFastqc() {
	// Método 1: Descargar y extraer .tar.gz
	fastqc_tar_url := "https://www.bioinformatics.babraham.ac.uk/projects/fastqc/fastqc_v0.12.1.zip"
	fastqc_cmd_tar := exec.Command("curl", "-L", fastqc_tar_url, "-o", "package.tar.gz")
	err := fastqc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastqc_zip_url := "https://www.bioinformatics.babraham.ac.uk/projects/fastqc/fastqc_v0.12.1.zip"
	fastqc_cmd_zip := exec.Command("curl", "-L", fastqc_zip_url, "-o", "package.zip")
	err = fastqc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastqc_bin_url := "https://www.bioinformatics.babraham.ac.uk/projects/fastqc/fastqc_v0.12.1.zip"
	fastqc_cmd_bin := exec.Command("curl", "-L", fastqc_bin_url, "-o", "binary.bin")
	err = fastqc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastqc_src_url := "https://www.bioinformatics.babraham.ac.uk/projects/fastqc/fastqc_v0.12.1.zip"
	fastqc_cmd_src := exec.Command("curl", "-L", fastqc_src_url, "-o", "source.tar.gz")
	err = fastqc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastqc_cmd_direct := exec.Command("./binary")
	err = fastqc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
