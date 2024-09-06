package main

// Kytea - Toolkit for analyzing text, especially Japanese and Chinese
// Homepage: https://www.phontron.com/kytea/

import (
	"fmt"
	
	"os/exec"
)

func installKytea() {
	// Método 1: Descargar y extraer .tar.gz
	kytea_tar_url := "https://www.phontron.com/kytea/download/kytea-0.4.7.tar.gz"
	kytea_cmd_tar := exec.Command("curl", "-L", kytea_tar_url, "-o", "package.tar.gz")
	err := kytea_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kytea_zip_url := "https://www.phontron.com/kytea/download/kytea-0.4.7.zip"
	kytea_cmd_zip := exec.Command("curl", "-L", kytea_zip_url, "-o", "package.zip")
	err = kytea_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kytea_bin_url := "https://www.phontron.com/kytea/download/kytea-0.4.7.bin"
	kytea_cmd_bin := exec.Command("curl", "-L", kytea_bin_url, "-o", "binary.bin")
	err = kytea_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kytea_src_url := "https://www.phontron.com/kytea/download/kytea-0.4.7.src.tar.gz"
	kytea_cmd_src := exec.Command("curl", "-L", kytea_src_url, "-o", "source.tar.gz")
	err = kytea_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kytea_cmd_direct := exec.Command("./binary")
	err = kytea_cmd_direct.Run()
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
