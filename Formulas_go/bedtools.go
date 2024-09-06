package main

// Bedtools - Tools for genome arithmetic (set theory on the genome)
// Homepage: https://github.com/arq5x/bedtools2

import (
	"fmt"
	
	"os/exec"
)

func installBedtools() {
	// Método 1: Descargar y extraer .tar.gz
	bedtools_tar_url := "https://github.com/arq5x/bedtools2/archive/refs/tags/v2.31.1.tar.gz"
	bedtools_cmd_tar := exec.Command("curl", "-L", bedtools_tar_url, "-o", "package.tar.gz")
	err := bedtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bedtools_zip_url := "https://github.com/arq5x/bedtools2/archive/refs/tags/v2.31.1.zip"
	bedtools_cmd_zip := exec.Command("curl", "-L", bedtools_zip_url, "-o", "package.zip")
	err = bedtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bedtools_bin_url := "https://github.com/arq5x/bedtools2/archive/refs/tags/v2.31.1.bin"
	bedtools_cmd_bin := exec.Command("curl", "-L", bedtools_bin_url, "-o", "binary.bin")
	err = bedtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bedtools_src_url := "https://github.com/arq5x/bedtools2/archive/refs/tags/v2.31.1.src.tar.gz"
	bedtools_cmd_src := exec.Command("curl", "-L", bedtools_src_url, "-o", "source.tar.gz")
	err = bedtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bedtools_cmd_direct := exec.Command("./binary")
	err = bedtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
