package main

// Advancecomp - Recompression utilities for .PNG, .MNG, .ZIP, and .GZ files
// Homepage: https://www.advancemame.it/comp-readme.html

import (
	"fmt"
	
	"os/exec"
)

func installAdvancecomp() {
	// Método 1: Descargar y extraer .tar.gz
	advancecomp_tar_url := "https://github.com/amadvance/advancecomp/releases/download/v2.6/advancecomp-2.6.tar.gz"
	advancecomp_cmd_tar := exec.Command("curl", "-L", advancecomp_tar_url, "-o", "package.tar.gz")
	err := advancecomp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	advancecomp_zip_url := "https://github.com/amadvance/advancecomp/releases/download/v2.6/advancecomp-2.6.zip"
	advancecomp_cmd_zip := exec.Command("curl", "-L", advancecomp_zip_url, "-o", "package.zip")
	err = advancecomp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	advancecomp_bin_url := "https://github.com/amadvance/advancecomp/releases/download/v2.6/advancecomp-2.6.bin"
	advancecomp_cmd_bin := exec.Command("curl", "-L", advancecomp_bin_url, "-o", "binary.bin")
	err = advancecomp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	advancecomp_src_url := "https://github.com/amadvance/advancecomp/releases/download/v2.6/advancecomp-2.6.src.tar.gz"
	advancecomp_cmd_src := exec.Command("curl", "-L", advancecomp_src_url, "-o", "source.tar.gz")
	err = advancecomp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	advancecomp_cmd_direct := exec.Command("./binary")
	err = advancecomp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
