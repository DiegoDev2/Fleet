package main

// Libforensic1394 - Live memory forensics over IEEE 1394 (\
// Homepage: https://freddie.witherden.org/tools/libforensic1394/

import (
	"fmt"
	
	"os/exec"
)

func installLibforensic1394() {
	// Método 1: Descargar y extraer .tar.gz
	libforensic1394_tar_url := "https://freddie.witherden.org/tools/libforensic1394/releases/libforensic1394-0.2.tar.gz"
	libforensic1394_cmd_tar := exec.Command("curl", "-L", libforensic1394_tar_url, "-o", "package.tar.gz")
	err := libforensic1394_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libforensic1394_zip_url := "https://freddie.witherden.org/tools/libforensic1394/releases/libforensic1394-0.2.zip"
	libforensic1394_cmd_zip := exec.Command("curl", "-L", libforensic1394_zip_url, "-o", "package.zip")
	err = libforensic1394_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libforensic1394_bin_url := "https://freddie.witherden.org/tools/libforensic1394/releases/libforensic1394-0.2.bin"
	libforensic1394_cmd_bin := exec.Command("curl", "-L", libforensic1394_bin_url, "-o", "binary.bin")
	err = libforensic1394_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libforensic1394_src_url := "https://freddie.witherden.org/tools/libforensic1394/releases/libforensic1394-0.2.src.tar.gz"
	libforensic1394_cmd_src := exec.Command("curl", "-L", libforensic1394_src_url, "-o", "source.tar.gz")
	err = libforensic1394_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libforensic1394_cmd_direct := exec.Command("./binary")
	err = libforensic1394_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
