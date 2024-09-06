package main

// Libgetdata - Reference implementation of the Dirfile Standards
// Homepage: https://getdata.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibgetdata() {
	// Método 1: Descargar y extraer .tar.gz
	libgetdata_tar_url := "https://downloads.sourceforge.net/project/getdata/getdata/0.11.0/getdata-0.11.0.tar.xz"
	libgetdata_cmd_tar := exec.Command("curl", "-L", libgetdata_tar_url, "-o", "package.tar.gz")
	err := libgetdata_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgetdata_zip_url := "https://downloads.sourceforge.net/project/getdata/getdata/0.11.0/getdata-0.11.0.tar.xz"
	libgetdata_cmd_zip := exec.Command("curl", "-L", libgetdata_zip_url, "-o", "package.zip")
	err = libgetdata_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgetdata_bin_url := "https://downloads.sourceforge.net/project/getdata/getdata/0.11.0/getdata-0.11.0.tar.xz"
	libgetdata_cmd_bin := exec.Command("curl", "-L", libgetdata_bin_url, "-o", "binary.bin")
	err = libgetdata_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgetdata_src_url := "https://downloads.sourceforge.net/project/getdata/getdata/0.11.0/getdata-0.11.0.tar.xz"
	libgetdata_cmd_src := exec.Command("curl", "-L", libgetdata_src_url, "-o", "source.tar.gz")
	err = libgetdata_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgetdata_cmd_direct := exec.Command("./binary")
	err = libgetdata_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
