package main

// Jhiccup - Measure pauses and stalls of an app's Java runtime platform
// Homepage: https://www.azul.com/products/components/jhiccup/

import (
	"fmt"
	
	"os/exec"
)

func installJhiccup() {
	// Método 1: Descargar y extraer .tar.gz
	jhiccup_tar_url := "https://www.azul.com/files/jHiccup-2.0.10-dist.zip"
	jhiccup_cmd_tar := exec.Command("curl", "-L", jhiccup_tar_url, "-o", "package.tar.gz")
	err := jhiccup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jhiccup_zip_url := "https://www.azul.com/files/jHiccup-2.0.10-dist.zip"
	jhiccup_cmd_zip := exec.Command("curl", "-L", jhiccup_zip_url, "-o", "package.zip")
	err = jhiccup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jhiccup_bin_url := "https://www.azul.com/files/jHiccup-2.0.10-dist.zip"
	jhiccup_cmd_bin := exec.Command("curl", "-L", jhiccup_bin_url, "-o", "binary.bin")
	err = jhiccup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jhiccup_src_url := "https://www.azul.com/files/jHiccup-2.0.10-dist.zip"
	jhiccup_cmd_src := exec.Command("curl", "-L", jhiccup_src_url, "-o", "source.tar.gz")
	err = jhiccup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jhiccup_cmd_direct := exec.Command("./binary")
	err = jhiccup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
