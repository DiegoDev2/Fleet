package main

// Ht - Viewer/editor/analyzer for executables
// Homepage: https://hte.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installHt() {
	// Método 1: Descargar y extraer .tar.gz
	ht_tar_url := "https://downloads.sourceforge.net/project/hte/ht-source/ht-2.1.0.tar.bz2"
	ht_cmd_tar := exec.Command("curl", "-L", ht_tar_url, "-o", "package.tar.gz")
	err := ht_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ht_zip_url := "https://downloads.sourceforge.net/project/hte/ht-source/ht-2.1.0.tar.bz2"
	ht_cmd_zip := exec.Command("curl", "-L", ht_zip_url, "-o", "package.zip")
	err = ht_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ht_bin_url := "https://downloads.sourceforge.net/project/hte/ht-source/ht-2.1.0.tar.bz2"
	ht_cmd_bin := exec.Command("curl", "-L", ht_bin_url, "-o", "binary.bin")
	err = ht_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ht_src_url := "https://downloads.sourceforge.net/project/hte/ht-source/ht-2.1.0.tar.bz2"
	ht_cmd_src := exec.Command("curl", "-L", ht_src_url, "-o", "source.tar.gz")
	err = ht_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ht_cmd_direct := exec.Command("./binary")
	err = ht_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
}
