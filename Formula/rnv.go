package main

// Rnv - Implementation of Relax NG Compact Syntax validator
// Homepage: https://sourceforge.net/projects/rnv/

import (
	"fmt"
	
	"os/exec"
)

func installRnv() {
	// Método 1: Descargar y extraer .tar.gz
	rnv_tar_url := "https://downloads.sourceforge.net/project/rnv/Sources/1.7.11/rnv-1.7.11.tar.bz2"
	rnv_cmd_tar := exec.Command("curl", "-L", rnv_tar_url, "-o", "package.tar.gz")
	err := rnv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rnv_zip_url := "https://downloads.sourceforge.net/project/rnv/Sources/1.7.11/rnv-1.7.11.tar.bz2"
	rnv_cmd_zip := exec.Command("curl", "-L", rnv_zip_url, "-o", "package.zip")
	err = rnv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rnv_bin_url := "https://downloads.sourceforge.net/project/rnv/Sources/1.7.11/rnv-1.7.11.tar.bz2"
	rnv_cmd_bin := exec.Command("curl", "-L", rnv_bin_url, "-o", "binary.bin")
	err = rnv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rnv_src_url := "https://downloads.sourceforge.net/project/rnv/Sources/1.7.11/rnv-1.7.11.tar.bz2"
	rnv_cmd_src := exec.Command("curl", "-L", rnv_src_url, "-o", "source.tar.gz")
	err = rnv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rnv_cmd_direct := exec.Command("./binary")
	err = rnv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: expat")
	exec.Command("latte", "install", "expat").Run()
}
