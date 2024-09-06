package main

// SblimSfcc - Project to enhance the manageability of GNU/Linux system
// Homepage: https://sourceforge.net/projects/sblim/

import (
	"fmt"
	
	"os/exec"
)

func installSblimSfcc() {
	// Método 1: Descargar y extraer .tar.gz
	sblimsfcc_tar_url := "https://downloads.sourceforge.net/project/sblim/sblim-sfcc/sblim-sfcc-2.2.8.tar.bz2"
	sblimsfcc_cmd_tar := exec.Command("curl", "-L", sblimsfcc_tar_url, "-o", "package.tar.gz")
	err := sblimsfcc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sblimsfcc_zip_url := "https://downloads.sourceforge.net/project/sblim/sblim-sfcc/sblim-sfcc-2.2.8.tar.bz2"
	sblimsfcc_cmd_zip := exec.Command("curl", "-L", sblimsfcc_zip_url, "-o", "package.zip")
	err = sblimsfcc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sblimsfcc_bin_url := "https://downloads.sourceforge.net/project/sblim/sblim-sfcc/sblim-sfcc-2.2.8.tar.bz2"
	sblimsfcc_cmd_bin := exec.Command("curl", "-L", sblimsfcc_bin_url, "-o", "binary.bin")
	err = sblimsfcc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sblimsfcc_src_url := "https://downloads.sourceforge.net/project/sblim/sblim-sfcc/sblim-sfcc-2.2.8.tar.bz2"
	sblimsfcc_cmd_src := exec.Command("curl", "-L", sblimsfcc_src_url, "-o", "source.tar.gz")
	err = sblimsfcc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sblimsfcc_cmd_direct := exec.Command("./binary")
	err = sblimsfcc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
