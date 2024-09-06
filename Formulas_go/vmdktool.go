package main

// Vmdktool - Converts raw filesystems to VMDK files and vice versa
// Homepage: https://manned.org/vmdktool

import (
	"fmt"
	
	"os/exec"
)

func installVmdktool() {
	// Método 1: Descargar y extraer .tar.gz
	vmdktool_tar_url := "https://people.freebsd.org/~brian/vmdktool/vmdktool-1.4.tar.gz"
	vmdktool_cmd_tar := exec.Command("curl", "-L", vmdktool_tar_url, "-o", "package.tar.gz")
	err := vmdktool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vmdktool_zip_url := "https://people.freebsd.org/~brian/vmdktool/vmdktool-1.4.zip"
	vmdktool_cmd_zip := exec.Command("curl", "-L", vmdktool_zip_url, "-o", "package.zip")
	err = vmdktool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vmdktool_bin_url := "https://people.freebsd.org/~brian/vmdktool/vmdktool-1.4.bin"
	vmdktool_cmd_bin := exec.Command("curl", "-L", vmdktool_bin_url, "-o", "binary.bin")
	err = vmdktool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vmdktool_src_url := "https://people.freebsd.org/~brian/vmdktool/vmdktool-1.4.src.tar.gz"
	vmdktool_cmd_src := exec.Command("curl", "-L", vmdktool_src_url, "-o", "source.tar.gz")
	err = vmdktool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vmdktool_cmd_direct := exec.Command("./binary")
	err = vmdktool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: groff")
exec.Command("latte", "install", "groff")
}
