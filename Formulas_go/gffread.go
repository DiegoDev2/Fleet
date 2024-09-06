package main

// Gffread - GFF/GTF format conversions, region filtering, FASTA sequence extraction
// Homepage: https://github.com/gpertea/gffread

import (
	"fmt"
	
	"os/exec"
)

func installGffread() {
	// Método 1: Descargar y extraer .tar.gz
	gffread_tar_url := "https://github.com/gpertea/gffread/releases/download/v0.12.7/gffread-0.12.7.tar.gz"
	gffread_cmd_tar := exec.Command("curl", "-L", gffread_tar_url, "-o", "package.tar.gz")
	err := gffread_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gffread_zip_url := "https://github.com/gpertea/gffread/releases/download/v0.12.7/gffread-0.12.7.zip"
	gffread_cmd_zip := exec.Command("curl", "-L", gffread_zip_url, "-o", "package.zip")
	err = gffread_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gffread_bin_url := "https://github.com/gpertea/gffread/releases/download/v0.12.7/gffread-0.12.7.bin"
	gffread_cmd_bin := exec.Command("curl", "-L", gffread_bin_url, "-o", "binary.bin")
	err = gffread_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gffread_src_url := "https://github.com/gpertea/gffread/releases/download/v0.12.7/gffread-0.12.7.src.tar.gz"
	gffread_cmd_src := exec.Command("curl", "-L", gffread_src_url, "-o", "source.tar.gz")
	err = gffread_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gffread_cmd_direct := exec.Command("./binary")
	err = gffread_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
