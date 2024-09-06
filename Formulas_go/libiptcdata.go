package main

// Libiptcdata - Virtual package provided by libiptcdata0
// Homepage: https://libiptcdata.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibiptcdata() {
	// Método 1: Descargar y extraer .tar.gz
	libiptcdata_tar_url := "https://downloads.sourceforge.net/project/libiptcdata/libiptcdata/1.0.4/libiptcdata-1.0.4.tar.gz"
	libiptcdata_cmd_tar := exec.Command("curl", "-L", libiptcdata_tar_url, "-o", "package.tar.gz")
	err := libiptcdata_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libiptcdata_zip_url := "https://downloads.sourceforge.net/project/libiptcdata/libiptcdata/1.0.4/libiptcdata-1.0.4.zip"
	libiptcdata_cmd_zip := exec.Command("curl", "-L", libiptcdata_zip_url, "-o", "package.zip")
	err = libiptcdata_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libiptcdata_bin_url := "https://downloads.sourceforge.net/project/libiptcdata/libiptcdata/1.0.4/libiptcdata-1.0.4.bin"
	libiptcdata_cmd_bin := exec.Command("curl", "-L", libiptcdata_bin_url, "-o", "binary.bin")
	err = libiptcdata_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libiptcdata_src_url := "https://downloads.sourceforge.net/project/libiptcdata/libiptcdata/1.0.4/libiptcdata-1.0.4.src.tar.gz"
	libiptcdata_cmd_src := exec.Command("curl", "-L", libiptcdata_src_url, "-o", "source.tar.gz")
	err = libiptcdata_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libiptcdata_cmd_direct := exec.Command("./binary")
	err = libiptcdata_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
