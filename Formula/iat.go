package main

// Iat - Converts many CD-ROM image formats to ISO9660
// Homepage: https://sourceforge.net/projects/iat.berlios/

import (
	"fmt"
	
	"os/exec"
)

func installIat() {
	// Método 1: Descargar y extraer .tar.gz
	iat_tar_url := "https://downloads.sourceforge.net/project/iat.berlios/iat-0.1.7.tar.bz2"
	iat_cmd_tar := exec.Command("curl", "-L", iat_tar_url, "-o", "package.tar.gz")
	err := iat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iat_zip_url := "https://downloads.sourceforge.net/project/iat.berlios/iat-0.1.7.tar.bz2"
	iat_cmd_zip := exec.Command("curl", "-L", iat_zip_url, "-o", "package.zip")
	err = iat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iat_bin_url := "https://downloads.sourceforge.net/project/iat.berlios/iat-0.1.7.tar.bz2"
	iat_cmd_bin := exec.Command("curl", "-L", iat_bin_url, "-o", "binary.bin")
	err = iat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iat_src_url := "https://downloads.sourceforge.net/project/iat.berlios/iat-0.1.7.tar.bz2"
	iat_cmd_src := exec.Command("curl", "-L", iat_src_url, "-o", "source.tar.gz")
	err = iat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iat_cmd_direct := exec.Command("./binary")
	err = iat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
