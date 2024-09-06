package main

// Fatsort - Sorts FAT16 and FAT32 partitions
// Homepage: https://fatsort.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installFatsort() {
	// Método 1: Descargar y extraer .tar.gz
	fatsort_tar_url := "https://downloads.sourceforge.net/project/fatsort/fatsort-1.6.5.640.tar.xz"
	fatsort_cmd_tar := exec.Command("curl", "-L", fatsort_tar_url, "-o", "package.tar.gz")
	err := fatsort_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fatsort_zip_url := "https://downloads.sourceforge.net/project/fatsort/fatsort-1.6.5.640.tar.xz"
	fatsort_cmd_zip := exec.Command("curl", "-L", fatsort_zip_url, "-o", "package.zip")
	err = fatsort_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fatsort_bin_url := "https://downloads.sourceforge.net/project/fatsort/fatsort-1.6.5.640.tar.xz"
	fatsort_cmd_bin := exec.Command("curl", "-L", fatsort_bin_url, "-o", "binary.bin")
	err = fatsort_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fatsort_src_url := "https://downloads.sourceforge.net/project/fatsort/fatsort-1.6.5.640.tar.xz"
	fatsort_cmd_src := exec.Command("curl", "-L", fatsort_src_url, "-o", "source.tar.gz")
	err = fatsort_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fatsort_cmd_direct := exec.Command("./binary")
	err = fatsort_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
}
