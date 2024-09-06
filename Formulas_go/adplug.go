package main

// Adplug - Free, hardware independent AdLib sound player library
// Homepage: https://adplug.github.io

import (
	"fmt"
	
	"os/exec"
)

func installAdplug() {
	// Método 1: Descargar y extraer .tar.gz
	adplug_tar_url := "https://github.com/adplug/adplug/releases/download/adplug-2.3.3/adplug-2.3.3.tar.bz2"
	adplug_cmd_tar := exec.Command("curl", "-L", adplug_tar_url, "-o", "package.tar.gz")
	err := adplug_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	adplug_zip_url := "https://github.com/adplug/adplug/releases/download/adplug-2.3.3/adplug-2.3.3.tar.bz2"
	adplug_cmd_zip := exec.Command("curl", "-L", adplug_zip_url, "-o", "package.zip")
	err = adplug_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	adplug_bin_url := "https://github.com/adplug/adplug/releases/download/adplug-2.3.3/adplug-2.3.3.tar.bz2"
	adplug_cmd_bin := exec.Command("curl", "-L", adplug_bin_url, "-o", "binary.bin")
	err = adplug_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	adplug_src_url := "https://github.com/adplug/adplug/releases/download/adplug-2.3.3/adplug-2.3.3.tar.bz2"
	adplug_cmd_src := exec.Command("curl", "-L", adplug_src_url, "-o", "source.tar.gz")
	err = adplug_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	adplug_cmd_direct := exec.Command("./binary")
	err = adplug_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libbinio")
exec.Command("latte", "install", "libbinio")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
