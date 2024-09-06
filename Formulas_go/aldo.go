package main

// Aldo - Morse code learning tool released under GPL
// Homepage: https://www.nongnu.org/aldo/

import (
	"fmt"
	
	"os/exec"
)

func installAldo() {
	// Método 1: Descargar y extraer .tar.gz
	aldo_tar_url := "https://savannah.nongnu.org/download/aldo/aldo-0.7.7.tar.bz2"
	aldo_cmd_tar := exec.Command("curl", "-L", aldo_tar_url, "-o", "package.tar.gz")
	err := aldo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aldo_zip_url := "https://savannah.nongnu.org/download/aldo/aldo-0.7.7.tar.bz2"
	aldo_cmd_zip := exec.Command("curl", "-L", aldo_zip_url, "-o", "package.zip")
	err = aldo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aldo_bin_url := "https://savannah.nongnu.org/download/aldo/aldo-0.7.7.tar.bz2"
	aldo_cmd_bin := exec.Command("curl", "-L", aldo_bin_url, "-o", "binary.bin")
	err = aldo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aldo_src_url := "https://savannah.nongnu.org/download/aldo/aldo-0.7.7.tar.bz2"
	aldo_cmd_src := exec.Command("curl", "-L", aldo_src_url, "-o", "source.tar.gz")
	err = aldo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aldo_cmd_direct := exec.Command("./binary")
	err = aldo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libao")
exec.Command("latte", "install", "libao")
}
