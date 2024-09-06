package main

// Libcsv - CSV library in ANSI C89
// Homepage: https://sourceforge.net/projects/libcsv/

import (
	"fmt"
	
	"os/exec"
)

func installLibcsv() {
	// Método 1: Descargar y extraer .tar.gz
	libcsv_tar_url := "https://downloads.sourceforge.net/project/libcsv/libcsv/libcsv-3.0.3/libcsv-3.0.3.tar.gz"
	libcsv_cmd_tar := exec.Command("curl", "-L", libcsv_tar_url, "-o", "package.tar.gz")
	err := libcsv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcsv_zip_url := "https://downloads.sourceforge.net/project/libcsv/libcsv/libcsv-3.0.3/libcsv-3.0.3.zip"
	libcsv_cmd_zip := exec.Command("curl", "-L", libcsv_zip_url, "-o", "package.zip")
	err = libcsv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcsv_bin_url := "https://downloads.sourceforge.net/project/libcsv/libcsv/libcsv-3.0.3/libcsv-3.0.3.bin"
	libcsv_cmd_bin := exec.Command("curl", "-L", libcsv_bin_url, "-o", "binary.bin")
	err = libcsv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcsv_src_url := "https://downloads.sourceforge.net/project/libcsv/libcsv/libcsv-3.0.3/libcsv-3.0.3.src.tar.gz"
	libcsv_cmd_src := exec.Command("curl", "-L", libcsv_src_url, "-o", "source.tar.gz")
	err = libcsv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcsv_cmd_direct := exec.Command("./binary")
	err = libcsv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
