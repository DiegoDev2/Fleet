package main

// Bam - Build system that uses Lua to describe the build process
// Homepage: https://matricks.github.io/bam/

import (
	"fmt"
	
	"os/exec"
)

func installBam() {
	// Método 1: Descargar y extraer .tar.gz
	bam_tar_url := "https://github.com/matricks/bam/archive/refs/tags/v0.5.1.tar.gz"
	bam_cmd_tar := exec.Command("curl", "-L", bam_tar_url, "-o", "package.tar.gz")
	err := bam_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bam_zip_url := "https://github.com/matricks/bam/archive/refs/tags/v0.5.1.zip"
	bam_cmd_zip := exec.Command("curl", "-L", bam_zip_url, "-o", "package.zip")
	err = bam_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bam_bin_url := "https://github.com/matricks/bam/archive/refs/tags/v0.5.1.bin"
	bam_cmd_bin := exec.Command("curl", "-L", bam_bin_url, "-o", "binary.bin")
	err = bam_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bam_src_url := "https://github.com/matricks/bam/archive/refs/tags/v0.5.1.src.tar.gz"
	bam_cmd_src := exec.Command("curl", "-L", bam_src_url, "-o", "source.tar.gz")
	err = bam_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bam_cmd_direct := exec.Command("./binary")
	err = bam_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
